// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// GetCheckerIpRanges still works, but we recommend that you download
// ip-ranges.json, which includes IP address ranges for all AWS services. For more
// information, see IP Address Ranges of Amazon Route 53 Servers
// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/route-53-ip-addresses.html)
// in the Amazon Route 53 Developer Guide.
func (c *Client) GetCheckerIpRanges(ctx context.Context, params *GetCheckerIpRangesInput, optFns ...func(*Options)) (*GetCheckerIpRangesOutput, error) {
	stack := middleware.NewStack("GetCheckerIpRanges", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetCheckerIpRangesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCheckerIpRanges(options.Region), middleware.Before)
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
			OperationName: "GetCheckerIpRanges",
			Err:           err,
		}
	}
	out := result.(*GetCheckerIpRangesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Empty request.
type GetCheckerIpRangesInput struct {
}

// A complex type that contains the CheckerIpRanges element.
type GetCheckerIpRangesOutput struct {

	// A complex type that contains sorted list of IP ranges in CIDR format for Amazon
	// Route 53 health checkers.
	//
	// This member is required.
	CheckerIpRanges []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetCheckerIpRangesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetCheckerIpRanges{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetCheckerIpRanges{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetCheckerIpRanges(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "GetCheckerIpRanges",
	}
}
