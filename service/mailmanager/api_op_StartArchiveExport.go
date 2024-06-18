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

// Initiates an export of emails from the specified archive.
func (c *Client) StartArchiveExport(ctx context.Context, params *StartArchiveExportInput, optFns ...func(*Options)) (*StartArchiveExportOutput, error) {
	if params == nil {
		params = &StartArchiveExportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartArchiveExport", params, optFns, c.addOperationStartArchiveExportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartArchiveExportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to initiate an export of emails from an archive.
type StartArchiveExportInput struct {

	// The identifier of the archive to export emails from.
	//
	// This member is required.
	ArchiveId *string

	// Details on where to deliver the exported email data.
	//
	// This member is required.
	ExportDestinationConfiguration types.ExportDestinationConfiguration

	// The start of the timestamp range to include emails from.
	//
	// This member is required.
	FromTimestamp *time.Time

	// The end of the timestamp range to include emails from.
	//
	// This member is required.
	ToTimestamp *time.Time

	// Criteria to filter which emails are included in the export.
	Filters *types.ArchiveFilters

	// The maximum number of email items to include in the export.
	MaxResults *int32

	noSmithyDocumentSerde
}

// The response from initiating an archive export.
type StartArchiveExportOutput struct {

	// The unique identifier for the initiated export job.
	ExportId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartArchiveExportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpStartArchiveExport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpStartArchiveExport{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartArchiveExport"); err != nil {
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
	if err = addOpStartArchiveExportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartArchiveExport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartArchiveExport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartArchiveExport",
	}
}
