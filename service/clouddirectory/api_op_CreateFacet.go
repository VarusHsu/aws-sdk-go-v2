// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new Facet in a schema. Facet creation is allowed only in development or
// applied schemas.
func (c *Client) CreateFacet(ctx context.Context, params *CreateFacetInput, optFns ...func(*Options)) (*CreateFacetOutput, error) {
	if params == nil {
		params = &CreateFacetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFacet", params, optFns, c.addOperationCreateFacetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFacetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFacetInput struct {

	// The name of the Facet, which is unique for a given schema.
	//
	// This member is required.
	Name *string

	// The schema ARN in which the new Facet will be created. For more information, see arns.
	//
	// This member is required.
	SchemaArn *string

	// The attributes that are associated with the Facet.
	Attributes []types.FacetAttribute

	// There are two different styles that you can define on any given facet, Static
	// and Dynamic . For static facets, all attributes must be defined in the schema.
	// For dynamic facets, attributes can be defined during data plane operations.
	FacetStyle types.FacetStyle

	// Specifies whether a given object created from this facet is of type node, leaf
	// node, policy or index.
	//
	//   - Node: Can have multiple children but one parent.
	//
	//   - Leaf node: Cannot have children but can have multiple parents.
	//
	//   - Policy: Allows you to store a policy document and policy type. For more
	//   information, see [Policies].
	//
	//   - Index: Can be created with the Index API.
	//
	// [Policies]: https://docs.aws.amazon.com/clouddirectory/latest/developerguide/key_concepts_directory.html#key_concepts_policies
	ObjectType types.ObjectType

	noSmithyDocumentSerde
}

type CreateFacetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFacetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateFacet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateFacet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateFacet"); err != nil {
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
	if err = addOpCreateFacetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFacet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateFacet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateFacet",
	}
}
