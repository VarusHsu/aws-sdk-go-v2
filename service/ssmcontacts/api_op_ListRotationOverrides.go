// Code generated by smithy-go-codegen DO NOT EDIT.

package ssmcontacts

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssmcontacts/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves a list of overrides currently specified for an on-call rotation.
func (c *Client) ListRotationOverrides(ctx context.Context, params *ListRotationOverridesInput, optFns ...func(*Options)) (*ListRotationOverridesOutput, error) {
	if params == nil {
		params = &ListRotationOverridesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRotationOverrides", params, optFns, c.addOperationListRotationOverridesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRotationOverridesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRotationOverridesInput struct {

	// The date and time for the end of a time range for listing overrides.
	//
	// This member is required.
	EndTime *time.Time

	// The Amazon Resource Name (ARN) of the rotation to retrieve information about.
	//
	// This member is required.
	RotationId *string

	// The date and time for the beginning of a time range for listing overrides.
	//
	// This member is required.
	StartTime *time.Time

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRotationOverridesOutput struct {

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// A list of rotation overrides in the specified time range.
	RotationOverrides []types.RotationOverride

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRotationOverridesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListRotationOverrides{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListRotationOverrides{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRotationOverrides"); err != nil {
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
	if err = addOpListRotationOverridesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRotationOverrides(options.Region), middleware.Before); err != nil {
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

// ListRotationOverridesPaginatorOptions is the paginator options for
// ListRotationOverrides
type ListRotationOverridesPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRotationOverridesPaginator is a paginator for ListRotationOverrides
type ListRotationOverridesPaginator struct {
	options   ListRotationOverridesPaginatorOptions
	client    ListRotationOverridesAPIClient
	params    *ListRotationOverridesInput
	nextToken *string
	firstPage bool
}

// NewListRotationOverridesPaginator returns a new ListRotationOverridesPaginator
func NewListRotationOverridesPaginator(client ListRotationOverridesAPIClient, params *ListRotationOverridesInput, optFns ...func(*ListRotationOverridesPaginatorOptions)) *ListRotationOverridesPaginator {
	if params == nil {
		params = &ListRotationOverridesInput{}
	}

	options := ListRotationOverridesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRotationOverridesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRotationOverridesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRotationOverrides page.
func (p *ListRotationOverridesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRotationOverridesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListRotationOverrides(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ListRotationOverridesAPIClient is a client that implements the
// ListRotationOverrides operation.
type ListRotationOverridesAPIClient interface {
	ListRotationOverrides(context.Context, *ListRotationOverridesInput, ...func(*Options)) (*ListRotationOverridesOutput, error)
}

var _ ListRotationOverridesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListRotationOverrides(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRotationOverrides",
	}
}
