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
		increment float64
		function  int
	)

	fmt.Println("press 1 for calculating days, 2 for calculating profit, 3 for calculating steps")
	_, err := fmt.Scanf("%d\n", &function)
	if err != nil {
		log.Fatal(err)
	}

	switch function {
	case 1:
		fmt.Println("please enter CAPITAL IN ACCOUNT and the PROFIT you want to gain")

		_, err2 := fmt.Scanf("%f %f\n", &inAccount, &profit)
		if err2 != nil {
			log.Fatal(err2)
		}
		fmt.Printf("%s\n\n", BcPay.InProfit(inAccount, profit))

	case 2:
		fmt.Println("please enter CAPITAL IN ACCOUNT and DAY COUNT you want to reach")

		_, err3 := fmt.Scanf("%f %f\n", &inAccount, &targetDay)
		if err3 != nil {
			log.Fatal(err3)
		}
		fmt.Printf("%s\n", BcPay.InDays(inAccount, targetDay*20))

	case 3:
		fmt.Println("Please type CAPITAL, TARGET and INCREMENTER")

		_, err4 := fmt.Scanf("%f %f %f\n", &increment, &profit, &inAccount)
		if err4 != nil {
			log.Fatal(err4)
		}
		BcPay.Steps(inAccount, profit, increment)

	default:
		log.Fatal("unknown command")
	}

}