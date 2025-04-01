package player

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
	"time"

	"gitlab.casinovip.tech/minigame_backend/c_engine/pkg/log"
)

var EventRouter *Router

type Entry func(c Player, req []byte, msg *Message) (ty MessageType)

type Wrapper func(e Entry) Entry

type Router struct {
	handler map[uint32]Entry
	lock    *sync.RWMutex
}

func RegisterEntry(event uint32, e Entry) {

	if EventRouter == nil {
		EventRouter = &Router{lock: new(sync.RWMutex), handler: make(map[uint32]Entry)}
	}

	EventRouter.lock.Lock()
	defer EventRouter.lock.Unlock()

	if _, ok := EventRouter.handler[event]; ok {
		panic("Entry is already register")
	}

	EventRouter.handler[event] = e
}

func getEntry(event uint32) (Entry Entry, err error) {

	var (
		found bool
	)

	if Entry, found = EventRouter.handler[event]; !found {
		return nil, errors.New(fmt.Sprintf("no Entry: %v", event))
	}

	return
}

type Middleware func(Entry Entry) Entry

var (
	middles = []Middleware{wrapper, jwt, maintain, recovery}
)

func MiddleInject(Entry Entry) Entry {
	for i := range middles {
		Entry = middles[i](Entry)
	}
	return Entry
}

func wrapper(e Entry) Entry {

	return func(c Actuator, req []byte, msg *Message) MessageType {
		now := time.Now()
		messageType := e(c, req, msg)

		if time.Since(now) > 1000*time.Millisecond {
			log.Engine().Error("wrapper",
				log.FieldReqAID(fmt.Sprintf("%d", msg.Seq)),
				log.FieldEvent(fmt.Sprintf("%d", msg.Event)),
				log.FieldCost(time.Since(now)),
			)
		}
		return messageType
	}
}

func jwt(e Entry) Entry {

	return func(c Actuator, req []byte, msg *Message) MessageType {

		now := time.Now()

		if _, ok := ExcludeJwtEventMap[int(msg.Event)]; ok {
			messageType := e(c, req, msg)
			return messageType
		}

		if c.GetId() == "" {
			log.Engine().Error("请先登陆",
				log.FieldReqAID(fmt.Sprintf("%d", msg.Seq)),
				log.FieldEvent(fmt.Sprintf("%d", msg.Event)),
				log.FieldCost(time.Since(now)),
			)
			ErrCode(msg, data.UnValidRequest)
			return 0
		} else {
			messageType := e(c, req, msg)
			return messageType
		}
	}
}

func maintain(e Entry) Entry {

	return func(c Actuator, req []byte, msg *Message) MessageType {

		now := time.Now()
		if _, ok := ExcludeMaintainEventMap[int(msg.Event)]; ok {
			messageType := e(c, req, msg)
			return messageType
		}

		status, err := GetGameCore().getGameStatus()
		//log.Engine().Sugar().Infof("GetGameCore().getGameStatus=================status:%d\n", status)
		//log.Engine().Sugar().Infof("GetGameCore().getGameStatus=========err:%+v\n", err)
		if err != nil {
			ErrRes(msg, err)
			return 0
		}

		if status == constant.ClosedStatus || status == constant.MaintainStatus {
			log.Engine().Error("关闭维护",
				log.FieldReqAID(fmt.Sprintf("%d", msg.Seq)),
				log.FieldEvent(fmt.Sprintf("%d", msg.Event)),
				log.FieldCost(time.Since(now)),
			)

			msg.Event = EventAvMaintain
			msg.Body = MaintainBoardRes{
				Status:         1,
				MaintainStatus: status,
			}

			return 0
		}

		messageType := e(c, req, msg)
		return messageType

	}
}

func recovery(e Entry) Entry {
	return func(c Actuator, req []byte, msg *Message) MessageType {
		defer func() {
			if err := recover(); err != nil {
				var buf [4096]byte
				n := runtime.Stack(buf[:], false)
				tmpStr := fmt.Sprintf("err=%v panic ==> %s\n", err, string(buf[:n]))
				log.Engine().Sugar().Errorf(tmpStr)
				ErrRes(msg, err.(error))
			}
		}()
		return e(c, req, msg)
	}
}
