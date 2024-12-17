package network

// OBS: cada nó vai atuar como servidor e como cliente, aqui estão as funções necessarias para executar a parte de "servidor" de cada nó

import (
	"bet/blockchain"
	"encoding/json"
	"fmt"
	"net"
)

// Função para iniciar o servidor do nó
func IniciarServ (porta string) {
  servidor, _:= net.Listen("tcp", porta)
  defer servidor.Close()
  fmt.Println("Nó atuando como servidor na porta: ", porta)

  for {
    conexao, _ := servidor.Accept()
    go processarConexao(conexao)
  }
}

// Função para processar uma conexão com outro nó
func processarConexao (conexao net.Conn) {
  defer conexao.Close()
  // Enviar e receber a blockchain que está ativa
  json.NewEncoder(conexao).Encode(blockchain.PegarBlockchain())

  var blockchainRecebida []blockchain.Bloco
  json.NewDecoder(conexao).Decode(&blockchainRecebida)
  blockchainAtualizada := blockchain.ResolverConflitos(blockchainRecebida)
  if blockchainAtualizada {
    fmt.Println("A blockchain foi atualizada com sucesso. ")
  }

}
