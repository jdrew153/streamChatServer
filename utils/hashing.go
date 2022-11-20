package utils

import (
	"fmt"
	"github.com/matthewhartstonge/argon2"
)

func HashPass(password string) (string, error) {
	argon := argon2.DefaultConfig()

	encoded, err := argon.HashEncoded([]byte(password))

	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func VerifyPass(hashedPass string, plainPass string) (bool, error) {
	//fmt.Println("hashedPass -->", []byte(plainPass))
	//fmt.Println("plainPass -->", plainPass)

	successful, err := argon2.VerifyEncoded([]byte(hashedPass), []byte(plainPass))

	fmt.Println("verification results ---> ", successful)
	if err != nil {
		return false, err
	}

	if successful {
		return successful, nil
	}

	return false, nil

}
