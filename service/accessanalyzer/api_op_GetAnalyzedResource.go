// Code generated by smithy-go-codegen DO NOT EDIT.

package accessanalyzer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about a resource that was analyzed.
func (c *Client) GetAnalyzedResource(ctx context.Context, params *GetAnalyzedResourceInput, optFns ...func(*Options)) (*GetAnalyzedResourceOutput, error) {
	stack := middleware.NewStack("GetAnalyzedResource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetAnalyzedResourceMiddlewares(stack)
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
	addOpGetAnalyzedResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAnalyzedResource(options.Region), middleware.Before)
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
			OperationName: "GetAnalyzedResource",
			Err:           err,
		}
	}
	out := result.(*GetAnalyzedResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Retrieves an analyzed resource.
type GetAnalyzedResourceInput struct {

	// The ARN of the analyzer to retrieve information from.
	//
	// This member is required.
	AnalyzerArn *string

	// The ARN of the resource to retrieve information about.
	//
	// This member is required.
	ResourceArn *string
}

// The response to the request.
type GetAnalyzedResourceOutput struct {

	// An AnalyedResource object that contains information that Access Analyzer found
	// when it analyzed the resource.
	Resource *types.AnalyzedResource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetAnalyzedResourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetAnalyzedResource{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAnalyzedResource{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetAnalyzedResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "access-analyzer",
		OperationName: "GetAnalyzedResource",
	}
}
