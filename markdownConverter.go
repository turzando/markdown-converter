package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
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

	if (len(os.Args) < 2) {
		reader := bufio.NewReader(os.Stdin)
		fullText, _ := reader.ReadString('\n')
		ConvertTextToBold("50", fullText)
	} else if (os.Args[1] == "-f") {
		ConvertTextToBold(os.Args[2], os.Args[3])
	}

}

func ConvertTextToBold(percentage string, fullText string) {
	caracteresToAdd := "**"
	intPercentage, _ := strconv.Atoi(percentage)

	var newTexts []string

	texts := strings.Fields(fullText)

	for _, text := range texts {
		indlexByPercentage := (intPercentage * len(text))/100

		part1 := text[:indlexByPercentage]
		part2 := text[indlexByPercentage:]
		newText := caracteresToAdd + part1 + caracteresToAdd + part2

		newTexts = append(newTexts, newText)
	}

	fmt.Println(strings.Join(newTexts, " "))
}