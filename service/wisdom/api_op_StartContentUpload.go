// Code generated by smithy-go-codegen DO NOT EDIT.

package wisdom

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Get a URL to upload content to a knowledge base. To upload content, first make
// a PUT request to the returned URL with your file, making sure to include the
// required headers. Then use [CreateContent]to finalize the content creation process or [UpdateContent] to
// modify an existing resource. You can only upload content to a knowledge base of
// type CUSTOM.
//
// [CreateContent]: https://docs.aws.amazon.com/wisdom/latest/APIReference/API_CreateContent.html
// [UpdateContent]: https://docs.aws.amazon.com/wisdom/latest/APIReference/API_UpdateContent.html
func (c *Client) StartContentUpload(ctx context.Context, params *StartContentUploadInput, optFns ...func(*Options)) (*StartContentUploadOutput, error) {
	if params == nil {
		params = &StartContentUploadInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartContentUpload", params, optFns, c.addOperationStartContentUploadMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartContentUploadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartContentUploadInput struct {

	// The type of content to upload.
	//
	// This member is required.
	ContentType *string

	// The identifier of the knowledge base. This should not be a QUICK_RESPONSES type
	// knowledge base if you're storing Wisdom Content resource to it. Can be either
	// the ID or the ARN. URLs cannot contain the ARN.
	//
	// This member is required.
	KnowledgeBaseId *string

	// The expected expiration time of the generated presigned URL, specified in
	// minutes.
	PresignedUrlTimeToLive *int32

	noSmithyDocumentSerde
}

type StartContentUploadOutput struct {

	// The headers to include in the upload.
	//
	// This member is required.
	HeadersToInclude map[string]string

	// The identifier of the upload.
	//
	// This member is required.
	UploadId *string

	// The URL of the upload.
	//
	// This member is required.
	Url *string

	// The expiration time of the URL as an epoch timestamp.
	//
	// This member is required.
	UrlExpiry *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartContentUploadMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartContentUpload{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartContentUpload{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartContentUpload"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpStartContentUploadValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartContentUpload(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opStartContentUpload(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartContentUpload",
	}
}
