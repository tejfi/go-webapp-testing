package Chromedriver

import (
	"fmt"
	"github.com/tebeka/selenium"
	"net"
	"os"
	"os/exec"
)

var driver selenium.WebDriver
var err error

type env struct {
	SEHUB_USERNAME string
	SEHUB_PW       string
	SEHUB_URL      string
	baseURl        string
}

func setUrl() string {
	env1 := env{SEHUB_USERNAME: "selenium", SEHUB_PW: "CoolCanvas19.", SEHUB_URL: "seleniumhub.codecool.codecanvas.hu/wd/hub", baseURl: "http://localhost:8081/articles"}

	hubUrl := "https://" + env1.SEHUB_USERNAME + ":" + env1.SEHUB_PW + "@" + env1.SEHUB_URL
	return hubUrl
}

func SeleniumDriver() selenium.WebDriver {

	port, err := pickUnusedPort()

	var opts []selenium.ServiceOption
	service, err := selenium.NewChromeDriverService("chromedriver",
		port, opts...)

	if err != nil {
		fmt.Printf("Error starting the ChromeDriver server: %v", err)
	}

	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	wd, err := selenium.NewRemote(caps, setUrl())
	if err != nil {
		panic(err)
	}

	wd.Refresh()

	defer service.Stop()
	return wd
}

func pickUnusedPort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	port := l.Addr().(*net.TCPAddr).Port
	if err := l.Close(); err != nil {
		return 0, err
	}
	return port, nil
}

func GetBrowserPath(browser string) string {
	if _, err := os.Stat(browser); err != nil {
		path, err := exec.LookPath(browser)
		if err != nil {
			panic("Browser binary path not found")
		}
		return path
	}
	return browser
}
