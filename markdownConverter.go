package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	percentage := 50
	jumps := 0
	shouldCreateOutputFile := false
	fullText := ""

	var outputFileName string

	systemCounter := 0
	for _, arg := range os.Args {

		index := systemCounter + 1

		if arg == "-f" {
			percentage, _ = strconv.Atoi(os.Args[index])
		} else if arg == "-j" {
			jumps, _ = strconv.Atoi(os.Args[index])
		} else if arg == "-o" {
			shouldCreateOutputFile = true
			outputFileName = os.Args[index]
		} else if strings.Contains(arg, ".txt") && os.Args[systemCounter-1] != "-o" {
			file, _ := os.Open(arg)
			defer file.Close()
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				fullText = fullText + scanner.Text()
			}
		}

		systemCounter++
	}

	if fullText == "" {
		reader := bufio.NewReader(os.Stdin)
		fullText, _ = reader.ReadString('\n')
	}

	convertedText := ConvertTextToBold(percentage, fullText, jumps)

	if shouldCreateOutputFile {
		file, _ := os.Create(outputFileName)
		file.WriteString(convertedText)
	} else {
		fmt.Println(convertedText)
	}
}

func ConvertTextToBold(percentage int, fullText string, jumps int) string {
	caracteresToAdd := "**"

	var newTexts []string

	texts := strings.Fields(fullText)

	quantity := 0

	for _, text := range texts {

		var newText string

		if quantity == jumps {
			//TODO reticencias
			indlexByPercentage := (percentage * len(text)) / 100

			part1 := text[:indlexByPercentage]
			part2 := text[indlexByPercentage:]
			newText = caracteresToAdd + part1 + caracteresToAdd + part2

			quantity = 0
		} else {
			newText = text
			quantity++
		}

		newTexts = append(newTexts, newText)
	}

	return strings.Join(newTexts, " ")
}
