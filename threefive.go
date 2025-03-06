package splicefu

import (
	"strings"
)

func Tothreefive(jason string) string{
	r := strings.NewReplacer(
	"InfoSection","info_section",
	"TableID" ,"table_id",
	"SectionSyntaxIndicator","section_syntax_indicator",
	"Private","private",
	"SapType","sap_type",
	"SapDetails","sap_details",
	"SectionLength","section_length",
	"ProtocolVersion","protocol_version",
	"EncryptedPacket","encrypted_packet",
	"EncryptionAlgorithm ","encryption_algorithm",
	"PtsAdjustment","pts_adjustment",
	"CwIndex","cw_index",
	"Tier","tier",
	"SpliceCommandLength","splice_command_length",
	"SpliceCommandType","splice_command_type",
	"DescriptorLoopLength" ,"descriptor_loop_length",
	"Crc","crc",
		"CommandLength","command_length",
		"CommandType","command_type",
		"TimeSpecifiedFlag","time_specified_flag",
		"SpliceEventID","splice_event_id",
		"SpliceEventCancelIndicator","splice_event_cancel_indicator",
		"OutOfNetworkIndicator","out_of_network_indicator",
		"ProgramSpliceFlag","program_splice_flag",
		"DurationFlag","duration_flag",
		"SpliceImmediateFlag","splice_immediate_flag",
		"EventIDComplianceFlag","event_id_compliance_flag",
		"UniqueProgramID","unique_program_id",
		"AvailNum","avail_num",
		"AvailExpected","avail_expected",
		"DescriptorLength", "descriptor_length",
		"SegmentationEventID", "segmentation_event_id",
		"SegmentationEventCancelIndicator","segmentation_event_cancel_indicator",
		"SegmentationEventIDComplianceIndicator","segmentation_event_id_compliance_indicator",
		"ProgramSegmentationFlag", "program_segmentation_flag",
		"SegmentationDurationFlag", "segmentation_duration_flag",
		"DeliveryNotRestrictedFlag", "delivery_not_restricted_flag",
		"WebDeliveryAllowedFlag",  "web_delivery_allowed_flag",
		"NoRegionalBlackoutFlag",  "no_regional_blackout_flag",
		"ArchiveAllowedFlag",  "archive_allowed_flag",
		"DeviceRestrictions",  "device_restrictions",
		"SegmentationMessage",  "segmentation_message",
		"UpidType", "segmentation_upid_type",
		"SegmentationUpidTypeName", "segmentation_upid_type_name",
		"SegmentationUpidLength",  "segmentation_upid_length",
		"Value",  "segmentation_upid",
		"SegmentationTypeID", "segmentation_type_id",
		"SegmentNum", "segment_num",
		"SegmentsExpected", "segments_expected",
		"Command","command",
		"Length","length",
		"Tag", "tag",
		"Name", "name",
		"Identifier", "identifier",
		"PTS","pts_time",
		"Descriptors","descriptors")
		return r.Replace(jason)

}

