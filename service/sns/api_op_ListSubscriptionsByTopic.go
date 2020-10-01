// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of the subscriptions to a specific topic. Each call returns a
// limited list of subscriptions, up to 100. If there are more subscriptions, a
// NextToken is also returned. Use the NextToken parameter in a new
// ListSubscriptionsByTopic call to get further results. This action is throttled
// at 30 transactions per second (TPS).
func (c *Client) ListSubscriptionsByTopic(ctx context.Context, params *ListSubscriptionsByTopicInput, optFns ...func(*Options)) (*ListSubscriptionsByTopicOutput, error) {
	stack := middleware.NewStack("ListSubscriptionsByTopic", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpListSubscriptionsByTopicMiddlewares(stack)
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
	addOpListSubscriptionsByTopicValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListSubscriptionsByTopic(options.Region), middleware.Before)
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
			OperationName: "ListSubscriptionsByTopic",
			Err:           err,
		}
	}
	out := result.(*ListSubscriptionsByTopicOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for ListSubscriptionsByTopic action.
type ListSubscriptionsByTopicInput struct {

	// Token returned by the previous ListSubscriptionsByTopic request.
	NextToken *string

	// The ARN of the topic for which you wish to find subscriptions.
	//
	// This member is required.
	TopicArn *string
}

// Response for ListSubscriptionsByTopic action.
type ListSubscriptionsByTopicOutput struct {

	// Token to pass along to the next ListSubscriptionsByTopic request. This element
	// is returned if there are more subscriptions to retrieve.
	NextToken *string

	// A list of subscriptions.
	Subscriptions []*types.Subscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpListSubscriptionsByTopicMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpListSubscriptionsByTopic{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpListSubscriptionsByTopic{}, middleware.After)
}

func newServiceMetadataMiddleware_opListSubscriptionsByTopic(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "ListSubscriptionsByTopic",
	}
}
