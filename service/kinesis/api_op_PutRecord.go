// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Writes a single data record into an Amazon Kinesis data stream. Call PutRecord
// to send data into the stream for real-time ingestion and subsequent processing,
// one record at a time. Each shard can support writes up to 1,000 records per
// second, up to a maximum data write total of 1 MB per second. You must specify
// the name of the stream that captures, stores, and transports the data; a
// partition key; and the data blob itself. The data blob can be any type of data;
// for example, a segment from a log file, geographic/location data, website
// clickstream data, and so on. The partition key is used by Kinesis Data Streams
// to distribute data across shards. Kinesis Data Streams segregates the data
// records that belong to a stream into multiple shards, using the partition key
// associated with each data record to determine the shard to which a given data
// record belongs. Partition keys are Unicode strings, with a maximum length limit
// of 256 characters for each key. An MD5 hash function is used to map partition
// keys to 128-bit integer values and to map associated data records to shards
// using the hash key ranges of the shards. You can override hashing the partition
// key to determine the shard by explicitly specifying a hash value using the
// ExplicitHashKey parameter. For more information, see Adding Data to a Stream
// (https://docs.aws.amazon.com/kinesis/latest/dev/developing-producers-with-sdk.html#kinesis-using-sdk-java-add-data-to-stream)
// in the Amazon Kinesis Data Streams Developer Guide. PutRecord returns the shard
// ID of where the data record was placed and the sequence number that was assigned
// to the data record. Sequence numbers increase over time and are specific to a
// shard within a stream, not across all shards within a stream. To guarantee
// strictly increasing ordering, write serially to a shard and use the
// SequenceNumberForOrdering parameter. For more information, see Adding Data to a
// Stream
// (https://docs.aws.amazon.com/kinesis/latest/dev/developing-producers-with-sdk.html#kinesis-using-sdk-java-add-data-to-stream)
// in the Amazon Kinesis Data Streams Developer Guide. If a PutRecord request
// cannot be processed because of insufficient provisioned throughput on the shard
// involved in the request, PutRecord throws
// ProvisionedThroughputExceededException. By default, data records are accessible
// for 24 hours from the time that they are added to a stream. You can use
// IncreaseStreamRetentionPeriod () or DecreaseStreamRetentionPeriod () to modify
// this retention period.
func (c *Client) PutRecord(ctx context.Context, params *PutRecordInput, optFns ...func(*Options)) (*PutRecordOutput, error) {
	stack := middleware.NewStack("PutRecord", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutRecordMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpPutRecordValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutRecord(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "PutRecord",
			Err:           err,
		}
	}
	out := result.(*PutRecordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for PutRecord.
type PutRecordInput struct {

	// Guarantees strictly increasing sequence numbers, for puts from the same client
	// and to the same partition key. Usage: set the SequenceNumberForOrdering of
	// record n to the sequence number of record n-1 (as returned in the result when
	// putting record n-1). If this parameter is not set, records are coarsely ordered
	// based on arrival time.
	SequenceNumberForOrdering *string

	// The name of the stream to put the data record into.
	//
	// This member is required.
	StreamName *string

	// The hash value used to explicitly determine the shard the data record is
	// assigned to by overriding the partition key hash.
	ExplicitHashKey *string

	// Determines which shard in the stream the data record is assigned to. Partition
	// keys are Unicode strings with a maximum length limit of 256 characters for each
	// key. Amazon Kinesis Data Streams uses the partition key as input to a hash
	// function that maps the partition key and associated data to a specific shard.
	// Specifically, an MD5 hash function is used to map partition keys to 128-bit
	// integer values and to map associated data records to shards. As a result of this
	// hashing mechanism, all data records with the same partition key map to the same
	// shard within the stream.
	//
	// This member is required.
	PartitionKey *string

	// The data blob to put into the record, which is base64-encoded when the blob is
	// serialized. When the data blob (the payload before base64-encoding) is added to
	// the partition key size, the total size must not exceed the maximum record size
	// (1 MB).
	//
	// This member is required.
	Data []byte
}

// Represents the output for PutRecord.
type PutRecordOutput struct {

	// The encryption type to use on the record. This parameter can be one of the
	// following values:
	//
	//     * NONE: Do not encrypt the records in the stream.
	//
	//     *
	// KMS: Use server-side encryption on the records in the stream using a
	// customer-managed AWS KMS key.
	EncryptionType types.EncryptionType

	// The shard ID of the shard where the data record was placed.
	//
	// This member is required.
	ShardId *string

	// The sequence number identifier that was assigned to the put data record. The
	// sequence number for the record is unique across all records in the stream. A
	// sequence number is the identifier associated with every record put into the
	// stream.
	//
	// This member is required.
	SequenceNumber *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutRecordMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutRecord{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutRecord{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutRecord(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "PutRecord",
	}
}
