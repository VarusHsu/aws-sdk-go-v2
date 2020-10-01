// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates a customer gateway with a device and optionally, with a link. If you
// specify a link, it must be associated with the specified device. You can only
// associate customer gateways that are connected to a VPN attachment on a transit
// gateway. The transit gateway must be registered in your global network. When you
// register a transit gateway, customer gateways that are connected to the transit
// gateway are automatically included in the global network. To list customer
// gateways that are connected to a transit gateway, use the DescribeVpnConnections
// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeVpnConnections.html)
// EC2 API and filter by transit-gateway-id. You cannot associate a customer
// gateway with more than one device and link.
func (c *Client) AssociateCustomerGateway(ctx context.Context, params *AssociateCustomerGatewayInput, optFns ...func(*Options)) (*AssociateCustomerGatewayOutput, error) {
	stack := middleware.NewStack("AssociateCustomerGateway", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpAssociateCustomerGatewayMiddlewares(stack)
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
	addOpAssociateCustomerGatewayValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateCustomerGateway(options.Region), middleware.Before)
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
			OperationName: "AssociateCustomerGateway",
			Err:           err,
		}
	}
	out := result.(*AssociateCustomerGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateCustomerGatewayInput struct {

	// The Amazon Resource Name (ARN) of the customer gateway. For more information,
	// see Resources Defined by Amazon EC2
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/list_amazonec2.html#amazonec2-resources-for-iam-policies).
	//
	// This member is required.
	CustomerGatewayArn *string

	// The ID of the global network.
	//
	// This member is required.
	GlobalNetworkId *string

	// The ID of the device.
	//
	// This member is required.
	DeviceId *string

	// The ID of the link.
	LinkId *string
}

type AssociateCustomerGatewayOutput struct {

	// The customer gateway association.
	CustomerGatewayAssociation *types.CustomerGatewayAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpAssociateCustomerGatewayMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpAssociateCustomerGateway{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateCustomerGateway{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssociateCustomerGateway(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "AssociateCustomerGateway",
	}
}
