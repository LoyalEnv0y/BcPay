package BcPay

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const fundamental float64 = 500

func InProfit(inAccount, targetProfit float64) string {
	if inAccount < fundamental {
		return fmt.Sprintf("ERROR: your capital => %f is lower than standard needed limit => %f", inAccount, fundamental)
	}

	var totalOrders float64

	for i := inAccount; i < targetProfit+inAccount; i += (0.2 / 100) * i {
		totalOrders++
	}

	return fmt.Sprintf("Order count:\t%d\nDay count:\t%f", int(totalOrders), totalOrders/20)
}

func InDays(inAccount, targetDay float64) string {
	if targetDay < 1 || inAccount < fundamental {
		return fmt.Sprintf("ERROR: target day is lower than 1 OR capital in account is less then fundamental [%f]", fundamental)
	}

	var endCapital float64

	for i := inAccount; targetDay >= 0; i += (0.2 / 100) * i {
		endCapital = i
		targetDay--
	}
	return fmt.Sprintf("your capital in account will be %d that is a total of %f increase", int(endCapital), endCapital-inAccount)
}

func Steps(inAccount, target, increment float64) {
	var (
		step    int
		logs    string
		saveLog string
	)

	for target > inAccount {
		logs += fmt.Sprintf("step [%d] \t↓\n%s\ntarget: \t%f\n\n", step, InProfit(inAccount, increment), inAccount+increment)
		inAccount += increment
		step++
	}
	saveLog = strings.Trim(saveLog, "\n")
	fmt.Print(logs)

	fmt.Println("do you want to save them into a file? Y/N")
	_, err := fmt.Scanf("%s\n", &saveLog)
	if err != nil {
		log.Fatal(err)
	}
	if saveLog == "Y" || saveLog == "y" {
		Logger(logs)
	}
}

func Logger(logs string) {
	var (
		timeStamp = time.Now().Format("02-Jan-06 15-04-05")
		raw       = "C:\\Users\\cetin\\Desktop\\Logs\\"
		path      = raw + timeStamp + ".txt"
	)

	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		file, err2 := os.Create(path)
		if err2 != nil {
			log.Fatal(err2)
		}
		defer file.Close()

		_, err2 = file.WriteString(logs)
		if err2 != nil {
			log.Fatal(err2)
		}

		err2 = file.Sync()
		if err2 != nil{
			log.Fatal(err2)
		}
	}

	fmt.Printf("Created and saved file at %s in %s\n", timeStamp, raw)
}
