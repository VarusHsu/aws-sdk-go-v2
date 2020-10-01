// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Delete an event destination. In Amazon Pinpoint, events include message sends,
// deliveries, opens, clicks, bounces, and complaints. Event destinations are
// places that you can send information about these events to. For example, you can
// send event data to Amazon SNS to receive notifications when you receive bounces
// or complaints, or you can use Amazon Kinesis Data Firehose to stream data to
// Amazon S3 for long-term storage.
func (c *Client) DeleteConfigurationSetEventDestination(ctx context.Context, params *DeleteConfigurationSetEventDestinationInput, optFns ...func(*Options)) (*DeleteConfigurationSetEventDestinationOutput, error) {
	stack := middleware.NewStack("DeleteConfigurationSetEventDestination", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteConfigurationSetEventDestinationMiddlewares(stack)
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
	addOpDeleteConfigurationSetEventDestinationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteConfigurationSetEventDestination(options.Region), middleware.Before)
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
			OperationName: "DeleteConfigurationSetEventDestination",
			Err:           err,
		}
	}
	out := result.(*DeleteConfigurationSetEventDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to delete an event destination from a configuration set.
type DeleteConfigurationSetEventDestinationInput struct {

	// The name of the configuration set that contains the event destination that you
	// want to delete.
	//
	// This member is required.
	ConfigurationSetName *string

	// The name of the event destination that you want to delete.
	//
	// This member is required.
	EventDestinationName *string
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type DeleteConfigurationSetEventDestinationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteConfigurationSetEventDestinationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteConfigurationSetEventDestination{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteConfigurationSetEventDestination{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteConfigurationSetEventDestination(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "DeleteConfigurationSetEventDestination",
	}
}
