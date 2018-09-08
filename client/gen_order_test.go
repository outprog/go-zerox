package client

import (
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/outprog/gozx/models"
	"github.com/outprog/gozx/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenOrder(t *testing.T) {
	_, err := TestClient.GenOrder(models.KovanTokens["WETH"], models.KovanTokens["ZRX"], "1", "1")
	require.NoError(t, err)
}

func TestSignOrder(t *testing.T) {
	order, err := TestClient.GenOrder(models.KovanTokens["WETH"], models.KovanTokens["ZRX"], "1", "1")
	require.NoError(t, err)

	sign, err := TestClient.SignOrder(order, utils.SIGNTYPE_EthSign)
	require.NoError(t, err)

	assert.Equal(t, 66, len(common.FromHex(sign)))
	assert.Equal(t, common.Hex2Bytes("0"+strconv.Itoa(utils.SIGNTYPE_EthSign)), common.FromHex(sign)[65:])
}
