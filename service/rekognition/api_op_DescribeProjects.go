// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists and gets information about your Amazon Rekognition Custom Labels projects.
// This operation requires permissions to perform the rekognition:DescribeProjects
// action.
func (c *Client) DescribeProjects(ctx context.Context, params *DescribeProjectsInput, optFns ...func(*Options)) (*DescribeProjectsOutput, error) {
	stack := middleware.NewStack("DescribeProjects", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeProjectsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeProjects(options.Region), middleware.Before)
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
			OperationName: "DescribeProjects",
			Err:           err,
		}
	}
	out := result.(*DescribeProjectsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeProjectsInput struct {

	// If the previous response was incomplete (because there is more results to
	// retrieve), Amazon Rekognition Custom Labels returns a pagination token in the
	// response. You can use this pagination token to retrieve the next set of results.
	NextToken *string

	// The maximum number of results to return per paginated call. The largest value
	// you can specify is 100. If you specify a value greater than 100, a
	// ValidationException error occurs. The default value is 100.
	MaxResults *int32
}

type DescribeProjectsOutput struct {

	// A list of project descriptions. The list is sorted by the date and time the
	// projects are created.
	ProjectDescriptions []*types.ProjectDescription

	// If the previous response was incomplete (because there is more results to
	// retrieve), Amazon Rekognition Custom Labels returns a pagination token in the
	// response. You can use this pagination token to retrieve the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeProjectsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeProjects{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeProjects{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeProjects(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "DescribeProjects",
	}
}
