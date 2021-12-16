package snapshot

import "io"

type SectionName string

const (
	SectionNameChainSnapshotHeader         SectionName = "roxe::chain::chain_snapshot_header"
	SectionNameBlockState                  SectionName = "roxe::chain::block_state"
	SectionNameAccountObject               SectionName = "roxe::chain::account_object"
	SectionNameAccountMetadataObject       SectionName = "roxe::chain::account_metadata_object"
	SectionNameAccountRamCorrectionObject  SectionName = "roxe::chain::account_ram_correction_object"
	SectionNameGlobalPropertyObject        SectionName = "roxe::chain::global_property_object"
	SectionNameProtocolStateObject         SectionName = "roxe::chain::protocol_state_object"
	SectionNameDynamicGlobalPropertyObject SectionName = "roxe::chain::dynamic_global_property_object"
	SectionNameBlockSummaryObject          SectionName = "roxe::chain::block_summary_object"
	SectionNameTransactionObject           SectionName = "roxe::chain::transaction_object"
	SectionNameGeneratedTransactionObject  SectionName = "roxe::chain::generated_transaction_object"
	SectionNameCodeObject                  SectionName = "roxe::chain::code_object"
	SectionNameContractTables              SectionName = "contract_tables"
	SectionNamePermissionObject            SectionName = "roxe::chain::permission_object"
	SectionNamePermissionLinkObject        SectionName = "roxe::chain::permission_link_object"
	SectionNameResourceLimitsObject        SectionName = "roxe::chain::resource_limits::resource_limits_object"
	SectionNameResourceUsageObject         SectionName = "roxe::chain::resource_limits::resource_usage_object"
	SectionNameResourceLimitsStateObject   SectionName = "roxe::chain::resource_limits::resource_limits_state_object"
	SectionNameResourceLimitsConfigObject  SectionName = "roxe::chain::resource_limits::resource_limits_config_object"
	SectionNameGenesisState                SectionName = "roxe::chain::genesis_state"

	// Ultra Specific
	SectionAccountFreeActionsObject SectionName = "roxe::chain::account_free_actions_object"
)

type Section struct {
	Name       SectionName
	Offset     uint64
	Size       uint64 // This includes the section name and row count
	BufferSize uint64 // This represents the bytes that are following the section header
	RowCount   uint64 // This is a count of rows packed in `Buffer`
	Buffer     io.Reader
}

type sectionHandlerFunc func(s *Section, f sectionCallbackFunc) error
type sectionCallbackFunc func(obj interface{}) error
