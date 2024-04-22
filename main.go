package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	var month, day, year, radar string
	var timeStart, timeEnd int

	fmt.Print("Enter Month: ")
	fmt.Scanln(&month)
	// fmt.Println("d3 =", month)

	fmt.Print("Enter Day: ")
	fmt.Scanln(&day)
	// fmt.Println("d3 =", day)

	fmt.Print("Enter Year: ")
	fmt.Scanln(&year)
	// fmt.Println("d3 =", year)

	fmt.Print("Enter Radar: ")
	fmt.Scanln(&radar)
	// fmt.Println("d3 =", timeNow)

	fmt.Print("Time Start in Zulu (HHMMSS): ")
	fmt.Scanln(&timeStart)

	fmt.Print("Time End in Zulu (HHMMSS): ")
	fmt.Scanln(&timeEnd)

	for x := timeStart; x <= timeEnd; x++ {
		timeComb := fmt.Sprintf("%06d", x)

		url := fmt.Sprintf("https://noaa-nexrad-level2.s3.amazonaws.com/%s/%s/%s/%s/%s%s%s%s_%s_V06", year, month, day, radar, radar, year, month, day, timeComb)

		folderLocation := fmt.Sprintf("C:\\Coding Projects\\Go\\Radar_Database\\%s_%s_%s_%s", day, month, year, radar)
		if _, err := os.Stat(folderLocation); os.IsNotExist(err) {
			os.MkdirAll(folderLocation, 0755)
		}

		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == 200 {
			fmt.Println("Fetching", url)
			body, _ := io.ReadAll(resp.Body)
			filePath := filepath.Join(folderLocation, fmt.Sprintf("%s_%s_%s_%s_%s", day, month, year, radar, timeComb))
			os.WriteFile(filePath, body, 0644)
		}

		if x == timeEnd {
			fmt.Println("Done...")
			return
		}
	}
}
