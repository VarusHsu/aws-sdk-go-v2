// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists information about your private certificate authority (CA). You specify the
// private CA on input by its ARN (Amazon Resource Name). The output contains the
// status of your CA. This can be any of the following:
//
//     * CREATING - ACM
// Private CA is creating your private certificate authority.
//
//     *
// PENDING_CERTIFICATE - The certificate is pending. You must use your ACM Private
// CA-hosted or on-premises root or subordinate CA to sign your private CA CSR and
// then import it into PCA.
//
//     * ACTIVE - Your private CA is active.
//
//     *
// DISABLED - Your private CA has been disabled.
//
//     * EXPIRED - Your private CA
// certificate has expired.
//
//     * FAILED - Your private CA has failed. Your CA can
// fail because of problems such a network outage or backend AWS failure or other
// errors. A failed CA can never return to the pending state. You must create a new
// CA.
//
//     * DELETED - Your private CA is within the restoration period, after
// which it is permanently deleted. The length of time remaining in the CA's
// restoration period is also included in this action's output.
func (c *Client) DescribeCertificateAuthority(ctx context.Context, params *DescribeCertificateAuthorityInput, optFns ...func(*Options)) (*DescribeCertificateAuthorityOutput, error) {
	stack := middleware.NewStack("DescribeCertificateAuthority", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeCertificateAuthorityMiddlewares(stack)
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
	addOpDescribeCertificateAuthorityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCertificateAuthority(options.Region), middleware.Before)
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
			OperationName: "DescribeCertificateAuthority",
			Err:           err,
		}
	}
	out := result.(*DescribeCertificateAuthorityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCertificateAuthorityInput struct {

	// The Amazon Resource Name (ARN) that was returned when you called
	// CreateCertificateAuthority (). This must be of the form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	// .
	//
	// This member is required.
	CertificateAuthorityArn *string
}

type DescribeCertificateAuthorityOutput struct {

	// A CertificateAuthority () structure that contains information about your private
	// CA.
	CertificateAuthority *types.CertificateAuthority

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeCertificateAuthorityMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCertificateAuthority{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCertificateAuthority{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeCertificateAuthority(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "DescribeCertificateAuthority",
	}
}
