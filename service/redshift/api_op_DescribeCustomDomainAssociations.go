// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Contains information about custom domain associations for a cluster.
func (c *Client) DescribeCustomDomainAssociations(ctx context.Context, params *DescribeCustomDomainAssociationsInput, optFns ...func(*Options)) (*DescribeCustomDomainAssociationsOutput, error) {
	if params == nil {
		params = &DescribeCustomDomainAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCustomDomainAssociations", params, optFns, c.addOperationDescribeCustomDomainAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCustomDomainAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCustomDomainAssociationsInput struct {

	// The certificate Amazon Resource Name (ARN) for the custom domain association.
	CustomDomainCertificateArn *string

	// The custom domain name for the custom domain association.
	CustomDomainName *string

	// The marker for the custom domain association.
	Marker *string

	// The maximum records setting for the associated custom domain.
	MaxRecords *int32

	noSmithyDocumentSerde
}

type DescribeCustomDomainAssociationsOutput struct {

	// The associations for the custom domain.
	Associations []types.Association

	// The marker for the custom domain association.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeCustomDomainAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeCustomDomainAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeCustomDomainAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeCustomDomainAssociations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCustomDomainAssociations(options.Region), middleware.Before); err != nil {
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

// DescribeCustomDomainAssociationsPaginatorOptions is the paginator options for
// DescribeCustomDomainAssociations
type DescribeCustomDomainAssociationsPaginatorOptions struct {
	// The maximum records setting for the associated custom domain.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeCustomDomainAssociationsPaginator is a paginator for
// DescribeCustomDomainAssociations
type DescribeCustomDomainAssociationsPaginator struct {
	options   DescribeCustomDomainAssociationsPaginatorOptions
	client    DescribeCustomDomainAssociationsAPIClient
	params    *DescribeCustomDomainAssociationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeCustomDomainAssociationsPaginator returns a new
// DescribeCustomDomainAssociationsPaginator
func NewDescribeCustomDomainAssociationsPaginator(client DescribeCustomDomainAssociationsAPIClient, params *DescribeCustomDomainAssociationsInput, optFns ...func(*DescribeCustomDomainAssociationsPaginatorOptions)) *DescribeCustomDomainAssociationsPaginator {
	if params == nil {
		params = &DescribeCustomDomainAssociationsInput{}
	}

	options := DescribeCustomDomainAssociationsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeCustomDomainAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeCustomDomainAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeCustomDomainAssociations page.
func (p *DescribeCustomDomainAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeCustomDomainAssociationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeCustomDomainAssociations(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// DescribeCustomDomainAssociationsAPIClient is a client that implements the
// DescribeCustomDomainAssociations operation.
type DescribeCustomDomainAssociationsAPIClient interface {
	DescribeCustomDomainAssociations(context.Context, *DescribeCustomDomainAssociationsInput, ...func(*Options)) (*DescribeCustomDomainAssociationsOutput, error)
}

var _ DescribeCustomDomainAssociationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeCustomDomainAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeCustomDomainAssociations",
	}
}
