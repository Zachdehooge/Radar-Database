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

// TODO: Handle user file path input and have radar data populate that folder, rather than the folder outside of it

func work(month, day, year, radar, filePathFolder string, xStart, xEnd int) {
	for x := xStart; x <= xEnd; x++ {

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

		//TODO: Handle blank input to evaluate in current working directory (IF STATEMENT)
		filePathFolder := fmt.Sprintf("./%s_%s_%s_%s", day, month, year, radar)
		os.MkdirAll(filePathFolder, 0755)

		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == 200 {
			fmt.Println("(+) FETCHING", url, filePathFolder)
			body, _ := io.ReadAll(resp.Body)
			filePath := filepath.Join(filePathFolder, fmt.Sprintf("%s_%s_%s_%s_%s", day, month, year, radar, timeComb))
			os.WriteFile(filePath, body, 0644)
		}

		resp2, err := http.Get(url2)
		if err == nil && resp2.StatusCode == 200 {
			fmt.Println("(+) FETCHING .GZ", url2, filePathFolder)
			body, _ := io.ReadAll(resp2.Body)
			filePath := filepath.Join(filePathFolder, fmt.Sprintf("%s_%s_%s_%s_%s", day, month, year, radar, timeComb))
			os.WriteFile(filePath, body, 0644)
		}

		// ! Uncomment for Debugging file download
		if err == nil && resp.StatusCode != 200 {
			fmt.Println("(-) CANNOT FETCH", url, resp.StatusCode)
		}

		/* 		if err == nil && resp.StatusCode != 200 {
			fmt.Println("(-) CANNOT FETCH .GZ", url2, resp2.StatusCode)
		} */

		if x == xEnd {
			fmt.Println("Done...")
			return
		}
	}
}

func main() {

	defer exeTime("main")()

	var month, day, year, radar, filePathFolder string
	var timeStart, timeEnd string
	var xStart, xEnd int

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

	/* 	fmt.Print("Folder Location: ")
	   	fmt.Scanln(&filePathFolder) */

	fmt.Print("Time Start in Zulu (HHMMSS)(Push Enter to Default to 000000): ")
	fmt.Scanln(&timeStart)
	if timeStart != "" {
		test := timeStart
		xStart, _ = strconv.Atoi(test)
	} else {
		test := "000000"
		xStart, _ = strconv.Atoi(test)
	}

	fmt.Print("Time End in Zulu (HHMMSS)(Push Enter to Default to 235959): ")
	fmt.Scanln(&timeEnd)
	if timeEnd != "" {
		test3 := timeEnd
		xEnd, _ = strconv.Atoi(test3)
	} else {
		test3 := "235959"
		xEnd, _ = strconv.Atoi(test3)
	}

	work(month, day, year, radar, filePathFolder, xStart, xEnd)
}
