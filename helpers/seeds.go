package helpers

import (
	"github.com/nats-io/jwt"
	"github.com/nats-io/nkeys"
	"github.com/sirupsen/logrus"
)

// Create creates an operator, an account and the key pairs for them
func Create() (nkeys.KeyPair, nkeys.KeyPair, error) {
	oSeed, err := createOperatorSeed()
	if err != nil {
		return nil, nil, err
	}
	operator, _, err := CreateOperator(oSeed)
	if err != nil {
		return nil, nil, err
	}

	aSeed, err := createAccountSeed()
	account, _, err := CreateAccount(aSeed, operator)
	if err != nil {
		return nil, nil, err
	}
	return operator, account, nil
}

// CreateOperator creates an operator based on oSeed
func CreateOperator(oSeed []byte) (nkeys.KeyPair, string, error) {

	okp, err := nkeys.FromSeed(oSeed)
	if err != nil {
		return nil, "", err
	}
	opub, err := okp.PublicKey()
	if err != nil {
		return nil, "", err
	}

	nac := jwt.NewOperatorClaims(opub)
	if err != nil {
		return nil, "", err
	}
	ojwt, err := nac.Encode(okp)
	if err != nil {
		return nil, "", err
	}
	return okp, ojwt, nil
}

// createOperatorSeed creates a seed for an operator
func createOperatorSeed() ([]byte, error) {
	// Create an operator
	// Needed to create a new seed -> run this once and set the output to OSeed to have the same seed every time
	okp, err := nkeys.CreateOperator()
	if err != nil {
		return nil, err
	}
	oseed, err := okp.Seed()
	if err != nil {
		return nil, err
	}
	return oseed, nil
}

// CreateAccount creates an account based on aSeed and the operator; returns the account, the jwt for the account and optional an error
func CreateAccount(aSeed []byte, okp nkeys.KeyPair) (nkeys.KeyPair, string, error) {
	akp, err := nkeys.FromSeed(aSeed)
	if err != nil {
		return nil, "", err
	}
	apub, err := akp.PublicKey()
	if err != nil {
		return nil, "", err
	}
	nac := jwt.NewAccountClaims(apub)
	if err != nil {
		return nil, "", err
	}
	ajwt, err := nac.Encode(okp)
	if err != nil {
		return nil, "", err
	}
	return akp, ajwt, nil
}

// createAccountSeed creates a seed for an account
func createAccountSeed() ([]byte, error) {
	// Create an account
	// Needed to create a new seed -> run this once and set the output to OSeed to have the same seed every time
	akp, err := nkeys.CreateAccount()
	if err != nil {
		return nil, err
	}
	aseed, err := akp.Seed()
	if err != nil {
		return nil, err
	}
	return aseed, nil
}

// createUserSeed creates a seed for a user
func createUserSeed() ([]byte, error) {
	// Create a user
	// Needed to create a new seed -> run this once and set the output to OSeed to have the same seed every time
	akp, err := nkeys.CreateUser()
	if err != nil {
		return nil, err
	}
	useed, err := akp.Seed()
	if err != nil {
		return nil, err
	}
	return useed, nil
}

// GetAccount reconstructs an account (KeyPair) from the operator and account seeds
func GetAccount(oSeed, aSeed []byte) (nkeys.KeyPair, error) {
	operator, _, err := CreateOperator(oSeed)
	if err != nil {
		return nil, err
	}
	account, _, err := CreateAccount(aSeed, operator)
	if err != nil {
		return nil, err
	}
	return account, nil
}

// CreateUser creates a new user with specific permissions
func CreateUser(oSeed, aSeed []byte, name string, allowPub, allowSub []string) (string, string, error) {
	operator, _, err := CreateOperator(oSeed)
	if err != nil {
		return "", "", err
	}
	account, _, err := CreateAccount(aSeed, operator)
	ukp, err := nkeys.CreateUser()
	if err != nil {
		return "", "", err
	}
	uSeed, err := ukp.Seed()
	if err != nil {
		return "", "", err
	}
	pub, err := ukp.PublicKey()
	if err != nil {
		return "", "", err
	}
	nuc := jwt.NewUserClaims(pub)
	nuc.Subject = pub
	nuc.BearerToken = true
	nuc.Name = name

	nuc.Pub.Allow = allowPub
	nuc.Sub.Allow = allowSub

	jwt, err := nuc.Encode(account)
	if err != nil {
		return "", "", err
	}
	logrus.Debugf("created a user jwt: %s\n", jwt)
	logrus.Debugf("created a user uSeed: %s\n", uSeed)

	return jwt, string(uSeed), err
}
