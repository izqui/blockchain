package core

import (
	"reflect"

	"github.com/izqui/helpers"
)

var (
	TRANSACTION_POW = helpers.ArrayOfBytes(TRANSACTION_POW_COMPLEXITY, POW_PREFIX)
	BLOCK_POW       = helpers.ArrayOfBytes(BLOCK_POW_COMPLEXITY, POW_PREFIX)

	TEST_TRANSACTION_POW = helpers.ArrayOfBytes(TEST_TRANSACTION_POW_COMPLEXITY, POW_PREFIX)
	TEST_BLOCK_POW       = helpers.ArrayOfBytes(TEST_BLOCK_POW_COMPLEXITY, POW_PREFIX)
)

func CheckProofOfWork(prefix []byte, hash []byte) bool {

	if len(prefix) > 0 {
		return reflect.DeepEqual(prefix, hash[:len(prefix)])
	}
	return true
}
