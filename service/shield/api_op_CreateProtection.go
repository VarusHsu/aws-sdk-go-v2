// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables AWS Shield Advanced for a specific AWS resource. The resource can be an
// Amazon CloudFront distribution, Elastic Load Balancing load balancer, AWS Global
// Accelerator accelerator, Elastic IP Address, or an Amazon Route 53 hosted zone.
// You can add protection to only a single resource with each CreateProtection
// request. If you want to add protection to multiple resources at once, use the
// AWS WAF console (https://console.aws.amazon.com/waf/). For more information see
// Getting Started with AWS Shield Advanced
// (https://docs.aws.amazon.com/waf/latest/developerguide/getting-started-ddos.html)
// and Add AWS Shield Advanced Protection to more AWS Resources
// (https://docs.aws.amazon.com/waf/latest/developerguide/configure-new-protection.html).
func (c *Client) CreateProtection(ctx context.Context, params *CreateProtectionInput, optFns ...func(*Options)) (*CreateProtectionOutput, error) {
	stack := middleware.NewStack("CreateProtection", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateProtectionMiddlewares(stack)
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
	addOpCreateProtectionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProtection(options.Region), middleware.Before)
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
			OperationName: "CreateProtection",
			Err:           err,
		}
	}
	out := result.(*CreateProtectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProtectionInput struct {

	// The ARN (Amazon Resource Name) of the resource to be protected. The ARN should
	// be in one of the following formats:
	//
	//     * For an Application Load Balancer:
	// arn:aws:elasticloadbalancing:region:account-id:loadbalancer/app/load-balancer-name/load-balancer-id
	//
	//
	// * For an Elastic Load Balancer (Classic Load Balancer):
	// arn:aws:elasticloadbalancing:region:account-id:loadbalancer/load-balancer-name
	//
	//
	// * For an AWS CloudFront distribution:
	// arn:aws:cloudfront::account-id:distribution/distribution-id
	//
	//     * For an AWS
	// Global Accelerator accelerator:
	// arn:aws:globalaccelerator::account-id:accelerator/accelerator-id
	//
	//     * For
	// Amazon Route 53: arn:aws:route53:::hostedzone/hosted-zone-id
	//
	//     * For an
	// Elastic IP address: arn:aws:ec2:region:account-id:eip-allocation/allocation-id
	//
	// This member is required.
	ResourceArn *string

	// Friendly name for the Protection you are creating.
	//
	// This member is required.
	Name *string
}

type CreateProtectionOutput struct {

	// The unique identifier (ID) for the Protection () object that is created.
	ProtectionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateProtectionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateProtection{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateProtection{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateProtection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "shield",
		OperationName: "CreateProtection",
	}
}
