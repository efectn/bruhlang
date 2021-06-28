package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

var Variables = make(map[string]int)

func checkCondition(condition string) bool {
	conditionParsed := strings.Split(condition, " ")

	if conditionParsed[1] == "momentum" && Variables[conditionParsed[0]] > Variables[conditionParsed[2]] {
		return true
	} else if conditionParsed[1] == "moment" && Variables[conditionParsed[0]] < Variables[conditionParsed[2]] {
		return true
	} else if Variables[conditionParsed[0]] == Variables[conditionParsed[1]] {
		return true
	}

	return false
}

func parse(content string) {
	contentParsed := strings.Split(content, " ")

	if contentParsed[0] == "bruh" {
		Variables[contentParsed[1]] = 0
	} else if contentParsed[1] == "momentum" {
		Variables[contentParsed[0]] += 1
	} else if contentParsed[1] == "moment" {
		Variables[contentParsed[0]] -= 1
	} else if contentParsed[0] == "moment" {
		fmt.Printf("%d", Variables[contentParsed[1]])
	} else if contentParsed[0] == "momentum" {
		fmt.Printf("%c", rune(Variables[contentParsed[1]]))
	} else if contentParsed[1] == "effect" {
		_, okFirst := Variables[contentParsed[0]]
		_, okLast := Variables[contentParsed[2]]

		if okFirst && okLast {
			Variables[contentParsed[0]] += Variables[contentParsed[2]]
		}
	} else if contentParsed[1] == "sound" {
		_, okFirst := Variables[contentParsed[0]]
		_, okLast := Variables[contentParsed[2]]

		if okFirst && okLast {
			Variables[contentParsed[0]] -= Variables[contentParsed[2]]
		}
	} else if contentParsed[0] == "sound" {
		expression := regexp.MustCompile(`sound\s-(?P<condition>.+)-\s>>\s(?P<statement>.+)\s+<<\ssound`).FindStringSubmatch(content)
		condition := checkCondition(expression[1])

		for condition {
			condition = checkCondition(expression[1])
			contents := strings.Split(expression[2], " |")

			for i := 0; i < len(contents); i++ {
				parse(strings.TrimSpace(contents[i]))
			}
		}
	} else {
		_, okFirst := Variables[contentParsed[0]]
		_, okLast := Variables[contentParsed[1]]

		if okFirst && okLast {
			Variables[contentParsed[0]] = Variables[contentParsed[1]]
		}
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("Exception: %v\n", err)
		}
	}()

	file, _ := ioutil.ReadFile(os.Args[1])
	contents := strings.Split(string(file), " ||")

	for i := 0; i < len(contents)-1; i++ {
		parse(strings.TrimSpace(contents[i]))
	}
}