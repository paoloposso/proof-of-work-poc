package server

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/paoloposso/wisdom/src/challenge"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func getSentences() []string {
	return []string{
		"Don't fear to be wrong",
		"Money is useful",
		"Dig where you stand",
		"Go further than you may think you can",
		"Put down the extra, unnecessary baggage you are carrying",
		"Be good to yourself in simple ways",
		"Be open to the idea that you can always learn",
	}
}

func GetData(chal challenge.Challenge) (string, error) {
	if validateChallenge(chal) {
		rand.Seed(time.Now().UnixNano())
		randomIx := rand.Intn(6)
		return getSentences()[randomIx], nil
	}
	return "", errors.New("challenge failed")
}

func GetChallenge() (chal challenge.Challenge) {
	return generateChallenge(5)
}

func generateChallenge(difficulty int32) (chal challenge.Challenge) {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(1000)

	h := sha256.New()
	h.Write([]byte(fmt.Sprint(randomNum)))

	hashData := h.Sum(nil)

	fmt.Printf("%x\n", hashData)

	prefix := ""

	for i := int32(0); i < difficulty; i++ {
		prefix = prefix + "0"
	}

	chal.Nonce = 0
	chal.DataHash = fmt.Sprintf("%x\n", hashData)
	chal.Prefix = prefix

	return chal
}

func validateChallenge(ch challenge.Challenge) bool {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", ch)))
	hashData := fmt.Sprintf("%x\n", h.Sum(nil))

	fmt.Println(hashData)

	return strings.HasPrefix(hashData, ch.Prefix)
}
