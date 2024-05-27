package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	dtos "github.com/chasinfo/multithreading-api/DTOs"
)

func getEnderecosViaCEP(cep string) (*dtos.ViaCEP, error) {
	resp, error := http.Get("https://viacep.com.br/ws/" + cep + "/json")
	if error != nil {
		return nil, error
	}

	defer resp.Body.Close()

	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}
	var c dtos.ViaCEP
	error = json.Unmarshal(body, &c)

	if error != nil {
		return nil, error
	}
	return &c, nil
}

func getEnderecoBrasilAPI(cep string) (*dtos.BrasilApiCep, error) {
	resp, error := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()

	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}
	var c dtos.BrasilApiCep
	error = json.Unmarshal(body, &c)

	if error != nil {
		return nil, error
	}
	return &c, nil
}

func main() {

	enderecoBrasilAPI := make(chan *dtos.BrasilApiCep)
	enderecoViaCEP := make(chan *dtos.ViaCEP)

	for _, cep := range os.Args[1:] {

		go func() {
			data, error := getEnderecoBrasilAPI(cep)

			if error != nil {
				fmt.Println(error)
			}

			enderecoBrasilAPI <- data
		}()

		go func() {
			data, error := getEnderecosViaCEP(cep)

			if error != nil {
				fmt.Println(error)
			}

			enderecoViaCEP <- data
		}()

		select {
		case data := <-enderecoBrasilAPI:
			fmt.Printf("CEP: %s, \nLogradouro: %s \nBairro: %s \nLocalidade: %s, \nUF: %s \napi: BrasilAPI\n", data.Cep, data.Logradouro, data.Bairro, data.Localidade, data.UF)
		case data := <-enderecoViaCEP:
			fmt.Printf("CEP: %s, \nLogradouro: %s \nBairro: %s \nLocalidade: %s, \nUF: %s \napi: ViaCEP\n", data.Cep, data.Logradouro, data.Bairro, data.Localidade, data.UF)
		case <-time.After(time.Second * 1):
			println("timeout")
		}
	}
}
