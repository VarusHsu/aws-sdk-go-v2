// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Advertises an IPv4 address range that is provisioned for use with your AWS
// resources through bring your own IP addresses (BYOIP). It can take a few minutes
// before traffic to the specified addresses starts routing to AWS because of
// propagation delays. To see an AWS CLI example of advertising an address range,
// scroll down to Example. To stop advertising the BYOIP address range, use
// WithdrawByoipCidr
// (https://docs.aws.amazon.com/global-accelerator/latest/api/WithdrawByoipCidr.html).
// For more information, see Bring Your Own IP Addresses (BYOIP)
// (https://docs.aws.amazon.com/global-accelerator/latest/dg/using-byoip.html) in
// the AWS Global Accelerator Developer Guide.
func (c *Client) AdvertiseByoipCidr(ctx context.Context, params *AdvertiseByoipCidrInput, optFns ...func(*Options)) (*AdvertiseByoipCidrOutput, error) {
	stack := middleware.NewStack("AdvertiseByoipCidr", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAdvertiseByoipCidrMiddlewares(stack)
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
	addOpAdvertiseByoipCidrValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAdvertiseByoipCidr(options.Region), middleware.Before)
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
			OperationName: "AdvertiseByoipCidr",
			Err:           err,
		}
	}
	out := result.(*AdvertiseByoipCidrOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AdvertiseByoipCidrInput struct {

	// The address range, in CIDR notation. This must be the exact range that you
	// provisioned. You can't advertise only a portion of the provisioned range.
	//
	// This member is required.
	Cidr *string
}

type AdvertiseByoipCidrOutput struct {

	// Information about the address range.
	ByoipCidr *types.ByoipCidr

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAdvertiseByoipCidrMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAdvertiseByoipCidr{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdvertiseByoipCidr{}, middleware.After)
}

func newServiceMetadataMiddleware_opAdvertiseByoipCidr(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "AdvertiseByoipCidr",
	}
}
