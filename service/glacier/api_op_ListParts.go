// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation lists the parts of an archive that have been uploaded in a
// specific multipart upload. You can make this request at any time during an
// in-progress multipart upload before you complete the upload (see
// CompleteMultipartUpload (). List Parts returns an error for completed uploads.
// The list returned in the List Parts response is sorted by part range.  <p>The
// List Parts operation supports pagination. By default, this operation returns up
// to 50 uploaded parts in the response. You should always check the response for a
// <code>marker</code> at which to continue the list; if there are no more items
// the <code>marker</code> is <code>null</code>. To return a list of parts that
// begins at a specific part, set the <code>marker</code> request parameter to the
// value you obtained from a previous List Parts request. You can also limit the
// number of parts returned in the response by specifying the <code>limit</code>
// parameter in the request. </p> <p>An AWS account has full permission to perform
// all operations (actions). However, AWS Identity and Access Management (IAM)
// users don't have any permissions by default. You must grant them explicit
// permission to perform specific actions. For more information, see <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html">Access
// Control Using AWS Identity and Access Management (IAM)</a>.</p> <p>For
// conceptual information and the underlying REST API, see <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/working-with-archives.html">Working
// with Archives in Amazon S3 Glacier</a> and <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/api-multipart-list-parts.html">List
// Parts</a> in the <i>Amazon Glacier Developer Guide</i>.</p>
func (c *Client) ListParts(ctx context.Context, params *ListPartsInput, optFns ...func(*Options)) (*ListPartsOutput, error) {
	stack := middleware.NewStack("ListParts", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListPartsMiddlewares(stack)
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
	addOpListPartsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListParts(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	glaciercust.AddTreeHashMiddleware(stack)
	glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion)
	glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID)

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
			OperationName: "ListParts",
			Err:           err,
		}
	}
	out := result.(*ListPartsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides options for retrieving a list of parts of an archive that have been
// uploaded in a specific multipart upload.
type ListPartsInput struct {

	// An opaque string used for pagination. This value specifies the part at which the
	// listing of parts should begin. Get the marker value from the response of a
	// previous List Parts response. You need only include the marker if you are
	// continuing the pagination of results started in a previous List Parts request.
	Marker *string

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen), in
	// which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// The upload ID of the multipart upload.
	//
	// This member is required.
	UploadId *string

	// The name of the vault.
	//
	// This member is required.
	VaultName *string

	// The maximum number of parts to be returned. The default limit is 50. The number
	// of parts returned might be fewer than the specified limit, but the number of
	// returned parts never exceeds the limit.
	Limit *string
}

// Contains the Amazon S3 Glacier response to your request.
type ListPartsOutput struct {

	// The Amazon Resource Name (ARN) of the vault to which the multipart upload was
	// initiated.
	VaultARN *string

	// The description of the archive that was specified in the Initiate Multipart
	// Upload request.
	ArchiveDescription *string

	// The ID of the upload to which the parts are associated.
	MultipartUploadId *string

	// An opaque string that represents where to continue pagination of the results.
	// You use the marker in a new List Parts request to obtain more jobs in the list.
	// If there are no more parts, this value is null.
	Marker *string

	// The UTC time at which the multipart upload was initiated.
	CreationDate *string

	// The part size in bytes. This is the same value that you specified in the
	// Initiate Multipart Upload request.
	PartSizeInBytes *int64

	// A list of the part sizes of the multipart upload. Each object in the array
	// contains a RangeBytes and sha256-tree-hash name/value pair.
	Parts []*types.PartListElement

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListPartsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListParts{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListParts{}, middleware.After)
}

func newServiceMetadataMiddleware_opListParts(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "ListParts",
	}
}
