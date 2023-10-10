package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	if (len(os.Args) < 2) {
		ProcessText()
	}
}

func ProcessText() {
	caracteresToAdd := "**"
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	fmt.Println("\nTexto recebido:", text)

	texts := strings.Fields(text)

	var newTexts []string

	for _, text := range texts {
		middleIndex:= len(text)/2

		part1 := text[:middleIndex]
		part2 := text[middleIndex:]
		newText := caracteresToAdd + part1 + caracteresToAdd + part2

		newTexts = append(newTexts, newText)
	}

	fmt.Println("Texto convertido para negrito:", strings.Join(newTexts, " "))
}