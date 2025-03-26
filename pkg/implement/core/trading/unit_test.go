package trading

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type Portfolio struct {
	shares     int     // 持有股数
	price      float64 // 当前股价
	totalValue float64 // 总市值
	cash       float64 // 持有现金
}

func TestTrading(t *testing.T) {
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	// 初始化投资组合
	portfolio := &Portfolio{
		shares:     100000, // 初始10万股
		price:      1.0,    // 初始股价1元
		totalValue: 100000, // 初始总市值10万
		cash:       0,      // 初始现金0元
	}

	// 模拟交易
	for i := 1; i <= 20; { // 模拟10次涨跌
		// 生成随机涨跌，0表示跌30%，1表示涨30%
		isUp := rand.Intn(2)
		fVal := rand.Float64()

		change := -(0.10*fVal + 0.2) // 默认跌30%
		if isUp == 1 {
			change = 0.10*fVal + 0.2 // 涨30%
		}
		newPrice := portfolio.price * (1 + change)

		fmt.Printf("\n第%d次交易:\n", i)
		fmt.Printf("随机数: %d\n", isUp)
		fmt.Printf("原始股价: %.4f\n", portfolio.price)
		fmt.Printf("涨跌幅: %.2f%%\n", change*100)
		fmt.Printf("新股价: %.4f\n", newPrice)

		if change < 0 { // 股价下跌
			// 计算需要补充的资金
			currentValue := float64(portfolio.shares) * newPrice
			needToAdd := 100000 - currentValue
			newShares := int(needToAdd / newPrice)

			portfolio.shares += newShares
			portfolio.price = newPrice
			portfolio.totalValue = float64(portfolio.shares) * newPrice

			fmt.Printf("下跌后总市值: %.2f\n", currentValue)
			fmt.Printf("补充股份: %d\n", newShares)
			fmt.Printf("补充金额: %.2f\n", needToAdd)
			fmt.Printf("现在持有股份: %d\n", portfolio.shares)
			fmt.Printf("现在总市值: %.2f\n", portfolio.totalValue)

		} else { // 股价上涨
			// 计算需要卖出的股份
			targetShares := int(100000 / newPrice)
			sharesToSell := portfolio.shares - targetShares
			saleProceeds := float64(sharesToSell) * newPrice

			portfolio.shares = targetShares
			portfolio.price = newPrice
			portfolio.totalValue = float64(portfolio.shares) * newPrice
			portfolio.cash += saleProceeds

			fmt.Printf("上涨后总市值: %.2f\n", float64(portfolio.shares)*newPrice+portfolio.cash)
			fmt.Printf("卖出股份: %d\n", sharesToSell)
			fmt.Printf("卖出收入: %.2f\n", saleProceeds)
			fmt.Printf("现在持有股份: %d\n", portfolio.shares)
			fmt.Printf("现在总市值: %.2f\n", portfolio.totalValue)
			fmt.Printf("累计现金: %.2f\n", portfolio.cash)
		}

		i++
		time.Sleep(time.Second) // 暂停1秒
	}

	// 输出最终结果
	fmt.Printf("\n最终结果:\n")
	fmt.Printf("最终持有股份: %d\n", portfolio.shares)
	fmt.Printf("最终股价: %.4f\n", portfolio.price)
	fmt.Printf("最终市值: %.2f\n", portfolio.totalValue)
	fmt.Printf("最终现金: %.2f\n", portfolio.cash)
	fmt.Printf("总资产: %.2f\n", portfolio.totalValue+portfolio.cash)
}
