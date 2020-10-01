// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new Facet () in a schema. Facet creation is allowed only in
// development or applied schemas.
func (c *Client) CreateFacet(ctx context.Context, params *CreateFacetInput, optFns ...func(*Options)) (*CreateFacetOutput, error) {
	stack := middleware.NewStack("CreateFacet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateFacetMiddlewares(stack)
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
	addOpCreateFacetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFacet(options.Region), middleware.Before)
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
			OperationName: "CreateFacet",
			Err:           err,
		}
	}
	out := result.(*CreateFacetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFacetInput struct {

	// The schema ARN in which the new Facet () will be created. For more information,
	// see arns ().
	//
	// This member is required.
	SchemaArn *string

	// There are two different styles that you can define on any given facet, Static
	// and Dynamic. For static facets, all attributes must be defined in the schema.
	// For dynamic facets, attributes can be defined during data plane operations.
	FacetStyle types.FacetStyle

	// The name of the Facet (), which is unique for a given schema.
	//
	// This member is required.
	Name *string

	// The attributes that are associated with the Facet ().
	Attributes []*types.FacetAttribute

	// Specifies whether a given object created from this facet is of type node, leaf
	// node, policy or index.
	//
	//     * Node: Can have multiple children but one parent.
	//
	//
	// * Leaf node: Cannot have children but can have multiple parents.
	//
	//     * Policy:
	// Allows you to store a policy document and policy type. For more information, see
	// Policies
	// (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/key_concepts_directory.html#key_concepts_policies).
	//
	//
	// * Index: Can be created with the Index API.
	ObjectType types.ObjectType
}

type CreateFacetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateFacetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateFacet{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateFacet{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateFacet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "CreateFacet",
	}
}
