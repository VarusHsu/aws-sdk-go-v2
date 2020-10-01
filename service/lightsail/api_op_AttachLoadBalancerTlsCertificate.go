// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Attaches a Transport Layer Security (TLS) certificate to your load balancer. TLS
// is just an updated, more secure version of Secure Socket Layer (SSL). Once you
// create and validate your certificate, you can attach it to your load balancer.
// You can also use this API to rotate the certificates on your account. Use the
// AttachLoadBalancerTlsCertificate action with the non-attached certificate, and
// it will replace the existing one and become the attached certificate. The
// AttachLoadBalancerTlsCertificate operation supports tag-based access control via
// resource tags applied to the resource identified by load balancer name. For more
// information, see the Lightsail Dev Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) AttachLoadBalancerTlsCertificate(ctx context.Context, params *AttachLoadBalancerTlsCertificateInput, optFns ...func(*Options)) (*AttachLoadBalancerTlsCertificateOutput, error) {
	stack := middleware.NewStack("AttachLoadBalancerTlsCertificate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAttachLoadBalancerTlsCertificateMiddlewares(stack)
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
	addOpAttachLoadBalancerTlsCertificateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAttachLoadBalancerTlsCertificate(options.Region), middleware.Before)
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
			OperationName: "AttachLoadBalancerTlsCertificate",
			Err:           err,
		}
	}
	out := result.(*AttachLoadBalancerTlsCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachLoadBalancerTlsCertificateInput struct {

	// The name of the load balancer to which you want to associate the SSL/TLS
	// certificate.
	//
	// This member is required.
	LoadBalancerName *string

	// The name of your SSL/TLS certificate.
	//
	// This member is required.
	CertificateName *string
}

type AttachLoadBalancerTlsCertificateOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request. These SSL/TLS certificates are only usable by Lightsail load balancers.
	// You can't get the certificate and use it for another purpose.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAttachLoadBalancerTlsCertificateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAttachLoadBalancerTlsCertificate{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAttachLoadBalancerTlsCertificate{}, middleware.After)
}

func newServiceMetadataMiddleware_opAttachLoadBalancerTlsCertificate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "AttachLoadBalancerTlsCertificate",
	}
}
