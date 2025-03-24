package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/joho/godotenv"
	"github.com/schollz/progressbar/v3"
)

type Response struct {
	Properties struct {
		RadarStation string `json:"radarStation"`
	} `json:"properties"`
}

func exeTime() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Execution time: %v\n", time.Since(start))
	}
}

// downloadFile downloads a file from url and saves it to the specified path
func downloadFile(url, filePath string, wg *sync.WaitGroup, bar *progressbar.ProgressBar) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err == nil && resp.StatusCode == 200 {
		fmt.Println(lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#67ff02")).
			Render("(+) FETCHING", url, filePath))

		body, _ := io.ReadAll(resp.Body)
		_ = os.WriteFile(filePath, body, 0644)
		err := resp.Body.Close()
		if err != nil {
			return
		}
	}

	err = bar.Add(1)
	if err != nil {
		return
	}
}

// isFolderEmpty checks if a folder is empty
func isFolderEmpty(folderPath string) (bool, error) {
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		return false, err
	}
	return len(entries) == 0, nil
}

// deleteEmptyFolder checks if a folder is empty and deletes it if it is
func deleteEmptyFolder(folderPath string) {
	isEmpty, err := isFolderEmpty(folderPath)
	if err != nil {
		fmt.Println(lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF0000")).
			Render("Error checking if folder is empty:", err.Error()))
		return
	}

	if isEmpty {
		err := os.Remove(folderPath)
		if err != nil {
			fmt.Println(lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#FF0000")).
				Render("Error deleting empty folder:", err.Error()))
			return
		}
		fmt.Println(lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF8C00")).
			Render("No data was downloaded. Empty folder deleted:", folderPath))
	} else {
		fmt.Println(lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#67ff02")).
			Render("Download complete. Files saved to:", folderPath))
	}
}

// downloadTimeRange downloads radar data for a specific time range
func downloadTimeRange(radar, year, month, day, filePathFolder string, timeStart, timeEnd int, bar *progressbar.ProgressBar) string {
	var wg sync.WaitGroup

	// Create base directory for files
	dirPath := filePathFolder
	if filePathFolder != "" {
		dirPath = fmt.Sprintf("%s\\%s_%s_%s_%s", filePathFolder, day, month, year, radar)
	} else {
		dirPath = fmt.Sprintf("%s_%s_%s_%s", day, month, year, radar)
	}

	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return ""
	}

	// Determine the file extension based on year
	var end string
	year1, _ := strconv.Atoi(year)
	if year1 >= 2009 && year1 <= 2012 {
		end = "V03"
	} else if year1 < 2009 {
		end = ""
	} else {
		end = "V06"
	}

	// Limit concurrent downloads to avoid overwhelming resources
	semaphore := make(chan struct{}, 10) // Limit to 10 concurrent downloads

	for x := timeStart; x <= timeEnd; x++ {
		timeComb := fmt.Sprintf("%06d", x)

		// Regular version
		url := fmt.Sprintf("https://noaa-nexrad-level2.s3.amazonaws.com/%s/%s/%s/%s/%s%s%s%s_%s_%s",
			year, month, day, radar, radar, year, month, day, timeComb, end)

		// GZ version
		url2 := fmt.Sprintf("https://noaa-nexrad-level2.s3.amazonaws.com/%s/%s/%s/%s/%s%s%s%s_%s_%s.gz",
			year, month, day, radar, radar, year, month, day, timeComb, end)

		filePath := filepath.Join(dirPath, fmt.Sprintf("%s_%s_%s_%s_%s", day, month, year, radar, timeComb))

		wg.Add(2) // One for regular URL, one for GZ URL

		// Use semaphore to limit concurrent downloads
		semaphore <- struct{}{} // Acquire
		go func(url, filePath string) {
			downloadFile(url, filePath, &wg, bar)
			<-semaphore // Release
		}(url, filePath)

		semaphore <- struct{}{} // Acquire
		go func(url, filePath string) {
			downloadFile(url, filePath+"_gz", &wg, bar)
			<-semaphore // Release
		}(url2, filePath)
	}

	wg.Wait() // Wait for all downloads to complete
	return dirPath
}

func main() {

	defer exeTime()()

	var month, day, year, radar, filePathFolder string
	var state string
	var choice int
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
		Render("Enter 1 for specific radar site or any key for city state entry: "))

	_, _ = fmt.Scanln(&choice)

	if choice == 1 {
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
		timeStartInt, _ := strconv.Atoi(timeStart)

		fmt.Println(lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#CC0000")).
			Render("Time End in Zulu (HHMMSS): "))
		_, _ = fmt.Scanln(&timeEnd)
		timeEndInt, _ := strconv.Atoi(timeEnd)

		fmt.Println(lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#67ff02")).
			Render("-----------------------------------------\n"))

		// We're making 2 requests per time value (regular and GZ)
		totalRequests := (timeEndInt - timeStartInt + 1) * 2
		bar := progressbar.Default(int64(totalRequests))

		// Call the function to download files with goroutines and get the download directory path
		downloadDir := downloadTimeRange(radar, year, month, day, filePathFolder, timeStartInt, timeEndInt, bar)

		// Check if the download directory is empty and delete if necessary
		if downloadDir != "" {
			deleteEmptyFolder(downloadDir)
		}

		fmt.Println(lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#00ff00")).
			Render("\nDone...\n"))
	}
	if choice != 1 {
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

		reader := bufio.NewReader(os.Stdin)
		fmt.Println(lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#ff00d0")).
			Render("Enter City (EX: Atlanta): "))
		city, _ := reader.ReadString('\n')
		city = strings.TrimSpace(city)

		fmt.Println(lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#ff00d0")).
			Render("Enter State (EX: GA): "))
		_, _ = fmt.Scanln(&state)

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
		timeStartInt, _ := strconv.Atoi(timeStart)

		fmt.Println(lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#CC0000")).
			Render("Time End in Zulu (HHMMSS): "))
		_, _ = fmt.Scanln(&timeEnd)
		timeEndInt, _ := strconv.Atoi(timeEnd)

		fmt.Println(lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#67ff02")).
			Render("-----------------------------------------\n"))

		// We're making 2 requests per time value (regular and GZ)
		totalRequests := (timeEndInt - timeStartInt + 1) * 2
		bar := progressbar.Default(int64(totalRequests))

		baseURL := "https://geocode.xyz"

		// Create URL parameters
		params := url.Values{}
		params.Add("locate", city+", "+state)
		params.Add("region", "US")
		params.Add("json", "1")

		// Build the request URL (note: Go's url.Values.Encode() is equivalent to Python's urlencode)

		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		// Get the API token
		apiToken := os.Getenv("API_TOKEN")
		if apiToken == "" {
			log.Fatal("API_TOKEN is not set")
		}

		reqURL := fmt.Sprintf("%s/?%s", baseURL, params.Encode()+"&auth="+apiToken)

		// Make the HTTP request
		resp, err := http.Get(reqURL)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Println("Error:", err)
			}
		}(resp.Body)

		// Check for HTTP errors
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("HTTP error: %s\n", resp.Status)
			os.Exit(1)
		}

		// Read and parse the JSON response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response:", err)
			os.Exit(1)
		}

		var geocodeData map[string]interface{}
		if err := json.Unmarshal(body, &geocodeData); err != nil {
			fmt.Println("Error parsing JSON:", err)
			os.Exit(1)
		}

		citystateURL := fmt.Sprintf("https://api.weather.gov/points/%v,%v", geocodeData["latt"], geocodeData["longt"])
		fmt.Println(citystateURL)

		resp, err = http.Get(citystateURL)
		if err != nil {
			log.Fatalf("Error fetching data: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Fatalf("API request failed with status: %d", resp.StatusCode)
		}

		var result Response
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			log.Fatalf("Error decoding JSON: %v", err)
		}

		fmt.Println("Radar Station:", result.Properties.RadarStation)

		radar := result.Properties.RadarStation

		downloadDir := downloadTimeRange(radar, year, month, day, filePathFolder, timeStartInt, timeEndInt, bar)

		// Check if the download directory is empty and delete if necessary
		if downloadDir != "" {
			deleteEmptyFolder(downloadDir)
		}

		fmt.Println(lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#00ff00")).
			Render("\nDone...\n"))
	}
}
