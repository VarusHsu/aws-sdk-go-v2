// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Configuration options for configure Cognito streams.
type CognitoStreams struct {

	// A string containing the reason why streaming of updates to the identity pool was
	// disabled.
	DisabledReason *string

	// The name of the Cognito stream to receive updates. This stream must be in the
	// developers account and in the same region as the identity pool.
	StreamName *string

	// Status of the Cognito streams. Valid values are:
	//
	//     * ENABLED - Streaming of
	// updates to identity pool is enabled.
	//
	//     * DISABLED - Streaming of updates to
	// identity pool is disabled. Bulk publish will also fail if StreamingStatus is
	// DISABLED.
	StreamingStatus StreamingStatus

	// The ARN of the role Amazon Cognito can assume in order to publish to the stream.
	// This role must grant access to Amazon Cognito (cognito-sync) to invoke PutRecord
	// on your Cognito stream.
	RoleArn *string
}

// A collection of data for an identity pool. An identity pool can have multiple
// datasets. A dataset is per identity and can be general or associated with a
// particular entity in an application (like a saved game). Datasets are
// automatically created if they don't exist. Data is synced by dataset, and a
// dataset can hold up to 1MB of key-value pairs.
type Dataset struct {

	// Number of records in this dataset.
	NumRecords *int64

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	IdentityId *string

	// The device that made the last change to this dataset.
	LastModifiedBy *string

	// Date when the dataset was last modified.
	LastModifiedDate *time.Time

	// Date on which the dataset was created.
	CreationDate *time.Time

	// Total size in bytes of the records in this dataset.
	DataStorage *int64

	// A string of up to 128 characters. Allowed characters are a-z, A-Z, 0-9, '_'
	// (underscore), '-' (dash), and '.' (dot).
	DatasetName *string
}

// Usage information for the identity pool.
type IdentityPoolUsage struct {

	// Date on which the identity pool was last modified.
	LastModifiedDate *time.Time

	// Data storage information for the identity pool.
	DataStorage *int64

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	IdentityPoolId *string

	// Number of sync sessions for the identity pool.
	SyncSessionsCount *int64
}

// Usage information for the identity.
type IdentityUsage struct {

	// Number of datasets for the identity.
	DatasetCount *int32

	// Date on which the identity was last modified.
	LastModifiedDate *time.Time

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	IdentityId *string

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	IdentityPoolId *string

	// Total data storage for this identity.
	DataStorage *int64
}

// Configuration options to be applied to the identity pool.
type PushSync struct {

	// List of SNS platform application ARNs that could be used by clients.
	ApplicationArns []*string

	// A role configured to allow Cognito to call SNS on behalf of the developer.
	RoleArn *string
}

// The basic data structure of a dataset.
type Record struct {

	// The last modified date of the client device.
	DeviceLastModifiedDate *time.Time

	// The server sync count for this record.
	SyncCount *int64

	// The date on which the record was last modified.
	LastModifiedDate *time.Time

	// The value for the record.
	Value *string

	// The key for the record.
	Key *string

	// The user/device that made the last change to this record.
	LastModifiedBy *string
}

// An update operation for a record.
type RecordPatch struct {

	// The key associated with the record patch.
	//
	// This member is required.
	Key *string

	// The last modified date of the client device.
	DeviceLastModifiedDate *time.Time

	// An operation, either replace or remove.
	//
	// This member is required.
	Op Operation

	// Last known server sync count for this record. Set to 0 if unknown.
	//
	// This member is required.
	SyncCount *int64

	// The value associated with the record patch.
	Value *string
}
