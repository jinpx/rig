package u_snowflake

import (
	"fmt"
	"strconv"
	"time"

	"github.com/sony/sonyflake"

	"rig/pkg/utils/u_rand"
)

const (
	retry    = 10
	idLength = 16
)

var gSnowflake *sonyflake.Sonyflake

func init() {

	var st sonyflake.Settings

	gSnowflake = sonyflake.NewSonyflake(st)
}

// 感觉效率不高.... try try
func SnowflakeId() int64 {

	var (
		id  uint64
		err error
	)

	for i := 0; i < retry; i++ {
		id, err = gSnowflake.NextID()
		if err != nil {
			continue
		}

		break
	}

	idStr := strconv.FormatUint(id, 10)

	if id == 0 {

		randomInt, _ := u_rand.GenerateRandomInt(1, 6)
		idStr = strconv.FormatInt(time.Now().Unix(), 10) + fmt.Sprintf("%+v", randomInt)
	} else {
		delta := idLength - len(idStr)
		if delta > 0 {
			randomInt, _ := u_rand.GenerateRandomInt(1, delta)
			idStr += fmt.Sprintf("%+v", randomInt)
		}
	}

	r, _ := strconv.ParseInt(idStr, 10, 64)
	return r
}

func SnowflakeIdStr() string {

	var (
		id  uint64
		err error
	)

	for i := 0; i < retry; i++ {
		id, err = gSnowflake.NextID()
		if err != nil {
			continue
		}

		break
	}

	idStr := strconv.FormatUint(id, 10)

	if id == 0 {

		randomInt, _ := u_rand.GenerateRandomInt(1, 6)
		idStr = strconv.FormatInt(time.Now().Unix(), 10) + fmt.Sprintf("%+v", randomInt)
	} else {
		delta := idLength - len(idStr)
		if delta > 0 {
			randomInt, _ := u_rand.GenerateRandomInt(1, delta)
			idStr += fmt.Sprintf("%+v", randomInt)
		}
	}

	return idStr
}
