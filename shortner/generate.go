package shortner

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/itchyny/base58-go"
)

// It implements hash function to map the url into bytes
func hash256(input string) []byte {
	algo := sha256.New()
	algo.Write([]byte(input))
	return algo.Sum(nil)
}

// It encodes the generated bytes from hash function to text
func base58encoder(bytes []byte) string {
	encode := base58.BitcoinEncoding

	encoded, err := encode.Encode(bytes)
	if err != nil {
		fmt.Println("error in encoding the url", err)
	}

	return string(encoded)
}

// It generates the short url from original url
func GenerateshorlUrl(originalurl string) string {
	byt := hash256(originalurl)

	// Here we have to convert into a number as genrated bytes are in array format , so unable to encode it
	generatenumber := new(big.Int).SetBytes(byt).Uint64()
	shorturl := base58encoder([]byte(fmt.Sprintf("%d", generatenumber)))

	fmt.Println("Generated short url is ", shorturl)
	return shorturl

}
