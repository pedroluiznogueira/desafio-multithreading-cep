package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type apiResult struct {
	source string
	body   string
}

func main() {
	cep := "01153000"
	if len(os.Args) > 1 {
		cep = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	results := make(chan apiResult, 2)

	go fetchCEP(ctx, results, "BrasilAPI", fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep))
	go fetchCEP(ctx, results, "ViaCEP", fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep))

	select {
	case res := <-results:
		fmt.Printf("Resposta mais rápida: %s\n%s\n", res.source, res.body)
	case <-ctx.Done():
		fmt.Println("Erro: timeout - nenhuma API respondeu em 1 segundo")
	}
}

func fetchCEP(ctx context.Context, out chan<- apiResult, source, url string) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var raw json.RawMessage
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return
	}

	out <- apiResult{source: source, body: string(raw)}
}
