// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the details for the input device
func (c *Client) DescribeInputDevice(ctx context.Context, params *DescribeInputDeviceInput, optFns ...func(*Options)) (*DescribeInputDeviceOutput, error) {
	stack := middleware.NewStack("DescribeInputDevice", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeInputDeviceMiddlewares(stack)
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
	addOpDescribeInputDeviceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInputDevice(options.Region), middleware.Before)
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
			OperationName: "DescribeInputDevice",
			Err:           err,
		}
	}
	out := result.(*DescribeInputDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for DescribeInputDeviceRequest
type DescribeInputDeviceInput struct {

	// The unique ID of this input device. For example, hd-123456789abcdef.
	//
	// This member is required.
	InputDeviceId *string
}

// Placeholder documentation for DescribeInputDeviceResponse
type DescribeInputDeviceOutput struct {

	// The network MAC address of the input device.
	MacAddress *string

	// The unique ARN of the input device.
	Arn *string

	// The state of the connection between the input device and AWS.
	ConnectionState types.InputDeviceConnectionState

	// The unique ID of the input device.
	Id *string

	// The network settings for the input device.
	NetworkSettings *types.InputDeviceNetworkSettings

	// The unique serial number of the input device.
	SerialNumber *string

	// The status of the action to synchronize the device configuration. If you change
	// the configuration of the input device (for example, the maximum bitrate),
	// MediaLive sends the new data to the device. The device might not update itself
	// immediately. SYNCED means the device has updated its configuration. SYNCING
	// means that it has not updated its configuration.
	DeviceSettingsSyncState types.DeviceSettingsSyncState

	// The type of the input device.
	Type types.InputDeviceType

	// Settings that describe an input device that is type HD.
	HdDeviceSettings *types.InputDeviceHdSettings

	// A name that you specify for the input device.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeInputDeviceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeInputDevice{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeInputDevice{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeInputDevice(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "DescribeInputDevice",
	}
}
