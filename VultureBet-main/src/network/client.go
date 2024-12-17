package network

import (
  "encoding/json"
  "fmt"
  "net"
)
// Função para se conectar a outro nó
func Conectar (endereco string) {
  conexao, err:=net.Dial("tcp", endereco)
  if err != nil {
    fmt.Println("Erro ao se conectar ao nó: ", err)
    return
  }

  defer conexao.Close()
  
  var mensagem string
  json.NewDecoder(conexao).Decode(&mensagem)
  fmt.Println("Mensagem recebida: ", mensagem)
}
