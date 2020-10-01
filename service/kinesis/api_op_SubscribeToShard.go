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

// Call this operation from your consumer after you call RegisterStreamConsumer ()
// to register the consumer with Kinesis Data Streams. If the call succeeds, your
// consumer starts receiving events of type SubscribeToShardEvent () for up to 5
// minutes, after which time you need to call SubscribeToShard again to renew the
// subscription if you want to continue to receive records. You can make one call
// to SubscribeToShard per second per ConsumerARN. If your call succeeds, and then
// you call the operation again less than 5 seconds later, the second call
// generates a ResourceInUseException (). If you call the operation a second time
// more than 5 seconds after the first call succeeds, the second call succeeds and
// the first connection gets shut down.
func (c *Client) SubscribeToShard(ctx context.Context, params *SubscribeToShardInput, optFns ...func(*Options)) (*SubscribeToShardOutput, error) {
	stack := middleware.NewStack("SubscribeToShard", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpSubscribeToShardMiddlewares(stack)
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
	addOpSubscribeToShardValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSubscribeToShard(options.Region), middleware.Before)
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
			OperationName: "SubscribeToShard",
			Err:           err,
		}
	}
	out := result.(*SubscribeToShardOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SubscribeToShardInput struct {

	// For this parameter, use the value you obtained when you called
	// RegisterStreamConsumer ().
	//
	// This member is required.
	ConsumerARN *string

	// The ID of the shard you want to subscribe to. To see a list of all the shards
	// for a given stream, use ListShards ().
	//
	// This member is required.
	ShardId *string

	StartingPosition *types.StartingPosition
}

type SubscribeToShardOutput struct {

	// The event stream that your consumer can use to read records from the shard.
	//
	// This member is required.
	EventStream types.SubscribeToShardEventStream

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpSubscribeToShardMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpSubscribeToShard{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpSubscribeToShard{}, middleware.After)
}

func newServiceMetadataMiddleware_opSubscribeToShard(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "SubscribeToShard",
	}
}
