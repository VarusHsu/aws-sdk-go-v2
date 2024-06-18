// Code generated by smithy-go-codegen DO NOT EDIT.

package connectparticipant

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides a pre-signed URL for download of a completed attachment. This is an
// asynchronous API for use with active contacts.
//
// ConnectionToken is used for invoking this API instead of ParticipantToken .
//
// The Amazon Connect Participant Service APIs do not use [Signature Version 4 authentication].
//
// [Signature Version 4 authentication]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
func (c *Client) GetAttachment(ctx context.Context, params *GetAttachmentInput, optFns ...func(*Options)) (*GetAttachmentOutput, error) {
	if params == nil {
		params = &GetAttachmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAttachment", params, optFns, c.addOperationGetAttachmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAttachmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAttachmentInput struct {

	// A unique identifier for the attachment.
	//
	// This member is required.
	AttachmentId *string

	// The authentication token associated with the participant's connection.
	//
	// This member is required.
	ConnectionToken *string

	noSmithyDocumentSerde
}

type GetAttachmentOutput struct {

	// This is the pre-signed URL that can be used for uploading the file to Amazon S3
	// when used in response to [StartAttachmentUpload].
	//
	// [StartAttachmentUpload]: https://docs.aws.amazon.com/connect-participant/latest/APIReference/API_StartAttachmentUpload.html
	Url *string

	// The expiration time of the URL in ISO timestamp. It's specified in ISO 8601
	// format: yyyy-MM-ddThh:mm:ss.SSSZ. For example, 2019-11-08T02:41:28.172Z.
	UrlExpiry *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAttachmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAttachment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAttachment{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAttachment"); err != nil {
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
	if err = addOpGetAttachmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAttachment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAttachment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAttachment",
	}
}
