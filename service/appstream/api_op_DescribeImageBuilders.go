// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a list that describes one or more specified image builders, if the
// image builder names are provided. Otherwise, all image builders in the account
// are described.
func (c *Client) DescribeImageBuilders(ctx context.Context, params *DescribeImageBuildersInput, optFns ...func(*Options)) (*DescribeImageBuildersOutput, error) {
	stack := middleware.NewStack("DescribeImageBuilders", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeImageBuildersMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeImageBuilders(options.Region), middleware.Before)
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
			OperationName: "DescribeImageBuilders",
			Err:           err,
		}
	}
	out := result.(*DescribeImageBuildersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeImageBuildersInput struct {

	// The maximum size of each page of results.
	MaxResults *int32

	// The pagination token to use to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string

	// The names of the image builders to describe.
	Names []*string
}

type DescribeImageBuildersOutput struct {

	// The pagination token to use to retrieve the next page of results for this
	// operation. If there are no more pages, this value is null.
	NextToken *string

	// Information about the image builders.
	ImageBuilders []*types.ImageBuilder

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeImageBuildersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeImageBuilders{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeImageBuilders{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeImageBuilders(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appstream",
		OperationName: "DescribeImageBuilders",
	}
}
