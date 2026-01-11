package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
)

const (
	baseURL = "https://adventofcode.com"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run fetch.go <year> <day_number>")
	}

	currentYear := time.Now().Year()
	currentMonth := time.Now().Month()
	currentDay := time.Now().Day()
	// fmt.Println(currentYear, currentMonth, currentDay)

	// Year validation
	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Please provide a valid year number (2015-%i)", currentYear)
	}

	if year < 2015 || year > currentYear {
		log.Fatal("Please provide a valid year number (2015-%i)", currentYear)
	}
	if year == currentYear && time.Now().Month() != time.December {
		log.Fatal("This year's problems are not out yet, wait until December")
	}

	// Day validation
	// TODO: add special case that starting 2025, only 12 problems
	day, err := strconv.Atoi(os.Args[2])
	if err != nil || day < 1 || day > 25 {
		log.Fatal("Please select a valid day (1-25)")
	}
	if year == currentYear && currentMonth == time.December {
		if day > currentDay {
			log.Fatal("This day's problem is not out yet, come back later")
		}
	}
	// fmt.Println("Valid date selected, here is the link:")
	url := fmt.Sprintf("%s/%d/day/%d", baseURL, year, day)
	fmt.Println("Link:", url)

	// Create the directory structure (part1.go, part2.go, test.txt, input.txt)
	// First, check if the dir structure doesnt already exists,

	// Check / Create the year directory
	dirYear := strconv.Itoa(year)
	if _, err := os.Stat(dirYear); os.IsNotExist(err) {
		err = os.Mkdir(dirYear, 0o755)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("year Directory created/exists")

	// Check / Create the day directory
	dirDay := strconv.Itoa(year) + "/day" + strconv.Itoa(day)
	if _, err := os.Stat(dirDay); os.IsNotExist(err) {
		err = os.Mkdir(dirDay, 0o755)
		if err != nil {
			log.Fatal(err)
		}
	}
	// fmt.Println("day Directory created/exists")

	// fmt.Println("___________________________")
	// 1. Get the page
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// 2. Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 3. Find and iterate over h2 elements (the day's title)
	doc.Find("h2").Each(func(i int, s *goquery.Selection) {
		// Get the text inside the code element
		title := s.Text()
		fmt.Printf("Title: %s\n", title)
	})

	// 4. Example detection and saving to test.txt
	sel := doc.Find("pre code")
	if sel.Length() > 0 {
		example := sel.First().Text()
		// fmt.Println("Found 1 Example")
		// save the example to text
		filePath := filepath.Join(dirDay, "test.txt")
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		file.WriteString(example)
		fmt.Println("Created test.txt file")

	} else {
		fmt.Println("Did not found an example for this day")
	}

	// 5. personal input retrieval and saving to input.txt
	client := &http.Client{}
	req, err := http.NewRequest("GET", url+"/input", nil)
	if err != nil {
		log.Fatal(fmt.Printf("Error creating HTTP request: %v", err))
	}

	// session cookie
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file. Make sure it exists and contains AOC_SESSION_COOKIE")
	}

	sessionCookie := os.Getenv("AOC_SESSION_COOKIE")
	if sessionCookie == "" {
		log.Fatal("session cookie not found in .env file")
	}
	sessionToken := "session=" + sessionCookie
	req.Header.Add("Cookie", sessionToken)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(fmt.Printf("Error creating HTTP request: %v", err))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected HTTP status: %d\n", resp.StatusCode)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	} else {
		filePath := filepath.Join(dirDay, "input.txt")
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		file.WriteString(string(body))
		fmt.Println("Created input.txt file")
	}

	// 6. Create main.go from template (Go) TODO: add option to create python template instead
	templatePath := filepath.Join("templates", "go", "main.go")
	targetPath := filepath.Join(dirDay, "main.go")

	if _, err := os.Stat(targetPath); err == nil {
		fmt.Println("main.go already exists, skipping template creation")
	} else if os.IsNotExist(err) {
		err := copyFile(templatePath, targetPath)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Created main.go from Go template")
	} else {
		log.Fatal(err)
	}
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return out.Sync()
}
