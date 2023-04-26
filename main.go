package main

import (
	"log"
	"time"

	"github.com/fedesog/webdriver"
	"github.com/tebeka/selenium"
)

func InitDriver(driver_path string) (*webdriver.ChromeDriver, error) {
	chromeDriver := webdriver.NewChromeDriver(driver_path)
	err := chromeDriver.Start()

	if err != nil {
		return nil, err
	}

	return chromeDriver, nil
}

func InitPlatformOption(chromeDriver *webdriver.ChromeDriver, playform string) (*webdriver.Session, error) {
	desired := webdriver.Capabilities{"Platform": playform}
	required := webdriver.Capabilities{}
	session, err := chromeDriver.NewSession(desired, required)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func main() {
	chromeDriver, err := InitDriver("C:\\Users\\사용자\\chromeDri\\chromedriver")
	if err != nil {
		log.Panicln(err.Error())
		return
	}

	session, err := InitPlatformOption(chromeDriver, "Windows")
	if err != nil {
		log.Panicln(err.Error())
		return
	}

	err = session.Url("http://golang.org")
	if err != nil {
		log.Println(err)
	}

	// 셀레니움 관련 제어 부분
	btn, _ := session.FindElement(selenium.ByCSSSelector, "#page > div > div.HomeContainer > section.HomeSection.Playground > div.Playground-controls > div > button")
	btn.Click()

	time.Sleep(60 * time.Second)
	session.Delete()
	chromeDriver.Stop()
}
