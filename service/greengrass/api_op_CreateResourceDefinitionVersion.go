// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a version of a resource definition that has already been defined.
func (c *Client) CreateResourceDefinitionVersion(ctx context.Context, params *CreateResourceDefinitionVersionInput, optFns ...func(*Options)) (*CreateResourceDefinitionVersionOutput, error) {
	stack := middleware.NewStack("CreateResourceDefinitionVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateResourceDefinitionVersionMiddlewares(stack)
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
	addOpCreateResourceDefinitionVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateResourceDefinitionVersion(options.Region), middleware.Before)
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
			OperationName: "CreateResourceDefinitionVersion",
			Err:           err,
		}
	}
	out := result.(*CreateResourceDefinitionVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateResourceDefinitionVersionInput struct {

	// A client token used to correlate requests and responses.
	AmznClientToken *string

	// The ID of the resource definition.
	//
	// This member is required.
	ResourceDefinitionId *string

	// A list of resources.
	Resources []*types.Resource
}

type CreateResourceDefinitionVersionOutput struct {

	// The ARN of the version.
	Arn *string

	// The time, in milliseconds since the epoch, when the version was created.
	CreationTimestamp *string

	// The ID of the parent definition that the version is associated with.
	Id *string

	// The ID of the version.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateResourceDefinitionVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateResourceDefinitionVersion{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateResourceDefinitionVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateResourceDefinitionVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "CreateResourceDefinitionVersion",
	}
}
