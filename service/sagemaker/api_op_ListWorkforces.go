// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Use this operation to list all private and vendor workforces in an AWS Region.
// Note that you can only have one private workforce per AWS Region.
func (c *Client) ListWorkforces(ctx context.Context, params *ListWorkforcesInput, optFns ...func(*Options)) (*ListWorkforcesOutput, error) {
	stack := middleware.NewStack("ListWorkforces", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListWorkforcesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListWorkforces(options.Region), middleware.Before)
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
			OperationName: "ListWorkforces",
			Err:           err,
		}
	}
	out := result.(*ListWorkforcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorkforcesInput struct {

	// A token to resume pagination.
	NextToken *string

	// Sort workforces using the workforce name or creation date.
	SortBy types.ListWorkforcesSortByOptions

	// Sort workforces in ascending or descending order.
	SortOrder types.SortOrder

	// The maximum number of workforces returned in the response.
	MaxResults *int32

	// A filter you can use to search for workforces using part of the workforce name.
	NameContains *string
}

type ListWorkforcesOutput struct {

	// A token to resume pagination.
	NextToken *string

	// A list containing information about your workforce.
	//
	// This member is required.
	Workforces []*types.Workforce

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListWorkforcesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListWorkforces{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListWorkforces{}, middleware.After)
}

func newServiceMetadataMiddleware_opListWorkforces(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListWorkforces",
	}
}
