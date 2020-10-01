// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a summary of all of the pipelines associated with your account.
func (c *Client) ListPipelines(ctx context.Context, params *ListPipelinesInput, optFns ...func(*Options)) (*ListPipelinesOutput, error) {
	stack := middleware.NewStack("ListPipelines", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListPipelinesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPipelines(options.Region), middleware.Before)
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
			OperationName: "ListPipelines",
			Err:           err,
		}
	}
	out := result.(*ListPipelinesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a ListPipelines action.
type ListPipelinesInput struct {

	// An identifier that was returned from the previous list pipelines call. It can be
	// used to return the next set of pipelines in the list.
	NextToken *string
}

// Represents the output of a ListPipelines action.
type ListPipelinesOutput struct {

	// The list of pipelines.
	Pipelines []*types.PipelineSummary

	// If the amount of returned information is significantly large, an identifier is
	// also returned. It can be used in a subsequent list pipelines call to return the
	// next set of pipelines in the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListPipelinesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListPipelines{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPipelines{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPipelines(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "ListPipelines",
	}
}
