package blockchain

// Estrutura de uma aposta
type Aposta struct {
  UserID              string      // O usuario que fez a aposta
  Valor               float64     // Quanto dinheiro foi passado na transacao
  Resultado           string      // Se o valor apostado foi em cara ou em coroa
}

