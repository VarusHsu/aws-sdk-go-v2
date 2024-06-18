// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the runtime configuration for the specified fleet. The runtime
// configuration tells Amazon GameLift how to launch server processes on computes
// in the fleet. For managed EC2 fleets, it determines what server processes to run
// on each fleet instance. For container fleets, it describes what server processes
// to run in each replica container group. You can update a fleet's runtime
// configuration at any time after the fleet is created; it does not need to be in
// ACTIVE status.
//
// To update runtime configuration, specify the fleet ID and provide a
// RuntimeConfiguration with an updated set of server process configurations.
//
// If successful, the fleet's runtime configuration settings are updated. Fleet
// computes that run game server processes regularly check for and receive updated
// runtime configurations. The computes immediately take action to comply with the
// new configuration by launching new server processes or by not replacing existing
// processes when they shut down. Updating a fleet's runtime configuration never
// affects existing server processes.
//
// # Learn more
//
// [Setting up Amazon GameLift fleets]
//
// [Setting up Amazon GameLift fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
func (c *Client) UpdateRuntimeConfiguration(ctx context.Context, params *UpdateRuntimeConfigurationInput, optFns ...func(*Options)) (*UpdateRuntimeConfigurationOutput, error) {
	if params == nil {
		params = &UpdateRuntimeConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRuntimeConfiguration", params, optFns, c.addOperationUpdateRuntimeConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRuntimeConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRuntimeConfigurationInput struct {

	// A unique identifier for the fleet to update runtime configuration for. You can
	// use either the fleet ID or ARN value.
	//
	// This member is required.
	FleetId *string

	// Instructions for launching server processes on fleet computes. Server processes
	// run either a custom game build executable or a Realtime Servers script. The
	// runtime configuration lists the types of server processes to run, how to launch
	// them, and the number of processes to run concurrently.
	//
	// This member is required.
	RuntimeConfiguration *types.RuntimeConfiguration

	noSmithyDocumentSerde
}

type UpdateRuntimeConfigurationOutput struct {

	// The runtime configuration currently in use by computes in the fleet. If the
	// update is successful, all property changes are shown.
	RuntimeConfiguration *types.RuntimeConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRuntimeConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateRuntimeConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateRuntimeConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateRuntimeConfiguration"); err != nil {
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
	if err = addOpUpdateRuntimeConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRuntimeConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRuntimeConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateRuntimeConfiguration",
	}
}
