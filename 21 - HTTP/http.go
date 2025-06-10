package main

import (
	"log"
	"net/http"
)

func main() {
	// HTTP é um protocolo de comunicação - base da comunicação Web

	// Cliente (Faz requisição) - Servidor (Processa requisição e envia resposta)

	// Request - Response

	// Rotas
	// URI - Identificador do Recurso
	// Método - GET, POST, PUT, DELETE

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá Mundo"))
	})

	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bem vindo a página de usuários"))
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
