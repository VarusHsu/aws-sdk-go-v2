// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrassv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/greengrassv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a list of client devices with a core device. Use this API operation
// to specify which client devices can discover a core device through cloud
// discovery. With cloud discovery, client devices connect to IoT Greengrass to
// retrieve associated core devices' connectivity information and certificates. For
// more information, see [Configure cloud discovery]in the IoT Greengrass V2 Developer Guide.
//
// Client devices are local IoT devices that connect to and communicate with an
// IoT Greengrass core device over MQTT. You can connect client devices to a core
// device to sync MQTT messages and data to Amazon Web Services IoT Core and
// interact with client devices in Greengrass components. For more information, see
// [Interact with local IoT devices]in the IoT Greengrass V2 Developer Guide.
//
// [Configure cloud discovery]: https://docs.aws.amazon.com/greengrass/v2/developerguide/configure-cloud-discovery.html
// [Interact with local IoT devices]: https://docs.aws.amazon.com/greengrass/v2/developerguide/interact-with-local-iot-devices.html
func (c *Client) BatchAssociateClientDeviceWithCoreDevice(ctx context.Context, params *BatchAssociateClientDeviceWithCoreDeviceInput, optFns ...func(*Options)) (*BatchAssociateClientDeviceWithCoreDeviceOutput, error) {
	if params == nil {
		params = &BatchAssociateClientDeviceWithCoreDeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchAssociateClientDeviceWithCoreDevice", params, optFns, c.addOperationBatchAssociateClientDeviceWithCoreDeviceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchAssociateClientDeviceWithCoreDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchAssociateClientDeviceWithCoreDeviceInput struct {

	// The name of the core device. This is also the name of the IoT thing.
	//
	// This member is required.
	CoreDeviceThingName *string

	// The list of client devices to associate.
	Entries []types.AssociateClientDeviceWithCoreDeviceEntry

	noSmithyDocumentSerde
}

type BatchAssociateClientDeviceWithCoreDeviceOutput struct {

	// The list of any errors for the entries in the request. Each error entry
	// contains the name of the IoT thing that failed to associate.
	ErrorEntries []types.AssociateClientDeviceWithCoreDeviceErrorEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchAssociateClientDeviceWithCoreDeviceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchAssociateClientDeviceWithCoreDevice{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchAssociateClientDeviceWithCoreDevice{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchAssociateClientDeviceWithCoreDevice"); err != nil {
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
	if err = addOpBatchAssociateClientDeviceWithCoreDeviceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchAssociateClientDeviceWithCoreDevice(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchAssociateClientDeviceWithCoreDevice(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchAssociateClientDeviceWithCoreDevice",
	}
}
