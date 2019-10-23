package main

import (
	"github.com/tebeka/selenium"
	Chromedriver "github.com/tejfi/go-webApp-testing/chromedriver"
	"github.com/tejfi/go-webApp-testing/page/page"
)

var driver selenium.WebDriver
var page page.Page

func setUP() {

	var driver = Chromedriver.SeleniumDriver()
	page := page.Page{Driver: driver}
	driver.Get("https://www.adevinta.com/")

}

func tearDown() {
	driver.Quit()

}
