// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Updates a robot application.
func (c *Client) UpdateRobotApplication(ctx context.Context, params *UpdateRobotApplicationInput, optFns ...func(*Options)) (*UpdateRobotApplicationOutput, error) {
	stack := middleware.NewStack("UpdateRobotApplication", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateRobotApplicationMiddlewares(stack)
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
	addOpUpdateRobotApplicationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRobotApplication(options.Region), middleware.Before)
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
			OperationName: "UpdateRobotApplication",
			Err:           err,
		}
	}
	out := result.(*UpdateRobotApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRobotApplicationInput struct {

	// The revision id for the robot application.
	CurrentRevisionId *string

	// The sources of the robot application.
	//
	// This member is required.
	Sources []*types.SourceConfig

	// The application information for the robot application.
	//
	// This member is required.
	Application *string

	// The robot software suite (ROS distribution) used by the robot application.
	//
	// This member is required.
	RobotSoftwareSuite *types.RobotSoftwareSuite
}

type UpdateRobotApplicationOutput struct {

	// The sources of the robot application.
	Sources []*types.Source

	// The time, in milliseconds since the epoch, when the robot application was last
	// updated.
	LastUpdatedAt *time.Time

	// The robot software suite (ROS distribution) used by the robot application.
	RobotSoftwareSuite *types.RobotSoftwareSuite

	// The revision id of the robot application.
	RevisionId *string

	// The version of the robot application.
	Version *string

	// The name of the robot application.
	Name *string

	// The Amazon Resource Name (ARN) of the updated robot application.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateRobotApplicationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateRobotApplication{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateRobotApplication{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateRobotApplication(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "UpdateRobotApplication",
	}
}
