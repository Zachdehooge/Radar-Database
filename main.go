package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func exeTime(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s execution time: %v\n", name, time.Since(start))
	}
}

type model struct {
	choices  []string         // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
}

func initialModel() model {
	return model{
		// Our to-do list is a grocery list
		choices: []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	s := "Please Select A Month\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func main(){
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	defer exeTime("main")()

	var month, day, year, radar, filePathFolder string
	var timeStart, timeEnd string

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

	fmt.Print("Folder Location (Paste directory path without the ending \"\\\" (C:\\Test)): ")
	fmt.Scanln(&filePathFolder)

	fmt.Print("Time Start in Zulu (HHMMSS): ")
	fmt.Scanln(&timeStart)
	test := timeStart
	test1, _ := strconv.Atoi(test)

	fmt.Print("Time End in Zulu (HHMMSS): ")
	fmt.Scanln(&timeEnd)
	test3 := timeEnd
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

		if filePathFolder != "" {
			filePathFolder := fmt.Sprintf("%s\\%s_%s_%s_%s", filePathFolder, day, month, year, radar)
			os.MkdirAll(filePathFolder, 0755)
		} else {
			filePathFolder := fmt.Sprintf("%s_%s_%s_%s", day, month, year, radar)
			os.MkdirAll(filePathFolder, 0755)
		}

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
		/* if err == nil && resp.StatusCode != 200 {
			fmt.Println("(-) CANNOT FETCH", url, resp.StatusCode)
		}

		if err == nil && resp.StatusCode != 200 {
			fmt.Println("(-) CANNOT FETCH .GZ", url2, resp2.StatusCode)
		}
		*/
		if x == test4 {
			fmt.Println("Done...")
			return
		}
	}
}
