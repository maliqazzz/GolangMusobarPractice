package main

import (
	"fmt"
	"time"
)

func main() {
	// t := time.Now()

	// fmt.Println(t)
	// fmt.Println(t.Year())
	// fmt.Println(t.Month())
	// fmt.Println(t.Day())

	fmt.Println("============================")

	timeString := "January 20, 2018"
	form := "January 2, 2006"
	// tamplate wajib seperti ini

	resTime, err := time.Parse(form, timeString)
	// fungsi parse time membutuhkan dua parameter
	// timeString dan form (form sebagai tamplate/perbandingan)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resTime)
	fmt.Println(resTime.Year())

	// t1 := time.Date(2015, 5, 31, 18, 0, 0, 0, time.UTC)
	// fmt.Println(t1)
	// untuk membuat waktu secara manual

	t1 := time.Date(2015, 5, 31, 18, 0, 0, 0, time.UTC)
	t2 := time.Date(2015, 5, 31, 18, 0, 0, 0, time.UTC)

	eq := t1.Equal(t2)

	fmt.Println(t1)
	fmt.Println(eq)

	//fungsi format, kebalikan dari parse
	// mengkonversi time menjadi string
	layout := "2006-02-01"
	// sebagai tamplate
	timeNow := time.Now()

	resString := timeNow.Format(layout)

	fmt.Println(resString)
}
