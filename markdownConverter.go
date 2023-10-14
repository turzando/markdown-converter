package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	// "io/ioutil"
)

func main() {

	// Ter um parâmetro para especificar um arquivo de saída
	// Exemplo: app -o novo_arquivo.txt meu_texto.txt

	// Ter um parâmetro para especificar a fixação (% mínimo de negrito por palavra)
	// Exemplo: app -f 70 meu_texto.txt

	// Ter um parâmetro para especificar os pulos (a cada quantas palavras tem um negrito)
	// Exemplo: app -j 2 meu_texto.txt

	// file prioridade

	// flags nativa pesquisar
	
	// se colocar entre aspas entende como apenas um

	// text in stdin or file | percentage, jump, fileOutput

	var percentage int
	var jumps int
	var fullText string

	reader := bufio.NewReader(os.Stdin)
	fullText, _ = reader.ReadString('\n')

	systemCounter := 0
	for _, arg := range os.Args {

		if (arg == "-f") {
			index := systemCounter+1
			percentage, _ = strconv.Atoi(os.Args[index])
		} else if (arg == "-j") {
			index := systemCounter+1
			jumps, _ = strconv.Atoi(os.Args[index])
		} else if (strings.Contains(arg, ".txt")) {
			file, _ := os.Open(arg)
			defer file.Close()
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				fullText = scanner.Text()
			}
		}

		systemCounter++
	}

	ConvertTextToBold(percentage, fullText, jumps)
}

func ConvertTextToBold(percentage int, fullText string, jumps int) {
	caracteresToAdd := "**"

	var newTexts []string

	texts := strings.Fields(fullText)

	quantity := 0

	for _, text := range texts {

		var newText string

		if (quantity == jumps) {
			indlexByPercentage := (percentage * len(text))/100
	
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

	fmt.Println(strings.Join(newTexts, " "))
}