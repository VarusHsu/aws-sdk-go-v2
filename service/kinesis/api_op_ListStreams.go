// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists your Kinesis data streams. The number of streams may be too large to
// return from a single call to ListStreams . You can limit the number of returned
// streams using the Limit parameter. If you do not specify a value for the Limit
// parameter, Kinesis Data Streams uses the default limit, which is currently 100.
// You can detect if there are more streams available to list by using the
// HasMoreStreams flag from the returned output. If there are more streams
// available, you can request more streams by using the name of the last stream
// returned by the ListStreams request in the ExclusiveStartStreamName parameter
// in a subsequent request to ListStreams . The group of stream names returned by
// the subsequent request is then added to the list. You can continue this process
// until all the stream names have been collected in the list. ListStreams has a
// limit of five transactions per second per account.
func (c *Client) ListStreams(ctx context.Context, params *ListStreamsInput, optFns ...func(*Options)) (*ListStreamsOutput, error) {
	if params == nil {
		params = &ListStreamsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStreams", params, optFns, c.addOperationListStreamsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStreamsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for ListStreams .
type ListStreamsInput struct {

	// The name of the stream to start the list with.
	ExclusiveStartStreamName *string

	// The maximum number of streams to list. The default value is 100. If you specify
	// a value greater than 100, at most 100 results are returned.
	Limit *int32

	//
	NextToken *string

	noSmithyDocumentSerde
}

// Represents the output for ListStreams .
type ListStreamsOutput struct {

	// If set to true , there are more streams available to list.
	//
	// This member is required.
	HasMoreStreams *bool

	// The names of the streams that are associated with the Amazon Web Services
	// account making the ListStreams request.
	//
	// This member is required.
	StreamNames []string

	//
	NextToken *string

	//
	StreamSummaries []types.StreamSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStreamsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListStreams{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListStreams{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListStreams"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStreams(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

// ListStreamsAPIClient is a client that implements the ListStreams operation.
type ListStreamsAPIClient interface {
	ListStreams(context.Context, *ListStreamsInput, ...func(*Options)) (*ListStreamsOutput, error)
}

var _ ListStreamsAPIClient = (*Client)(nil)

// ListStreamsPaginatorOptions is the paginator options for ListStreams
type ListStreamsPaginatorOptions struct {
	// The maximum number of streams to list. The default value is 100. If you specify
	// a value greater than 100, at most 100 results are returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStreamsPaginator is a paginator for ListStreams
type ListStreamsPaginator struct {
	options   ListStreamsPaginatorOptions
	client    ListStreamsAPIClient
	params    *ListStreamsInput
	nextToken *string
	firstPage bool
}

// NewListStreamsPaginator returns a new ListStreamsPaginator
func NewListStreamsPaginator(client ListStreamsAPIClient, params *ListStreamsInput, optFns ...func(*ListStreamsPaginatorOptions)) *ListStreamsPaginator {
	if params == nil {
		params = &ListStreamsInput{}
	}

	options := ListStreamsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStreamsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStreamsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStreams page.
func (p *ListStreamsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStreamsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.ListStreams(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListStreams(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListStreams",
	}
}
