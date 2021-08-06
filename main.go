package main

import (
	"fmt"
	"log"
	BcPay "main.go/src"
)

func main() {
	var (
		capital      float64
		targetProfit float64
		targetDay    float64
		interestRate = 0.22
		dailyOrders  = 13.0
		function     int
	)

	fmt.Println("press 1 for calculating days, 2 for calculating Profit, 3 for database recording")
	_, err := fmt.Scanf("%d\n", &function)
	if err != nil {
		log.Fatal(err)
	}

	switch function {
	case 1:
		fmt.Println("please enter CAPITAL IN ACCOUNT and the PROFIT you want to gain")

		_, err = fmt.Scanf("%f %f\n", &capital, &targetProfit)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n\n", BcPay.InProfit(capital, targetProfit, interestRate, dailyOrders))
	case 2:
		fmt.Println("please enter CAPITAL IN ACCOUNT and DAY COUNT you want to reach")

		_, err = fmt.Scanf("%f %f\n", &capital, &targetDay)
		if err != nil {
			log.Fatal(err)
		}

		result, _ := BcPay.InDays(capital, targetDay*dailyOrders, interestRate, false)
		fmt.Printf("%s\n", result)
	case 3:
		fmt.Println("Please enter the CAPITAL IN ACCOUNT")
		_, err = fmt.Scanf("%f", &capital)
		if err != nil {
			log.Fatal(err)
		}
		BcPay.DataRecorder(capital, interestRate, dailyOrders)
		fmt.Println("Successfully inserted into database")
	default:
		log.Fatal("unknown command")
	}

}
