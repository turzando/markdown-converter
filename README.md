# Markdown-converter
Aplicação para converter palavras em negrito no formato de markdown. Projetinho simples apenas para praticar e aprender a codar em golang. :D

#### Para rodar a aplicação é necessário ter Golang instalado.
# Como rodar
 
É possível passar o input do texto através do stdin ou de um arquivo. Para pegar do stdin basta utilizar o "echo", mas caso queira pegar de um arquivo é só passar como parâmetro o nome do arquivo que esteja no diretório, sem nenhum "comando como prefixo".

##### Obs: Caso passe stdin junto com um arquivo, o input será feito pelo arquivo. Exemplo logo abaixo.
~~~
$ go build markdownConverter.go

$ echo "maturidade maturidade maturidade" | ./markdownConverter

$ ./markdownConverter teste.txt

$ echo "maturidade maturidade maturidade" | ./markdownConverter teste.txt

~~~

Exemplo de todos parâmetros:
~~~
$ ./markdownConverter -f 70 -j 1 -o output.txt teste.txt
~~~

O resultado será um arquivo chamando output.txt contendo:
~~~
maturidade **maturid**ade maturidade
~~~

# Parâmetros

### -f [valor] (serve para especificar a porcentagem em que cada palavra ficará em negrito)

#### Exemplo
~~~
$ echo "maturidade maturidade" ./markdownConverter -f 70
~~~

#### Resultado
~~~
**maturid**ade **maturid**ade **maturid**ade
~~~

### -j [valor] (serve para espeficicar de quantas em quantas palavras ficará em negrito)

#### Exemplo
~~~
$ echo "maturidade maturidade" | ./markdownConverter -j 1
~~~

#### Resultado
~~~
maturidade **matur**idade maturidade
~~~

### -o [nome do arquivo.txt] (serve para informar o arquivo de saída que conterá o texto convertido)

#### Exemplo
~~~
$ echo "maturidade maturidade" | ./markdownConverter -o output.txt
~~~

#### Resultado

~~~
Será criado um arquivo, contendo o texto convertido, no diretório com o nome passado. (Como no exemplo em #Como rodar)
~~~
## Default

 * Cada palavra terá 50% convertida para negrito;
 * Não terá pulo de palavras para ter a conversão;
 * Não será criado um arquivo, mas sim um output no terminal mostrando o resultado da conversão.
