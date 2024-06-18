// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalogappregistry

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Disassociates a resource from application. Both the resource and the
//
// application can be specified either by ID or name.
//
// # Minimum permissions
//
// You must have the following permissions to remove a resource that's been
// associated with an application using the APPLY_APPLICATION_TAG option for [AssociateResource].
//
//   - tag:GetResources
//
//   - tag:UntagResources
//
// You must also have the following permissions if you don't use the
// AWSServiceCatalogAppRegistryFullAccess policy. For more information, see [AWSServiceCatalogAppRegistryFullAccess] in
// the AppRegistry Administrator Guide.
//
//   - resource-groups:DisassociateResource
//
//   - cloudformation:UpdateStack
//
//   - cloudformation:DescribeStacks
//
// In addition, you must have the tagging permission defined by the Amazon Web
// Services service that creates the resource. For more information, see [UntagResources]in the
// Resource Groups Tagging API Reference.
//
// [UntagResources]: https://docs.aws.amazon.com/resourcegroupstagging/latest/APIReference/API_UntTagResources.html
// [AWSServiceCatalogAppRegistryFullAccess]: https://docs.aws.amazon.com/servicecatalog/latest/arguide/full.html
// [AssociateResource]: https://docs.aws.amazon.com/servicecatalog/latest/dg/API_app-registry_AssociateResource.html
func (c *Client) DisassociateResource(ctx context.Context, params *DisassociateResourceInput, optFns ...func(*Options)) (*DisassociateResourceOutput, error) {
	if params == nil {
		params = &DisassociateResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateResource", params, optFns, c.addOperationDisassociateResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateResourceInput struct {

	// The name or ID of the application.
	//
	// This member is required.
	Application *string

	// The name or ID of the resource.
	//
	// This member is required.
	Resource *string

	// The type of the resource that is being disassociated.
	//
	// This member is required.
	ResourceType types.ResourceType

	noSmithyDocumentSerde
}

type DisassociateResourceOutput struct {

	// The Amazon resource name (ARN) that specifies the application.
	ApplicationArn *string

	// The Amazon resource name (ARN) that specifies the resource.
	ResourceArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisassociateResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociateResource{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DisassociateResource"); err != nil {
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
	if err = addOpDisassociateResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateResource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DisassociateResource",
	}
}
