package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	fileName    string
	fullURLFile string

	month     string
	day       string
	year      string
	radar     string
	timeStart int
	timeEnd   int
)

func main() {

	fmt.Println("Enter Month: ")
	reader := bufio.NewReader(os.Stdin)
	month, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read line: %s-\n", month)

	fmt.Println("Enter Day: ")
	reader1 := bufio.NewReader(os.Stdin)
	day, err := reader1.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read line: %s-\n", day)

	fmt.Println("Enter Year: ")
	reader2 := bufio.NewReader(os.Stdin)
	year, err := reader2.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read line: %s-\n", year)

	fmt.Println("Enter Radar: ")
	// TODO: Enter Scanner for user input
	fmt.Println("Time Start in Zulu: ")
	// TODO: Enter Scanner for user input
	fmt.Println("Time End in Zulu: ")
	// TODO: Enter Scanner for user input
	// TODO: Calculate time Combined from start and end to make For loop to exit on finish

	fullURLFile = "https://noaa-nexrad-level2.s3.amazonaws.com/2024/04/16/KHTX/KHTX20240416_000302_V06"

	// Build fileName from fullPath
	fileURL, err := url.Parse(fullURLFile)
	if err != nil {
		log.Fatal(err)
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName = segments[len(segments)-1]

	// Create blank file
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(fullURLFile)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	fmt.Printf("Downloaded a file %s with size %d", fileName, size)

}
