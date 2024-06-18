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

//	Associates a resource with an application. The resource can be specified by
//
// its ARN or name. The application can be specified by ARN, ID, or name.
//
// # Minimum permissions
//
// You must have the following permissions to associate a resource using the
// OPTIONS parameter set to APPLY_APPLICATION_TAG .
//
//   - tag:GetResources
//
//   - tag:TagResources
//
// You must also have these additional permissions if you don't use the
// AWSServiceCatalogAppRegistryFullAccess policy. For more information, see [AWSServiceCatalogAppRegistryFullAccess] in
// the AppRegistry Administrator Guide.
//
//   - resource-groups:AssociateResource
//
//   - cloudformation:UpdateStack
//
//   - cloudformation:DescribeStacks
//
// In addition, you must have the tagging permission defined by the Amazon Web
// Services service that creates the resource. For more information, see [TagResources]in the
// Resource Groups Tagging API Reference.
//
// [TagResources]: https://docs.aws.amazon.com/resourcegroupstagging/latest/APIReference/API_TagResources.html
// [AWSServiceCatalogAppRegistryFullAccess]: https://docs.aws.amazon.com/servicecatalog/latest/arguide/full.html
func (c *Client) AssociateResource(ctx context.Context, params *AssociateResourceInput, optFns ...func(*Options)) (*AssociateResourceOutput, error) {
	if params == nil {
		params = &AssociateResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateResource", params, optFns, c.addOperationAssociateResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateResourceInput struct {

	//  The name, ID, or ARN of the application.
	//
	// This member is required.
	Application *string

	// The name or ID of the resource of which the application will be associated.
	//
	// This member is required.
	Resource *string

	// The type of resource of which the application will be associated.
	//
	// This member is required.
	ResourceType types.ResourceType

	//  Determines whether an application tag is applied or skipped.
	Options []types.AssociationOption

	noSmithyDocumentSerde
}

type AssociateResourceOutput struct {

	// The Amazon resource name (ARN) of the application that was augmented with
	// attributes.
	ApplicationArn *string

	//  Determines whether an application tag is applied or skipped.
	Options []types.AssociationOption

	// The Amazon resource name (ARN) that specifies the resource.
	ResourceArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateResource{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociateResource"); err != nil {
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
	if err = addOpAssociateResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateResource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociateResource",
	}
}
