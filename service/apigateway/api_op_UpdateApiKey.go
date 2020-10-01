// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Changes information about an ApiKey () resource.
func (c *Client) UpdateApiKey(ctx context.Context, params *UpdateApiKeyInput, optFns ...func(*Options)) (*UpdateApiKeyOutput, error) {
	stack := middleware.NewStack("UpdateApiKey", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateApiKeyMiddlewares(stack)
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
	addOpUpdateApiKeyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateApiKey(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)

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
			OperationName: "UpdateApiKey",
			Err:           err,
		}
	}
	out := result.(*UpdateApiKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to change information about an ApiKey () resource.
type UpdateApiKeyInput struct {
	Title *string

	Name *string

	// [Required] The identifier of the ApiKey () resource to be updated.
	//
	// This member is required.
	ApiKey *string

	Template *bool

	TemplateSkipList []*string

	// A list of update operations to be applied to the specified resource and in the
	// order specified in this list.
	PatchOperations []*types.PatchOperation
}

// A resource that can be distributed to callers for executing Method () resources
// that require an API key. API keys can be mapped to any Stage () on any RestApi
// (), which indicates that the callers with the API key can make requests to that
// stage. Use API Keys
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-api-keys.html)
type UpdateApiKeyOutput struct {

	// The name of the API Key.
	Name *string

	// A list of Stage () resources that are associated with the ApiKey () resource.
	StageKeys []*string

	// The timestamp when the API Key was created.
	CreatedDate *time.Time

	// The identifier of the API Key.
	Id *string

	// An AWS Marketplace customer identifier , when integrating with the AWS SaaS
	// Marketplace.
	CustomerId *string

	// The timestamp when the API Key was last updated.
	LastUpdatedDate *time.Time

	// Specifies whether the API Key can be used by callers.
	Enabled *bool

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]*string

	// The description of the API Key.
	Description *string

	// The value of the API Key.
	Value *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateApiKeyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateApiKey{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateApiKey{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateApiKey(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "UpdateApiKey",
	}
}
