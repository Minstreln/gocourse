package intermediate

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {

	password := "password123"

	// hash := sha256.Sum256([]byte(password))
	// fmt.Println(password)
	// fmt.Println(hash)
	// fmt.Printf("sha-256 hex value: %x\n", hash)

	salt, err := generateSalt()
	if err != nil {
		fmt.Println("Error generating salt:", err)
		return
	}

	// hash the password with salt
	hash := hashPassword(password, salt)

	// store the salt and hash
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt (base64):", saltStr)
	fmt.Println("Hash (base64):", hash)

	// verify the password
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Error decoding salt:", err)
		return
	}

	loginHash := hashPassword(password, decodedSalt)

	// compare the hashes
	if hash == loginHash {
		fmt.Println("Password verified!")
	} else {
		fmt.Println("Invalid password.")
		return
	}

}

// generate a random salt
func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// function to hash password
func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}
