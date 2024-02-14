package main

//https://dzone.com/articles/building-a-concurrent-chat-app-with-go-and-websock Made by following this as shell

import ( // crypto import
	"crypto/aes"
	"encoding/hex"
	"os"

	"fmt" //general import

	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	//net import
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Print("Server starting at localhost:4444")
	err := godotenv.Load()
	if err != nil {
		  log.Fatal("Error loading .env file")}	
	port := os.Getenv("PORT")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
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
