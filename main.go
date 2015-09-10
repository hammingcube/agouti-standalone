package main

import (
	"github.com/sclevine/agouti"
	"log"
)

func TestSomething() {
	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		log.Fatal("Failed to start chrome driver:", err)
	}
	page, err := driver.NewPage()
	if err != nil {
		log.Fatal("Failed to create new page:", err)
	}

	if err := page.Navigate("http://www.google.com"); err != nil {
		log.Fatal("Failed to navigate:", err)
	}

}

func main() {
	TestSomething()
}
