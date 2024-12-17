package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// VultureBetContract representa a interface para interação com o contrato
type VultureBetContract struct {
	address  common.Address
	instance *bind.BoundContract
	client   *ethclient.Client
	auth     *bind.TransactOpts
}

// Nova função para configurar conexão com o contrato
func newVultureBetContract(
	rpcURL,
	contractAddressHex,
	privateKeyHex string,
) (*VultureBetContract, error) {
	// Conectar ao cliente Ethereum
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("falha ao conectar: %v", err)
	}

	// Converter chave privada
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("falha ao carregar chave privada: %v", err)
	}

	// Obter endereço público da chave
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("erro ao converter chave pública")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Preparar opções de transação
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("falha ao obter chain ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar transactor: %v", err)
	}
	auth.From = fromAddress

	// Converter endereço do contrato
	contractAddress := common.HexToAddress(contractAddressHex)

	// Parsear ABI do contrato
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return nil, fmt.Errorf("falha ao parsear ABI: %v", err)
	}

	// Criar instância do contrato
	instance := bind.NewBoundContract(contractAddress, parsedABI, client, client, client)

	return &VultureBetContract{
		address:  contractAddress,
		instance: instance,
		client:   client,
		auth:     auth,
	}, nil
}

// Métodos corrigidos
func (vc *VultureBetContract) criarAposta() (*types.Transaction, error) {
	tx, err := vc.instance.Transact(vc.auth, "criarAposta")
	resetAuthValue(vc.auth)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar aposta: %v", err)
	}
	return tx, nil
}

func (vc *VultureBetContract) registrarAposta(apostaID uint64, escolha string, valor float64) (*types.Transaction, error) {
	vc.auth.Value = toWei(valor)
	tx, err := vc.instance.Transact(vc.auth, "registrarAposta", big.NewInt(int64(apostaID)), escolha)
	resetAuthValue(vc.auth)
	if err != nil {
		return nil, fmt.Errorf("erro ao registrar aposta: %v", err)
	}
	return tx, nil
}

// Resetar auth.Value após uso
func resetAuthValue(auth *bind.TransactOpts) {
	auth.Value = big.NewInt(0)
}

// Funções de conversão
func toWei(eth float64) *big.Int {
	wei := new(big.Float).Mul(big.NewFloat(eth), big.NewFloat(1e18))
	weiInt, _ := wei.Int(nil)
	return weiInt
}

// Placeholder ABI
const contractABI = `[{
	"inputs": [],
	"name": "criarAposta",
	"outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}],
	"stateMutability": "nonpayable",
	"type": "function"
}]`

func main() {
	rpcURL := "http://localhost:8545"
	contractAddress := ""
	privateKey := ""

	contract, err := newVultureBetContract(rpcURL, contractAddress, privateKey)
	if err != nil {
		log.Fatalf("Erro ao configurar contrato: %v", err)
	}
	defer contract.client.Close()

	// Criar aposta como exemplo
	tx, err := contract.criarAposta()
	if err != nil {
		log.Fatalf("Erro ao criar aposta: %v", err)
	}
	fmt.Printf("Aposta criada. Transaction Hash: %s\n", tx.Hash().Hex())
}
