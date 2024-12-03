package blockchain

import (
	"crypto/sha256"
	"fmt"
  "time"
  "math/rand"
)

// Função para calcular um hash de um novo bloco baseado em suas informações
func CalcularHash (bloco Bloco) string {
  registro := fmt.Sprintf("%d%s%v%s", bloco.Indice, bloco.HorarioDeCriacao, bloco.Apostas, bloco.Anterior)
  hash     := sha256.Sum256([]byte(registro)) // calcula o Hash para o bloco baseado nos dados guardados em registro
  return fmt.Sprintf("%x", hash)              // o hash é retornado em formato hexadecimal
}

func AdicionarBloco (bloco Bloco) {
  blockchainLocal = append(blockchainLocal, bloco)
}

func pegarUltimoBloco () Bloco {
  return blockchainLocal[len(blockchainLocal)-1]
}

// Função para finalizar a aposta e distribuir o premio
func FinalizarAposta (apostas []Aposta) string {
  rand.Seed(time.Now().UnixNano())

  // Inicia o resultado como cara, e se escolhendo aleatoriamente entre 1 e 2 o resultado for 1, então vira coroa
  resultado := "cara"
  if rand.Intn(2) == 1 {
    resultado = "coroa"
  }

  // Definindo a soma dos valores apostados em um resultado como "cara" e "coroa"
  totalCara, totalCoroa := 0.0, 0.0
  for _, t:= range apostas {
    if t.Resultado == "cara" {
      totalCara += t.Valor
    } else {
      totalCoroa += t.Valor
    }
  }

  // O pote representa o valor total a ser repartido pelos vencedores
  pote := totalCara + totalCoroa

  // O total dos vencedores é usado para determinar a proporção da divisao do pote
  totalDosVencedores := totalCara
  if resultado == "coroa" {
    totalDosVencedores = totalCoroa
  }

  for _, i:= range(apostas){
    if i.Resultado == resultado {
      montante := (i.Valor/totalDosVencedores)*pote
      fmt.Printf("Usuário %s ganhou %.2f\n", i.UserID, montante)
      // Colocar algum algoritmo para somar ao saldo do usuario "i"
    }
  }

  // Esse return está aqui só pra confirmar, precisa alterar depois
  return resultado
}

// Blockchain rodando localmente em memória
var blockchainLocal []Bloco
