package blockchain

import (
  "time"
)

// Estrutura de um bloco da blockchain
type Bloco struct {
  Indice              int         // Posicao do bloco na sequencia
  HorarioDeCriacao    string      // Horario de criacao do bloco
  Apostas             []Aposta    // Lista das transacoes que estarao presentes nesse bloco
  Hash                string      // Hash do bloco
  Anterior            string      // Hash do bloco anterior
}

// Função para criar o primeiro bloco da blockchain, ou o 'gênesis'
func CriarPrimeiroBloco() Bloco {
  genesis := Bloco {
    Indice: 0,
    HorarioDeCriacao: time.Now().String(),
    Apostas: []Aposta{},
    Hash: "0",
    Anterior: "",
  }

  return genesis
}

// Função para criar um bloco par a blockchain
func criarNovoBloco (novasApostas []Aposta, blocoAnterior Bloco) Bloco {
  novoBloco := Bloco {
    Indice: blocoAnterior.Indice+1,
    HorarioDeCriacao: time.Now().String(),
    Apostas: novasApostas,
    Anterior: blocoAnterior.Hash,
  }

  novoBloco.Hash = CalcularHash(novoBloco)

  return novoBloco
}

