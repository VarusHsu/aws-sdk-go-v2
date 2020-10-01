// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This API starts recording the contact when the agent joins the call.
// StartContactRecording is a one-time action. For example, if you use
// StopContactRecording to stop recording an ongoing call, you can't use
// StartContactRecording to restart it. For scenarios where the recording has
// started and you want to suspend and resume it, such as when collecting sensitive
// information (for example, a credit card number), use SuspendContactRecording and
// ResumeContactRecording. You can use this API to override the recording behavior
// configured in the Set recording behavior
// (https://docs.aws.amazon.com/connect/latest/adminguide/set-recording-behavior.html)
// block. Only voice recordings are supported at this time.
func (c *Client) StartContactRecording(ctx context.Context, params *StartContactRecordingInput, optFns ...func(*Options)) (*StartContactRecordingOutput, error) {
	stack := middleware.NewStack("StartContactRecording", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpStartContactRecordingMiddlewares(stack)
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
	addOpStartContactRecordingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartContactRecording(options.Region), middleware.Before)
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
			OperationName: "StartContactRecording",
			Err:           err,
		}
	}
	out := result.(*StartContactRecordingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartContactRecordingInput struct {

	// The identifier of the contact.
	//
	// This member is required.
	ContactId *string

	// The identifier of the contact. This is the identifier of the contact associated
	// with the first interaction with the contact center.
	//
	// This member is required.
	InitialContactId *string

	// Who is being recorded.
	//
	// This member is required.
	VoiceRecordingConfiguration *types.VoiceRecordingConfiguration

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string
}

type StartContactRecordingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpStartContactRecordingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpStartContactRecording{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpStartContactRecording{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartContactRecording(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "StartContactRecording",
	}
}
