package entity

import "github.com/Le-BlitzZz/onchain-auth-app/pkg/password"

type User struct {
	ID           int `gorm:"primary_key"`
	UserName     string
	PasswordHash string
}

func newUser(name, pwd string) *User {
	passwordHash, err := password.GeneratePasswordHash(pwd)

	if err != nil {
		log.Errorf("Error generating password hash: %v", err)
	}

	return &User{
		UserName:     name,
		PasswordHash: passwordHash,
	}
}

func CreateUser(name, pwd string) {
	// Check if a user already exists
	existingUser := FindUserByName(name)
	if existingUser != nil {
		// User already exists, don't create a duplicate
		log.Infof("User %s already exists, skipping creation", name)
		return
	}

	// User doesn't exist, create a new one
	Db().Create(newUser(name, pwd))
}

func FindUserByName(userName string) *User {
	user := &User{}
	if err := Db().Where("user_name = ?", userName).First(user).Error; err != nil {
		return nil
	}

	return user
}
