// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of inventory configurations for the bucket. You can have up to
// 1,000 analytics configurations per bucket.  <p>This operation supports list
// pagination and does not return more than 100 configurations at a time. Always
// check the <code>IsTruncated</code> element in the response. If there are no more
// configurations to list, <code>IsTruncated</code> is set to false. If there are
// more configurations to list, <code>IsTruncated</code> is set to true, and there
// is a value in <code>NextContinuationToken</code>. You use the
// <code>NextContinuationToken</code> value to continue the pagination of the list
// by passing the value in continuation-token in the request to <code>GET</code>
// the next page.</p> <p> To use this operation, you must have permissions to
// perform the <code>s3:GetInventoryConfiguration</code> action. The bucket owner
// has this permission by default. The bucket owner can grant this permission to
// others. For more information about permissions, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources">Permissions
// Related to Bucket Subresource Operations</a> and <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html">Managing
// Access Permissions to Your Amazon S3 Resources</a>.</p> <p>For information about
// the Amazon S3 inventory feature, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html">Amazon
// S3 Inventory</a> </p> <p>The following operations are related to
// <code>ListBucketInventoryConfigurations</code>:</p> <ul> <li> <p>
// <a>GetBucketInventoryConfiguration</a> </p> </li> <li> <p>
// <a>DeleteBucketInventoryConfiguration</a> </p> </li> <li> <p>
// <a>PutBucketInventoryConfiguration</a> </p> </li> </ul>
func (c *Client) ListBucketInventoryConfigurations(ctx context.Context, params *ListBucketInventoryConfigurationsInput, optFns ...func(*Options)) (*ListBucketInventoryConfigurationsOutput, error) {
	stack := middleware.NewStack("ListBucketInventoryConfigurations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpListBucketInventoryConfigurationsMiddlewares(stack)
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
	addOpListBucketInventoryConfigurationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListBucketInventoryConfigurations(options.Region), middleware.Before)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	addMetadataRetrieverMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)

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
			OperationName: "ListBucketInventoryConfigurations",
			Err:           err,
		}
	}
	out := result.(*ListBucketInventoryConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBucketInventoryConfigurationsInput struct {

	// The marker used to continue an inventory configuration listing that has been
	// truncated. Use the NextContinuationToken from a previously truncated list
	// response to continue the listing. The continuation token is an opaque value that
	// Amazon S3 understands.
	ContinuationToken *string

	// The name of the bucket containing the inventory configurations to retrieve.
	//
	// This member is required.
	Bucket *string
}

type ListBucketInventoryConfigurationsOutput struct {

	// The list of inventory configurations for a bucket.
	InventoryConfigurationList []*types.InventoryConfiguration

	// The marker used to continue this inventory configuration listing. Use the
	// NextContinuationToken from this response to continue the listing in a subsequent
	// request. The continuation token is an opaque value that Amazon S3 understands.
	NextContinuationToken *string

	// If sent in the request, the marker that is used as a starting point for this
	// inventory configuration list response.
	ContinuationToken *string

	// Tells whether the returned list of inventory configurations is complete. A value
	// of true indicates that the list is not complete and the NextContinuationToken is
	// provided for a subsequent request.
	IsTruncated *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpListBucketInventoryConfigurationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpListBucketInventoryConfigurations{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpListBucketInventoryConfigurations{}, middleware.After)
}

func newServiceMetadataMiddleware_opListBucketInventoryConfigurations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "ListBucketInventoryConfigurations",
	}
}
