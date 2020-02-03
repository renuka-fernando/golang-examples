package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	username, password, err := Credentials()
	fmt.Println("GOT", username, password, err)
}

func getInputString(printText string, defaultVal string, validRegex string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(printText)
	value, err := reader.ReadString('\n')
	value = strings.TrimSpace(value)

	if err != nil {
		return "", err
	}
	if value == "" {
		return defaultVal, nil
	}

	var reg = regexp.MustCompile(validRegex)
	if !reg.MatchString(value) {
		return "", errors.New("input validation failed")
	}

	return value, nil
}

func getInputPassword(printText string) (string, error) {
	if printText == "" {
		printText = "Enter Password: "
	}
	fmt.Print(printText)

	password, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	return string(password), nil
}

// credentials reads username and password from terminal and return them
func Credentials() (string, string, error) {
	username, err := getInputString("Enter Username: ", "", `^[\w\d\-]+$`)
	if err != nil {
		return "", "", err
	}

	password, err := getInputPassword("")
	if err != nil {
		return "", "", err
	}
	return strings.TrimSpace(username), strings.TrimSpace(password), nil
}
