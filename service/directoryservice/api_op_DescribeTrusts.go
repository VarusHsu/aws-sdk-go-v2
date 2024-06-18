// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Obtains information about the trust relationships for this account.
//
// If no input parameters are provided, such as DirectoryId or TrustIds, this
// request describes all the trust relationships belonging to the account.
func (c *Client) DescribeTrusts(ctx context.Context, params *DescribeTrustsInput, optFns ...func(*Options)) (*DescribeTrustsOutput, error) {
	if params == nil {
		params = &DescribeTrustsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTrusts", params, optFns, c.addOperationDescribeTrustsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTrustsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Describes the trust relationships for a particular Managed Microsoft AD
// directory. If no input parameters are provided, such as directory ID or trust
// ID, this request describes all the trust relationships.
type DescribeTrustsInput struct {

	// The Directory ID of the Amazon Web Services directory that is a part of the
	// requested trust relationship.
	DirectoryId *string

	// The maximum number of objects to return.
	Limit *int32

	// The DescribeTrustsResult.NextToken value from a previous call to DescribeTrusts. Pass null if
	// this is the first call.
	NextToken *string

	// A list of identifiers of the trust relationships for which to obtain the
	// information. If this member is null, all trust relationships that belong to the
	// current account are returned.
	//
	// An empty list results in an InvalidParameterException being thrown.
	TrustIds []string

	noSmithyDocumentSerde
}

// The result of a DescribeTrust request.
type DescribeTrustsOutput struct {

	// If not null, more results are available. Pass this value for the NextToken
	// parameter in a subsequent call to DescribeTruststo retrieve the next set of items.
	NextToken *string

	// The list of Trust objects that were retrieved.
	//
	// It is possible that this list contains less than the number of items specified
	// in the Limit member of the request. This occurs if there are less than the
	// requested number of items left to retrieve, or if the limitations of the
	// operation have been exceeded.
	Trusts []types.Trust

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTrustsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTrusts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTrusts{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeTrusts"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTrusts(options.Region), middleware.Before); err != nil {
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

// DescribeTrustsPaginatorOptions is the paginator options for DescribeTrusts
type DescribeTrustsPaginatorOptions struct {
	// The maximum number of objects to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeTrustsPaginator is a paginator for DescribeTrusts
type DescribeTrustsPaginator struct {
	options   DescribeTrustsPaginatorOptions
	client    DescribeTrustsAPIClient
	params    *DescribeTrustsInput
	nextToken *string
	firstPage bool
}

// NewDescribeTrustsPaginator returns a new DescribeTrustsPaginator
func NewDescribeTrustsPaginator(client DescribeTrustsAPIClient, params *DescribeTrustsInput, optFns ...func(*DescribeTrustsPaginatorOptions)) *DescribeTrustsPaginator {
	if params == nil {
		params = &DescribeTrustsInput{}
	}

	options := DescribeTrustsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeTrustsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeTrustsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeTrusts page.
func (p *DescribeTrustsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeTrustsOutput, error) {
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

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeTrusts(ctx, &params, optFns...)
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

// DescribeTrustsAPIClient is a client that implements the DescribeTrusts
// operation.
type DescribeTrustsAPIClient interface {
	DescribeTrusts(context.Context, *DescribeTrustsInput, ...func(*Options)) (*DescribeTrustsOutput, error)
}

var _ DescribeTrustsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeTrusts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeTrusts",
	}
}
