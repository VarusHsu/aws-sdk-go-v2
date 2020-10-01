// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of deployment jobs for a fleet. You can optionally provide
// filters to retrieve specific deployment jobs.
func (c *Client) ListDeploymentJobs(ctx context.Context, params *ListDeploymentJobsInput, optFns ...func(*Options)) (*ListDeploymentJobsOutput, error) {
	stack := middleware.NewStack("ListDeploymentJobs", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListDeploymentJobsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDeploymentJobs(options.Region), middleware.Before)
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
			OperationName: "ListDeploymentJobs",
			Err:           err,
		}
	}
	out := result.(*ListDeploymentJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDeploymentJobsInput struct {

	// Optional filters to limit results. The filter names status and fleetName are
	// supported. When filtering, you must use the complete value of the filtered item.
	// You can use up to three filters, but they must be for the same named item. For
	// example, if you are looking for items with the status InProgress or the status
	// Pending.
	Filters []*types.Filter

	// When this parameter is used, ListDeploymentJobs only returns maxResults results
	// in a single page along with a nextToken response element. The remaining results
	// of the initial request can be seen by sending another ListDeploymentJobs request
	// with the returned nextToken value. This value can be between 1 and 200. If this
	// parameter is not used, then ListDeploymentJobs returns up to 200 results and a
	// nextToken value if applicable.
	MaxResults *int32

	// The nextToken value returned from a previous paginated ListDeploymentJobs
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value.
	NextToken *string
}

type ListDeploymentJobsOutput struct {

	// A list of deployment jobs that meet the criteria of the request.
	DeploymentJobs []*types.DeploymentJob

	// The nextToken value to include in a future ListDeploymentJobs request. When the
	// results of a ListDeploymentJobs request exceed maxResults, this value can be
	// used to retrieve the next page of results. This value is null when there are no
	// more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListDeploymentJobsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListDeploymentJobs{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListDeploymentJobs{}, middleware.After)
}

func newServiceMetadataMiddleware_opListDeploymentJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "ListDeploymentJobs",
	}
}
