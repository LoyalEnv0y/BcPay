package BcPay

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Data struct {
	Capital      float64
	Profit       float64
	Day          float64
	InterestRate float64
	DailyOrders  float64
	TotalOrders  float64
	function     int
}

const fundamental float64 = 500

var (
	lowCapital  = errors.New(fmt.Sprintf("Your capital in account is lower then standard needed limit => %f", fundamental))
	lowDayCount = errors.New(fmt.Sprintf("target day is lower than 1 OR capital in account is less then fundamental [%f]", fundamental))
)

func Init(capitalInAccount, targetProfit, targetDay float64) Data {
	customer := Data{
		Capital:      capitalInAccount,
		Profit:       targetProfit,
		Day:          targetDay,
		InterestRate: 0.22,
		DailyOrders:  13.0,
	}

	return customer
}

func InProfit(customer *Data) error {
	if customer.Capital < fundamental {
		return lowCapital
	}

	for i := customer.Capital; i < customer.Profit+customer.Capital; i += (customer.InterestRate / 100) * i {
		customer.TotalOrders++
	}
	return nil
}

func InDays(customer *Data) (float64, error) {
	if customer.Day < 1 {
		return 0, lowDayCount
	} else if customer.Capital < fundamental {
		return 0, lowCapital
	}

	adjusted := customer.Day * customer.DailyOrders
	startingCapital := customer.Capital

	for i := customer.Capital; adjusted >= 0; i += (customer.InterestRate / 100) * i {
		adjusted--
		customer.Capital = i
	}
	return startingCapital, nil
}

func DataRecorder(customer *Data) error {
	if customer.Capital < fundamental {
		return lowCapital
	}

	startingCapital, err := InDays(customer)
	if err != nil {
		return err
	}

	db, err := sql.Open("mysql", "root:testPassword@tcp(127.0.0.1:3306)/bcpay")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func(db *sql.DB) {
		err2 := db.Close()
		if err2 != nil {
			log.Fatal(err2.Error())
		}
	}(db)

	insert, err3 := db.Query("INSERT INTO records (started, target, daily_gain) VALUES(?,?, CONCAT(? - ?, '$'))", startingCapital, int(customer.Capital), int(customer.Capital), int(startingCapital))
	if err3 != nil {
		log.Fatal(err3.Error())
	}
	defer func(db *sql.DB) {
		err4 := insert.Close()
		if err4 != nil {
			log.Fatal(err4.Error())
		}
	}(db)
	return nil
}

func Stringer(function int, customer Data, startingCapital float64) string {
	switch function {
	case 1:
		return fmt.Sprintf("Order count:\t%d\nDay count:\t%f\nFinal capital:\t%d", int(customer.TotalOrders), customer.TotalOrders/customer.DailyOrders, int(customer.Capital+customer.Profit))
	case 2:
		return fmt.Sprintf("Final capital:\t%d\nincrease:\t%f", int(customer.Capital), customer.Capital-startingCapital)
	case 3:
		return fmt.Sprintf("Successfully inserted into database")
	}
	return ""
}
