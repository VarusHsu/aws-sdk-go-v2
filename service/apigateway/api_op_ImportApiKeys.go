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
)

// Import API keys from an external source, such as a CSV-formatted file.
func (c *Client) ImportApiKeys(ctx context.Context, params *ImportApiKeysInput, optFns ...func(*Options)) (*ImportApiKeysOutput, error) {
	stack := middleware.NewStack("ImportApiKeys", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpImportApiKeysMiddlewares(stack)
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
	addOpImportApiKeysValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opImportApiKeys(options.Region), middleware.Before)
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
			OperationName: "ImportApiKeys",
			Err:           err,
		}
	}
	out := result.(*ImportApiKeysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The POST request to import API keys from an external source, such as a
// CSV-formatted file.
type ImportApiKeysInput struct {
	Name *string

	// A query parameter to specify the input format to imported API keys. Currently,
	// only the csv format is supported.
	//
	// This member is required.
	Format types.ApiKeysFormat

	Template *bool

	TemplateSkipList []*string

	// A query parameter to indicate whether to rollback ApiKey () importation (true)
	// or not (false) when error is encountered.
	FailOnWarnings *bool

	Title *string
}

// The identifier of an ApiKey () used in a UsagePlan ().
type ImportApiKeysOutput struct {

	// A list of all the ApiKey () identifiers.
	Ids []*string

	// A list of warning messages.
	Warnings []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpImportApiKeysMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpImportApiKeys{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpImportApiKeys{}, middleware.After)
}

func newServiceMetadataMiddleware_opImportApiKeys(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "ImportApiKeys",
	}
}
