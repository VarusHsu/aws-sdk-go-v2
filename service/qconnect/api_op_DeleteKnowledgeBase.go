// Code generated by smithy-go-codegen DO NOT EDIT.

package qconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the knowledge base.
//
// When you use this API to delete an external knowledge base such as Salesforce
// or ServiceNow, you must also delete the [Amazon AppIntegrations]DataIntegration. This is because you
// can't reuse the DataIntegration after it's been associated with an external
// knowledge base. However, you can delete and recreate it. See [DeleteDataIntegration]and [CreateDataIntegration] in the Amazon
// AppIntegrations API Reference.
//
// [Amazon AppIntegrations]: https://docs.aws.amazon.com/appintegrations/latest/APIReference/Welcome.html
// [DeleteDataIntegration]: https://docs.aws.amazon.com/appintegrations/latest/APIReference/API_DeleteDataIntegration.html
// [CreateDataIntegration]: https://docs.aws.amazon.com/appintegrations/latest/APIReference/API_CreateDataIntegration.html
func (c *Client) DeleteKnowledgeBase(ctx context.Context, params *DeleteKnowledgeBaseInput, optFns ...func(*Options)) (*DeleteKnowledgeBaseOutput, error) {
	if params == nil {
		params = &DeleteKnowledgeBaseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteKnowledgeBase", params, optFns, c.addOperationDeleteKnowledgeBaseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteKnowledgeBaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteKnowledgeBaseInput struct {

	// The knowledge base to delete content from. Can be either the ID or the ARN.
	// URLs cannot contain the ARN.
	//
	// This member is required.
	KnowledgeBaseId *string

	noSmithyDocumentSerde
}

type DeleteKnowledgeBaseOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteKnowledgeBaseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteKnowledgeBase{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteKnowledgeBase{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteKnowledgeBase"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
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
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpDeleteKnowledgeBaseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteKnowledgeBase(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opDeleteKnowledgeBase(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteKnowledgeBase",
	}
}
