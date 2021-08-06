package BcPay

import (
	"fmt"
	"log"
)

func IO() {
	var (
		capital  float64
		profit   float64
		day      float64
		function int
	)

	fmt.Println("press 1 for calculating days, 2 for calculating profit, 3 for database recording")
	_, err := fmt.Scanf("%d\n", &function)
	if err != nil {
		log.Fatal(err)
	}

	switch function {
	case 1:
		fmt.Println("please enter CAPITAL IN ACCOUNT and the PROFIT you want to gain")

		_, err = fmt.Scanf("%f %f\n", &capital, &profit)
		if err != nil {
			log.Fatal(err)
		}

		customer := Init(capital, profit, 0)
		err = InProfit(&customer)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(Stringer(function, customer, 0))
	case 2:
		fmt.Println("please enter CAPITAL IN ACCOUNT and DAY COUNT you want to reach")

		_, err = fmt.Scanf("%f %f\n", &capital, &day)
		if err != nil {
			log.Fatal(err)
		}

		customer := Init(capital, 0, day)
		startingCapital, err2 := InDays(&customer)
		if err2 != nil {
			log.Fatal(err2)
		}

		fmt.Printf(Stringer(function, customer, startingCapital))
	case 3:
		fmt.Println("Please enter the CAPITAL IN ACCOUNT")

		_, err = fmt.Scanf("%f", &capital)
		if err != nil {
			log.Fatal(err)
		}

		customer := Init(capital, 0, 1)
		err = DataRecorder(&customer)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(Stringer(function, customer, 0))
	default:
		log.Fatal("unknown command")
	}
}
