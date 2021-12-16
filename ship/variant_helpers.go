package ship

import (
	"github.com/dwjpeng/roxe_go"
)

// Request
var RequestVariant = roxe.NewVariantDefinition([]roxe.VariantType{
	{"get_status_request_v0", (*GetStatusRequestV0)(nil)},
	{"get_blocks_request_v0", (*GetBlocksRequestV0)(nil)},
	{"get_blocks_ack_request_v0", (*GetBlocksAckRequestV0)(nil)},
})

type Request struct {
	roxe.BaseVariant
}

func (r *Request) UnmarshalBinary(decoder *roxe.Decoder) error {
	return r.BaseVariant.UnmarshalBinaryVariant(decoder, RequestVariant)
}

// Result
var ResultVariant = roxe.NewVariantDefinition([]roxe.VariantType{
	{"get_status_result_v0", (*GetStatusResultV0)(nil)},
	{"get_blocks_result_v0", (*GetBlocksResultV0)(nil)},
})

type Result struct {
	roxe.BaseVariant
}

func (r *Result) UnmarshalBinary(decoder *roxe.Decoder) error {
	return r.BaseVariant.UnmarshalBinaryVariant(decoder, ResultVariant)
}

// TransactionTrace
var TransactionTraceVariant = roxe.NewVariantDefinition([]roxe.VariantType{
	{"transaction_trace_v0", (*TransactionTraceV0)(nil)},
})

type TransactionTrace struct {
	roxe.BaseVariant
}

type TransactionTraceArray struct {
	Elem []*TransactionTrace
}

func (t *TransactionTraceArray) AsTransactionTracesV0() (out []*TransactionTraceV0) {
	if t == nil || t.Elem == nil {
		return nil
	}
	for _, e := range t.Elem {
		switch v := e.Impl.(type) {
		case *TransactionTraceV0:
			out = append(out, v)

		default:
			panic("wrong type for conversion")
		}
	}
	return out
}

func (r *TransactionTraceArray) UnmarshalBinary(decoder *roxe.Decoder) error {
	data, err := decoder.ReadByteArray()
	if err != nil {
		return err
	}
	return roxe.UnmarshalBinary(data, &r.Elem)
}

func (r *TransactionTrace) UnmarshalBinary(decoder *roxe.Decoder) error {
	return r.BaseVariant.UnmarshalBinaryVariant(decoder, TransactionTraceVariant)
}

// ActionTrace
var ActionTraceVariant = roxe.NewVariantDefinition([]roxe.VariantType{
	{"action_trace_v0", (*ActionTraceV0)(nil)},
})

type ActionTrace struct {
	roxe.BaseVariant
}

func (r *ActionTrace) UnmarshalBinary(decoder *roxe.Decoder) error {
	return r.BaseVariant.UnmarshalBinaryVariant(decoder, ActionTraceVariant)
}

// PartialTransaction
var PartialTransactionVariant = roxe.NewVariantDefinition([]roxe.VariantType{
	{"partial_transaction_v0", (*PartialTransactionV0)(nil)},
})

type PartialTransaction struct {
	roxe.BaseVariant
}

func (r *PartialTransaction) UnmarshalBinary(decoder *roxe.Decoder) error {
	return r.BaseVariant.UnmarshalBinaryVariant(decoder, PartialTransactionVariant)
}

// TableDelta
var TableDeltaVariant = roxe.NewVariantDefinition([]roxe.VariantType{
	{"table_delta_v0", (*TableDeltaV0)(nil)},
})

type TableDelta struct {
	roxe.BaseVariant
}

func (d *TableDelta) UnmarshalBinary(decoder *roxe.Decoder) error {
	return d.BaseVariant.UnmarshalBinaryVariant(decoder, TableDeltaVariant)
}

type TableDeltaArray struct {
	Elem []*TableDelta
}

func (d *TableDeltaArray) UnmarshalBinary(decoder *roxe.Decoder) error {
	data, err := decoder.ReadByteArray()
	if err != nil {
		return err
	}
	return roxe.UnmarshalBinary(data, &d.Elem)
}

func (t *TableDeltaArray) AsTableDeltasV0() (out []*TableDeltaV0) {
	if t == nil || t.Elem == nil {
		return nil
	}
	for _, e := range t.Elem {
		switch v := e.Impl.(type) {
		case *TableDeltaV0:
			out = append(out, v)

		default:
			panic("wrong type for conversion")
		}
	}
	return out
}

// Transaction
var TransactionVariant = roxe.NewVariantDefinition([]roxe.VariantType{
	{"transaction_id", (*roxe.Checksum256)(nil)},
	{"packed_transaction", (*roxe.PackedTransaction)(nil)},
})

type Transaction struct {
	roxe.BaseVariant
}

func (d *Transaction) UnmarshalBinary(decoder *roxe.Decoder) error {
	return d.BaseVariant.UnmarshalBinaryVariant(decoder, TransactionVariant)
}

// ActionReceipt
var ActionReceiptVariant = roxe.NewVariantDefinition([]roxe.VariantType{
	{"action_receipt_v0", (*ActionReceiptV0)(nil)},
})

type ActionReceipt struct {
	roxe.BaseVariant
}

func (r *ActionReceipt) UnmarshalBinary(decoder *roxe.Decoder) error {
	return r.BaseVariant.UnmarshalBinaryVariant(decoder, ActionReceiptVariant)
}
