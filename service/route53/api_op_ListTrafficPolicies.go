// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about the latest version for every traffic policy that is
// associated with the current AWS account. Policies are listed in the order that
// they were created in.
func (c *Client) ListTrafficPolicies(ctx context.Context, params *ListTrafficPoliciesInput, optFns ...func(*Options)) (*ListTrafficPoliciesOutput, error) {
	stack := middleware.NewStack("ListTrafficPolicies", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpListTrafficPoliciesMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTrafficPolicies(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "ListTrafficPolicies",
			Err:           err,
		}
	}
	out := result.(*ListTrafficPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains the information about the request to list the
// traffic policies that are associated with the current AWS account.
type ListTrafficPoliciesInput struct {

	// (Conditional) For your first request to ListTrafficPolicies, don't include the
	// TrafficPolicyIdMarker parameter. If you have more traffic policies than the
	// value of MaxItems, ListTrafficPolicies returns only the first MaxItems traffic
	// policies. To get the next group of policies, submit another request to
	// ListTrafficPolicies. For the value of TrafficPolicyIdMarker, specify the value
	// of TrafficPolicyIdMarker that was returned in the previous response.
	TrafficPolicyIdMarker *string

	// (Optional) The maximum number of traffic policies that you want Amazon Route 53
	// to return in response to this request. If you have more than MaxItems traffic
	// policies, the value of IsTruncated in the response is true, and the value of
	// TrafficPolicyIdMarker is the ID of the first traffic policy that Route 53 will
	// return if you submit another request.
	MaxItems *string
}

// A complex type that contains the response information for the request.
type ListTrafficPoliciesOutput struct {

	// If the value of IsTruncated is true, TrafficPolicyIdMarker is the ID of the
	// first traffic policy in the next group of MaxItems traffic policies.
	//
	// This member is required.
	TrafficPolicyIdMarker *string

	// A flag that indicates whether there are more traffic policies to be listed. If
	// the response was truncated, you can get the next group of traffic policies by
	// submitting another ListTrafficPolicies request and specifying the value of
	// TrafficPolicyIdMarker in the TrafficPolicyIdMarker request parameter.
	//
	// This member is required.
	IsTruncated *bool

	// The value that you specified for the MaxItems parameter in the
	// ListTrafficPolicies request that produced the current response.
	//
	// This member is required.
	MaxItems *string

	// A list that contains one TrafficPolicySummary element for each traffic policy
	// that was created by the current AWS account.
	//
	// This member is required.
	TrafficPolicySummaries []*types.TrafficPolicySummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpListTrafficPoliciesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpListTrafficPolicies{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpListTrafficPolicies{}, middleware.After)
}

func newServiceMetadataMiddleware_opListTrafficPolicies(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "ListTrafficPolicies",
	}
}
