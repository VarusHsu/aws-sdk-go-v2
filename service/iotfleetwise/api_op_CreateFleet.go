// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleetwise

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotfleetwise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a fleet that represents a group of vehicles. You must create both a
// signal catalog and vehicles before you can create a fleet. For more information,
// see Fleets
// (https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/fleets.html) in
// the Amazon Web Services IoT FleetWise Developer Guide.
func (c *Client) CreateFleet(ctx context.Context, params *CreateFleetInput, optFns ...func(*Options)) (*CreateFleetOutput, error) {
	if params == nil {
		params = &CreateFleetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFleet", params, optFns, c.addOperationCreateFleetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFleetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFleetInput struct {

	// The unique ID of the fleet to create.
	//
	// This member is required.
	FleetId *string

	// The Amazon Resource Name (ARN) of a signal catalog.
	//
	// This member is required.
	SignalCatalogArn *string

	// A brief description of the fleet to create.
	Description *string

	// Metadata that can be used to manage the fleet.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateFleetOutput struct {

	// The ARN of the created fleet.
	//
	// This member is required.
	Arn *string

	// The ID of the created fleet.
	//
	// This member is required.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFleetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateFleet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateFleet{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateFleetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFleet(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opCreateFleet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotfleetwise",
		OperationName: "CreateFleet",
	}
}