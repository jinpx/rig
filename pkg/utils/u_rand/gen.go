package u_rand

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// WeightedRandomIndex /**
/**
 * @author 权重随机索引
 **/
func WeightedRandomIndex(weights map[int]int, minKey, maxKey int) (int, error) {
	// 权重计算
	// 1. 加总权重
	totalWeight := int64(0)
	for _, weight := range weights {
		if weight < 0 {
			return 0, fmt.Errorf("negative weight error")
		}
		totalWeight += int64(weight)
	}
	if totalWeight == 0 {
		return 0, fmt.Errorf("total weight is zero")
	}

	// 2. 产出随机数
	randValue := 0
	n, err := rand.Int(rand.Reader, big.NewInt(totalWeight))
	if err != nil {
		return 0, fmt.Errorf("rand error: %v", err)
	}

	// 随机值
	randValue = int(n.Int64())

	// 3. 选出此次权重指标
	for i := minKey; i <= maxKey; i++ {
		weight, ok := weights[i]
		if !ok {
			continue
		}
		if randValue < weight {
			return i, nil
		} else {
			randValue -= weight
		}
	}

	return 0, fmt.Errorf("unable to calculate weight index")
}
