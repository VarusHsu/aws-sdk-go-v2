// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers a robot with a fleet.
//
// This API is no longer supported and will throw an error if used.
//
// Deprecated: AWS RoboMaker is unable to process this request as the support for
// the AWS RoboMaker application deployment feature has ended. For additional
// information, see https://docs.aws.amazon.com/robomaker/latest/dg/fleets.html.
func (c *Client) RegisterRobot(ctx context.Context, params *RegisterRobotInput, optFns ...func(*Options)) (*RegisterRobotOutput, error) {
	if params == nil {
		params = &RegisterRobotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterRobot", params, optFns, c.addOperationRegisterRobotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterRobotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterRobotInput struct {

	// The Amazon Resource Name (ARN) of the fleet.
	//
	// This member is required.
	Fleet *string

	// The Amazon Resource Name (ARN) of the robot.
	//
	// This member is required.
	Robot *string

	noSmithyDocumentSerde
}

type RegisterRobotOutput struct {

	// The Amazon Resource Name (ARN) of the fleet that the robot will join.
	Fleet *string

	// Information about the robot registration.
	Robot *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterRobotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRegisterRobot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRegisterRobot{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RegisterRobot"); err != nil {
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
	if err = addOpRegisterRobotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterRobot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterRobot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RegisterRobot",
	}
}
