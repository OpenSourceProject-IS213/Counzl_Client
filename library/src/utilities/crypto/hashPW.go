package crypto

// Kode som jeg har tatt utgangspunkt i er hentet fra:
// https://golang.org/pkg/crypto/cipher/#example_NewCFBDecrypter
import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
