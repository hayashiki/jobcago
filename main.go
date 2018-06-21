package main

import (
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/sclevine/agouti"
)

// Config is Jobcan Config
type Config struct {
	email    string
	company  string
	password string
}

func main() {
	log.Printf("start")
	config, _ := ConfigParse("./jobcan.cfg")
	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--headless",
		}),
		agouti.Debug,
	)
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}
	if err := page.Navigate("https://ssl.jobcan.jp/login/pc-employee"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
	// fill
	page.FindByID("client_id").Fill(config.company)
	page.FindByID("email").Fill(config.email)
	page.FindByID("password").Fill(config.password)

	// submit
	if err := page.Find("button.btn.btn-info.btn-block").Click(); err != nil {
		log.Fatalf("Failed to login:%v", err)
	}
	time.Sleep(1 * time.Second)
	page.FindByID("adit-button-push").Click()
	time.Sleep(1 * time.Second)
	log.Printf("end")
}

// ConfigParse for Parse JobCan Setting.
func ConfigParse(filename string) (*Config, error) {
	c := Config{}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fields := strings.Split(string(data), "\n")

	c.company = fields[0]
	c.email = fields[1]
	c.password = fields[2]
	return &c, nil
}
