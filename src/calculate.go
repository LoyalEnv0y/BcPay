package BcPay

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const fundamental float64 = 500

func InProfit(inAccount, targetProfit float64) string {
	if inAccount < fundamental {
		return fmt.Sprintf("ERROR: your capital => %f is lower than standard needed limit => %f", inAccount, fundamental)
	}

	var totalOrders float64

	for i := inAccount; i < targetProfit+inAccount; i += (0.22 / 100) * i {
		totalOrders++
	}

	return fmt.Sprintf("Order count:\t%d\nDay count:\t%f\nFinal capital:\t%d", int(totalOrders), totalOrders/15, int(inAccount+targetProfit))
}

func InDays(inAccount, targetDay float64, database bool) (string, float64) {
	if targetDay < 1 || inAccount < fundamental {
		return fmt.Sprintf("ERROR: target day is lower than 1 OR capital in account is less then fundamental [%f]", fundamental), 0
	}

	var endCapital float64

	for i := inAccount; targetDay >= 0; i += (0.22 / 100) * i {
		endCapital = i
		targetDay--
	}

	if database {
		return "", endCapital
	}

	return fmt.Sprintf("Final capital:\t%d\nincrease:\t%f", int(endCapital), endCapital-inAccount), 0
}

func DataRecorder(inAccount float64) {
	_, endCapital := InDays(inAccount, 15, true)

	db, err := sql.Open("mysql", "root:Ct145353.@tcp(127.0.0.1:3306)/bcpay")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func(db *sql.DB) {
		err2 := db.Close()
		if err2 != nil {
			log.Fatal(err2.Error())
		}
	}(db)

	insert, err3 := db.Query("INSERT INTO records (started, target) VALUES(?,?)", inAccount, endCapital)
	if err3 != nil {
		log.Fatal(err3.Error())
	}
	defer func(db *sql.DB) {
		err4 := insert.Close()
		if err4 != nil {
			log.Fatal(err4.Error())
		}
	}(db)
}
