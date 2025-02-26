package main

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/schollz/progressbar/v3"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// TODO: Refactor Var Names

func exeTime() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Execution time: %v\n", time.Since(start))
	}
}

func main() {

	defer exeTime()()

	var month, day, year, radar, filePathFolder string
	var timeStart, timeEnd string

	initialText := "Radar Database"

	// Create a Lipgloss style with a border
	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(0, 2).
		Margin(0).
		BorderForeground(lipgloss.Color("#FF69B4")) // Optional: Color the border

	// Apply the style to the text
	styledText := borderStyle.Render(initialText)

	// Print the styled text
	fmt.Println(styledText)

	fmt.Println(lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#ff00d0")).
		Render("Enter Month: "))

	_, _ = fmt.Scanln(&month)
	num1, _ := strconv.Atoi(month)
	if num1 < 10 {
		month = fmt.Sprintf("%02d", num1)
	}

	fmt.Println(lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#ff00d0")).
		Render("Enter Day: "))

	_, _ = fmt.Scanln(&day)
	num2, _ := strconv.Atoi(day)
	if num2 < 10 {
		day = fmt.Sprintf("%02d", num2)
	}

	fmt.Println(lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#ff00d0")).
		Render("Enter Year: "))
	_, _ = fmt.Scanln(&year)

	fmt.Println(lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#ff00d0")).
		Render("Enter Radar: "))
	_, _ = fmt.Scanln(&radar)

	fmt.Println(lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#ff00d0")).
		Render("Folder Location (Paste directory path without the ending \"\\\" (C:\\Test)): "))
	_, _ = fmt.Scanln(&filePathFolder)

	fmt.Println(lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#67ff02")).
		Render("Time Start in Zulu (HHMMSS): "))
	_, _ = fmt.Scanln(&timeStart)
	test := timeStart
	test1, _ := strconv.Atoi(test)

	fmt.Println(lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#CC0000")).
		Render("Time End in Zulu (HHMMSS): "))
	_, _ = fmt.Scanln(&timeEnd)
	test3 := timeEnd
	test4, _ := strconv.Atoi(test3)

	fmt.Println(lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#67ff02")).
		Render("-----------------------------------------\n"))

	bar := progressbar.Default(int64(test4 - test1))

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

		if filePathFolder != "" {
			filePathFolder := fmt.Sprintf("%s\\%s_%s_%s_%s", filePathFolder, day, month, year, radar)
			err := os.MkdirAll(filePathFolder, 0755)
			if err != nil {
				return
			}
		} else {
			filePathFolder := fmt.Sprintf("%s_%s_%s_%s", day, month, year, radar)
			err := os.MkdirAll(filePathFolder, 0755)
			if err != nil {
				return
			}
		}

		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == 200 {

			fmt.Println(lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#67ff02")).
				Render("(+) FETCHING", url, filePathFolder))

			body, _ := io.ReadAll(resp.Body)
			filePath := filepath.Join(filePathFolder, fmt.Sprintf("%s_%s_%s_%s_%s", day, month, year, radar, timeComb))
			_ = os.WriteFile(filePath, body, 0644)
		}

		resp2, err := http.Get(url2)
		if err == nil && resp2.StatusCode == 200 {
			fmt.Println("(+) FETCHING .GZ", url2, filePathFolder)
			body, _ := io.ReadAll(resp2.Body)
			filePath := filepath.Join(filePathFolder, fmt.Sprintf("%s_%s_%s_%s_%s", day, month, year, radar, timeComb))
			_ = os.WriteFile(filePath, body, 0644)
		}

		// Uncomment below for Debugging file download
		/*if err == nil && resp.StatusCode != 200 {
			fmt.Println(lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#FF0000")).
				Render("(-) CANNOT FETCH", url))
		}*/

		/*if err == nil && resp.StatusCode != 200 {
			fmt.Println(lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF0000")).
			Render("(-) CANNOT FETCH", url))
		}
		*/

		if x == test4 {
			fmt.Println(lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#00ff00")).
				Render("\nDone...\n"))
			return
		}
		_ = bar.Add(1)
	}
}
