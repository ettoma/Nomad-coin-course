package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/ettoretoma/Nomad-coin-course/utils"
)

func Start() {

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	utils.HandleErr(err)

	hashedMessage := utils.Hash("I love Jinseo")

	hashAsByte, err := hex.DecodeString(hashedMessage)

	utils.HandleErr(err)

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsByte)

	utils.HandleErr(err)

	ok := ecdsa.Verify(&privateKey.PublicKey, hashAsByte, r, s)

	fmt.Println(ok)
}
