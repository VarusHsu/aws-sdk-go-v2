// Code generated by smithy-go-codegen DO NOT EDIT.

package honeycode

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/honeycode/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The InvokeScreenAutomation API allows invoking an action defined in a screen in
// a Honeycode app. The API allows setting local variables, which can then be used
// in the automation being invoked. This allows automating the Honeycode app
// interactions to write, update or delete data in the workbook.
func (c *Client) InvokeScreenAutomation(ctx context.Context, params *InvokeScreenAutomationInput, optFns ...func(*Options)) (*InvokeScreenAutomationOutput, error) {
	stack := middleware.NewStack("InvokeScreenAutomation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpInvokeScreenAutomationMiddlewares(stack)
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
	addOpInvokeScreenAutomationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opInvokeScreenAutomation(options.Region), middleware.Before)
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
			OperationName: "InvokeScreenAutomation",
			Err:           err,
		}
	}
	out := result.(*InvokeScreenAutomationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InvokeScreenAutomationInput struct {

	// The ID of the automation action to be performed.
	//
	// This member is required.
	ScreenAutomationId *string

	// The ID of the app that contains the screen automation.
	//
	// This member is required.
	AppId *string

	// The request token for performing the automation action. Request tokens help to
	// identify duplicate requests. If a call times out or fails due to a transient
	// error like a failed network connection, you can retry the call with the same
	// request token. The service ensures that if the first call using that request
	// token is successfully performed, the second call will return the response of the
	// previous call rather than performing the action again. Note that request tokens
	// are valid only for a few minutes. You cannot use request tokens to dedupe
	// requests spanning hours or days.
	ClientRequestToken *string

	// The ID of the workbook that contains the screen automation.
	//
	// This member is required.
	WorkbookId *string

	// Variables are optional and are needed only if the screen requires them to render
	// correctly. Variables are specified as a map where the key is the name of the
	// variable as defined on the screen. The value is an object which currently has
	// only one property, rawValue, which holds the value of the variable to be passed
	// to the screen.
	Variables map[string]*types.VariableValue

	// The row ID for the automation if the automation is defined inside a block with
	// source or list.
	RowId *string

	// The ID of the screen that contains the screen automation.
	//
	// This member is required.
	ScreenId *string
}

type InvokeScreenAutomationOutput struct {

	// The updated workbook cursor after performing the automation action.
	//
	// This member is required.
	WorkbookCursor *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpInvokeScreenAutomationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpInvokeScreenAutomation{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpInvokeScreenAutomation{}, middleware.After)
}

func newServiceMetadataMiddleware_opInvokeScreenAutomation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "honeycode",
		OperationName: "InvokeScreenAutomation",
	}
}
