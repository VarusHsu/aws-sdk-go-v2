// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Dissociates a package from the Amazon ES domain.
func (c *Client) DissociatePackage(ctx context.Context, params *DissociatePackageInput, optFns ...func(*Options)) (*DissociatePackageOutput, error) {
	stack := middleware.NewStack("DissociatePackage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDissociatePackageMiddlewares(stack)
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
	addOpDissociatePackageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDissociatePackage(options.Region), middleware.Before)
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
			OperationName: "DissociatePackage",
			Err:           err,
		}
	}
	out := result.(*DissociatePackageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for request parameters to DissociatePackage () operation.
type DissociatePackageInput struct {

	// Internal ID of the package that you want to associate with a domain. Use
	// DescribePackages to find this value.
	//
	// This member is required.
	PackageID *string

	// Name of the domain that you want to associate the package with.
	//
	// This member is required.
	DomainName *string
}

// Container for response returned by DissociatePackage () operation.
type DissociatePackageOutput struct {

	// DomainPackageDetails
	DomainPackageDetails *types.DomainPackageDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDissociatePackageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDissociatePackage{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDissociatePackage{}, middleware.After)
}

func newServiceMetadataMiddleware_opDissociatePackage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "DissociatePackage",
	}
}
