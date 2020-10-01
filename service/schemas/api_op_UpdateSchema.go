// Code generated by smithy-go-codegen DO NOT EDIT.

package schemas

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/schemas/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Updates the schema definition Inactive schemas will be deleted after two years.
func (c *Client) UpdateSchema(ctx context.Context, params *UpdateSchemaInput, optFns ...func(*Options)) (*UpdateSchemaOutput, error) {
	stack := middleware.NewStack("UpdateSchema", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateSchemaMiddlewares(stack)
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
	addIdempotencyToken_opUpdateSchemaMiddleware(stack, options)
	addOpUpdateSchemaValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSchema(options.Region), middleware.Before)
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
			OperationName: "UpdateSchema",
			Err:           err,
		}
	}
	out := result.(*UpdateSchemaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSchemaInput struct {

	// The source of the schema definition.
	Content *string

	// The name of the registry.
	//
	// This member is required.
	RegistryName *string

	// The name of the schema.
	//
	// This member is required.
	SchemaName *string

	// The description of the schema.
	Description *string

	// The schema type for the events schema.
	Type types.Type

	// The ID of the client token.
	ClientTokenId *string
}

type UpdateSchemaOutput struct {

	// The description of the schema.
	Description *string

	// The date and time that schema was modified.
	LastModified *time.Time

	// The ARN of the schema.
	SchemaArn *string

	// The name of the schema.
	SchemaName *string

	// The version number of the schema
	SchemaVersion *string

	// The date the schema version was created.
	VersionCreatedDate *time.Time

	// Key-value pairs associated with a resource.
	Tags map[string]*string

	// The type of the schema.
	Type *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateSchemaMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateSchema{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateSchema{}, middleware.After)
}

type idempotencyToken_initializeOpUpdateSchema struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateSchema) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateSchema) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateSchemaInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateSchemaInput ")
	}

	if input.ClientTokenId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientTokenId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateSchemaMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpUpdateSchema{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateSchema(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "schemas",
		OperationName: "UpdateSchema",
	}
}
