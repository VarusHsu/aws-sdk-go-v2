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

// Returns the tag-set of an object. You send the GET request against the tagging
// subresource associated with the object.  <p>To use this operation, you must have
// permission to perform the <code>s3:GetObjectTagging</code> action. By default,
// the GET operation returns information about current version of an object. For a
// versioned bucket, you can have multiple versions of an object in your bucket. To
// retrieve tags of any other version, use the versionId query parameter. You also
// need permission for the <code>s3:GetObjectVersionTagging</code> action.</p> <p>
// By default, the bucket owner has this permission and can grant this permission
// to others.</p> <p> For information about the Amazon S3 object tagging feature,
// see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/object-tagging.html">Object
// Tagging</a>.</p> <p>The following operation is related to
// <code>GetObjectTagging</code>:</p> <ul> <li> <p> <a>PutObjectTagging</a> </p>
// </li> </ul>
func (c *Client) GetObjectTagging(ctx context.Context, params *GetObjectTaggingInput, optFns ...func(*Options)) (*GetObjectTaggingOutput, error) {
	stack := middleware.NewStack("GetObjectTagging", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetObjectTaggingMiddlewares(stack)
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
	addOpGetObjectTaggingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetObjectTagging(options.Region), middleware.Before)
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
			OperationName: "GetObjectTagging",
			Err:           err,
		}
	}
	out := result.(*GetObjectTaggingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetObjectTaggingInput struct {

	// The versionId of the object for which to get the tagging information.
	VersionId *string

	// The bucket name containing the object for which to get the tagging information.
	// When using this API with an access point, you must direct requests to the access
	// point hostname. The access point hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// operation using an access point through the AWS SDKs, you provide the access
	// point ARN in place of the bucket name. For more information about access point
	// ARNs, see Using Access Points
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) in
	// the Amazon Simple Storage Service Developer Guide.
	//
	// This member is required.
	Bucket *string

	// Object key for which to get the tagging information.
	//
	// This member is required.
	Key *string
}

type GetObjectTaggingOutput struct {

	// Contains the tag set.
	//
	// This member is required.
	TagSet []*types.Tag

	// The versionId of the object for which you got the tagging information.
	VersionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetObjectTaggingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetObjectTagging{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetObjectTagging{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetObjectTagging(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetObjectTagging",
	}
}
