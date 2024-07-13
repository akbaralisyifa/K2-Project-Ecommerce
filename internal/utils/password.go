package utils

import "golang.org/x/crypto/bcrypt";


type HashingPwInterface interface{
	GeneretePassword(currentPw string) ([]byte, error)
	CheckPassword(inputPw []byte, currentPw []byte) error
}

type hashingPassword struct{};

func NewHashingPassword() HashingPwInterface {
	return &hashingPassword{}
}

func (pw *hashingPassword) GeneretePassword(currentPw string) ([]byte, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(currentPw), bcrypt.DefaultCost);

	if err != nil {
		return nil, err
	};

	return result, nil;
}

func (pw *hashingPassword) CheckPassword(inputPw []byte, currentPw []byte) error {
	return bcrypt.CompareHashAndPassword(currentPw, inputPw)
}