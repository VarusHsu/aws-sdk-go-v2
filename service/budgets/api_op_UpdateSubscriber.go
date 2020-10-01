// Code generated by smithy-go-codegen DO NOT EDIT.

package budgets

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/budgets/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a subscriber.
func (c *Client) UpdateSubscriber(ctx context.Context, params *UpdateSubscriberInput, optFns ...func(*Options)) (*UpdateSubscriberOutput, error) {
	stack := middleware.NewStack("UpdateSubscriber", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateSubscriberMiddlewares(stack)
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
	addOpUpdateSubscriberValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSubscriber(options.Region), middleware.Before)
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
			OperationName: "UpdateSubscriber",
			Err:           err,
		}
	}
	out := result.(*UpdateSubscriberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request of UpdateSubscriber
type UpdateSubscriberInput struct {

	// The name of the budget whose subscriber you want to update.
	//
	// This member is required.
	BudgetName *string

	// The accountId that is associated with the budget whose subscriber you want to
	// update.
	//
	// This member is required.
	AccountId *string

	// The notification whose subscriber you want to update.
	//
	// This member is required.
	Notification *types.Notification

	// The previous subscriber that is associated with a budget notification.
	//
	// This member is required.
	OldSubscriber *types.Subscriber

	// The updated subscriber that is associated with a budget notification.
	//
	// This member is required.
	NewSubscriber *types.Subscriber
}

// Response of UpdateSubscriber
type UpdateSubscriberOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateSubscriberMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateSubscriber{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateSubscriber{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateSubscriber(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "budgets",
		OperationName: "UpdateSubscriber",
	}
}
