package helpers

import (
	"testing"

	"github.com/nats-io/jwt"
	"github.com/nats-io/nkeys"
	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	operator, account, err := Create()
	assert.Nil(t, err)
	assert.NotNil(t, operator)
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

func Test_GetAccount(t *testing.T) {
	operator, account, err := Create()
	assert.Nil(t, err)
	assert.NotNil(t, operator)
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
	pubkSrc, err := opSrc.PublicKey()
	ocSrc := jwt.NewOperatorClaims(pubkSrc)
	jwtSrc, err := ocSrc.Encode(opSrc)

	opDest, err := nkeys.FromSeed(oSeed)
	assert.Nil(t, err)
	assert.NotNil(t, opDest)
	pkDest, err := opDest.PrivateKey()
	pubkDest, err := opDest.PublicKey()
	ocDest := jwt.NewOperatorClaims(pubkDest)
	jwtDest, err := ocDest.Encode(opDest)
	assert.Equal(t, pkSrc, pkDest)
	assert.Equal(t, pubkSrc, pubkDest)
	assert.Equal(t, jwtSrc, jwtDest)
}
