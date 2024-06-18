// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a VPC connection.
func (c *Client) UpdateVPCConnection(ctx context.Context, params *UpdateVPCConnectionInput, optFns ...func(*Options)) (*UpdateVPCConnectionOutput, error) {
	if params == nil {
		params = &UpdateVPCConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateVPCConnection", params, optFns, c.addOperationUpdateVPCConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateVPCConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateVPCConnectionInput struct {

	// The Amazon Web Services account ID of the account that contains the VPC
	// connection that you want to update.
	//
	// This member is required.
	AwsAccountId *string

	// The display name for the VPC connection.
	//
	// This member is required.
	Name *string

	// An IAM role associated with the VPC connection.
	//
	// This member is required.
	RoleArn *string

	// A list of security group IDs for the VPC connection.
	//
	// This member is required.
	SecurityGroupIds []string

	// A list of subnet IDs for the VPC connection.
	//
	// This member is required.
	SubnetIds []string

	// The ID of the VPC connection that you're updating. This ID is a unique
	// identifier for each Amazon Web Services Region in an Amazon Web Services
	// account.
	//
	// This member is required.
	VPCConnectionId *string

	// A list of IP addresses of DNS resolver endpoints for the VPC connection.
	DnsResolvers []string

	noSmithyDocumentSerde
}

type UpdateVPCConnectionOutput struct {

	// The Amazon Resource Name (ARN) of the VPC connection.
	Arn *string

	// The availability status of the VPC connection.
	AvailabilityStatus types.VPCConnectionAvailabilityStatus

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// The update status of the VPC connection's last update.
	UpdateStatus types.VPCConnectionResourceStatus

	// The ID of the VPC connection that you are updating. This ID is a unique
	// identifier for each Amazon Web Services Region in anAmazon Web Services account.
	VPCConnectionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateVPCConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateVPCConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateVPCConnection{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateVPCConnection"); err != nil {
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
	if err = addOpUpdateVPCConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateVPCConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateVPCConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateVPCConnection",
	}
}
