// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackage

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediapackage/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a collection of HarvestJob records.
func (c *Client) ListHarvestJobs(ctx context.Context, params *ListHarvestJobsInput, optFns ...func(*Options)) (*ListHarvestJobsOutput, error) {
	stack := middleware.NewStack("ListHarvestJobs", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListHarvestJobsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListHarvestJobs(options.Region), middleware.Before)
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
			OperationName: "ListHarvestJobs",
			Err:           err,
		}
	}
	out := result.(*ListHarvestJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListHarvestJobsInput struct {

	// The upper bound on the number of records to return.
	MaxResults *int32

	// When specified, the request will return only HarvestJobs in the given status.
	IncludeStatus *string

	// A token used to resume pagination from the end of a previous request.
	NextToken *string

	// When specified, the request will return only HarvestJobs associated with the
	// given Channel ID.
	IncludeChannelId *string
}

type ListHarvestJobsOutput struct {

	// A list of HarvestJob records.
	HarvestJobs []*types.HarvestJob

	// A token that can be used to resume pagination from the end of the collection.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListHarvestJobsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListHarvestJobs{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListHarvestJobs{}, middleware.After)
}

func newServiceMetadataMiddleware_opListHarvestJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediapackage",
		OperationName: "ListHarvestJobs",
	}
}
