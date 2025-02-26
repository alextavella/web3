package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/alextavella/web3/pkg/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/YOUR_INFURA_PROJECT_ID")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("YOUR_PRIVATE_KEY")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error getting public key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111)) // Sepolia Chain ID
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // Sem ETH enviado
	auth.GasLimit = uint64(3000000) // Definir um limite de gás adequado
	auth.GasPrice = gasPrice

	address, tx, _, err := contracts.DeployContracts(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Contrato implantado em:", address.Hex())
	fmt.Println("Transaction Hash:", tx.Hash().Hex())
}

func updateContract(client *ethclient.Client, privateKey *ecdsa.PrivateKey) {
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
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	tx, err := instance.Increment(auth)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transação enviada:", tx.Hash().Hex())
}
