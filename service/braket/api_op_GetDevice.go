// Code generated by smithy-go-codegen DO NOT EDIT.

package braket

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/braket/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the devices available in Amazon Braket.
//
// For backwards compatibility with older versions of BraketSchemas, OpenQASM
// information is omitted from GetDevice API calls. To get this information the
// user-agent needs to present a recent version of the BraketSchemas (1.8.0 or
// later). The Braket SDK automatically reports this for you. If you do not see
// OpenQASM results in the GetDevice response when using a Braket SDK, you may need
// to set AWS_EXECUTION_ENV environment variable to configure user-agent. See the
// code examples provided below for how to do this for the AWS CLI, Boto3, and the
// Go, Java, and JavaScript/TypeScript SDKs.
func (c *Client) GetDevice(ctx context.Context, params *GetDeviceInput, optFns ...func(*Options)) (*GetDeviceOutput, error) {
	if params == nil {
		params = &GetDeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDevice", params, optFns, c.addOperationGetDeviceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDeviceInput struct {

	// The ARN of the device to retrieve.
	//
	// This member is required.
	DeviceArn *string

	noSmithyDocumentSerde
}

type GetDeviceOutput struct {

	// The ARN of the device.
	//
	// This member is required.
	DeviceArn *string

	// Details about the capabilities of the device.
	//
	// This value conforms to the media type: application/json
	//
	// This member is required.
	DeviceCapabilities *string

	// The name of the device.
	//
	// This member is required.
	DeviceName *string

	// The status of the device.
	//
	// This member is required.
	DeviceStatus types.DeviceStatus

	// The type of the device.
	//
	// This member is required.
	DeviceType types.DeviceType

	// The name of the partner company for the device.
	//
	// This member is required.
	ProviderName *string

	// List of information about tasks and jobs queued on a device.
	DeviceQueueInfo []types.DeviceQueueInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDeviceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDevice{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDevice{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetDevice"); err != nil {
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
	if err = addOpGetDeviceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDevice(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDevice(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetDevice",
	}
}
