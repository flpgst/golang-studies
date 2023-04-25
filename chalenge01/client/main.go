package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	var usdPrice string
	err = json.Unmarshal(body, &usdPrice)
	if err != nil {
		panic(err)
	}
	writeUSDPrice(usdPrice)

}

func writeUSDPrice(usdPrice string) {
	file, err := os.Create("cotacao.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %x\n", err)
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("Dólar:{%s}", usdPrice))

}
