package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	configs "github.com/alextavella/web3/config"
	"github.com/alextavella/web3/pkg/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("erro ao carregar configuração: %w", err))
	}

	client, err := ethclient.Dial(fmt.Sprintf("%s/%s", config.INFURA.URL, config.INFURA.API_KEY))
	if err != nil {
		log.Fatal(fmt.Errorf("erro ao conectar ao cliente: %w", err))
	}

	privateKey, err := crypto.HexToECDSA(config.WALLET.PRIVATE_KEY[2:])
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error getting public key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(fmt.Errorf("erro ao obter nonce: %w", err))
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(fmt.Errorf("erro ao sugerir preço do gás: %w", err))
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(config.INFURA.NETWORKING)) // Sepolia Chain ID
	if err != nil {
		log.Fatal(fmt.Errorf("erro ao criar transactor: %w", err))
	}

	balance, err := client.BalanceAt(context.Background(), auth.From, nil)
	if err != nil {
		log.Fatal(fmt.Errorf("erro ao obter saldo da carteira: %w", err))
	}

	fmt.Println("💰 Saldo da carteira:", balance) // O saldo está em WEI

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // Sem ETH enviado
	auth.GasLimit = uint64(3000000) // Definir um limite de gás adequado
	auth.GasPrice = gasPrice

	fmt.Println("⛽️ Custo da transação: ", gasPrice)

	address, tx, _, err := contracts.DeployContracts(auth, client)
	if err != nil {
		log.Fatal(fmt.Errorf("erro ao implantar contrato: %w", err))
	}

	fmt.Println("Contrato implantado em:", address.Hex())
	fmt.Println("Transaction Hash:", tx.Hash().Hex())
}

func updateContract(config *configs.Conf, client *ethclient.Client, privateKey *ecdsa.PrivateKey) {
	contractAddress := common.HexToAddress("0xYourContractAddress")
	instance, err := contracts.NewContracts(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// Chamando a função `getCount` do contrato (view)
	count, err := instance.GetCount(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Valor atual de count:", count)

	// Interagindo com o contrato
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(config.INFURA.NETWORKING))
	tx, err := instance.Increment(auth)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transação enviada:", tx.Hash().Hex())
}
