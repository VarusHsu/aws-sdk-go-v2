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

// Registers a transit gateway in your global network. Not all Regions support
// transit gateways for global networks. For a list of the supported Regions, see [Region Availability]
// in the Amazon Web Services Transit Gateways for Global Networks User Guide. The
// transit gateway can be in any of the supported Amazon Web Services Regions, but
// it must be owned by the same Amazon Web Services account that owns the global
// network. You cannot register a transit gateway in more than one global network.
//
// [Region Availability]: https://docs.aws.amazon.com/network-manager/latest/tgwnm/what-are-global-networks.html#nm-available-regions
func (c *Client) RegisterTransitGateway(ctx context.Context, params *RegisterTransitGatewayInput, optFns ...func(*Options)) (*RegisterTransitGatewayOutput, error) {
	if params == nil {
		params = &RegisterTransitGatewayInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterTransitGateway", params, optFns, c.addOperationRegisterTransitGatewayMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterTransitGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterTransitGatewayInput struct {

	// The ID of the global network.
	//
	// This member is required.
	GlobalNetworkId *string

	// The Amazon Resource Name (ARN) of the transit gateway.
	//
	// This member is required.
	TransitGatewayArn *string

	noSmithyDocumentSerde
}

type RegisterTransitGatewayOutput struct {

	// Information about the transit gateway registration.
	TransitGatewayRegistration *types.TransitGatewayRegistration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterTransitGatewayMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRegisterTransitGateway{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRegisterTransitGateway{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RegisterTransitGateway"); err != nil {
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
	if err = addOpRegisterTransitGatewayValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterTransitGateway(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterTransitGateway(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RegisterTransitGateway",
	}
}
