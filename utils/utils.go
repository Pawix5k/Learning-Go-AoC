package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type Config struct {
	Session string `json:"cookie"`
}

func loadConfigFile() Config {
	jsonFile, err := os.Open("config.json")
	defer jsonFile.Close()
	if err != nil {
		panic(err)
	}
	byteValue, _ := io.ReadAll(jsonFile)

	var config Config

	json.Unmarshal(byteValue, &config)

	return config
}

func getInput(session, day string) []byte {
	var client http.Client

	cookie := &http.Cookie{
		Name:  "session",
		Value: session,
	}

	url := fmt.Sprintf("https://adventofcode.com/2022/day/%s/input", day)

	req, _ := http.NewRequest("GET", url, nil)
	req.AddCookie(cookie)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	resp.Body.Close()

	return body
}

func GetFilePath(day int) string {
	return fmt.Sprintf("day%02s/input.txt", strconv.Itoa(day))
}

func DownloadInput(day int) {
	dayStr := strconv.Itoa(day)
	filePath := GetFilePath(day)

	if _, err := os.Stat(filePath); err == nil {
		return
	}

	config := loadConfigFile()

	content := getInput(config.Session, dayStr)

	outputFile, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	outputFile.Write(content)
}

func IsLetter(ch byte) bool {
	if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
		return true
	}
	return false
}
