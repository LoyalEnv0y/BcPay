package main

import (
	"fmt"
	"log"
	BcPay "main.go/src"
)

func main() {
	var (
		inAccount float64
		profit    float64
		targetDay float64
		function  int
	)

	fmt.Println("press 1 for calculating days, 2 for calculating profit, 3 for database recording")
	_, err := fmt.Scanf("%d\n", &function)
	if err != nil {
		log.Fatal(err)
	}

	switch function {
	case 1:
		fmt.Println("please enter CAPITAL IN ACCOUNT and the PROFIT you want to gain")

		_, err = fmt.Scanf("%f %f\n", &inAccount, &profit)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n\n", BcPay.InProfit(inAccount, profit))
	case 2:
		fmt.Println("please enter CAPITAL IN ACCOUNT and DAY COUNT you want to reach")

		_, err = fmt.Scanf("%f %f\n", &inAccount, &targetDay)
		if err != nil {
			log.Fatal(err)
		}

		result, _ := BcPay.InDays(inAccount, targetDay*15, false)
		fmt.Printf("%s\n", result)
	case 3:
		fmt.Println("Please enter the CAPITAL IN ACCOUNT")
		_, err = fmt.Scanf("%f", &inAccount)
		if err != nil {
			log.Fatal(err)
		}
		BcPay.DataRecorder(inAccount)
		fmt.Println("Successfully inserted into database")
	default:
		log.Fatal("unknown command")
	}

}
