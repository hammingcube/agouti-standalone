package main

import (
	"github.com/sclevine/agouti"
	"log"
)

func TestSomething(done chan bool) {
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
	done <- true

}

func main() {
	ch := make(chan bool)
	go TestSomething(ch)
	go TestSomething(ch)
	<-ch
	<-ch
}
