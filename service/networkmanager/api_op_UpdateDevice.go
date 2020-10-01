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

// Updates the details for an existing device. To remove information for any of the
// parameters, specify an empty string.
func (c *Client) UpdateDevice(ctx context.Context, params *UpdateDeviceInput, optFns ...func(*Options)) (*UpdateDeviceOutput, error) {
	stack := middleware.NewStack("UpdateDevice", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateDeviceMiddlewares(stack)
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
	addOpUpdateDeviceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDevice(options.Region), middleware.Before)
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
			OperationName: "UpdateDevice",
			Err:           err,
		}
	}
	out := result.(*UpdateDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDeviceInput struct {

	// The serial number of the device. Length Constraints: Maximum length of 128
	// characters.
	SerialNumber *string

	// The ID of the site.
	SiteId *string

	// The type of the device.
	Type *string

	// The ID of the global network.
	//
	// This member is required.
	GlobalNetworkId *string

	// The vendor of the device. Length Constraints: Maximum length of 128 characters.
	Vendor *string

	// Describes a location.
	Location *types.Location

	// A description of the device. Length Constraints: Maximum length of 256
	// characters.
	Description *string

	// The model of the device. Length Constraints: Maximum length of 128 characters.
	Model *string

	// The ID of the device.
	//
	// This member is required.
	DeviceId *string
}

type UpdateDeviceOutput struct {

	// Information about the device.
	Device *types.Device

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateDeviceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDevice{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDevice{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateDevice(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "UpdateDevice",
	}
}
