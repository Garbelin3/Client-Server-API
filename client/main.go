package main

import (
	"context"
	"github.com/Garbelin3/client-server-api/client/exchangerate"
)

func main() {

	cli := exchangerate.NewClient()
	repo := exchangerate.NewRepository("cotacao.txt")

	resp, err := cli.Get(context.Background())
	if err != nil {
		panic(err)
	}

	if err = repo.Save(resp); err != nil {
		panic(err)
	}
}
