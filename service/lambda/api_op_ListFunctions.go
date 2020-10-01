// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of Lambda functions, with the version-specific configuration of
// each. Lambda returns up to 50 functions per call. Set FunctionVersion to ALL to
// include all published versions of each function in addition to the unpublished
// version. To get more information about a function or version, use GetFunction
// ().
func (c *Client) ListFunctions(ctx context.Context, params *ListFunctionsInput, optFns ...func(*Options)) (*ListFunctionsOutput, error) {
	stack := middleware.NewStack("ListFunctions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListFunctionsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListFunctions(options.Region), middleware.Before)
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
			OperationName: "ListFunctions",
			Err:           err,
		}
	}
	out := result.(*ListFunctionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFunctionsInput struct {

	// Set to ALL to include entries for all published versions of each function.
	FunctionVersion types.FunctionVersion

	// The maximum number of functions to return.
	MaxItems *int32

	// For Lambda@Edge functions, the AWS Region of the master function. For example,
	// us-east-1 filters the list of functions to only include Lambda@Edge functions
	// replicated from a master function in US East (N. Virginia). If specified, you
	// must set FunctionVersion to ALL.
	MasterRegion *string

	// Specify the pagination token that's returned by a previous request to retrieve
	// the next page of results.
	Marker *string
}

// A list of Lambda functions.
type ListFunctionsOutput struct {

	// A list of Lambda functions.
	Functions []*types.FunctionConfiguration

	// The pagination token that's included if more results are available.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListFunctionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListFunctions{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListFunctions{}, middleware.After)
}

func newServiceMetadataMiddleware_opListFunctions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "ListFunctions",
	}
}
