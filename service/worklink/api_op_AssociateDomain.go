// Code generated by smithy-go-codegen DO NOT EDIT.

package worklink

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Specifies a domain to be associated to Amazon WorkLink.
func (c *Client) AssociateDomain(ctx context.Context, params *AssociateDomainInput, optFns ...func(*Options)) (*AssociateDomainOutput, error) {
	stack := middleware.NewStack("AssociateDomain", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpAssociateDomainMiddlewares(stack)
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
	addOpAssociateDomainValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateDomain(options.Region), middleware.Before)
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
			OperationName: "AssociateDomain",
			Err:           err,
		}
	}
	out := result.(*AssociateDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateDomainInput struct {

	// The fully qualified domain name (FQDN).
	//
	// This member is required.
	DomainName *string

	// The ARN of an issued ACM certificate that is valid for the domain being
	// associated.
	//
	// This member is required.
	AcmCertificateArn *string

	// The Amazon Resource Name (ARN) of the fleet.
	//
	// This member is required.
	FleetArn *string

	// The name to display.
	DisplayName *string
}

type AssociateDomainOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpAssociateDomainMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpAssociateDomain{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateDomain{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssociateDomain(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "worklink",
		OperationName: "AssociateDomain",
	}
}
