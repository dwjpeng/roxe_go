package snapshot

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/dwjpeng/roxe_go"
)

type TableIDObject struct {
	Code      string
	Scope     string
	TableName string
	Payer     string
	Count     uint32 // represents the number of rows & indices for a given table
}

type ContractRow struct {
	PrimKey string
	Payer   string
}

type KeyValueObject struct {
	ContractRow
	Value roxe.HexBytes
}

type Index64Object struct {
	ContractRow
	SecondaryKey roxe.Name
}

type Index128Object struct {
	ContractRow
	SecondaryKey roxe.Uint128
}

type Index256Object struct {
	ContractRow
	SecondaryKey roxe.Checksum256
}

type IndexDoubleObject struct {
	ContractRow
	SecondaryKey roxe.Float64
}

type IndexLongDoubleObject struct {
	ContractRow
	SecondaryKey roxe.Float128
}

func readContractTables(section *Section, f sectionCallbackFunc) error {
	fl := section.Buffer

	bufSize := section.BufferSize
	bytesBuf := make([]byte, bufSize)
	slurped, err := fl.Read(bytesBuf)
	if err != nil {
		return err
	}
	if slurped != int(bufSize) {
		slurped2, err := fl.Read(bytesBuf[slurped:])
		if err != nil {
			return err
		}
		if slurped+slurped2 != int(bufSize) {
			return fmt.Errorf("read less than section size: %d of %d", slurped+slurped2, section.BufferSize)
		}
	}
	buf := bytes.NewBuffer(bytesBuf)

	for {
		head := make([]byte, 8+8+8+8+4)
		readz, err := buf.Read(head)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("reading table id: %w", err)
		}
		if readz != 8+8+8+8+4 {
			return fmt.Errorf("incomplete read for table_id_object: %d out of %d", readz, 8+8+8+8+4)
		}

		t := &TableIDObject{
			Code:      roxe.NameToString(binary.LittleEndian.Uint64(head[0:8])),
			Scope:     roxe.NameToString(binary.LittleEndian.Uint64(head[8:16])),
			TableName: roxe.NameToString(binary.LittleEndian.Uint64(head[16:24])),
			Payer:     roxe.NameToString(binary.LittleEndian.Uint64(head[24:32])),
			Count:     binary.LittleEndian.Uint32(head[32:36]),
		}

		if err := f(t); err != nil {
			return err
		}

		for idxType := 0; idxType < 6; idxType++ {
			size, err := readUvarint(buf)
			if err != nil {
				return fmt.Errorf("reading index type size: %w", err)
			}

			for i := 0; i < int(size); i++ {
				head := make([]byte, 8+8)
				readz, err := buf.Read(head)
				if err != nil {
					return fmt.Errorf("reading key value head: %w", err)
				}
				if readz != 16 {
					return fmt.Errorf("incomplete read for row header: %d out of 16", readz)
				}

				contractRow := ContractRow{
					PrimKey: roxe.NameToString(binary.LittleEndian.Uint64(head[0:8])),
					Payer:   roxe.NameToString(binary.LittleEndian.Uint64(head[8:16])),
				}

				var row interface{}
				switch idxType {

				case 0: /* key_value_object */
					obj := &KeyValueObject{ContractRow: contractRow}

					valueSize, err := readUvarint(buf)
					if err != nil {
						return err
					}
					val := make([]byte, valueSize)
					readz, err = buf.Read(val)
					if err != nil {
						return err
					}
					if readz != int(valueSize) {
						return fmt.Errorf("incomplete read key_value_object: %d out of %d", readz, valueSize)
					}

					obj.Value = val
					row = obj

				case 1: /* index64_object */
					obj := &Index64Object{ContractRow: contractRow}
					val := make([]byte, 8)
					readz, err := buf.Read(val)
					if err != nil {
						return err
					}
					if readz != 8 {
						return fmt.Errorf("incomplete read index64_object: %d out of 8", readz)
					}
					if err := roxe.UnmarshalBinary(val, &obj.SecondaryKey); err != nil {
						return err
					}
					row = obj

				case 2: /* index128_object */
					obj := &Index128Object{ContractRow: contractRow}
					val := make([]byte, 16)
					if _, err = buf.Read(val); err != nil {
						return err
					}
					if err := roxe.UnmarshalBinary(val, &obj.SecondaryKey); err != nil {
						return err
					}
					row = obj
				case 3: /* index256_object */
					obj := &Index256Object{ContractRow: contractRow}
					val := make([]byte, 32)
					if _, err = buf.Read(val); err != nil {
						return err
					}
					if err := roxe.UnmarshalBinary(val, &obj.SecondaryKey); err != nil {
						return err
					}
					row = obj
				case 4: /* index_double_object */
					obj := &IndexDoubleObject{ContractRow: contractRow}
					val := make([]byte, 8)
					if _, err = buf.Read(val); err != nil {
						return err
					}
					if err := roxe.UnmarshalBinary(val, &obj.SecondaryKey); err != nil {
						return err
					}
					row = obj
				case 5: /* index_long_double_object */
					obj := &IndexLongDoubleObject{ContractRow: contractRow}
					val := make([]byte, 16)
					if _, err = buf.Read(val); err != nil {
						return err
					}
					if err := roxe.UnmarshalBinary(val, &obj.SecondaryKey); err != nil {
						return err
					}
					row = obj
				}

				if err := f(row); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

type BlockState struct {
	/// from block_header_state_common
	BlockNum                         uint32                          `json:"block_num"`
	DposProposedIrreversibleBlocknum uint32                          `json:"dpos_proposed_irreversible_blocknum"`
	DposIrreversibleBlocknum         uint32                          `json:"dpos_irreversible_blocknum"`
	ActiveSchedule                   *roxe.ProducerAuthoritySchedule `json:"active_schedule"`
	BlockrootMerkle                  *roxe.MerkleRoot                `json:"blockroot_merkle"`
	ProducerToLastProduced           []roxe.PairAccountNameBlockNum  `json:"producer_to_last_produced"`
	ProducerToLastImpliedIrb         []roxe.PairAccountNameBlockNum  `json:"producer_to_last_implied_irb"`
	BlockSigningKey                  *roxe.BlockSigningAuthority     `json:"block_signing_key"`
	ConfirmCount                     []uint8                         `json:"confirm_count"`

	// from block_header_state
	BlockID                   roxe.Checksum256                   `json:"id"`
	Header                    *roxe.SignedBlockHeader            `json:"header"`
	PendingSchedule           *ScheduleInfo                      `json:"pending_schedule"`
	ActivatedProtocolFeatures *roxe.ProtocolFeatureActivationSet `json:"activated_protocol_features"`
}

type ScheduleInfo struct {
	ScheduleLIBNum uint32                          `json:"schedule_lib_num"`
	ScheduleHash   roxe.Checksum256                `json:"schedule_hash"`
	Schedule       *roxe.ProducerAuthoritySchedule `json:"schedule"`
}

func readBlockState(section *Section, f sectionCallbackFunc) (err error) {
	cnt := make([]byte, section.BufferSize)
	_, err = section.Buffer.Read(cnt)
	if err != nil {
		return
	}

	var state BlockState
	err = roxe.UnmarshalBinary(cnt, &state)
	if err != nil {
		return
	}

	if err := f(state); err != nil {
		return err
	}

	return nil
}

////

type AccountObject struct {
	Name         roxe.AccountName
	CreationDate roxe.BlockTimestamp
	RawABI       []byte
}

func readAccountObjects(section *Section, f sectionCallbackFunc) error {
	for i := uint64(0); i < section.RowCount; i++ {
		a := AccountObject{}
		cnt := make([]byte, 12)
		_, err := section.Buffer.Read(cnt)
		if err != nil {
			return err
		}

		if err := roxe.UnmarshalBinary(cnt[:8], &a.Name); err != nil {
			return err
		}

		if err := roxe.UnmarshalBinary(cnt[8:12], &a.CreationDate); err != nil {
			return err
		}

		val, err := readByteArray(section.Buffer)
		if err != nil {
			return err
		}

		a.RawABI = val

		if err := f(a); err != nil {
			return err
		}
	}
	return nil
}

////

type AccountMetadataObject struct {
	Name           roxe.AccountName //< name should not be changed within a chainbase modifier lambda
	RecvSequence   roxe.Uint64
	AuthSequence   roxe.Uint64
	CodeSequence   roxe.Uint64
	ABISequence    roxe.Uint64
	CodeHash       roxe.Checksum256
	LastCodeUpdate roxe.TimePoint
	Flags          uint32 // First flag means "privileged".
	VMType         byte
	VMVersion      byte
}

func readAccountMetadataObjects(section *Section, f sectionCallbackFunc) error {
	for i := uint64(0); i < section.RowCount; i++ {
		a := AccountMetadataObject{}
		cnt := make([]byte, 86) // account_metadata_object is fixed size 86 bytes
		_, err := section.Buffer.Read(cnt)
		if err != nil {
			return err
		}

		if err := roxe.UnmarshalBinary(cnt, &a); err != nil {
			return err
		}

		if err := f(a); err != nil {
			return err
		}
	}
	return nil
}

////

type ChainSnapshotHeader struct {
	Version uint32
}

func readChainSnapshotHeader(section *Section, f sectionCallbackFunc) error {
	cnt := make([]byte, section.BufferSize)
	_, err := section.Buffer.Read(cnt)
	if err != nil {
		return err
	}

	var header ChainSnapshotHeader
	err = roxe.UnmarshalBinary(cnt, &header)
	if err != nil {
		return err
	}

	if err := f(header); err != nil {
		return err
	}

	return nil
}

type GlobalPropertyObject struct {
	ProposedScheduleBlockNum uint32 `roxe:"optional"`
	ProposedSchedule         *roxe.ProducerAuthoritySchedule
	Configuration            ChainConfig
	ChainID                  roxe.Checksum256
}

func readGlobalPropertyObject(section *Section, f sectionCallbackFunc) error {
	cnt := make([]byte, section.BufferSize)
	_, err := section.Buffer.Read(cnt)
	if err != nil {
		return err
	}

	var obj GlobalPropertyObject
	err = roxe.UnmarshalBinary(cnt, &obj)
	if err != nil {
		return err
	}

	if err := f(obj); err != nil {
		return err
	}

	return nil
}

//

type ProtocolStateObject struct {
	ActivatedProtocolFeatures    []*ActivatedProtocolFeature
	PreactivatedProtocolFeatures []roxe.Checksum256
	WhitelistedIntrinsics        []string
	NumSupportedKeyTypes         uint32
}

type ActivatedProtocolFeature struct {
	FeatureDigest      roxe.Checksum256
	ActivationBlockNum uint32
}

func readProtocolStateObject(section *Section, f sectionCallbackFunc) error {
	cnt := make([]byte, section.BufferSize)
	_, err := section.Buffer.Read(cnt)
	if err != nil {
		return err
	}

	var obj ProtocolStateObject
	err = roxe.UnmarshalBinary(cnt, &obj)
	if err != nil {
		return err
	}

	if err := f(obj); err != nil {
		return err
	}

	return nil
}

//

type DynamicGlobalPropertyObject struct {
	GlobalActionSequence roxe.Uint64
}

func readDynamicGlobalPropertyObject(section *Section, f sectionCallbackFunc) error {
	cnt := make([]byte, section.BufferSize)
	_, err := section.Buffer.Read(cnt)
	if err != nil {
		return err
	}

	var obj DynamicGlobalPropertyObject
	err = roxe.UnmarshalBinary(cnt, &obj)
	if err != nil {
		return err
	}

	if err := f(obj); err != nil {
		return err
	}

	return nil
}

//

type AccountRAMCorrectionObject struct {
	Name          roxe.AccountName
	RAMCorrection roxe.Uint64
}

func readAccountRAMCorrectionObject(section *Section, f sectionCallbackFunc) error {
	for i := uint64(0); i < section.RowCount; i++ {
		a := AccountRAMCorrectionObject{}
		cnt := make([]byte, 16) // fixed size of account_ram_correction_object
		_, err := section.Buffer.Read(cnt)
		if err != nil {
			return err
		}

		if err := roxe.UnmarshalBinary(cnt, &a); err != nil {
			return err
		}

		if err := f(a); err != nil {
			return err
		}
	}
	return nil
}

//

type BlockSummary struct {
	BlockID roxe.Checksum256
}

func readBlockSummary(section *Section, f sectionCallbackFunc) error {
	for i := uint64(0); i < section.RowCount; i++ {
		a := BlockSummary{}
		cnt := make([]byte, 32) // fixed size of block_summary
		_, err := section.Buffer.Read(cnt)
		if err != nil {
			return err
		}

		if err := roxe.UnmarshalBinary(cnt, &a); err != nil {
			return err
		}

		if err := f(a); err != nil {
			return err
		}
	}
	return nil
}

///

type PermissionObject struct { /* special snapshot version of the object */
	Parent      roxe.PermissionName ///< parent permission
	Owner       roxe.AccountName    ///< the account this permission belongs to
	Name        roxe.PermissionName ///< human-readable name for the permission
	LastUpdated roxe.TimePoint      ///< the last time this authority was updated
	LastUsed    roxe.TimePoint      ///< when this permission was last used
	Auth        roxe.Authority      ///< authority required to execute this permission
}

func readPermissionObject(section *Section, f sectionCallbackFunc) error {
	cnt := make([]byte, section.BufferSize)
	_, err := section.Buffer.Read(cnt)
	if err != nil {
		return err
	}

	for pos := 0; pos < int(section.BufferSize); {
		d := roxe.NewDecoder(cnt[pos:])
		var po PermissionObject
		err = d.Decode(&po)
		if err != nil {
			return err
		}

		if err := f(po); err != nil {
			return err
		}

		pos += d.LastPos()
	}

	// _ = ioutil.WriteFile("/tmp/test.dat", cnt, 0664)

	return nil
}

///

type PermissionLinkObject struct {
	/// The account which is defining its permission requirements
	Account roxe.AccountName
	/// The contract which account requires @ref required_permission to invoke
	Code roxe.AccountName
	/// The message type which account requires @ref required_permission to invoke
	/// May be empty; if so, it sets a default @ref required_permission for all messages to @ref code
	MessageType roxe.ActionName
	/// The permission level which @ref account requires for the specified message types
	/// all of the above fields should not be changed within a chainbase modifier lambda
	RequiredPermission roxe.PermissionName
}

func readPermissionLinkObject(section *Section, f sectionCallbackFunc) error {
	cnt := make([]byte, section.BufferSize)
	_, err := section.Buffer.Read(cnt)
	if err != nil {
		return err
	}

	for pos := 0; pos < int(section.BufferSize); {
		d := roxe.NewDecoder(cnt[pos:])
		var po PermissionLinkObject
		err = d.Decode(&po)
		if err != nil {
			return err
		}

		if err := f(po); err != nil {
			return err
		}

		pos += d.LastPos()
	}

	return nil
}

////

type ResourceLimitsObject struct {
	Owner roxe.AccountName //<  should not be changed within a chainbase modifier lambda

	NetWeight roxe.Int64
	CPUWeight roxe.Int64
	RAMBytes  roxe.Int64
}

func readResourceLimitsObject(section *Section, f sectionCallbackFunc) error {
	for i := uint64(0); i < section.RowCount; i++ {
		a := ResourceLimitsObject{}
		cnt := make([]byte, 8+8+8+8) // fixed size of resource_limits_object
		_, err := section.Buffer.Read(cnt)
		if err != nil {
			return err
		}

		if err := roxe.UnmarshalBinary(cnt, &a); err != nil {
			return err
		}

		if err := f(a); err != nil {
			return err
		}
	}
	return nil
}

////

type ResourceUsageObject struct {
	Owner roxe.AccountName //< owner should not be changed within a chainbase modifier lambda

	NetUsage UsageAccumulator
	CPUUsage UsageAccumulator

	RAMUsage roxe.Uint64
}

func readResourceUsageObject(section *Section, f sectionCallbackFunc) error {
	for i := uint64(0); i < section.RowCount; i++ {
		a := ResourceUsageObject{}
		cnt := make([]byte, 8+20+20+8) // fixed size of resource_limits_object
		_, err := section.Buffer.Read(cnt)
		if err != nil {
			return err
		}

		if err := roxe.UnmarshalBinary(cnt, &a); err != nil {
			return err
		}

		if err := f(a); err != nil {
			return err
		}
	}
	return nil
}

////

type ResourceLimitsStateObject struct {
	/**
	 * Track the average netusage for blocks
	 */
	AverageBlockNetUsage UsageAccumulator

	/**
	 * Track the average cpu usage for blocks
	 */
	AverageBlockCPUUsage UsageAccumulator

	PendingNetUsage roxe.Uint64
	PendingCPUUsage roxe.Uint64

	TotalNetWeight roxe.Uint64
	TotalCPUWeight roxe.Uint64
	TotalRAMBytes  roxe.Uint64

	/**
	 * The virtual number of bytes that would be consumed over blocksize_average_window_ms
	 * if all blocks were at their maximum virtual size. This is virtual because the
	 * real maximum block is less, this virtual number is only used for rate limiting users.
	 *
	 * It's lowest possible value is max_block_size * blocksize_average_window_ms / block_interval
	 * It's highest possible value is config::maximum_elastic_resource_multiplier (1000) times its lowest possible value
	 *
	 * This means that the most an account can consume during idle periods is 1000x the bandwidth
	 * it is gauranteed under congestion.
	 *
	 * Increases when average_block_size < target_block_size, decreases when
	 * average_block_size > target_block_size, with a cap at 1000x max_block_size
	 * and a floor at max_block_size;
	 **/
	VirtualNetLimit roxe.Uint64

	/**
	 *  Increases when average_bloc
	 */
	VirtualCPULimit roxe.Uint64
}

type UsageAccumulator struct {
	LastOrdinal uint32      ///< The ordinal of the last period which has contributed to the average
	ValueEx     roxe.Uint64 ///< The current average pre-multiplied by Precision
	Consumed    roxe.Uint64 ///< The last periods average + the current periods contribution so far
}

func readResourceLimitsStateObject(section *Section, f sectionCallbackFunc) error {
	cnt := make([]byte, section.BufferSize)
	_, err := section.Buffer.Read(cnt)
	if err != nil {
		return err
	}

	var obj ResourceLimitsStateObject
	if err = roxe.UnmarshalBinary(cnt, &obj); err != nil {
		return fmt.Errorf("unmarshal binary: %w", err)
	}

	return f(obj)
}

////

type ResourceLimitsConfigObject struct {
	CPULimitParameters ElasticLimitParameters
	NetLimitParameters ElasticLimitParameters

	AccountCPUUsageAverageWindow uint32
	AccountNetUsageAverageWindow uint32
}

type ElasticLimitParameters struct {
	Target  roxe.Uint64 // the desired usage
	Max     roxe.Uint64 // the maximum usage
	Periods uint32      // the number of aggregation periods that contribute to the average usage

	MaxMultiplier uint32      // the multiplier by which virtual space can oversell usage when uncongested
	ContractRate  roxe.Uint64 // the rate at which a congested resource contracts its limit
	ExpandRate    roxe.Uint64 // the rate at which an uncongested resource expands its limits
}

func readResourceLimitsConfigObject(section *Section, f sectionCallbackFunc) error {
	cnt := make([]byte, section.BufferSize)
	_, err := section.Buffer.Read(cnt)
	if err != nil {
		return err
	}

	var obj ResourceLimitsConfigObject
	if err = roxe.UnmarshalBinary(cnt, &obj); err != nil {
		return fmt.Errorf("unmarshal binary: %w", err)
	}

	return f(obj)
}

////

type CodeObject struct {
	CodeHash       roxe.Checksum256 //< code_hash should not be changed within a chainbase modifier lambda
	Code           roxe.HexBytes
	CodeRefCount   roxe.Uint64
	FirstBlockUsed uint32
	VMType         uint8 //< vm_type should not be changed within a chainbase modifier lambda
	VMVersion      uint8 //< vm_version should not be changed within a chainbase modifier lambda
}

func readCodeObject(section *Section, f sectionCallbackFunc) error {
	cnt := make([]byte, section.BufferSize)
	readz, err := section.Buffer.Read(cnt)
	if err != nil {
		return err
	}
	if readz != len(cnt) {
		return fmt.Errorf("failed reading the whole code object section: %d of %d", readz, len(cnt))
	}

	for pos := 0; pos < int(section.BufferSize); {
		d := roxe.NewDecoder(cnt[pos:])
		var co CodeObject
		err = d.Decode(&co)
		if err != nil {
			return err
		}

		if err := f(co); err != nil {
			return err
		}

		pos += d.LastPos()
	}

	return nil
}

////

type GeneratedTransactionObject struct {
	TrxID      roxe.Checksum256 //< trx_id should not be changed within a chainbase modifier lambda
	Sender     roxe.AccountName //< sender should not be changed within a chainbase modifier lambda
	SenderID   roxe.Uint128     /// ID given this transaction by the sender (should not be changed within a chainbase modifier lambda)
	Payer      roxe.AccountName
	DelayUntil roxe.TimePoint /// this generated transaction will not be applied until the specified time
	Expiration roxe.TimePoint /// this generated transaction will not be applied after  time
	Published  roxe.TimePoint
	PackedTrx  roxe.HexBytes
}

func readGeneratedTransactionObject(section *Section, f sectionCallbackFunc) error {
	cnt := make([]byte, section.BufferSize)
	readz, err := section.Buffer.Read(cnt)
	if err != nil {
		return err
	}
	if readz != len(cnt) {
		return fmt.Errorf("failed reading the whole code object section: %d of %d", readz, len(cnt))
	}

	for pos := 0; pos < int(section.BufferSize); {
		d := roxe.NewDecoder(cnt[pos:])
		var gto GeneratedTransactionObject
		err = d.Decode(&gto)
		if err != nil {
			return err
		}

		if err := f(gto); err != nil {
			return err
		}

		pos += d.LastPos()
	}

	return nil
}

////

type TransactionObject struct {
	Expiration roxe.TimePointSec
	TrxID      roxe.Checksum256 //< trx_id shou
}

func readTransactionObject(section *Section, f sectionCallbackFunc) error {
	cnt := make([]byte, section.BufferSize)
	readz, err := section.Buffer.Read(cnt)
	if err != nil {
		return err
	}
	if readz != len(cnt) {
		return fmt.Errorf("failed reading the whole code object section: %d of %d", readz, len(cnt))
	}

	for pos := 0; pos < int(section.BufferSize); {
		d := roxe.NewDecoder(cnt[pos:])
		var to TransactionObject
		err = d.Decode(&to)
		if err != nil {
			return err
		}

		if err := f(to); err != nil {
			return err
		}

		pos += d.LastPos()
	}

	return nil
}

/// Ultra Specific

type FreeObjectUsage struct {
	UserSize   roxe.Uint64
	UserCount  roxe.Uint64
	UltraSize  roxe.Uint64
	UltraCount roxe.Uint64
}

type AccountFreeActionsObject struct {
	Name                 roxe.AccountName
	PermissionObject     FreeObjectUsage
	SharedKey            FreeObjectUsage
	PermissionLevel      FreeObjectUsage
	Wait                 FreeObjectUsage
	PermissionLinkObject FreeObjectUsage
}

func readAccountFreeActionsObject(section *Section, f sectionCallbackFunc) error {
	for i := uint64(0); i < section.RowCount; i++ {
		a := AccountFreeActionsObject{}
		cnt := make([]byte, 8+(8*4)*5) // fixed size of resource_limits_object
		_, err := section.Buffer.Read(cnt)
		if err != nil {
			return err
		}

		if err := roxe.UnmarshalBinary(cnt, &a); err != nil {
			return err
		}

		if err := f(a); err != nil {
			return err
		}
	}
	return nil
}
