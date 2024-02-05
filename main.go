package main

import ( // crypto import
	"crypto/aes"
	"encoding/hex"

	"fmt"//general import

    
	"log"
	"net/http"
	"github.com/gorilla/mux" 
	//net import
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})
	log.Print("Server starting at localhost:4444")
	http.ListenAndServe(":4444", r)
}

func cipher() {

	// cipher key
	key := "thisis32bitlongpassphraseimusing"

	// plaintext
	pt := "This is a secret"

	c := EncryptAES([]byte(key), pt)

	// plaintext
	fmt.Println(pt)

	// ciphertext
	fmt.Println(c)

	// decrypt
	DecryptAES([]byte(key), c)
}

func EncryptAES(key []byte, plaintext string) string {

	c, err := aes.NewCipher(key)
	CheckError(err)

	out := make([]byte, len(plaintext))

	c.Encrypt(out, []byte(plaintext))

	return hex.EncodeToString(out)
}

func DecryptAES(key []byte, ct string) {
	ciphertext, _ := hex.DecodeString(ct)

	c, err := aes.NewCipher(key)
	CheckError(err)

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)

	s := string(pt[:])
	fmt.Println("DECRYPTED:", s)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
