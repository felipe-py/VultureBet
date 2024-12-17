package blockchain

import (
  // "fmt"
)

//Função para validar um bloco
func ValidarBloco (novoBloco Bloco, blocoAnterior Bloco) bool {
  //Verificando se o indice está correto
  if novoBloco.Indice != blocoAnterior.Indice+1 {
    return false
  }

  // Verificando se o hash anterior esta correto
  if novoBloco.Anterior != blocoAnterior.Hash {
    return false
  }

  // Verificando se o hash calculado para o bloco era válido
  h := CalcularHash(novoBloco)
  if novoBloco.Hash != h {
    return false
  }

  return true
}

// Função para resolver conflitos
func ResolverConflitos (blockchainRecebida []Bloco) bool {
  // Quem tiver mais blocos deve estar mais atualizado
  if len(blockchainRecebida) > len(blockchainLocal) {
    blockchainLocal = blockchainRecebida
    return true
  }
  return false
}

// Função para validar os blocos da blockchain
func ValidarBlockchain (blockchainRecebida []Bloco) bool {
  for i:=1; i<len(blockchainRecebida); i++ {
    if !ValidarBloco(blockchainRecebida[i], blockchainRecebida[i-1]) {
      return false
    }
  }
  return true
}
