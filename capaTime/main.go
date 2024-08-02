package main

import (
	"fmt"
	"http"
	"time"
)

// Define a channel to carry quotes
var quotes chan string

func fetchQuote(quote string) {
	quotes <- quote
}

func printQuote() {
	quote := <-quotes
	fmt.Println(quote)
}

func main() {
	quotes = make(chan string)
	go fetchQuote("Una buena conexión es algo que tenemos en mente")
	printQuote()

	start := time.Now()

	apis := []string{
		"https://api.github.com/",
		"https://management.azure.com/",
		"https://api.github.com/authorizations",
		"https://api.github.com/search/code?q={query}{&page,per_page,sort,order}",
		"https://api.github.com/search/commits?q={query}{&page,per_page,sort,order}",
		"https://pokeapi.co/api/v2/pokemon/ditto",
	}

	for _, api := range apis {
		chekApi(api)
	}

	elapsed := time.Since(start)
	fmt.Println("!Listo¡, Tomó %v Segundos\n", elapsed.Seconds())
}

func chekApi(api string) {
	if _, err := http.Get(api); err != nil {
		fmt.Println("Error: !%s está caído! \n", api)
		return
	}
	fmt.Println("Success: !%s está funcionando! \n", api)
}
