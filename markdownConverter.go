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
			fileContent, _ := os.ReadFile(arg)

			fullText = string(fileContent)
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

	var convertedWords []string

	quantity := 0

	treatedWords := TreatLineBreak(fullText)

	for _, word := range treatedWords {
		var caractersIndex int

		for i := len(word) - 1; i >= 0; i-- {
			if !ItsLyrics(word[i]) {
				caractersIndex = i
			} else if ItsLyrics(word[i]) && word[i-1] == 92 {
				caractersIndex = i
			} else {
				break
			}
		}

		var caractersFromWord string

		if caractersIndex != 0 {
			caractersFromWord = word[caractersIndex:]
			word = word[:caractersIndex]
		}

		var convertedWord string

		if quantity == jumps {
			indlexByPercentage := (percentage * len(word)) / 100

			part1 := word[:indlexByPercentage]
			part2 := word[indlexByPercentage:]
			convertedWord = caracteresToAdd + part1 + caracteresToAdd + part2 + caractersFromWord
			quantity = 0
		} else {
			convertedWord = word
			quantity++
		}

		convertedWords = append(convertedWords, convertedWord)
	}

	return strings.Join(convertedWords, " ")
}

func ItsLyrics(byteCaracter byte) bool {
	return byteCaracter >= 65 && byteCaracter <= 90 || byteCaracter >= 97 && byteCaracter <= 122
}

func TreatLineBreak(fullText string) []string {

	words := strings.Split(fullText, " ")
	var treatedWords string

	for _, word := range words {
		for i := len(word) - 1; i >= 0; i-- {
			if word[i] == 10 {
				parte1 := word[:i]
				parte2 := word[i+1:]
				word = parte1 + "\\n " + parte2
			} else {
				continue
			}
		}
		word = word + " "

		treatedWords = treatedWords + word
	}

	return strings.Fields(treatedWords)
}
