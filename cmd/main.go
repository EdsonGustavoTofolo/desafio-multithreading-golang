package main

import (
	"errors"
	"github.com/EdsonGustavoTofolo/desafio-multithreading-golang/internal/usecase"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		cep := os.Args[1]
		usecase.NewGetCep(cep).Execute()
	} else {
		panic(errors.New("missing cep"))
	}
}
