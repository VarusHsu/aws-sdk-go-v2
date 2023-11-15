// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified flow. You can also create and update flows using the
// Amazon Connect Flow language (https://docs.aws.amazon.com/connect/latest/APIReference/flow-language.html)
// .
func (c *Client) UpdateContactFlowContent(ctx context.Context, params *UpdateContactFlowContentInput, optFns ...func(*Options)) (*UpdateContactFlowContentOutput, error) {
	if params == nil {
		params = &UpdateContactFlowContentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateContactFlowContent", params, optFns, c.addOperationUpdateContactFlowContentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateContactFlowContentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateContactFlowContentInput struct {

	// The identifier of the flow.
	//
	// This member is required.
	ContactFlowId *string

	// The JSON string that represents the content of the flow. For an example, see
	// Example flow in Amazon Connect Flow language (https://docs.aws.amazon.com/connect/latest/APIReference/flow-language-example.html)
	// . Length Constraints: Minimum length of 1. Maximum length of 256000.
	//
	// This member is required.
	Content *string

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	noSmithyDocumentSerde
}

type UpdateContactFlowContentOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateContactFlowContentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateContactFlowContent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateContactFlowContent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateContactFlowContent"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateContactFlowContentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateContactFlowContent(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateContactFlowContent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateContactFlowContent",
	}
}
