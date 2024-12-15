package storage

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"os"
	"bet/blockchain"
)

// Nome do arquivo onde o blockchain será salvo
const arquivoBlockchain = "blockchain.json"

// Salva o blockchain em um arquivo local
func SalvarBlockchain(cadeia []blockchain.Bloco) error {
	// Codifica o blockchain para JSON
	dados, err := json.MarshalIndent(cadeia, "", "  ")
	if err != nil {
		return fmt.Errorf("erro ao codificar blockchain: %v", err)
	}

	// Salva os dados no arquivo
	// err = ioutil.WriteFile(arquivoBlockchain, dados, 0644)
  err = os.WriteFile(arquivoBlockchain, dados, 0644)
	if err != nil {
		return fmt.Errorf("erro ao salvar blockchain: %v", err)
	}
	return nil
}

// Carrega o blockchain de um arquivo local
func CarregarBlockchain() ([]blockchain.Bloco, error) {
	// Verifica se o arquivo existe
	if _, err := os.Stat(arquivoBlockchain); os.IsNotExist(err) {
		// Retorna uma cadeia vazia se o arquivo não existir
		return []blockchain.Bloco{}, nil
	}

	// Lê os dados do arquivo
	// dados, err := ioutil.ReadFile(arquivoBlockchain)
	dados, err := os.ReadFile(arquivoBlockchain)

	if err != nil {
		return nil, fmt.Errorf("erro ao ler arquivo blockchain: %v", err)
	}

	// Decodifica os dados JSON para uma cadeia de blocos
	var cadeia []blockchain.Bloco
	err = json.Unmarshal(dados, &cadeia)
	if err != nil {
		return nil, fmt.Errorf("erro ao decodificar blockchain: %v", err)
	}
	return cadeia, nil
}

