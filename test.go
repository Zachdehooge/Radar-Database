package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func getMonth() int {
	var month string

	fmt.Print("Enter Month: ")
	fmt.Scanln(&month)
	num1, _ := strconv.Atoi(month)
	if num1 < 10 {
		month = fmt.Sprintf("%02d", num1)
	}

	return num1
}

func getDay() int {
	var day string

	fmt.Print("Enter Day: ")
	fmt.Scanln(&day)
	num2, _ := strconv.Atoi(day)
	if num2 < 10 {
		day = fmt.Sprintf("%02d", num2)
	}

	return num2
}

func getYear() string {
	var year string

	fmt.Print("Enter Year: ")
	fmt.Scanln(&year)

	return year
}

func getRadar() string {
	var radar string

	fmt.Print("Enter Radar: ")
	fmt.Scanln(&radar)

	return radar
}

func getTimeStart() string {
	var timeStart string

	fmt.Print("Time Start in Zulu (HHMMSS): ")
	fmt.Scanln(&timeStart)

	test := timeStart

	return test
}

func getTimeEnd() string {
	var timeEnd string

	fmt.Print("Time End in Zulu (HHMMSS): ")
	fmt.Scanln(&timeEnd)
	test3 := timeEnd

	return test3
}

func testing(test, test3, month, day string, year, radar string) {

	test1, _ := strconv.Atoi(test)
	test4, _ := strconv.Atoi(test3)

	for x := test1; x <= test4; x++ {

		var end string
		year2 := year
		year1, _ := strconv.Atoi(year2)

		if year1 >= 2009 && year1 <= 2012 {
			end = "V03"
		} else if year1 < 2009 {
			end = ""
		} else {
			end = "V06"
		}

		timeComb := fmt.Sprintf("%06d", x)

		url := fmt.Sprintf("https://noaa-nexrad-level2.s3.amazonaws.com/%s/%s/%s/%s/%s%s%s%s_%s_%s", year, month, day, radar, radar, year, month, day, timeComb, end)

		url2 := fmt.Sprintf("https://noaa-nexrad-level2.s3.amazonaws.com/%s/%s/%s/%s/%s%s%s%s_%s_%s.gz", year, month, day, radar, radar, year, month, day, timeComb, end)

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

		resp2, err := http.Get(url2)
		if err == nil && resp2.StatusCode == 200 {
			fmt.Println("(+) FETCHING .GZ", url2, folderLocation)
			body, _ := io.ReadAll(resp2.Body)
			filePath := filepath.Join(folderLocation, fmt.Sprintf("%s_%s_%s_%s_%s", day, month, year, radar, timeComb))
			os.WriteFile(filePath, body, 0644)
		}

		// ! Uncomment for Debugging file download
		if err == nil && resp.StatusCode != 200 {
			fmt.Println("(-) CANNOT FETCH", url, resp.StatusCode)
		}

		/* if err == nil && resp.StatusCode != 200 {
			fmt.Println("(-) CANNOT FETCH .GZ", url2, resp2.StatusCode)
		}*/

		if x == test4 {
			fmt.Println("Done...")
			return
		}
	}
}

func main() {

	/* fmt.Printf("Month: %d \n", getMonth())
	fmt.Printf("Test1: %d\n", getTimeStart())
	fmt.Printf("Test4: %d\n", getTimeEnd())
	fmt.Printf("Day: %d\n", getDay())
	fmt.Printf("Year: %d\n", getYear())
	fmt.Printf("Radar: %d\n", getRadar()) */

	go testing("000000", "000100", "05", "20", "2024", "KHTX")
	// TODO: Test main testing func for why time start and time end are not being carried over as intended
}
