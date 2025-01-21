package Infrastructure

// Functions for hashing and comparing passwords to ensure secure storage of user credentials.
import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password using bcrypt
func (*Infrastructure) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// ComparePasswords compares a hashed password with a plaintext password
func (*Infrastructure) ComparePasswords(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
