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

// This implementation of the GET operation uses the accelerate subresource to
// return the Transfer Acceleration state of a bucket, which is either Enabled or
// Suspended. Amazon S3 Transfer Acceleration is a bucket-level feature that
// enables you to perform faster data transfers to and from Amazon S3. To use this
// operation, you must have permission to perform the s3:GetAccelerateConfiguration
// action. The bucket owner has this permission by default. The bucket owner can
// grant this permission to others. For more information about permissions, see
// Permissions Related to Bucket Subresource Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to your Amazon S3 Resources
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html) in the
// Amazon Simple Storage Service Developer Guide. You set the Transfer Acceleration
// state of an existing bucket to Enabled or Suspended by using the
// PutBucketAccelerateConfiguration () operation. A GET accelerate request does not
// return a state value for a bucket that has no transfer acceleration state. A
// bucket has no Transfer Acceleration state if a state has never been set on the
// bucket.  <p>For more information about transfer acceleration, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html">Transfer
// Acceleration</a> in the Amazon Simple Storage Service Developer Guide.</p> <p
// class="title"> <b>Related Resources</b> </p> <ul> <li> <p>
// <a>PutBucketAccelerateConfiguration</a> </p> </li> </ul>
func (c *Client) GetBucketAccelerateConfiguration(ctx context.Context, params *GetBucketAccelerateConfigurationInput, optFns ...func(*Options)) (*GetBucketAccelerateConfigurationOutput, error) {
	stack := middleware.NewStack("GetBucketAccelerateConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetBucketAccelerateConfigurationMiddlewares(stack)
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
	addOpGetBucketAccelerateConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketAccelerateConfiguration(options.Region), middleware.Before)
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
			OperationName: "GetBucketAccelerateConfiguration",
			Err:           err,
		}
	}
	out := result.(*GetBucketAccelerateConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketAccelerateConfigurationInput struct {

	// Name of the bucket for which the accelerate configuration is retrieved.
	//
	// This member is required.
	Bucket *string
}

type GetBucketAccelerateConfigurationOutput struct {

	// The accelerate configuration of the bucket.
	Status types.BucketAccelerateStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetBucketAccelerateConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetBucketAccelerateConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketAccelerateConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetBucketAccelerateConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketAccelerateConfiguration",
	}
}
