package main

import (
	"bytes"
	"context"
	"github.com/carloseduribeiro/busca-cep-desafio/clients/apicep"
	"github.com/carloseduribeiro/busca-cep-desafio/clients/viacep"
	"io"
	"log"
	"os"
	"time"
)

const cep string = "89036-370"

func main() {
	ctx := context.TODO()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	apicepResults := make(chan []byte)
	viacepResults := make(chan []byte)

	go func(ctx context.Context, cep string, results chan<- []byte) {
		result, err := apicep.FindCEP(ctx, cep)
		if err != nil {
			log.Println("error fetching apicep", err)
			return
		}
		results <- result
	}(ctx, cep, apicepResults)

	go func(ctx context.Context, cep string, results chan<- []byte) {
		result, err := viacep.FindCEP(ctx, cep)
		if err != nil {
			log.Println("error fetching apvicap", err)
			return
		}
		results <- result
	}(ctx, cep, viacepResults)

	select {
	case r := <-apicepResults:
		log.Print("apicep results:")
		if _, err := io.Copy(os.Stdout, bytes.NewReader(r)); err != nil {
			log.Println(err)
		}
	case r := <-viacepResults:
		log.Println("viacep results:")
		if _, err := io.Copy(os.Stdout, bytes.NewReader(r)); err != nil {
			log.Println(err)
		}
	case <-ctx.Done():
		log.Fatalf("timeout")
	}
}
