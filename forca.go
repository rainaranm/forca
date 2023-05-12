//TODO: fazer testes

package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// resultado irá formar um novo tracejado com todas as letras dadas pelo usuário
func resultado(palavraSecreta string, entradasValidas []string) (resposta string, finalizado bool) {
	//Variáveis
	tracejado := " "
	todasEntradas := ""

	// Convertendo toda a slice "entradasValidas" para uma única string "todasEntradas"
	for _, char := range entradasValidas {
		todasEntradas += char
	}

	//Desta forma, sempre terá ao menos um caractere em "todasEntradas"
	todasEntradas += "@"

	//Matriz para comparar todos os caracteres da palavra secreta com cada entrada do usuário
	for i := 0; i < len(palavraSecreta); i++ {
		for j := 0; j < len(entradasValidas); {

			/* Se forem diferentes, então irá comparar o próximo carac. do usuário com carac. da palavra atual:
			- Se forem iguais, então irá seguir para o próximo carac. do usuário;
			- Se não, um traçado "_" será impresso e retornará à matriz para comparar o próximo carac. do usuário
			*/
			if entradasValidas[j] != string(palavraSecreta[i]) {
				sum := j + 1
				if strings.Contains(palavraSecreta, string(todasEntradas[sum])) {
					j++

				} else {
					tracejado += "_"
					break
				}

				//Se forem iguais, o caractere será impressa, então retonará à matriz para comparar o próximo carac. da palavra
			} else {
				tracejado += string(palavraSecreta[i])
				break
			}
		}
	}

	//Retornará verdadeiro se a forca estiver completa
	if strings.Contains(tracejado, "_") {
		return tracejado, false
	} else {
		return tracejado, true
	}

}

// removerDuplicacao devolve a slice entradasValidas sem caracteres duplicados
func removerDuplicacao(palavraSecreta string, entradas []string) (entradasValidas []string) {

	keys := make(map[string]bool)

	for _, entrada := range entradas {
		if _, value := keys[entrada]; !value {
			keys[entrada] = true
			if strings.Contains(palavraSecreta, entrada) {
				entradasValidas = append(entradasValidas, entrada)
			}
		}
	}

	return
}

// entrada irá pedir ao usuário que digite uma letra para a forca ou "1" para sair do jogo
func entrada(tracejado, palavraSecreta string, entradas []string, erros, nTentativas int) {

	//FIXME: fazer variáveis para cada frase

	//CaractereAtual será a resposta do usuário para cada vez que a função "entrada" é chamada
	caractereAtual := ""

	fmt.Println("-------------------------------------------------")
	fmt.Print("Digite uma letra para a forca (ou 1 para sair): ")
	fmt.Scan(&caractereAtual)
	fmt.Println("-------------------------------------------------")

	caractereAtual = strings.ToTitle(caractereAtual)

	//entradas obterá todas as tentativas do usuário
	entradas = append(entradas, caractereAtual)

	//O caractere "1" irá finalizar o jogo. FIXME: alterar para tecla "esc"
	if caractereAtual == "1" {
		fmt.Print("saiu")

		//Se o usuário digitar algo além de "1"
	} else {

		//Se a palavra secreta tiver o caractere atual
		if strings.Contains(palavraSecreta, caractereAtual) {

			fmt.Println("\n| Tentativas:", entradas)
			forca(palavraSecreta, erros)

			//Nova slice sem caracteres duplicados
			entradasValidas := removerDuplicacao(palavraSecreta, entradas)

			//Função resultado devolve o tracejado atualizado e booleano para responder se o jogo está finalizado
			tracejado, finalizado := (resultado(palavraSecreta, entradasValidas))
			fmt.Println(tracejado + "\n")

			if finalizado {
				fmt.Print(`------------------------------------------------

                                      Obrigado!
                                    O/ /
| Parabéns, vc completou a forca!! /|
                                   / \
								   
`)

			} else {
				entrada(tracejado, palavraSecreta, entradas, erros, nTentativas)
			}

			//Se o usuário escrever mais de um caractere
		} else if len(caractereAtual) > 1 {

			if nTentativas == 1 {
				fmt.Print("\nPor favor, digite apenas uma letra. Essa é sua última chance!\n")
			} else {
				fmt.Printf("Por favor, digite apenas uma letra. Você tem mais %d tentativas\n", nTentativas)
			}

			fmt.Println("\n| Tentativas: ", entradas)
			forca(palavraSecreta, erros)
			fmt.Println(tracejado)

			entrada(tracejado, palavraSecreta, entradas, erros, nTentativas)

			//Se a palavra secreta não tiver o caractere atual
		} else {
			if erros < 5 {
				erros++
				nTentativas--

				if nTentativas == 1 {
					fmt.Printf("\nOps, não tem %s na palavra secreta :( Essa é sua última chance!\n", caractereAtual)
				} else {
					fmt.Printf("\nOps, não tem %s na palavra secreta :( Você tem mais %d tentativas\n", caractereAtual, nTentativas)
				}

				fmt.Println("\n| Tentativas: ", entradas)
				forca(palavraSecreta, erros)
				fmt.Println(tracejado)

				entrada(tracejado, palavraSecreta, entradas, erros, nTentativas)

				//Acabou todas a tentativas
			} else {
				erros++

				fmt.Printf("\n Não tem %s na palavra secreta e suas chances acabaram );\n", caractereAtual)

				forca(palavraSecreta, erros)
				fmt.Println(tracejado)
				fmt.Println("")
			}

		}

	}

}

// forca irá apresentar quantidade de caracteres da palavra secreta e o desenho da forca correpondente a quantidade de erros
func forca(palavraSecreta string, erros int) {
	forca := ""

	if erros < 6 {
		fmt.Printf("\nA palavra secreta tem %d letras\n", len(palavraSecreta))
	}

	//FIXME: Refatorar -> deixar apenas um desenho da forca padrão e ir acrecentando o bonequinho

	switch erros {
	case 0:
		forca = ` ________
|        |
|
|
|
|
|
|`
	case 1:
		forca = ` ________
|        |
|        O
|
|
|
|
|`

	case 2:
		forca = ` ________
|        |
|        O
|        |
|
|
|
|`

	case 3:
		forca = ` ________
|        |
|        O
|       /|
|
|
|
|`

	case 4:
		forca = ` ________
|        |
|        O
|       /|\
|       
|
|
|`

	case 5:
		forca = ` ________
|        | 
|        O 
|       /|\
|       /
|
|
|`

	case 6:
		forca = ` ________
|        |  aaaaaaa
|       \O/ /
|        |
|       / \ 
|
|
|`

	}

	fmt.Println(forca)

}

func main() {

	//Slices vazias
	entradas := make([]string, 0)
	palavrasSecretas := make([]string, 0)

	//TODO: fazer listas de palavras de acordo com temas (animais, objetos...) para o usuário escolher

	//Implementando as palavras secretas
	palavrasSecretas = append(palavrasSecretas, "casamento", "salada", "cinema",
		"cadeira", "coelho", "janela", "pescador", "biblioteca", "semente", "olho", "cachorro", "esqueleto", "bode", "chuveiro")

	//Variáveis
	erros, nTentativas := 0, 6
	index := len(palavrasSecretas) - 1
	palavraSecreta := strings.ToTitle(palavrasSecretas[rand.Intn(index)]) //strings.ToTitle converterá todas as letras para maiúsculas 
	tracejado := " "

	//Interface inicial
	fmt.Println(`
 ---------------
* JOGO DA FORCA *
 ---------------`)
	forca(palavraSecreta, erros)

	// Tracejado inicial
	for i := 0; i < len(palavraSecreta); i++ {
		tracejado += "_"
	}

	fmt.Print(tracejado + "\n\n")

	//Chamada da função
	entrada(tracejado, palavraSecreta, entradas, erros, nTentativas)

}