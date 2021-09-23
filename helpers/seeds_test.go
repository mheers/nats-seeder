package helpers

import (
	"testing"

	jwt "github.com/nats-io/jwt/v2"
	"github.com/nats-io/nkeys"
	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	operator, sysAccount, account, err := Create()
	assert.Nil(t, err)
	assert.NotNil(t, operator)
	assert.NotNil(t, sysAccount)
	assert.NotNil(t, account)
}

func Test_createOperatorSeed(t *testing.T) {
	oSeed, err := createOperatorSeed()
	assert.Nil(t, err)
	assert.NotNil(t, oSeed)
}

func Test_createAccountSeed(t *testing.T) {
	aSeed, err := createAccountSeed()
	assert.Nil(t, err)
	assert.NotNil(t, aSeed)
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

func Test_GetAccount(t *testing.T) {
	operator, sysAccount, account, err := Create()
	assert.Nil(t, err)
	assert.NotNil(t, operator)
	assert.NotNil(t, sysAccount)
	assert.NotNil(t, account)

	oSeed, err := operator.Seed()
	assert.NotNil(t, oSeed)
	assert.Nil(t, err)
	aSeed, err := account.Seed()
	assert.Nil(t, err)
	assert.NotNil(t, aSeed)

	accountGot, err := GetAccount(oSeed, aSeed)
	assert.Nil(t, err)
	assert.NotNil(t, aSeed)
	assert.Equal(t, account, accountGot)
}

// Used to prove that nkeys can be recreated by seeds
func TestNkeys(t *testing.T) {
	opSrc, err := nkeys.CreateOperator()
	assert.Nil(t, err)
	oSeed, err := opSrc.Seed()
	assert.Nil(t, err)
	pkSrc, err := opSrc.PrivateKey()
	assert.Nil(t, err)
	pubkSrc, err := opSrc.PublicKey()
	assert.Nil(t, err)
	ocSrc := jwt.NewOperatorClaims(pubkSrc)
	jwtSrc, err := ocSrc.Encode(opSrc)
	assert.Nil(t, err)

	opDest, err := nkeys.FromSeed(oSeed)
	assert.Nil(t, err)
	assert.NotNil(t, opDest)
	pkDest, err := opDest.PrivateKey()
	assert.Nil(t, err)
	pubkDest, err := opDest.PublicKey()
	assert.Nil(t, err)
	ocDest := jwt.NewOperatorClaims(pubkDest)
	jwtDest, err := ocDest.Encode(opDest)
	assert.Nil(t, err)
	assert.Equal(t, pkSrc, pkDest)
	assert.Equal(t, pubkSrc, pubkDest)
	assert.Equal(t, jwtSrc, jwtDest)
}

func TestCreateUser(t *testing.T) {
	operator, sysAccount, account, err := Create()
	assert.Nil(t, err)
	assert.NotNil(t, operator)
	assert.NotNil(t, sysAccount)
	assert.NotNil(t, account)

	oSeed, err := operator.Seed()
	assert.NotNil(t, oSeed)
	assert.Nil(t, err)
	aSeed, err := account.Seed()
	assert.Nil(t, err)
	assert.NotNil(t, aSeed)

	accountGot, err := GetAccount(oSeed, aSeed)
	assert.Nil(t, err)
	assert.NotNil(t, aSeed)
	assert.Equal(t, account, accountGot)

	jwt, uSeed, err := CreateUser(oSeed, aSeed, "demo", []string{}, []string{})
	assert.Nil(t, err)
	assert.NotNil(t, uSeed)
	assert.NotEmpty(t, jwt)
}
