package trading

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

type Portfolio struct {
	shares        int     // 持有股数
	price         float64 // 当前股价
	totalValue    float64 // 总市值
	cash          float64 // 持有现金
	dropDownCount int     // 连续下跌次数
	upCount       int     // 连续上涨次数
}

func TestTrading(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	portfolio := &Portfolio{
		shares:        100000,
		price:         1.0,
		totalValue:    100000,
		cash:          1000000, // 初始现金充足
		dropDownCount: 0,       // 初始下跌次数为0
		upCount:       0,       // 初始上涨次数为0
	}

	fmt.Println("交易次数,随机数,原始股价,涨跌幅(%),新股价,下跌次数,上涨次数,补仓目标金额,补充股份,补充金额,现在持有股份,现在总市值,卖出股份,卖出收入,累计现金,总资产")

	for i := 1; i <= 30; i++ {
		isUp := rand.Intn(2)
		prevPrice := portfolio.price

		change := -0.20 // 跌30%
		if isUp == 1 {
			change = 0.20 // 涨30%
		}
		newPrice := portfolio.price * (1 + change)
		if newPrice <= 0 {
			newPrice = 0.01
		}

		var needToAdd, currentValue, saleProceeds float64
		var newShares, targetShares, sharesToSell int
		var targetValue float64 = 100000 // 基础目标金额

		if change < 0 {
			portfolio.dropDownCount++ // 增加下跌计数
			portfolio.upCount = 0     // 重置上涨计数
			currentValue = float64(portfolio.shares) * newPrice

			// 下跌时逐步增加补仓金额
			if portfolio.dropDownCount == 1 {
				targetValue = 100000 // 第一次跌，补到10万
			} else {
				// 从第二次开始，每次增加10%
				targetValue = 100000 * (1 + float64(portfolio.dropDownCount-1)*0.1)
			}

			needToAdd = targetValue - currentValue
			newShares = int(math.Ceil(needToAdd / newPrice))
			cost := float64(newShares) * newPrice

			if portfolio.cash < cost {
				newShares = int(portfolio.cash / newPrice)
				cost = float64(newShares) * newPrice
			}

			portfolio.shares += newShares
			portfolio.cash -= cost
			portfolio.price = newPrice
			portfolio.totalValue = float64(portfolio.shares) * newPrice
		} else {
			portfolio.upCount++         // 增加上涨计数
			portfolio.dropDownCount = 0 // 重置下跌计数

			// 上涨时逐步增加卖出比例
			baseTarget := 100000.0
			if portfolio.upCount > 1 {
				// 从第二次上涨开始，每次多卖出10%
				baseTarget = baseTarget * (1 - float64(portfolio.upCount-1)*0.1)
				if baseTarget < 50000 { // 设置最低持仓金额
					baseTarget = 50000
				}
			}

			targetShares = int(math.Floor(baseTarget / newPrice))
			sharesToSell = portfolio.shares - targetShares
			if sharesToSell < 0 {
				sharesToSell = 0
			}
			saleProceeds = float64(sharesToSell) * newPrice

			portfolio.shares = targetShares
			portfolio.cash += saleProceeds
			portfolio.price = newPrice
			portfolio.totalValue = float64(portfolio.shares) * newPrice
		}

		fmt.Printf("%d,%d,%.4f,%.2f,%.4f,%d,%d,%.2f,%d,%.2f,%d,%.2f,%d,%.2f,%.2f,%.2f\n",
			i, isUp, prevPrice, change*100, newPrice, portfolio.dropDownCount, portfolio.upCount, targetValue,
			newShares, needToAdd, portfolio.shares, portfolio.totalValue,
			sharesToSell, saleProceeds, portfolio.cash, portfolio.totalValue+portfolio.cash,
		)

		time.Sleep(time.Millisecond * 10)
	}

	fmt.Printf("\n最终结果,,%.4f,,,,,,,%d,%.2f,,%.2f,%.2f\n",
		portfolio.price, portfolio.shares, portfolio.totalValue, portfolio.cash, portfolio.totalValue+portfolio.cash,
	)
}
