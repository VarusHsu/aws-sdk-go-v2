// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ChannelLatencyMode string

// Enum values for ChannelLatencyMode
const (
	ChannelLatencyModeNormalLatency ChannelLatencyMode = "NORMAL"
	ChannelLatencyModeLowLatency    ChannelLatencyMode = "LOW"
)

// Values returns all known values for ChannelLatencyMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ChannelLatencyMode) Values() []ChannelLatencyMode {
	return []ChannelLatencyMode{
		"NORMAL",
		"LOW",
	}
}

type ChannelType string

// Enum values for ChannelType
const (
	ChannelTypeBasicChannelType      ChannelType = "BASIC"
	ChannelTypeStandardChannelType   ChannelType = "STANDARD"
	ChannelTypeAdvancedSDChannelType ChannelType = "ADVANCED_SD"
	ChannelTypeAdvancedHDChannelType ChannelType = "ADVANCED_HD"
)

// Values returns all known values for ChannelType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ChannelType) Values() []ChannelType {
	return []ChannelType{
		"BASIC",
		"STANDARD",
		"ADVANCED_SD",
		"ADVANCED_HD",
	}
}

type RecordingConfigurationState string

// Enum values for RecordingConfigurationState
const (
	RecordingConfigurationStateCreating     RecordingConfigurationState = "CREATING"
	RecordingConfigurationStateCreateFailed RecordingConfigurationState = "CREATE_FAILED"
	RecordingConfigurationStateActive       RecordingConfigurationState = "ACTIVE"
)

// Values returns all known values for RecordingConfigurationState. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RecordingConfigurationState) Values() []RecordingConfigurationState {
	return []RecordingConfigurationState{
		"CREATING",
		"CREATE_FAILED",
		"ACTIVE",
	}
}

type RecordingMode string

// Enum values for RecordingMode
const (
	RecordingModeDisabled RecordingMode = "DISABLED"
	RecordingModeInterval RecordingMode = "INTERVAL"
)

// Values returns all known values for RecordingMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RecordingMode) Values() []RecordingMode {
	return []RecordingMode{
		"DISABLED",
		"INTERVAL",
	}
}

type RenditionConfigurationRendition string

// Enum values for RenditionConfigurationRendition
const (
	RenditionConfigurationRenditionSd               RenditionConfigurationRendition = "SD"
	RenditionConfigurationRenditionHd               RenditionConfigurationRendition = "HD"
	RenditionConfigurationRenditionFullHd           RenditionConfigurationRendition = "FULL_HD"
	RenditionConfigurationRenditionLowestResolution RenditionConfigurationRendition = "LOWEST_RESOLUTION"
)

// Values returns all known values for RenditionConfigurationRendition. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RenditionConfigurationRendition) Values() []RenditionConfigurationRendition {
	return []RenditionConfigurationRendition{
		"SD",
		"HD",
		"FULL_HD",
		"LOWEST_RESOLUTION",
	}
}

type RenditionConfigurationRenditionSelection string

// Enum values for RenditionConfigurationRenditionSelection
const (
	RenditionConfigurationRenditionSelectionAll    RenditionConfigurationRenditionSelection = "ALL"
	RenditionConfigurationRenditionSelectionNone   RenditionConfigurationRenditionSelection = "NONE"
	RenditionConfigurationRenditionSelectionCustom RenditionConfigurationRenditionSelection = "CUSTOM"
)

// Values returns all known values for RenditionConfigurationRenditionSelection.
// Note that this can be expanded in the future, and so it is only as up to date as
// the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RenditionConfigurationRenditionSelection) Values() []RenditionConfigurationRenditionSelection {
	return []RenditionConfigurationRenditionSelection{
		"ALL",
		"NONE",
		"CUSTOM",
	}
}

type StreamHealth string

// Enum values for StreamHealth
const (
	StreamHealthStreamHealthy StreamHealth = "HEALTHY"
	StreamHealthStarving      StreamHealth = "STARVING"
	StreamHealthUnknown       StreamHealth = "UNKNOWN"
)

// Values returns all known values for StreamHealth. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (StreamHealth) Values() []StreamHealth {
	return []StreamHealth{
		"HEALTHY",
		"STARVING",
		"UNKNOWN",
	}
}

type StreamState string

// Enum values for StreamState
const (
	StreamStateStreamLive    StreamState = "LIVE"
	StreamStateStreamOffline StreamState = "OFFLINE"
)

// Values returns all known values for StreamState. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (StreamState) Values() []StreamState {
	return []StreamState{
		"LIVE",
		"OFFLINE",
	}
}

type ThumbnailConfigurationResolution string

// Enum values for ThumbnailConfigurationResolution
const (
	ThumbnailConfigurationResolutionSd               ThumbnailConfigurationResolution = "SD"
	ThumbnailConfigurationResolutionHd               ThumbnailConfigurationResolution = "HD"
	ThumbnailConfigurationResolutionFullHd           ThumbnailConfigurationResolution = "FULL_HD"
	ThumbnailConfigurationResolutionLowestResolution ThumbnailConfigurationResolution = "LOWEST_RESOLUTION"
)

// Values returns all known values for ThumbnailConfigurationResolution. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ThumbnailConfigurationResolution) Values() []ThumbnailConfigurationResolution {
	return []ThumbnailConfigurationResolution{
		"SD",
		"HD",
		"FULL_HD",
		"LOWEST_RESOLUTION",
	}
}

type ThumbnailConfigurationStorage string

// Enum values for ThumbnailConfigurationStorage
const (
	ThumbnailConfigurationStorageSequential ThumbnailConfigurationStorage = "SEQUENTIAL"
	ThumbnailConfigurationStorageLatest     ThumbnailConfigurationStorage = "LATEST"
)

// Values returns all known values for ThumbnailConfigurationStorage. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ThumbnailConfigurationStorage) Values() []ThumbnailConfigurationStorage {
	return []ThumbnailConfigurationStorage{
		"SEQUENTIAL",
		"LATEST",
	}
}

type TranscodePreset string

// Enum values for TranscodePreset
const (
	TranscodePresetHigherBandwidthTranscodePreset      TranscodePreset = "HIGHER_BANDWIDTH_DELIVERY"
	TranscodePresetConstrainedBandwidthTranscodePreset TranscodePreset = "CONSTRAINED_BANDWIDTH_DELIVERY"
)

// Values returns all known values for TranscodePreset. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TranscodePreset) Values() []TranscodePreset {
	return []TranscodePreset{
		"HIGHER_BANDWIDTH_DELIVERY",
		"CONSTRAINED_BANDWIDTH_DELIVERY",
	}
}
