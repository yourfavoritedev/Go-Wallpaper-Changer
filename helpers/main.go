package helpers

import (
	"io/ioutil"
	"log"
	"strings"
)

func ReadFile(filename string) ([]byte, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
		return nil, err
	}
	return content, nil
}

func WriteFile(filename string, data []byte) error {
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
		return err
	}
	return nil
}

func ReplaceText(content, old, new string) string {
	return strings.Replace(content, old, new, 1)
}

func GetStringInBetween(str, left, right string) (result string) {
	s := strings.Index(str, left)
	if s == -1 {
		return
	}
	s += len(left)
	e := strings.Index(str[s:], right)
	if e == -1 {
		return
	}
	e += s
	return str[s:e]
}
