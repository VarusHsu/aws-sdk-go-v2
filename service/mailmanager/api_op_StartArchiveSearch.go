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

// Initiates a search across emails in the specified archive.
func (c *Client) StartArchiveSearch(ctx context.Context, params *StartArchiveSearchInput, optFns ...func(*Options)) (*StartArchiveSearchOutput, error) {
	if params == nil {
		params = &StartArchiveSearchInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartArchiveSearch", params, optFns, c.addOperationStartArchiveSearchMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartArchiveSearchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to initiate a search across emails in an archive.
type StartArchiveSearchInput struct {

	// The identifier of the archive to search emails in.
	//
	// This member is required.
	ArchiveId *string

	// The start timestamp of the range to search emails from.
	//
	// This member is required.
	FromTimestamp *time.Time

	// The maximum number of search results to return.
	//
	// This member is required.
	MaxResults *int32

	// The end timestamp of the range to search emails from.
	//
	// This member is required.
	ToTimestamp *time.Time

	// Criteria to filter which emails are included in the search results.
	Filters *types.ArchiveFilters

	noSmithyDocumentSerde
}

// The response from initiating an archive search.
type StartArchiveSearchOutput struct {

	// The unique identifier for the initiated search job.
	SearchId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartArchiveSearchMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpStartArchiveSearch{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpStartArchiveSearch{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartArchiveSearch"); err != nil {
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
	if err = addOpStartArchiveSearchValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartArchiveSearch(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartArchiveSearch(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartArchiveSearch",
	}
}
