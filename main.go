package main

import "fmt"

func main() {

	var stocks = []int{17, 3, 6, 9, 15, 8, 6, 1, 10}
	stockPicker(stocks)
}

func stockPicker(stocks []int) {
	profit := 0
	var profitDays = [2]int{}
	for i := 0; i < len(stocks); i++ {

		for j := i + 1; j < len(stocks); j++ {
			if stocks[j] > stocks[i] {

				if (stocks[j] - stocks[i]) > profit {
					profit = stocks[j] - stocks[i]
					profitDays[0] = i
					profitDays[1] = j
				}
			}

		}
	}

	fmt.Println("ProfitDay:= ", profitDays)
}
