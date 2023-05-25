package cli

import "golang.org/x/crypto/bcrypt"

// Hashes a password using bcrypt
func hashPassword(password string) (string, error) {
	// Generate a hash of the password with a cost factor of 14
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// Verifies a password against a hashed password using bcrypt
func verifyPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
