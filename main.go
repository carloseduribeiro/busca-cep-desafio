package main

import (
	"context"
	"encoding/json"
	"github.com/carloseduribeiro/busca-cep-desafio/clients/viacep"
	"log"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancel()
	r, err := viacep.FindCEP(ctx, "89036-370")
	if err != nil {
		log.Fatal(err)
	}
	if err = json.NewEncoder(os.Stdout).Encode(r); err != nil {
		log.Fatal(err)
	}
}
