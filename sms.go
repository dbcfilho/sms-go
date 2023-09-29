package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
)

func main() {
    // Obter os parâmetros da linha de comando
    args := os.Args[1:]
    if len(args) != 2 {
        fmt.Println("Uso: script <número de telefone> <mensagem>")
        return
    }

    // Obter o número de telefone e a mensagem
    number := args[0]
    message := args[1]

    // Criar uma solicitação HTTP
    url := "https://api.smsapi.com.br/v1/send"
    payload := strings.NewReader(fmt.Sprintf("number=%s&message=%s", number, message))
    req, err := http.NewRequest("POST", url, payload)
    if err != nil {
        log.Fatal(err)
    }

    // Adicionar cabeçalhos de autenticação
    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("SMSAPI_TOKEN")))

    // Enviar a solicitação HTTP
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }

    // Verificar o status da resposta
    if resp.StatusCode != 200 {
        log.Fatal(resp.Status)
    }

    // Imprimir o status da resposta
    fmt.Println(resp.Status)
}