package tests

import (
	"testing"
	"bet/blockchain"
)

// Testa a criação do bloco gênesis
func TestCriarBlocoGenesis(t *testing.T) {
  bloco := blockchain.CriarPrimeiroBloco()
	if bloco.Indice != 0 {
		t.Errorf("Índice esperado: 0, obtido: %d", bloco.Indice)
	}
	if bloco.Anterior != "0" {
		t.Errorf("HashAnterior esperado: '0', obtido: '%s'", bloco.Anterior)
	}
}

// Testa a criação de um novo bloco
// func TestCriarNovoBloco(t *testing.T) {
// 	genesis := blockchain.CriarPrimeiroBloco()
// 	apostas := []blockchain.Aposta{
// 		{UserID: "user1", Valor: 100, Resultado: "cara"},
// 	}
// 	novoBloco := blockchain.CriarNovoBloco(transacoes, genesis)
//
// 	if novoBloco.Indice != genesis.Indice+1 {
// 		t.Errorf("Índice esperado: %d, obtido: %d", genesis.Indice+1, novoBloco.Indice)
// 	}
// 	if novoBloco.HashAnterior != genesis.Hash {
// 		t.Errorf("HashAnterior esperado: '%s', obtido: '%s'", genesis.Hash, novoBloco.HashAnterior)
// 	}
// }
//
