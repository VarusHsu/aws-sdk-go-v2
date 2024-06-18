// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a customer gateway with a device and optionally, with a link. If you
// specify a link, it must be associated with the specified device.
//
// You can only associate customer gateways that are connected to a VPN attachment
// on a transit gateway or core network registered in your global network. When you
// register a transit gateway or core network, customer gateways that are connected
// to the transit gateway are automatically included in the global network. To list
// customer gateways that are connected to a transit gateway, use the [DescribeVpnConnections]EC2 API and
// filter by transit-gateway-id .
//
// You cannot associate a customer gateway with more than one device and link.
//
// [DescribeVpnConnections]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeVpnConnections.html
func (c *Client) AssociateCustomerGateway(ctx context.Context, params *AssociateCustomerGatewayInput, optFns ...func(*Options)) (*AssociateCustomerGatewayOutput, error) {
	if params == nil {
		params = &AssociateCustomerGatewayInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateCustomerGateway", params, optFns, c.addOperationAssociateCustomerGatewayMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateCustomerGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateCustomerGatewayInput struct {

	// The Amazon Resource Name (ARN) of the customer gateway.
	//
	// This member is required.
	CustomerGatewayArn *string

	// The ID of the device.
	//
	// This member is required.
	DeviceId *string

	// The ID of the global network.
	//
	// This member is required.
	GlobalNetworkId *string

	// The ID of the link.
	LinkId *string

	noSmithyDocumentSerde
}

type AssociateCustomerGatewayOutput struct {

	// The customer gateway association.
	CustomerGatewayAssociation *types.CustomerGatewayAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateCustomerGatewayMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateCustomerGateway{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateCustomerGateway{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociateCustomerGateway"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpAssociateCustomerGatewayValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateCustomerGateway(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opAssociateCustomerGateway(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociateCustomerGateway",
	}
}
