// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrassv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/greengrassv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a paginated list of deployment jobs that IoT Greengrass sends to
// Greengrass core devices.
func (c *Client) ListEffectiveDeployments(ctx context.Context, params *ListEffectiveDeploymentsInput, optFns ...func(*Options)) (*ListEffectiveDeploymentsOutput, error) {
	if params == nil {
		params = &ListEffectiveDeploymentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEffectiveDeployments", params, optFns, c.addOperationListEffectiveDeploymentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEffectiveDeploymentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEffectiveDeploymentsInput struct {

	// The name of the core device. This is also the name of the IoT thing.
	//
	// This member is required.
	CoreDeviceThingName *string

	// The maximum number of results to be returned per paginated request.
	MaxResults *int32

	// The token to be used for the next set of paginated results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListEffectiveDeploymentsOutput struct {

	// A list that summarizes each deployment on the core device.
	EffectiveDeployments []types.EffectiveDeployment

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEffectiveDeploymentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEffectiveDeployments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEffectiveDeployments{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListEffectiveDeployments"); err != nil {
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
	if err = addOpListEffectiveDeploymentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEffectiveDeployments(options.Region), middleware.Before); err != nil {
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

// ListEffectiveDeploymentsPaginatorOptions is the paginator options for
// ListEffectiveDeployments
type ListEffectiveDeploymentsPaginatorOptions struct {
	// The maximum number of results to be returned per paginated request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEffectiveDeploymentsPaginator is a paginator for ListEffectiveDeployments
type ListEffectiveDeploymentsPaginator struct {
	options   ListEffectiveDeploymentsPaginatorOptions
	client    ListEffectiveDeploymentsAPIClient
	params    *ListEffectiveDeploymentsInput
	nextToken *string
	firstPage bool
}

// NewListEffectiveDeploymentsPaginator returns a new
// ListEffectiveDeploymentsPaginator
func NewListEffectiveDeploymentsPaginator(client ListEffectiveDeploymentsAPIClient, params *ListEffectiveDeploymentsInput, optFns ...func(*ListEffectiveDeploymentsPaginatorOptions)) *ListEffectiveDeploymentsPaginator {
	if params == nil {
		params = &ListEffectiveDeploymentsInput{}
	}

	options := ListEffectiveDeploymentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEffectiveDeploymentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEffectiveDeploymentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEffectiveDeployments page.
func (p *ListEffectiveDeploymentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEffectiveDeploymentsOutput, error) {
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
	result, err := p.client.ListEffectiveDeployments(ctx, &params, optFns...)
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

// ListEffectiveDeploymentsAPIClient is a client that implements the
// ListEffectiveDeployments operation.
type ListEffectiveDeploymentsAPIClient interface {
	ListEffectiveDeployments(context.Context, *ListEffectiveDeploymentsInput, ...func(*Options)) (*ListEffectiveDeploymentsOutput, error)
}

var _ ListEffectiveDeploymentsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListEffectiveDeployments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListEffectiveDeployments",
	}
}
