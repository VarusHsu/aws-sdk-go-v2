// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describe the content of a hub.
//
// Hub APIs are only callable through SageMaker Studio.
func (c *Client) DescribeHubContent(ctx context.Context, params *DescribeHubContentInput, optFns ...func(*Options)) (*DescribeHubContentOutput, error) {
	if params == nil {
		params = &DescribeHubContentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeHubContent", params, optFns, c.addOperationDescribeHubContentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeHubContentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeHubContentInput struct {

	// The name of the content to describe.
	//
	// This member is required.
	HubContentName *string

	// The type of content in the hub.
	//
	// This member is required.
	HubContentType types.HubContentType

	// The name of the hub that contains the content to describe.
	//
	// This member is required.
	HubName *string

	// The version of the content to describe.
	HubContentVersion *string

	noSmithyDocumentSerde
}

type DescribeHubContentOutput struct {

	// The date and time that hub content was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The document schema version for the hub content.
	//
	// This member is required.
	DocumentSchemaVersion *string

	// The Amazon Resource Name (ARN) of the hub that contains the content.
	//
	// This member is required.
	HubArn *string

	// The Amazon Resource Name (ARN) of the hub content.
	//
	// This member is required.
	HubContentArn *string

	// The hub content document that describes information about the hub content such
	// as type, associated containers, scripts, and more.
	//
	// This member is required.
	HubContentDocument *string

	// The name of the hub content.
	//
	// This member is required.
	HubContentName *string

	// The status of the hub content.
	//
	// This member is required.
	HubContentStatus types.HubContentStatus

	// The type of hub content.
	//
	// This member is required.
	HubContentType types.HubContentType

	// The version of the hub content.
	//
	// This member is required.
	HubContentVersion *string

	// The name of the hub that contains the content.
	//
	// This member is required.
	HubName *string

	// The failure reason if importing hub content failed.
	FailureReason *string

	// The location of any dependencies that the hub content has, such as scripts,
	// model artifacts, datasets, or notebooks.
	HubContentDependencies []types.HubContentDependency

	// A description of the hub content.
	HubContentDescription *string

	// The display name of the hub content.
	HubContentDisplayName *string

	// A string that provides a description of the hub content. This string can
	// include links, tables, and standard markdown formating.
	HubContentMarkdown *string

	// The searchable keywords for the hub content.
	HubContentSearchKeywords []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeHubContentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeHubContent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeHubContent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeHubContent"); err != nil {
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
	if err = addOpDescribeHubContentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeHubContent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeHubContent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeHubContent",
	}
}
