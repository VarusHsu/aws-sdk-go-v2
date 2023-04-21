// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Cancels a pending VPC peering authorization for the specified VPC. If you need
// to delete an existing VPC peering connection, use DeleteVpcPeeringConnection (https://docs.aws.amazon.com/gamelift/latest/apireference/API_DeleteVpcPeeringConnection.html)
// . Related actions All APIs by task (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) DeleteVpcPeeringAuthorization(ctx context.Context, params *DeleteVpcPeeringAuthorizationInput, optFns ...func(*Options)) (*DeleteVpcPeeringAuthorizationOutput, error) {
	if params == nil {
		params = &DeleteVpcPeeringAuthorizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteVpcPeeringAuthorization", params, optFns, c.addOperationDeleteVpcPeeringAuthorizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteVpcPeeringAuthorizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteVpcPeeringAuthorizationInput struct {

	// A unique identifier for the Amazon Web Services account that you use to manage
	// your Amazon GameLift fleet. You can find your Account ID in the Amazon Web
	// Services Management Console under account settings.
	//
	// This member is required.
	GameLiftAwsAccountId *string

	// A unique identifier for a VPC with resources to be accessed by your Amazon
	// GameLift fleet. The VPC must be in the same Region as your fleet. To look up a
	// VPC ID, use the VPC Dashboard (https://console.aws.amazon.com/vpc/) in the
	// Amazon Web Services Management Console. Learn more about VPC peering in VPC
	// Peering with Amazon GameLift Fleets (https://docs.aws.amazon.com/gamelift/latest/developerguide/vpc-peering.html)
	// .
	//
	// This member is required.
	PeerVpcId *string

	noSmithyDocumentSerde
}

type DeleteVpcPeeringAuthorizationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteVpcPeeringAuthorizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteVpcPeeringAuthorization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteVpcPeeringAuthorization{}, middleware.After)
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
	if err = addOpDeleteVpcPeeringAuthorizationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteVpcPeeringAuthorization(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opDeleteVpcPeeringAuthorization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DeleteVpcPeeringAuthorization",
	}
}
