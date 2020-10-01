// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns metadata about your copy jobs.
func (c *Client) ListCopyJobs(ctx context.Context, params *ListCopyJobsInput, optFns ...func(*Options)) (*ListCopyJobsOutput, error) {
	stack := middleware.NewStack("ListCopyJobs", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListCopyJobsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListCopyJobs(options.Region), middleware.Before)
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
			OperationName: "ListCopyJobs",
			Err:           err,
		}
	}
	out := result.(*ListCopyJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCopyJobsInput struct {

	// The account ID to list the jobs from. Returns only copy jobs associated with the
	// specified account ID.
	ByAccountId *string

	// Returns only copy jobs that match the specified resource Amazon Resource Name
	// (ARN).
	ByResourceArn *string

	// The maximum number of items to be returned.
	MaxResults *int32

	// Returns only copy jobs that are in the specified state.
	ByState types.CopyJobState

	// Returns only copy jobs that were created after the specified date.
	ByCreatedAfter *time.Time

	// Returns only backup jobs for the specified resources:
	//
	//     * DynamoDB for Amazon
	// DynamoDB
	//
	//     * EBS for Amazon Elastic Block Store
	//
	//     * EC2 for Amazon Elastic
	// Compute Cloud
	//
	//     * EFS for Amazon Elastic File System
	//
	//     * RDS for Amazon
	// Relational Database Service
	//
	//     * Storage Gateway for AWS Storage Gateway
	ByResourceType *string

	// The next item following a partial list of returned items. For example, if a
	// request is made to return maxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string

	// Returns only copy jobs that were created before the specified date.
	ByCreatedBefore *time.Time

	// An Amazon Resource Name (ARN) that uniquely identifies a source backup vault to
	// copy from; for example,
	// arn:aws:backup:us-east-1:123456789012:vault:aBackupVault.
	ByDestinationVaultArn *string
}

type ListCopyJobsOutput struct {

	// An array of structures containing metadata about your copy jobs returned in JSON
	// format.
	CopyJobs []*types.CopyJob

	// The next item following a partial list of returned items. For example, if a
	// request is made to return maxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListCopyJobsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListCopyJobs{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListCopyJobs{}, middleware.After)
}

func newServiceMetadataMiddleware_opListCopyJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "ListCopyJobs",
	}
}
