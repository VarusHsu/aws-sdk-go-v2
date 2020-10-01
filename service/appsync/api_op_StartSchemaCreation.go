// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds a new schema to your GraphQL API. This operation is asynchronous. Use to
// determine when it has completed.
func (c *Client) StartSchemaCreation(ctx context.Context, params *StartSchemaCreationInput, optFns ...func(*Options)) (*StartSchemaCreationOutput, error) {
	stack := middleware.NewStack("StartSchemaCreation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpStartSchemaCreationMiddlewares(stack)
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
	addOpStartSchemaCreationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartSchemaCreation(options.Region), middleware.Before)
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
			OperationName: "StartSchemaCreation",
			Err:           err,
		}
	}
	out := result.(*StartSchemaCreationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartSchemaCreationInput struct {

	// The schema definition, in GraphQL schema language format.
	//
	// This member is required.
	Definition []byte

	// The API ID.
	//
	// This member is required.
	ApiId *string
}

type StartSchemaCreationOutput struct {

	// The current state of the schema (PROCESSING, FAILED, SUCCESS, or
	// NOT_APPLICABLE). When the schema is in the ACTIVE state, you can add data.
	Status types.SchemaStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpStartSchemaCreationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpStartSchemaCreation{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpStartSchemaCreation{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartSchemaCreation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appsync",
		OperationName: "StartSchemaCreation",
	}
}
