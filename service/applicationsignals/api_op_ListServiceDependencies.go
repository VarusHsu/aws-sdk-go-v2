// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationsignals

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/applicationsignals/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a list of service dependencies of the service that you specify. A
// dependency is an infrastructure component that an operation of this service
// connects with. Dependencies can include Amazon Web Services services, Amazon Web
// Services resources, and third-party services.
func (c *Client) ListServiceDependencies(ctx context.Context, params *ListServiceDependenciesInput, optFns ...func(*Options)) (*ListServiceDependenciesOutput, error) {
	if params == nil {
		params = &ListServiceDependenciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListServiceDependencies", params, optFns, c.addOperationListServiceDependenciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListServiceDependenciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListServiceDependenciesInput struct {

	// The end of the time period to retrieve information about. When used in a raw
	// HTTP Query API, it is formatted as be epoch time in seconds. For example:
	// 1698778057
	//
	// This member is required.
	EndTime *time.Time

	// Use this field to specify which service you want to retrieve information for.
	// You must specify at least the Type , Name , and Environment attributes.
	//
	// This is a string-to-string map. It can include the following fields.
	//
	//   - Type designates the type of object this is.
	//
	//   - ResourceType specifies the type of the resource. This field is used only
	//   when the value of the Type field is Resource or AWS::Resource .
	//
	//   - Name specifies the name of the object. This is used only if the value of the
	//   Type field is Service , RemoteService , or AWS::Service .
	//
	//   - Identifier identifies the resource objects of this resource. This is used
	//   only if the value of the Type field is Resource or AWS::Resource .
	//
	//   - Environment specifies the location where this object is hosted, or what it
	//   belongs to.
	//
	// This member is required.
	KeyAttributes map[string]string

	// The start of the time period to retrieve information about. When used in a raw
	// HTTP Query API, it is formatted as be epoch time in seconds. For example:
	// 1698778057
	//
	// This member is required.
	StartTime *time.Time

	// The maximum number of results to return in one operation. If you omit this
	// parameter, the default of 50 is used.
	MaxResults *int32

	// Include this value, if it was returned by the previous operation, to get the
	// next set of service dependencies.
	NextToken *string

	noSmithyDocumentSerde
}

type ListServiceDependenciesOutput struct {

	// The end of the time period that the returned information applies to. When used
	// in a raw HTTP Query API, it is formatted as be epoch time in seconds. For
	// example: 1698778057
	//
	// This member is required.
	EndTime *time.Time

	// An array, where each object in the array contains information about one of the
	// dependencies of this service.
	//
	// This member is required.
	ServiceDependencies []types.ServiceDependency

	// The start of the time period that the returned information applies to. When
	// used in a raw HTTP Query API, it is formatted as be epoch time in seconds. For
	// example: 1698778057
	//
	// This member is required.
	StartTime *time.Time

	// Include this value in your next use of this API to get next set of service
	// dependencies.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListServiceDependenciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListServiceDependencies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListServiceDependencies{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListServiceDependencies"); err != nil {
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
	if err = addOpListServiceDependenciesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListServiceDependencies(options.Region), middleware.Before); err != nil {
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

// ListServiceDependenciesPaginatorOptions is the paginator options for
// ListServiceDependencies
type ListServiceDependenciesPaginatorOptions struct {
	// The maximum number of results to return in one operation. If you omit this
	// parameter, the default of 50 is used.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListServiceDependenciesPaginator is a paginator for ListServiceDependencies
type ListServiceDependenciesPaginator struct {
	options   ListServiceDependenciesPaginatorOptions
	client    ListServiceDependenciesAPIClient
	params    *ListServiceDependenciesInput
	nextToken *string
	firstPage bool
}

// NewListServiceDependenciesPaginator returns a new
// ListServiceDependenciesPaginator
func NewListServiceDependenciesPaginator(client ListServiceDependenciesAPIClient, params *ListServiceDependenciesInput, optFns ...func(*ListServiceDependenciesPaginatorOptions)) *ListServiceDependenciesPaginator {
	if params == nil {
		params = &ListServiceDependenciesInput{}
	}

	options := ListServiceDependenciesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListServiceDependenciesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListServiceDependenciesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListServiceDependencies page.
func (p *ListServiceDependenciesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListServiceDependenciesOutput, error) {
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
	result, err := p.client.ListServiceDependencies(ctx, &params, optFns...)
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

// ListServiceDependenciesAPIClient is a client that implements the
// ListServiceDependencies operation.
type ListServiceDependenciesAPIClient interface {
	ListServiceDependencies(context.Context, *ListServiceDependenciesInput, ...func(*Options)) (*ListServiceDependenciesOutput, error)
}

var _ ListServiceDependenciesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListServiceDependencies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListServiceDependencies",
	}
}
