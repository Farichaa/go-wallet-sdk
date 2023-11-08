package kaspa

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTransfer(t *testing.T) {
	var txInputs []*TxInput
	txInputs = append(txInputs, &TxInput{
		TxId:       "120c5410cc4512f29da50a8befc67c1cfbf7bb4f594ef91c14741150d8dadd24",
		Index:      0,
		Address:    "kaspa:qrcnkrtrjptghtrntvyqkqafj06f9tamn0pnqvelmt2vmz68yp4gqj5lnal2h",
		Amount:     "900000",
		PrivateKey: "b827bb46d921bde498a535999d7554071045f02e4fdfdebb10b08583f1c6afbe",
	})
	txData := &TxData{
		TxInputs:      txInputs,
		ToAddress:     "kaspa:qqvxjssnw024e93vykhzd8d7t6dua2sx8ak4mj7xm8s9370yevxcv0jgl2xfj", // 443642da97444e52af9eb35e3d32d6270f47d255854b63299b29f21c1ded4c7c
		Amount:        "100000",
		Fee:           "10000",
		ChangeAddress: "kaspa:qrcnkrtrjptghtrntvyqkqafj06f9tamn0pnqvelmt2vmz68yp4gqj5lnal2h",
		MinOutput:     "546",
	}

	signedTx, err := Transfer(txData)
	require.NoError(t, err)
	res := &struct {
		TxId string `json:"txId"`
	}{}
	err = json.Unmarshal([]byte(signedTx), res)
	require.NoError(t, err)
	expected := "0dcbc57ae8b4be9c0769bdfffd54db09f7b36f048ba19d43f09c14b323a2b0d8"
	require.Equal(t, expected, res.TxId)
}
