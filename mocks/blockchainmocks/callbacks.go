// Code generated by mockery v1.0.0. DO NOT EDIT.

package blockchainmocks

import (
	blockchain "github.com/hyperledger-labs/firefly/pkg/blockchain"
	fftypes "github.com/hyperledger-labs/firefly/pkg/fftypes"

	mock "github.com/stretchr/testify/mock"
)

// Callbacks is an autogenerated mock type for the Callbacks type
type Callbacks struct {
	mock.Mock
}

// BatchPinComplete provides a mock function with given fields: batch, signingIdentity, protocolTxID, additionalInfo
func (_m *Callbacks) BatchPinComplete(batch *blockchain.BatchPin, signingIdentity string, protocolTxID string, additionalInfo fftypes.JSONObject) error {
	ret := _m.Called(batch, signingIdentity, protocolTxID, additionalInfo)

	var r0 error
	if rf, ok := ret.Get(0).(func(*blockchain.BatchPin, string, string, fftypes.JSONObject) error); ok {
		r0 = rf(batch, signingIdentity, protocolTxID, additionalInfo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TxSubmissionUpdate provides a mock function with given fields: tx, txState, protocolTxID, errorMessage, additionalInfo
func (_m *Callbacks) TxSubmissionUpdate(tx string, txState fftypes.OpStatus, protocolTxID string, errorMessage string, additionalInfo fftypes.JSONObject) error {
	ret := _m.Called(tx, txState, protocolTxID, errorMessage, additionalInfo)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, fftypes.OpStatus, string, string, fftypes.JSONObject) error); ok {
		r0 = rf(tx, txState, protocolTxID, errorMessage, additionalInfo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
