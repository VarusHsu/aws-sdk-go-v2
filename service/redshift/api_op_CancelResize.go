// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Cancels a resize operation for a cluster.
func (c *Client) CancelResize(ctx context.Context, params *CancelResizeInput, optFns ...func(*Options)) (*CancelResizeOutput, error) {
	stack := middleware.NewStack("CancelResize", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCancelResizeMiddlewares(stack)
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
	addOpCancelResizeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCancelResize(options.Region), middleware.Before)
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
			OperationName: "CancelResize",
			Err:           err,
		}
	}
	out := result.(*CancelResizeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelResizeInput struct {

	// The unique identifier for the cluster that you want to cancel a resize operation
	// for.
	//
	// This member is required.
	ClusterIdentifier *string
}

// Describes the result of a cluster resize operation.
type CancelResizeOutput struct {

	// An enum with possible values of ClassicResize and ElasticResize. These values
	// describe the type of resize operation being performed.
	ResizeType *string

	// The number of nodes that the cluster will have after the resize operation is
	// complete.
	TargetNumberOfNodes *int32

	// The average rate of the resize operation over the last few minutes, measured in
	// megabytes per second. After the resize operation completes, this value shows the
	// average rate of the entire resize operation.
	AvgResizeRateInMegaBytesPerSecond *float64

	// The estimated time remaining, in seconds, until the resize operation is
	// complete. This value is calculated based on the average resize rate and the
	// estimated amount of data remaining to be processed. Once the resize operation is
	// complete, this value will be 0.
	EstimatedTimeToCompletionInSeconds *int64

	// The names of tables that have not been yet imported. Valid Values: List of table
	// names
	ImportTablesNotStarted []*string

	// The amount of seconds that have elapsed since the resize operation began. After
	// the resize operation completes, this value shows the total actual time, in
	// seconds, for the resize operation.
	ElapsedTimeInSeconds *int64

	// The cluster type after the resize operation is complete. Valid Values:
	// multi-node | single-node
	TargetClusterType *string

	// The percent of data transferred from source cluster to target cluster.
	DataTransferProgressPercent *float64

	// The node type that the cluster will have after the resize operation is complete.
	TargetNodeType *string

	// The names of tables that have been completely imported . Valid Values: List of
	// table names.
	ImportTablesCompleted []*string

	// An optional string to provide additional details about the resize action.
	Message *string

	// The estimated total amount of data, in megabytes, on the cluster before the
	// resize operation began.
	TotalResizeDataInMegaBytes *int64

	// The type of encryption for the cluster after the resize is complete. Possible
	// values are KMS and None. In the China region possible values are: Legacy and
	// None.
	TargetEncryptionType *string

	// While the resize operation is in progress, this value shows the current amount
	// of data, in megabytes, that has been processed so far. When the resize operation
	// is complete, this value shows the total amount of data, in megabytes, on the
	// cluster, which may be more or less than TotalResizeDataInMegaBytes (the
	// estimated total amount of data before resize).
	ProgressInMegaBytes *int64

	// The names of tables that are being currently imported. Valid Values: List of
	// table names.
	ImportTablesInProgress []*string

	// The status of the resize operation. Valid Values: NONE | IN_PROGRESS | FAILED |
	// SUCCEEDED | CANCELLING
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCancelResizeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCancelResize{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCancelResize{}, middleware.After)
}

func newServiceMetadataMiddleware_opCancelResize(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "CancelResize",
	}
}
