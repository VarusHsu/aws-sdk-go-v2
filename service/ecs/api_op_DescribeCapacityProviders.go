// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes one or more of your capacity providers.
func (c *Client) DescribeCapacityProviders(ctx context.Context, params *DescribeCapacityProvidersInput, optFns ...func(*Options)) (*DescribeCapacityProvidersOutput, error) {
	if params == nil {
		params = &DescribeCapacityProvidersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCapacityProviders", params, optFns, c.addOperationDescribeCapacityProvidersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCapacityProvidersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCapacityProvidersInput struct {

	// The short name or full Amazon Resource Name (ARN) of one or more capacity
	// providers. Up to 100 capacity providers can be described in an action.
	CapacityProviders []string

	// Specifies whether or not you want to see the resource tags for the capacity
	// provider. If TAGS is specified, the tags are included in the response. If this
	// field is omitted, tags aren't included in the response.
	Include []types.CapacityProviderField

	// The maximum number of account setting results returned by
	// DescribeCapacityProviders in paginated output. When this parameter is used,
	// DescribeCapacityProviders only returns maxResults results in a single page
	// along with a nextToken response element. The remaining results of the initial
	// request can be seen by sending another DescribeCapacityProviders request with
	// the returned nextToken value. This value can be between 1 and 10. If this
	// parameter is not used, then DescribeCapacityProviders returns up to 10 results
	// and a nextToken value if applicable.
	MaxResults *int32

	// The nextToken value returned from a previous paginated DescribeCapacityProviders
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value.
	//
	// This token should be treated as an opaque identifier that is only used to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeCapacityProvidersOutput struct {

	// The list of capacity providers.
	CapacityProviders []types.CapacityProvider

	// Any failures associated with the call.
	Failures []types.Failure

	// The nextToken value to include in a future DescribeCapacityProviders request.
	// When the results of a DescribeCapacityProviders request exceed maxResults , this
	// value can be used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeCapacityProvidersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCapacityProviders{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCapacityProviders{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeCapacityProviders"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCapacityProviders(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeCapacityProviders(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeCapacityProviders",
	}
}
