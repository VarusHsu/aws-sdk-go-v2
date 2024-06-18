// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Import hub content.
//
// Hub APIs are only callable through SageMaker Studio.
func (c *Client) ImportHubContent(ctx context.Context, params *ImportHubContentInput, optFns ...func(*Options)) (*ImportHubContentOutput, error) {
	if params == nil {
		params = &ImportHubContentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportHubContent", params, optFns, c.addOperationImportHubContentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportHubContentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportHubContentInput struct {

	// The version of the hub content schema to import.
	//
	// This member is required.
	DocumentSchemaVersion *string

	// The hub content document that describes information about the hub content such
	// as type, associated containers, scripts, and more.
	//
	// This member is required.
	HubContentDocument *string

	// The name of the hub content to import.
	//
	// This member is required.
	HubContentName *string

	// The type of hub content to import.
	//
	// This member is required.
	HubContentType types.HubContentType

	// The name of the hub to import content into.
	//
	// This member is required.
	HubName *string

	// A description of the hub content to import.
	HubContentDescription *string

	// The display name of the hub content to import.
	HubContentDisplayName *string

	// A string that provides a description of the hub content. This string can
	// include links, tables, and standard markdown formating.
	HubContentMarkdown *string

	// The searchable keywords of the hub content.
	HubContentSearchKeywords []string

	// The version of the hub content to import.
	HubContentVersion *string

	// Any tags associated with the hub content.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type ImportHubContentOutput struct {

	// The ARN of the hub that the content was imported into.
	//
	// This member is required.
	HubArn *string

	// The ARN of the hub content that was imported.
	//
	// This member is required.
	HubContentArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportHubContentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpImportHubContent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpImportHubContent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ImportHubContent"); err != nil {
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
	if err = addOpImportHubContentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportHubContent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opImportHubContent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ImportHubContent",
	}
}
