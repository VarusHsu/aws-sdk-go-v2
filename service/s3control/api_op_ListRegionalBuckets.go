// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of all Outposts buckets in an Outposts that are owned by the
// authenticated sender of the request. For more information, see Using Amazon S3
// on Outposts (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html)
// in the Amazon Simple Storage Service Developer Guide. For an example of the
// request syntax for Amazon S3 on Outposts that uses the S3 on Outposts endpoint
// hostname prefix and outpost-id in your API request, see the  Example
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_ListRegionalBuckets.html#API_control_ListRegionalBuckets_Examples)
// section below.
func (c *Client) ListRegionalBuckets(ctx context.Context, params *ListRegionalBucketsInput, optFns ...func(*Options)) (*ListRegionalBucketsOutput, error) {
	stack := middleware.NewStack("ListRegionalBuckets", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpListRegionalBucketsMiddlewares(stack)
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
	addOpListRegionalBucketsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListRegionalBuckets(options.Region), middleware.Before)
	addResponseErrorMiddleware(stack)
	addMetadataRetrieverMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)

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
			OperationName: "ListRegionalBuckets",
			Err:           err,
		}
	}
	out := result.(*ListRegionalBucketsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRegionalBucketsInput struct {

	// The AWS account ID of the Outposts bucket.
	//
	// This member is required.
	AccountId *string

	//
	MaxResults *int32

	// The ID of the AWS Outposts. This is required by Amazon S3 on Outposts buckets.
	OutpostId *string

	//
	NextToken *string
}

type ListRegionalBucketsOutput struct {

	// NextToken is sent when isTruncated is true, which means there are more buckets
	// that can be listed. The next list requests to Amazon S3 can be continued with
	// this NextToken. NextToken is obfuscated and is not a real key.
	NextToken *string

	//
	RegionalBucketList []*types.RegionalBucket

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpListRegionalBucketsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpListRegionalBuckets{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpListRegionalBuckets{}, middleware.After)
}

func newServiceMetadataMiddleware_opListRegionalBuckets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "ListRegionalBuckets",
	}
}
