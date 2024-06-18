// Code generated by smithy-go-codegen DO NOT EDIT.

package mailmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mailmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the full details and current state of a specified email archive.
func (c *Client) GetArchive(ctx context.Context, params *GetArchiveInput, optFns ...func(*Options)) (*GetArchiveOutput, error) {
	if params == nil {
		params = &GetArchiveInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetArchive", params, optFns, c.addOperationGetArchiveMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetArchiveOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to retrieve details of an email archive.
type GetArchiveInput struct {

	// The identifier of the archive to retrieve.
	//
	// This member is required.
	ArchiveId *string

	noSmithyDocumentSerde
}

// The response containing details of the requested archive.
type GetArchiveOutput struct {

	// The Amazon Resource Name (ARN) of the archive.
	//
	// This member is required.
	ArchiveArn *string

	// The unique identifier of the archive.
	//
	// This member is required.
	ArchiveId *string

	// The unique name assigned to the archive.
	//
	// This member is required.
	ArchiveName *string

	// The current state of the archive:
	//
	//   - ACTIVE – The archive is ready and available for use.
	//
	//   - PENDING_DELETION – The archive has been marked for deletion and will be
	//   permanently deleted in 30 days. No further modifications can be made in this
	//   state.
	//
	// This member is required.
	ArchiveState types.ArchiveState

	// The retention period for emails in this archive.
	//
	// This member is required.
	Retention types.ArchiveRetention

	// The timestamp of when the archive was created.
	CreatedTimestamp *time.Time

	// The Amazon Resource Name (ARN) of the KMS key used to encrypt the archive.
	KmsKeyArn *string

	// The timestamp of when the archive was modified.
	LastUpdatedTimestamp *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetArchiveMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetArchive{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetArchive{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetArchive"); err != nil {
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
	if err = addOpGetArchiveValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetArchive(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetArchive(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetArchive",
	}
}
