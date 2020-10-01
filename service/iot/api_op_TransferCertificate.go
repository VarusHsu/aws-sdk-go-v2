// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Transfers the specified certificate to the specified AWS account. You can cancel
// the transfer until it is acknowledged by the recipient. No notification is sent
// to the transfer destination's account. It is up to the caller to notify the
// transfer target. The certificate being transferred must not be in the ACTIVE
// state. You can use the UpdateCertificate API to deactivate it. The certificate
// must not have any policies attached to it. You can use the DetachPrincipalPolicy
// API to detach them.
func (c *Client) TransferCertificate(ctx context.Context, params *TransferCertificateInput, optFns ...func(*Options)) (*TransferCertificateOutput, error) {
	stack := middleware.NewStack("TransferCertificate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpTransferCertificateMiddlewares(stack)
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
	addOpTransferCertificateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTransferCertificate(options.Region), middleware.Before)
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
			OperationName: "TransferCertificate",
			Err:           err,
		}
	}
	out := result.(*TransferCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the TransferCertificate operation.
type TransferCertificateInput struct {

	// The ID of the certificate. (The last part of the certificate ARN contains the
	// certificate ID.)
	//
	// This member is required.
	CertificateId *string

	// The transfer message.
	TransferMessage *string

	// The AWS account.
	//
	// This member is required.
	TargetAwsAccount *string
}

// The output from the TransferCertificate operation.
type TransferCertificateOutput struct {

	// The ARN of the certificate.
	TransferredCertificateArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpTransferCertificateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpTransferCertificate{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpTransferCertificate{}, middleware.After)
}

func newServiceMetadataMiddleware_opTransferCertificate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "TransferCertificate",
	}
}
