// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about reactions to a specified comment ID. Reactions from
// users who have been deleted will not be included in the count.
func (c *Client) GetCommentReactions(ctx context.Context, params *GetCommentReactionsInput, optFns ...func(*Options)) (*GetCommentReactionsOutput, error) {
	if params == nil {
		params = &GetCommentReactionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCommentReactions", params, optFns, c.addOperationGetCommentReactionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCommentReactionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCommentReactionsInput struct {

	// The ID of the comment for which you want to get reactions information.
	//
	// This member is required.
	CommentId *string

	// A non-zero, non-negative integer used to limit the number of returned results.
	// The default is the same as the allowed maximum, 1,000.
	MaxResults *int32

	// An enumeration token that, when provided in a request, returns the next batch
	// of the results.
	NextToken *string

	// Optional. The Amazon Resource Name (ARN) of the user or identity for which you
	// want to get reaction information.
	ReactionUserArn *string

	noSmithyDocumentSerde
}

type GetCommentReactionsOutput struct {

	// An array of reactions to the specified comment.
	//
	// This member is required.
	ReactionsForComment []types.ReactionForComment

	// An enumeration token that can be used in a request to return the next batch of
	// the results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCommentReactionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCommentReactions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCommentReactions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetCommentReactions"); err != nil {
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
	if err = addOpGetCommentReactionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCommentReactions(options.Region), middleware.Before); err != nil {
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

// GetCommentReactionsPaginatorOptions is the paginator options for
// GetCommentReactions
type GetCommentReactionsPaginatorOptions struct {
	// A non-zero, non-negative integer used to limit the number of returned results.
	// The default is the same as the allowed maximum, 1,000.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetCommentReactionsPaginator is a paginator for GetCommentReactions
type GetCommentReactionsPaginator struct {
	options   GetCommentReactionsPaginatorOptions
	client    GetCommentReactionsAPIClient
	params    *GetCommentReactionsInput
	nextToken *string
	firstPage bool
}

// NewGetCommentReactionsPaginator returns a new GetCommentReactionsPaginator
func NewGetCommentReactionsPaginator(client GetCommentReactionsAPIClient, params *GetCommentReactionsInput, optFns ...func(*GetCommentReactionsPaginatorOptions)) *GetCommentReactionsPaginator {
	if params == nil {
		params = &GetCommentReactionsInput{}
	}

	options := GetCommentReactionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetCommentReactionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetCommentReactionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetCommentReactions page.
func (p *GetCommentReactionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetCommentReactionsOutput, error) {
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
	result, err := p.client.GetCommentReactions(ctx, &params, optFns...)
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

// GetCommentReactionsAPIClient is a client that implements the
// GetCommentReactions operation.
type GetCommentReactionsAPIClient interface {
	GetCommentReactions(context.Context, *GetCommentReactionsInput, ...func(*Options)) (*GetCommentReactionsOutput, error)
}

var _ GetCommentReactionsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetCommentReactions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetCommentReactions",
	}
}
