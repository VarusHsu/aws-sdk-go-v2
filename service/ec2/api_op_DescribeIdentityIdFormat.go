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

// Describes the ID format settings for resources for the specified IAM user, IAM
// role, or root user. For example, you can view the resource types that are
// enabled for longer IDs. This request only returns information about resource
// types whose ID formats can be modified; it does not return information about
// other resource types. For more information, see Resource IDs
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/resource-ids.html) in the
// Amazon Elastic Compute Cloud User Guide. The following resource types support
// longer IDs: bundle | conversion-task | customer-gateway | dhcp-options |
// elastic-ip-allocation | elastic-ip-association | export-task | flow-log | image
// | import-task | instance | internet-gateway | network-acl |
// network-acl-association | network-interface | network-interface-attachment |
// prefix-list | reservation | route-table | route-table-association |
// security-group | snapshot | subnet | subnet-cidr-block-association | volume |
// vpc | vpc-cidr-block-association | vpc-endpoint | vpc-peering-connection |
// vpn-connection | vpn-gateway. These settings apply to the principal specified in
// the request. They do not apply to the principal that makes the request.
func (c *Client) DescribeIdentityIdFormat(ctx context.Context, params *DescribeIdentityIdFormatInput, optFns ...func(*Options)) (*DescribeIdentityIdFormatOutput, error) {
	stack := middleware.NewStack("DescribeIdentityIdFormat", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeIdentityIdFormatMiddlewares(stack)
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
	addOpDescribeIdentityIdFormatValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeIdentityIdFormat(options.Region), middleware.Before)
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
			OperationName: "DescribeIdentityIdFormat",
			Err:           err,
		}
	}
	out := result.(*DescribeIdentityIdFormatOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeIdentityIdFormatInput struct {

	// The type of resource: bundle | conversion-task | customer-gateway | dhcp-options
	// | elastic-ip-allocation | elastic-ip-association | export-task | flow-log |
	// image | import-task | instance | internet-gateway | network-acl |
	// network-acl-association | network-interface | network-interface-attachment |
	// prefix-list | reservation | route-table | route-table-association |
	// security-group | snapshot | subnet | subnet-cidr-block-association | volume |
	// vpc | vpc-cidr-block-association | vpc-endpoint | vpc-peering-connection |
	// vpn-connection | vpn-gateway
	Resource *string

	// The ARN of the principal, which can be an IAM role, IAM user, or the root user.
	//
	// This member is required.
	PrincipalArn *string
}

type DescribeIdentityIdFormatOutput struct {

	// Information about the ID format for the resources.
	Statuses []*types.IdFormat

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeIdentityIdFormatMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeIdentityIdFormat{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeIdentityIdFormat{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeIdentityIdFormat(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeIdentityIdFormat",
	}
}
