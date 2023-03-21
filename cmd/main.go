package main

import "github.com/EdsonGustavoTofolo/desafio-multithreading-golang/internal/usecase"

func main() {
	usecase.NewGetCep("89803250").Execute()
}
