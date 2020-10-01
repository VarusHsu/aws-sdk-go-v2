// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an AWS IoT OTAUpdate on a target group of things or groups.
func (c *Client) CreateOTAUpdate(ctx context.Context, params *CreateOTAUpdateInput, optFns ...func(*Options)) (*CreateOTAUpdateOutput, error) {
	stack := middleware.NewStack("CreateOTAUpdate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateOTAUpdateMiddlewares(stack)
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
	addOpCreateOTAUpdateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateOTAUpdate(options.Region), middleware.Before)
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
			OperationName: "CreateOTAUpdate",
			Err:           err,
		}
	}
	out := result.(*CreateOTAUpdateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateOTAUpdateInput struct {

	// Configuration for the rollout of OTA updates.
	AwsJobExecutionsRolloutConfig *types.AwsJobExecutionsRolloutConfig

	// The IAM role that grants AWS IoT access to the Amazon S3, AWS IoT jobs and AWS
	// Code Signing resources to create an OTA update job.
	//
	// This member is required.
	RoleArn *string

	// The devices targeted to receive OTA updates.
	//
	// This member is required.
	Targets []*string

	// The protocol used to transfer the OTA update image. Valid values are [HTTP],
	// [MQTT], [HTTP, MQTT]. When both HTTP and MQTT are specified, the target device
	// can choose the protocol.
	Protocols []types.Protocol

	// Configuration information for pre-signed URLs.
	AwsJobPresignedUrlConfig *types.AwsJobPresignedUrlConfig

	// The criteria that determine when and how a job abort takes place.
	AwsJobAbortConfig *types.AwsJobAbortConfig

	// A list of additional OTA update parameters which are name-value pairs.
	AdditionalParameters map[string]*string

	// The description of the OTA update.
	Description *string

	// Specifies the amount of time each device has to finish its execution of the job.
	// A timer is started when the job execution status is set to IN_PROGRESS. If the
	// job execution status is not set to another terminal state before the timer
	// expires, it will be automatically set to TIMED_OUT.
	AwsJobTimeoutConfig *types.AwsJobTimeoutConfig

	// The files to be streamed by the OTA update.
	//
	// This member is required.
	Files []*types.OTAUpdateFile

	// Metadata which can be used to manage updates.
	Tags []*types.Tag

	// Specifies whether the update will continue to run (CONTINUOUS), or will be
	// complete after all the things specified as targets have completed the update
	// (SNAPSHOT). If continuous, the update may also be run on a thing when a change
	// is detected in a target. For example, an update will run on a thing when the
	// thing is added to a target group, even after the update was completed by all
	// things originally in the group. Valid values: CONTINUOUS | SNAPSHOT.
	TargetSelection types.TargetSelection

	// The ID of the OTA update to be created.
	//
	// This member is required.
	OtaUpdateId *string
}

type CreateOTAUpdateOutput struct {

	// The OTA update ID.
	OtaUpdateId *string

	// The AWS IoT job ID associated with the OTA update.
	AwsIotJobId *string

	// The OTA update ARN.
	OtaUpdateArn *string

	// The OTA update status.
	OtaUpdateStatus types.OTAUpdateStatus

	// The AWS IoT job ARN associated with the OTA update.
	AwsIotJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateOTAUpdateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateOTAUpdate{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateOTAUpdate{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateOTAUpdate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateOTAUpdate",
	}
}
