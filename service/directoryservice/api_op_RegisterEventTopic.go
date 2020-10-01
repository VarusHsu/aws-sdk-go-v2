// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates a directory with an SNS topic. This establishes the directory as a
// publisher to the specified SNS topic. You can then receive email or text (SMS)
// messages when the status of your directory changes. You get notified if your
// directory goes from an Active status to an Impaired or Inoperable status. You
// also receive a notification when the directory returns to an Active status.
func (c *Client) RegisterEventTopic(ctx context.Context, params *RegisterEventTopicInput, optFns ...func(*Options)) (*RegisterEventTopicOutput, error) {
	stack := middleware.NewStack("RegisterEventTopic", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRegisterEventTopicMiddlewares(stack)
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
	addOpRegisterEventTopicValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterEventTopic(options.Region), middleware.Before)
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
			OperationName: "RegisterEventTopic",
			Err:           err,
		}
	}
	out := result.(*RegisterEventTopicOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Registers a new event topic.
type RegisterEventTopicInput struct {

	// The SNS topic name to which the directory will publish status messages. This SNS
	// topic must be in the same region as the specified Directory ID.
	//
	// This member is required.
	TopicName *string

	// The Directory ID that will publish status messages to the SNS topic.
	//
	// This member is required.
	DirectoryId *string
}

// The result of a RegisterEventTopic request.
type RegisterEventTopicOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRegisterEventTopicMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterEventTopic{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterEventTopic{}, middleware.After)
}

func newServiceMetadataMiddleware_opRegisterEventTopic(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "RegisterEventTopic",
	}
}
