package domain

import (
	"encoding/json"
	"math/big"
)

const (

	// AccountForExecutorTypeNone ...
	AccountForExecutorTypeNone = AccountForExecutorType("None")

	// AccountForExecutorTypeUninit ...
	AccountForExecutorTypeUninit = AccountForExecutorType("Uninit")

	// AccountForExecutorTypeAccount ...
	AccountForExecutorTypeAccount = AccountForExecutorType("Account")
)

type (
	// ExecutionOptions ...
	ExecutionOptions struct {
		BlockchainConfig string `json:"blockchain_config,omitempty"`
		BlockTime        int    `json:"block_time,omitempty"`
		BlockLt          string `json:"block_lt,omitempty"`
		TransactionLt    string `json:"transaction_lt,omitempty"`
	}

	// AccountForExecutorType ...
	AccountForExecutorType string

	// AccountForExecutorNone ...
	AccountForExecutorNone struct {
		Type AccountForExecutorType `json:"type"`
	}

	// AccountForExecutorUninit ...
	AccountForExecutorUninit struct {
		Type AccountForExecutorType `json:"type"`
	}

	// AccountForExecutorAccount ...
	AccountForExecutorAccount struct {
		Type             AccountForExecutorType `json:"type"`
		Boc              string                 `json:"boc"`
		UnlimitedBalance bool                   `json:"unlimited_balance,omitempty"`
	}

	// ParamsOfRunExecutor ...
	ParamsOfRunExecutor struct {
		Message              string            `json:"message"`
		Account              interface{}       `json:"account"` // AccountForExecutor
		ExecutionOptions     *ExecutionOptions `json:"execution_options,omitempty"`
		Abi                  interface{}       `json:"abi,omitempty"` //ABI??? AbiS and AbiH
		SkipTransactionCheck bool              `json:"skip_transaction_check,omitempty"`
	}

	// ResultOfRunExecuteMessage ...
	ResultOfRunExecuteMessage struct {
		Transaction json.RawMessage `json:"transaction,omitempty"`
		OutMessages []string        `json:"out_messages"`
		Decoded     DecodedOutput   `json:"decoded,omitempty"`
		Account     string          `json:"account"`
		Fees        TransactionFees `json:"fees"`
	}

	// ParamsOfRunTvm ...
	ParamsOfRunTvm struct {
		Message          string            `json:"message"`
		Account          string            `json:"account"`
		Abi              interface{}       `json:"abi,omitempty"` //ABI??? AbiS and AbiH
		ExecutionOptions *ExecutionOptions `json:"execution_options,omitempty"`
	}

	// ResultOfRunTvm ...
	ResultOfRunTvm struct {
		OutMessages []string       `json:"out_messages"`
		Decoded     *DecodedOutput `json:"decoded,omitempty"`
		Account     string         `json:"account"`
	}

	// ParamsOfRunGet ...
	ParamsOfRunGet struct {
		Account          string            `json:"account"`
		FunctionName     string            `json:"function_name,omitempty"`
		Input            interface{}       `json:"input,omitempty"`
		ExecutionOptions *ExecutionOptions `json:"execution_options,omitempty"`
	}

	// ResultOfRunGet ...
	ResultOfRunGet struct {
		Output json.RawMessage `json:"output"`
	}

	// TransactionFees ...
	TransactionFees struct {
		InMsgFwdFee      big.Int `json:"in_msg_fwd_fee"`
		StorageFee       big.Int `json:"storage_fee"`
		GasFee           big.Int `json:"gas_fee"`
		OutMsgsFwdFee    big.Int `json:"out_msgs_fwd_fee"`
		TotalAccountFees big.Int `json:"total_account_fees"`
		TotalOutput      big.Int `json:"total_output"`
	}

	// TvmUseCase ...
	TvmUseCase interface {
		RunExecutor(pORE ParamsOfRunExecutor) (*ResultOfRunExecuteMessage, error)
		RunTvm(pORT ParamsOfRunTvm) (*ResultOfRunTvm, error)
		RunGet(pORG ParamsOfRunGet) (*ResultOfRunGet, error)
	}
)
