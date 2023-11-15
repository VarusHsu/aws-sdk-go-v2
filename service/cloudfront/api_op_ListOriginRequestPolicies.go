// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of origin request policies. You can optionally apply a filter to
// return only the managed policies created by Amazon Web Services, or only the
// custom policies created in your Amazon Web Services account. You can optionally
// specify the maximum number of items to receive in the response. If the total
// number of items in the list exceeds the maximum that you specify, or the default
// maximum, the response is paginated. To get the next page of items, send a
// subsequent request that specifies the NextMarker value from the current
// response as the Marker value in the subsequent request.
func (c *Client) ListOriginRequestPolicies(ctx context.Context, params *ListOriginRequestPoliciesInput, optFns ...func(*Options)) (*ListOriginRequestPoliciesOutput, error) {
	if params == nil {
		params = &ListOriginRequestPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOriginRequestPolicies", params, optFns, c.addOperationListOriginRequestPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOriginRequestPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOriginRequestPoliciesInput struct {

	// Use this field when paginating results to indicate where to begin in your list
	// of origin request policies. The response includes origin request policies in the
	// list that occur after the marker. To get the next page of the list, set this
	// field's value to the value of NextMarker from the current page's response.
	Marker *string

	// The maximum number of origin request policies that you want in the response.
	MaxItems *int32

	// A filter to return only the specified kinds of origin request policies. Valid
	// values are:
	//   - managed – Returns only the managed policies created by Amazon Web Services.
	//   - custom – Returns only the custom policies created in your Amazon Web
	//   Services account.
	Type types.OriginRequestPolicyType

	noSmithyDocumentSerde
}

type ListOriginRequestPoliciesOutput struct {

	// A list of origin request policies.
	OriginRequestPolicyList *types.OriginRequestPolicyList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListOriginRequestPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpListOriginRequestPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListOriginRequestPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListOriginRequestPolicies"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOriginRequestPolicies(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListOriginRequestPolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListOriginRequestPolicies",
	}
}
