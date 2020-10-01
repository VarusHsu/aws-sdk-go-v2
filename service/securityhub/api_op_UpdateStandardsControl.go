// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Used to control whether an individual security standard control is enabled or
// disabled.
func (c *Client) UpdateStandardsControl(ctx context.Context, params *UpdateStandardsControlInput, optFns ...func(*Options)) (*UpdateStandardsControlOutput, error) {
	stack := middleware.NewStack("UpdateStandardsControl", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateStandardsControlMiddlewares(stack)
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
	addOpUpdateStandardsControlValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateStandardsControl(options.Region), middleware.Before)
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
			OperationName: "UpdateStandardsControl",
			Err:           err,
		}
	}
	out := result.(*UpdateStandardsControlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateStandardsControlInput struct {

	// The updated status of the security standard control.
	ControlStatus types.ControlStatus

	// A description of the reason why you are disabling a security standard control.
	// If you are disabling a control, then this is required.
	DisabledReason *string

	// The ARN of the security standard control to enable or disable.
	//
	// This member is required.
	StandardsControlArn *string
}

type UpdateStandardsControlOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateStandardsControlMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateStandardsControl{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateStandardsControl{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateStandardsControl(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "UpdateStandardsControl",
	}
}
