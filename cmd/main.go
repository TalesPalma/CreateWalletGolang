package main

import (
	"fmt"
	"log"
	"os"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

func main() {

	infosTestNet := criarCarteiraTestNet3()
	saveTestNetInfosFileTxt(infosTestNet)
	fmt.Println("Arquivo com info da nova carteira salvo com sucesso")
}

func saveTestNetInfosFileTxt(infosTestNet string) {
	file, err := os.Create("testeNetInfos.txt")
	if err != nil {
		log.Fatal("Erro with create arquivo", err)
	}

	// Escrevendo no arquivo
	file.WriteString(infosTestNet)
	file.Close()
}

func criarCarteiraTestNet3() string {

	// Definindo parâmetros da rede de teste
	netParams := &chaincfg.TestNet3Params

	// Gerando uma nova chave privada
	privateKey, err := btcec.NewPrivateKey()
	if err != nil {
		log.Fatal(err)
	}

	// Criando um novo WIF (Wallet Import Format) a partir da chave privada
	wif, err := btcutil.NewWIF(privateKey, netParams, true)
	if err != nil {
		log.Fatal(err)
	}

	// Obtendo a chave pública a partir da chave privada
	pubKey := privateKey.PubKey()

	// Calculando o hash da chave pública
	witnessProg := btcutil.Hash160(pubKey.SerializeCompressed())

	// Criando um novo endereço a partir do hash da chave pública
	addre, err := btcutil.NewAddressWitnessPubKeyHash(witnessProg, netParams)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Carteira Gerada com sucesso!")

	return fmt.Sprintf("PrivateKey: %s \n Address : %s\n", wif.String(), addre.EncodeAddress())

}
