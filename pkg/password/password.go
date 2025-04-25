package password

import "golang.org/x/crypto/bcrypt"

var defaultPasswordCost = bcrypt.DefaultCost

func GeneratePasswordHash(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), defaultPasswordCost)
	if err != nil {
		return "", err
	}

	return string(passwordHash), nil
}

func HashAndPasswordMatched(passwordHash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)) == nil
}
