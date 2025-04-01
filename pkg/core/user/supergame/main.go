package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// Random Int returns a uniform random value in [0, max). It panics if max <= 0.
func Random(max int64) int64 {
	n, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		panic(err)
	}
	return n.Int64()
}

/* func main() {

	var (
		bet    = 0
		amount = 0
		// oddsSecond = [6]int{10, 15, 30, 45, 60, 150}
		hitRate = [3]int{0, 0, 0}
	)

	for i := 0; i < 100000; i++ {
		for p1 := 0; p1 < 6; p1++ {
			for p2 := 0; p2 < 6; p2++ {
				for p3 := 0; p3 < 6; p3++ {
					bet += 6000
					if p1 == p2 && p2 == p3 {
						amount += 1000 * 5 // 3个中5倍
						hitRate[2]++

						if p1 == 0 {
							amount += 1000 * 51
						}
					} else {
						if p1 == p2 || p2 == p3 || p1 == p3 {
							amount += 1000 * 3 // 2个中3倍
							amount += 1000 * 2 // 1个中2倍
							hitRate[1]++
						} else {
							amount += (1000 * 2) * 3 // 3个中2倍
							hitRate[0]++
						}
					}
				}
			}
		}
	}

	fmt.Println("bet:", bet, "amount:", amount)
	fmt.Println("hit1-rate:", hitRate[0])
	fmt.Println("hit2-rate:", hitRate[1])
	fmt.Println("hit3-rate:", hitRate[2])
	fmt.Println(float64(amount) / float64(bet))
} */

func main() {
	var (
		// 各科目得分
		scores = [10]float64{71, 53, 40, 67, 74, 71, 47, 75, 80, 0}
		// 各科目权重
		weights = [10]float64{120, 120, 100, 60, 60, 70, 50, 40, 40, 0}
		// 各科目加权得分
		weightedScores = [10]float64{}
		// 总分
		totalScore = float64(0)
	)

	for i := 0; i < 10; i++ {
		// 计算加权得分：(得分 + 15) * 权重 / 100
		weightedScores[i] = (scores[i] + 15) * weights[i] / 100
		totalScore += weightedScores[i]
	}

	// 加上体育40分
	totalScore += 40
	fmt.Printf("%-8f \n", totalScore)
}
