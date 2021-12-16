package system

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	roxe "github.com/dwjpeng/roxe_go"
)

func NewSetContract(account roxe.AccountName, wasmPath, abiPath string) (out []*roxe.Action, err error) {
	codeAction, err := NewSetCode(account, wasmPath)
	if err != nil {
		return nil, err
	}

	abiAction, err := NewSetABI(account, abiPath)
	if err != nil {
		return nil, err
	}

	return []*roxe.Action{codeAction, abiAction}, nil
}

func NewSetContractContent(account roxe.AccountName, wasmContent, abiContent []byte) (out []*roxe.Action, err error) {
	codeAction := NewSetCodeContent(account, wasmContent)

	abiAction, err := NewSetAbiContent(account, abiContent)
	if err != nil {
		return nil, err
	}

	return []*roxe.Action{codeAction, abiAction}, nil
}

func NewSetCode(account roxe.AccountName, wasmPath string) (out *roxe.Action, err error) {
	codeContent, err := ioutil.ReadFile(wasmPath)
	if err != nil {
		return nil, err
	}
	return NewSetCodeContent(account, codeContent), nil
}

func NewSetCodeContent(account roxe.AccountName, codeContent []byte) *roxe.Action {
	return &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("setcode"),
		Authorization: []roxe.PermissionLevel{
			{
				Actor:      account,
				Permission: roxe.PermissionName("active"),
			},
		},
		ActionData: roxe.NewActionData(SetCode{
			Account:   account,
			VMType:    0,
			VMVersion: 0,
			Code:      roxe.HexBytes(codeContent),
		}),
	}
}

func NewSetABI(account roxe.AccountName, abiPath string) (out *roxe.Action, err error) {
	abiContent, err := ioutil.ReadFile(abiPath)
	if err != nil {
		return nil, err
	}

	return NewSetAbiContent(account, abiContent)
}

func NewSetAbiContent(account roxe.AccountName, abiContent []byte) (out *roxe.Action, err error) {
	var abiPacked []byte
	if len(abiContent) > 0 {
		var abiDef roxe.ABI
		if err := json.Unmarshal(abiContent, &abiDef); err != nil {
			return nil, fmt.Errorf("unmarshal ABI file: %w", err)
		}

		abiPacked, err = roxe.MarshalBinary(abiDef)
		if err != nil {
			return nil, fmt.Errorf("packing ABI: %w", err)
		}
	}

	return &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("setabi"),
		Authorization: []roxe.PermissionLevel{
			{
				Actor:      account,
				Permission: roxe.PermissionName("active"),
			},
		},
		ActionData: roxe.NewActionData(SetABI{
			Account: account,
			ABI:     roxe.HexBytes(abiPacked),
		}),
	}, nil
}

func NewSetAbiFromAbi(account roxe.AccountName, abi roxe.ABI) (out *roxe.Action, err error) {
	var abiPacked []byte
	abiPacked, err = roxe.MarshalBinary(abi)
	if err != nil {
		return nil, fmt.Errorf("packing ABI: %w", err)
	}

	return &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("setabi"),
		Authorization: []roxe.PermissionLevel{
			{
				Actor:      account,
				Permission: roxe.PermissionName("active"),
			},
		},
		ActionData: roxe.NewActionData(SetABI{
			Account: account,
			ABI:     roxe.HexBytes(abiPacked),
		}),
	}, nil
}

// NewSetCodeTx is _deprecated_. Use NewSetContract instead, and build
// your transaction yourself.
func NewSetCodeTx(account roxe.AccountName, wasmPath, abiPath string) (out *roxe.Transaction, err error) {
	actions, err := NewSetContract(account, wasmPath, abiPath)
	if err != nil {
		return nil, err
	}
	return &roxe.Transaction{Actions: actions}, nil
}

// SetCode represents the hard-coded `setcode` action.
type SetCode struct {
	Account   roxe.AccountName `json:"account"`
	VMType    byte             `json:"vmtype"`
	VMVersion byte             `json:"vmversion"`
	Code      roxe.HexBytes    `json:"code"`
}

// SetABI represents the hard-coded `setabi` action.
type SetABI struct {
	Account roxe.AccountName `json:"account"`
	ABI     roxe.HexBytes    `json:"abi"`
}
