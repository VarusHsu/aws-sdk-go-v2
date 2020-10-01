// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the list of event trackers associated with the account. The response
// provides the properties for each event tracker, including the Amazon Resource
// Name (ARN) and tracking ID. For more information on event trackers, see
// CreateEventTracker ().
func (c *Client) ListEventTrackers(ctx context.Context, params *ListEventTrackersInput, optFns ...func(*Options)) (*ListEventTrackersOutput, error) {
	stack := middleware.NewStack("ListEventTrackers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListEventTrackersMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListEventTrackers(options.Region), middleware.Before)
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
			OperationName: "ListEventTrackers",
			Err:           err,
		}
	}
	out := result.(*ListEventTrackersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEventTrackersInput struct {

	// The maximum number of event trackers to return.
	MaxResults *int32

	// The ARN of a dataset group used to filter the response.
	DatasetGroupArn *string

	// A token returned from the previous call to ListEventTrackers for getting the
	// next set of event trackers (if they exist).
	NextToken *string
}

type ListEventTrackersOutput struct {

	// A token for getting the next set of event trackers (if they exist).
	NextToken *string

	// A list of event trackers.
	EventTrackers []*types.EventTrackerSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListEventTrackersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListEventTrackers{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEventTrackers{}, middleware.After)
}

func newServiceMetadataMiddleware_opListEventTrackers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "ListEventTrackers",
	}
}
