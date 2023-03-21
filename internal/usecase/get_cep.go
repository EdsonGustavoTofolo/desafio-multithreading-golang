package usecase

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	cdnApiCepUrl = "https://cdn.apicep.com/file/apicep/%s.json"
	viaCepUrl    = "http://viacep.com.br/ws/%s/json/"
)

type getCep struct {
	cep string
}

func NewGetCep(cep string) *getCep {
	return &getCep{cep}
}

func (c getCep) Execute() {
	chanCdnApiCep := make(chan []byte)
	chanViaCep := make(chan []byte)

	go func(value string) {
		formattedCep := value[:5] + "-" + value[5:]

		body, err := requestGetCep(cdnApiCepUrl, formattedCep)

		if err != nil {
			fmt.Println(err.Error())
		}

		chanCdnApiCep <- body
	}(c.cep)

	go func(value string) {
		body, err := requestGetCep(viaCepUrl, value)

		if err != nil {
			fmt.Println(err.Error())
		}

		chanViaCep <- body
	}(c.cep)

	select {
	case resp := <-chanCdnApiCep:
		fmt.Printf("Cdn api cep response %v\n", string(resp))
	case resp := <-chanViaCep:
		fmt.Printf("Via cep response %v\n", string(resp))
	case <-time.After(1 * time.Second):
		fmt.Printf("Timeout\n")
	}
}

func requestGetCep(url, cep string) ([]byte, error) {
	url = fmt.Sprintf(url, cep)

	log.Printf("[GET] %s\n", url)

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
