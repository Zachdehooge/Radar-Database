package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func exeTime(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s execution time: %v\n", name, time.Since(start))
	}
}

func main() {

	defer exeTime("main")()

	var month, day, year, radar string
	var timeStart, timeEnd string
	//var fullDay string

	fmt.Print("Enter Month: ")
	fmt.Scanln(&month)
	num1, _ := strconv.Atoi(month)
	if num1 < 10 {
		month = fmt.Sprintf("%02d", num1)
	}

	fmt.Print("Enter Day: ")
	fmt.Scanln(&day)
	num2, _ := strconv.Atoi(day)
	if num2 < 10 {
		day = fmt.Sprintf("%02d", num2)
	}

	fmt.Print("Enter Year: ")
	fmt.Scanln(&year)

	fmt.Print("Enter Radar: ")
	fmt.Scanln(&radar)

	// TODO: Have user prompted with Full day download
	// TODO: Handle if else where y timeStart and timeEnd are 000000 and 235959 respectively for y (yes) and continue prompts for user as normal if no or anything other than y (yes)
	/* 	fmt.Print("Download Full Day?[y/n]")
	   	fmt.Scanln(&fullDay)
	   	fullDay1 := strings.ToLower(fullDay)
	   	if fullDay1 == "y" {
	   		timeStart = "000000"
	   		timeEnd = "000000"
	   	} else {
	   		fmt.Print("Time Start in Zulu (HHMMSS): ")
	   		fmt.Scanln(&timeStart)
	   		// Fixed octal issue by dereferencing the pointer and converting to integer from string that is passed to the CLI
	   		test := timeStart
	   		timeStart = strconv.Atoi(timeStart)

	   		fmt.Print("Time End in Zulu (HHMMSS): ")
	   		fmt.Scanln(&timeEnd)
	   		test3 := timeEnd
	   		test4, _ := strconv.Atoi(test3)
	   	} */

	fmt.Print("Time Start in Zulu (HHMMSS): ")
	fmt.Scanln(&timeStart)
	// Fixed octal issue by dereferencing the pointer and converting to integer from string that is passed to the CLI
	test := timeStart
	test1, _ := strconv.Atoi(test)

	fmt.Print("Time End in Zulu (HHMMSS): ")
	fmt.Scanln(&timeEnd)
	test3 := timeEnd
	test4, _ := strconv.Atoi(test3)

	for x := test1; x <= test4; x++ {

		timeComb := fmt.Sprintf("%06d", x)

		url := fmt.Sprintf("https://noaa-nexrad-level2.s3.amazonaws.com/%s/%s/%s/%s/%s%s%s%s_%s_V06", year, month, day, radar, radar, year, month, day, timeComb)

		folderLocation := fmt.Sprintf("%s_%s_%s_%s", day, month, year, radar)
		if _, err := os.Stat(folderLocation); os.IsNotExist(err) {
			os.MkdirAll(folderLocation, 0755)
		}

		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == 200 {
			fmt.Println("(+) FETCHING", url, folderLocation)
			body, _ := io.ReadAll(resp.Body)
			filePath := filepath.Join(folderLocation, fmt.Sprintf("%s_%s_%s_%s_%s", day, month, year, radar, timeComb))
			os.WriteFile(filePath, body, 0644)
		}

		// ! Uncomment for Debugging file download
		/* if err == nil && resp.StatusCode != 200 {
			fmt.Println("(-) CANNOT FETCH", url, resp.StatusCode)
		} */

		if x == test4 {
			fmt.Println("Done...")
			return
		}
	}
}
