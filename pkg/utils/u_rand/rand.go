package u_rand

import (
	"crypto/rand"
	"fmt"
	"math/big"
	mathRand "math/rand"
	"time"
)

// GenerateRandomNumbers /**
/**
 * @author 大菠萝
 * @description //指定范围内获取指定数量的不重复随机整数
 * @date 2:30 pm 8/23/23
 * @param 随机数数量 最小值 最大值
 * @return 整型切片
 **/
func GenerateRandomNumbers(count, min, max int) ([]int, error) {
	if count > max-min+1 {
		return nil, fmt.Errorf("count must be less than or equal to the range size")
	}

	rangeSize := max - min + 1
	//创建一个包含指定范围内所有整数的切片
	availableNumbers := make([]int, rangeSize)
	for i := 0; i < rangeSize; i++ {
		availableNumbers[i] = i + min
	}

	randomNumbers := make([]int, count)
	for i := 0; i < count; i++ {
		//从中随机选择一个整数
		randomIdx, err := rand.Int(rand.Reader, big.NewInt(int64(rangeSize-i)))
		if err != nil {
			return nil, err
		}
		idx := int(randomIdx.Int64())
		//添加到结果列表
		randomNumbers[i] = availableNumbers[idx]
		//将其从 availableNumbers 中移除
		availableNumbers[idx] = availableNumbers[rangeSize-i-1]
	}

	return randomNumbers, nil
}

func GenerateRandomNumbersFromZero(count, max int) ([]int, error) {
	randomSlice, err := GenerateRandomNumbers(count, 0, max)
	if err != nil {
		return nil, err
	}
	return randomSlice, nil
}

// GenerateRandomInt /**
/**
 * @author 大菠萝
 * @description // 这个函数首先计算出范围内的差值，然后生成随机的差值，最后将其加上 min 得到最终的随机整数
 * @date 2:13 pm 8/23/23
 * @param min 随机数的qid
 * @return
 **/
func GenerateRandomInt(min, max int) (int, error) {
	if min >= max {
		return 0, fmt.Errorf("min must be less than max")
	}

	diff := big.NewInt(int64(max - min))
	randomDiff, err := rand.Int(rand.Reader, diff)
	if err != nil {
		return 0, err
	}

	randomInt := int(randomDiff.Int64()) + min
	return randomInt, nil
}

func GenerateRandomIntFromZero(max int) (int, error) {
	randomInt, err := GenerateRandomInt(0, max)
	if err != nil {
		return 0, err
	}
	return randomInt, nil
}

func GenerateRandomPassword(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	mathRand.Seed(time.Now().UnixNano())
	password := make([]byte, length)
	for i := range password {
		password[i] = charset[mathRand.Intn(len(charset))]
	}
	return string(password)
}
