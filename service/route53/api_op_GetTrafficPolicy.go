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

// Gets information about a specific traffic policy version.
func (c *Client) GetTrafficPolicy(ctx context.Context, params *GetTrafficPolicyInput, optFns ...func(*Options)) (*GetTrafficPolicyOutput, error) {
	stack := middleware.NewStack("GetTrafficPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetTrafficPolicyMiddlewares(stack)
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
	addOpGetTrafficPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetTrafficPolicy(options.Region), middleware.Before)
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
			OperationName: "GetTrafficPolicy",
			Err:           err,
		}
	}
	out := result.(*GetTrafficPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Gets information about a specific traffic policy version.
type GetTrafficPolicyInput struct {

	// The version number of the traffic policy that you want to get information about.
	//
	// This member is required.
	Version *int32

	// The ID of the traffic policy that you want to get information about.
	//
	// This member is required.
	Id *string
}

// A complex type that contains the response information for the request.
type GetTrafficPolicyOutput struct {

	// A complex type that contains settings for the specified traffic policy.
	//
	// This member is required.
	TrafficPolicy *types.TrafficPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetTrafficPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetTrafficPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetTrafficPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetTrafficPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "GetTrafficPolicy",
	}
}
