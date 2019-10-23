package main

import (
	"github.com/tebeka/selenium"
	"github.com/tejfi/restapi-testing/Chromedriver"
	"github.com/tejfi/restapi-testing/page"
)

var driver selenium.WebDriver

func Adevinta() {

	var driver = Chromedriver.SeleniumDriver()
	page := page.Page{Driver: driver}
	driver.Get("https://www.adevinta.com/")
	carr := page.FindElementByXpath("//ul[@id='menu-top-menu']/li[2]//a[contains(text(),'Careers')]")
	carr.Click()

}

func tearDown() {
	driver.Quit()

}
