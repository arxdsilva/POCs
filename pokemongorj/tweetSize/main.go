package main

import "fmt"

func main() {
	fmt.Println(len("👕 New avatar item alert! 👕 Add some spring to your step with these fresh spring-themed looks, now available in the Style Shop! 🍃 Don’t forget to show off your new look by sharing a screenshot of your freshly dressed avatar with #PokemonGO! https://t.co/GMcvv2IUdh"))
	f := "👕 Novo alerta de item de avatar! 👕 Adicione um pouco de primavera ao seu passo com estes novos looks com tema de primavera, agora disponíveis na Loja de Estilo! 🍃 Não se esqueça de mostrar sua nova aparência compartilhando uma captura de tela do seu avatar recém-vestido com #PokemonGO! https://t.co/GMcvv2IUdh"
	fmt.Println(len(f[:280]))
	fmt.Println(f[:277] + "...")
}
