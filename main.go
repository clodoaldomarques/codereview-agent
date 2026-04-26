package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ollama/ollama/api"
)

// Estrutura para receber a requisição
type ReviewRequest struct {
	Code     string `json:"code"`
	Language string `json:"language"`
}

// Função para chamar o Ollama
func callOllama(prompt string) (string, error) {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	req := &api.GenerateRequest{
		Model:  "qwen2.5-coder:latest",
		Prompt: prompt,
	}

	var fullResponse string
	respFunc := func(resp api.GenerateResponse) error {
		fullResponse += resp.Response
		return nil
	}

	err = client.Generate(ctx, req, respFunc)
	if err != nil {
		return "", err
	}
	return fullResponse, nil
}

// Handler da nossa API
func reviewHandler(w http.ResponseWriter, r *http.Request) {
	var req ReviewRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Engenharia de Prompt: Dando "personalidade" e contexto ao modelo
	prompt := fmt.Sprintf(`Você é um especialista em revisão de código Go.
    Por favor, analise o código Go abaixo e sugira melhorias ou gere um teste unitário simples para ele.
    Código:
    ` + "```go\n" + req.Code + "\n```")

	review, err := callOllama(prompt)
	if err != nil {
		http.Error(w, "AI Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"review": review})
}

func main() {
	http.HandleFunc("/review", reviewHandler)
	log.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
