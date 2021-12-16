package ship

import (
	"github.com/dwjpeng/roxe_go"
	"github.com/dwjpeng/roxe_go/ecc"
)

// State History Plugin Requests

type GetStatusRequestV0 struct {
}

type GetBlocksAckRequestV0 struct {
	NumMessages uint32
}

type GetBlocksRequestV0 struct {
	StartBlockNum       uint32
	EndBlockNum         uint32
	MaxMessagesInFlight uint32
	HavePositions       []*BlockPosition
	IrreversibleOnly    bool
	FetchBlock          bool
	FetchTraces         bool
	FetchDeltas         bool
}

// State History Plugin Results
type GetStatusResultV0 struct {
	Head                 *BlockPosition
	LastIrreversible     *BlockPosition
	TraceBeginBlock      uint32
	TraceEndBlock        uint32
	ChainStateBeginBlock uint32
	ChainStateEndBlock   uint32
}

type GetBlocksResultV0 struct {
	Head             *BlockPosition
	LastIrreversible *BlockPosition
	ThisBlock        *BlockPosition         `roxe:"optional"`
	PrevBlock        *BlockPosition         `roxe:"optional"`
	Block            *SignedBlockBytes      `roxe:"optional"`
	Traces           *TransactionTraceArray `roxe:"optional"`
	Deltas           *TableDeltaArray       `roxe:"optional"`
}

// State History Plugin version of ROXE structs
type BlockPosition struct {
	BlockNum uint32
	BlockID  roxe.Checksum256
}

type Row struct {
	Present bool
	Data    []byte
}

type ActionTraceV0 struct {
	ActionOrdinal        roxe.Varuint32
	CreatorActionOrdinal roxe.Varuint32
	Receipt              *ActionReceipt `roxe:"optional"`
	Receiver             roxe.Name
	Act                  *Action
	ContextFree          bool
	Elapsed              int64
	Console              roxe.SafeString
	AccountRamDeltas     []*roxe.AccountRAMDelta
	Except               string `roxe:"optional"`
	ErrorCode            uint64 `roxe:"optional"`
}

type Action struct {
	Account       roxe.AccountName
	Name          roxe.ActionName
	Authorization []roxe.PermissionLevel
	Data          []byte
}

type ActionReceiptV0 struct {
	Receiver       roxe.Name
	ActDigest      roxe.Checksum256
	GlobalSequence uint64
	RecvSequence   uint64
	AuthSequence   []AccountAuthSequence
	CodeSequence   roxe.Varuint32
	ABISequence    roxe.Varuint32
}

type AccountAuthSequence struct {
	Account  roxe.Name
	Sequence uint64
}

type TableDeltaV0 struct {
	Name string
	Rows []Row
}

type PartialTransactionV0 struct {
	Expiration            uint32
	RefBlockNum           uint16
	RefBlockPrefix        uint32
	MaxNetUsageWords      roxe.Varuint32
	MaxCpuUsageMs         uint8
	DelaySec              roxe.Varuint32
	TransactionExtensions []*Extension
	Signatures            []ecc.Signature
	ContextFreeData       []byte
}

type TransactionTraceV0 struct {
	ID              roxe.Checksum256 `json:"id"`
	Status          roxe.TransactionStatus
	CPUUsageUS      uint32                `json:"cpu_usage_us"`
	NetUsageWords   roxe.Varuint32        `json:"net_usage_words"`
	Elapsed         roxe.Int64            `json:"elapsed"`
	NetUsage        uint64                `json:"net_usage"`
	Scheduled       bool                  `json:"scheduled"`
	ActionTraces    []*ActionTrace        `json:"action_traces"`
	AccountDelta    *roxe.AccountRAMDelta `json:"account_delta" roxe:"optional"`
	Except          string                `json:"except" roxe:"optional"`
	ErrorCode       uint64                `json:"error_code" roxe:"optional"`
	FailedDtrxTrace *TransactionTrace     `json:"failed_dtrx_trace" roxe:"optional"`
	Partial         *PartialTransaction   `json:"partial" roxe:"optional"`
}

type SignedBlockHeader struct {
	roxe.BlockHeader
	ProducerSignature ecc.Signature // no pointer!!
}

type TransactionReceipt struct {
	roxe.TransactionReceiptHeader
	Trx *Transaction
}

//type TransactionID roxe.Checksum256

type SignedBlock struct {
	SignedBlockHeader
	Transactions    []*TransactionReceipt
	BlockExtensions []*Extension
}

type SignedBlockBytes SignedBlock

func (s *SignedBlockBytes) AsSignedBlock() *SignedBlock {
	if s == nil {
		return nil
	}
	ss := SignedBlock(*s)
	return &ss
}

func (s *SignedBlockBytes) UnmarshalBinary(decoder *roxe.Decoder) error {
	data, err := decoder.ReadByteArray()
	if err != nil {
		return err
	}
	return roxe.UnmarshalBinary(data, s)
}

type Extension struct {
	Type uint16
	Data []byte
}
