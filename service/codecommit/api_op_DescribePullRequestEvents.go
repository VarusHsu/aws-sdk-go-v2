// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about one or more pull request events.
func (c *Client) DescribePullRequestEvents(ctx context.Context, params *DescribePullRequestEventsInput, optFns ...func(*Options)) (*DescribePullRequestEventsOutput, error) {
	stack := middleware.NewStack("DescribePullRequestEvents", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribePullRequestEventsMiddlewares(stack)
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
	addOpDescribePullRequestEventsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePullRequestEvents(options.Region), middleware.Before)
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
			OperationName: "DescribePullRequestEvents",
			Err:           err,
		}
	}
	out := result.(*DescribePullRequestEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePullRequestEventsInput struct {

	// An enumeration token that, when provided in a request, returns the next batch of
	// the results.
	NextToken *string

	// The Amazon Resource Name (ARN) of the user whose actions resulted in the event.
	// Examples include updating the pull request with more commits or changing the
	// status of a pull request.
	ActorArn *string

	// Optional. The pull request event type about which you want to return
	// information.
	PullRequestEventType types.PullRequestEventType

	// A non-zero, non-negative integer used to limit the number of returned results.
	// The default is 100 events, which is also the maximum number of events that can
	// be returned in a result.
	MaxResults *int32

	// The system-generated ID of the pull request. To get this ID, use
	// ListPullRequests ().
	//
	// This member is required.
	PullRequestId *string
}

type DescribePullRequestEventsOutput struct {

	// Information about the pull request events.
	//
	// This member is required.
	PullRequestEvents []*types.PullRequestEvent

	// An enumeration token that can be used in a request to return the next batch of
	// the results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribePullRequestEventsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePullRequestEvents{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePullRequestEvents{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribePullRequestEvents(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "DescribePullRequestEvents",
	}
}
