// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a network ACL in a VPC. Network ACLs provide an optional layer of
// security (in addition to security groups) for the instances in your VPC. For
// more information, see Network ACLs
// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_ACLs.html) in the Amazon
// Virtual Private Cloud User Guide.
func (c *Client) CreateNetworkAcl(ctx context.Context, params *CreateNetworkAclInput, optFns ...func(*Options)) (*CreateNetworkAclOutput, error) {
	stack := middleware.NewStack("CreateNetworkAcl", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpCreateNetworkAclMiddlewares(stack)
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
	addOpCreateNetworkAclValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNetworkAcl(options.Region), middleware.Before)
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
			OperationName: "CreateNetworkAcl",
			Err:           err,
		}
	}
	out := result.(*CreateNetworkAclOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateNetworkAclInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The ID of the VPC.
	//
	// This member is required.
	VpcId *string

	// The tags to assign to the network ACL.
	TagSpecifications []*types.TagSpecification
}

type CreateNetworkAclOutput struct {

	// Information about the network ACL.
	NetworkAcl *types.NetworkAcl

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpCreateNetworkAclMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpCreateNetworkAcl{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpCreateNetworkAcl{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateNetworkAcl(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateNetworkAcl",
	}
}
