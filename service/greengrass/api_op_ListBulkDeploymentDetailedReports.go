// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a paginated list of the deployments that have been started in a bulk
// deployment operation, and their current deployment status.
func (c *Client) ListBulkDeploymentDetailedReports(ctx context.Context, params *ListBulkDeploymentDetailedReportsInput, optFns ...func(*Options)) (*ListBulkDeploymentDetailedReportsOutput, error) {
	stack := middleware.NewStack("ListBulkDeploymentDetailedReports", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListBulkDeploymentDetailedReportsMiddlewares(stack)
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
	addOpListBulkDeploymentDetailedReportsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListBulkDeploymentDetailedReports(options.Region), middleware.Before)
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
			OperationName: "ListBulkDeploymentDetailedReports",
			Err:           err,
		}
	}
	out := result.(*ListBulkDeploymentDetailedReportsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBulkDeploymentDetailedReportsInput struct {

	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string

	// The maximum number of results to be returned per request.
	MaxResults *string

	// The ID of the bulk deployment.
	//
	// This member is required.
	BulkDeploymentId *string
}

type ListBulkDeploymentDetailedReportsOutput struct {

	// A list of the individual group deployments in the bulk deployment operation.
	Deployments []*types.BulkDeploymentResult

	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListBulkDeploymentDetailedReportsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListBulkDeploymentDetailedReports{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListBulkDeploymentDetailedReports{}, middleware.After)
}

func newServiceMetadataMiddleware_opListBulkDeploymentDetailedReports(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "ListBulkDeploymentDetailedReports",
	}
}
