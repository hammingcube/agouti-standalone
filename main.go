package main

import (
	"fmt"
	"github.com/sclevine/agouti"
	"log"
	_ "time"
)

func TestSomething(driver *agouti.WebDriver, done chan bool) {
	var username, password string
	var d int

	page, err := agouti.NewPage("http://192.168.33.1:8075")
	if err != nil {
		log.Fatal("Failed to create new page:", err)
	}
	if err := page.Navigate("http://www.github.com"); err != nil {
		log.Fatal("Failed to navigate:", err)
	}
	page.FindByLink("Sign in").Click()
	fmt.Scanf("%s", &username)
	page.Find("#login_field").Fill(username)
	fmt.Scanf("%s", &password)
	page.Find("#password").Fill(password)
	fmt.Printf("Entered %s as username and %s as password", username, password)
	fmt.Scanf("%d", &d)
	page.Find("#login > form").Submit()
	done <- true
}

func NewChromeDriver() (*agouti.WebDriver, error) {
	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		log.Fatal("Failed to start chrome driver:", err)
		return nil, err
	}
	return driver, nil
}

func main() {
	ch := make(chan bool)
	//driver := agouti.NewWebDriver("http://127.0.0.1:8075", []string{"ls"}, agouti.Timeout(600))
	//log.Fatal(driver.Start())
	TestSomething(nil, ch)
	//driver.Stop()
	fmt.Println("Done")
}
