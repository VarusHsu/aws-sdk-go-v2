// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation aborts a multipart upload identified by the upload ID.  <p>After
// the Abort Multipart Upload request succeeds, you cannot upload any more parts to
// the multipart upload or complete the multipart upload. Aborting a completed
// upload fails. However, aborting an already-aborted upload will succeed, for a
// short time. For more information about uploading a part and completing a
// multipart upload, see <a>UploadMultipartPart</a> and
// <a>CompleteMultipartUpload</a>.</p> <p>This operation is idempotent.</p> <p>An
// AWS account has full permission to perform all operations (actions). However,
// AWS Identity and Access Management (IAM) users don't have any permissions by
// default. You must grant them explicit permission to perform specific actions.
// For more information, see <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html">Access
// Control Using AWS Identity and Access Management (IAM)</a>.</p> <p> For
// conceptual information and underlying REST API, see <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/working-with-archives.html">Working
// with Archives in Amazon S3 Glacier</a> and <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/api-multipart-abort-upload.html">Abort
// Multipart Upload</a> in the <i>Amazon Glacier Developer Guide</i>. </p>
func (c *Client) AbortMultipartUpload(ctx context.Context, params *AbortMultipartUploadInput, optFns ...func(*Options)) (*AbortMultipartUploadOutput, error) {
	stack := middleware.NewStack("AbortMultipartUpload", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpAbortMultipartUploadMiddlewares(stack)
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
	addOpAbortMultipartUploadValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAbortMultipartUpload(options.Region), middleware.Before)
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
			OperationName: "AbortMultipartUpload",
			Err:           err,
		}
	}
	out := result.(*AbortMultipartUploadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides options to abort a multipart upload identified by the upload ID.
// <p>For information about the underlying REST API, see <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/api-multipart-abort-upload.html">Abort
// Multipart Upload</a>. For conceptual information, see <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/working-with-archives.html">Working
// with Archives in Amazon S3 Glacier</a>.</p>
type AbortMultipartUploadInput struct {

	// The name of the vault.
	//
	// This member is required.
	VaultName *string

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen), in
	// which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// The upload ID of the multipart upload to delete.
	//
	// This member is required.
	UploadId *string
}

type AbortMultipartUploadOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpAbortMultipartUploadMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpAbortMultipartUpload{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpAbortMultipartUpload{}, middleware.After)
}

func newServiceMetadataMiddleware_opAbortMultipartUpload(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "AbortMultipartUpload",
	}
}
