package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryptID(t *testing.T) {
	userID := uint64(30010008)

	str := EncryptID(userID)
	t.Logf("EncryptID(%d) = %s", userID, str)
	decodeUserID, err := DecryptID(str)
	assert.Equal(t, userID, decodeUserID)
	t.Logf("%d", decodeUserID)
	assert.Nil(t, err)

}
