// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Syncrhonizes robots in a fleet to the latest deployment. This is helpful if
// robots were added after a deployment.
//
// This API will no longer be supported as of May 2, 2022. Use it to remove
// resources that were created for Deployment Service.
//
// Deprecated: Support for the AWS RoboMaker application deployment feature has
// ended. For additional information, see
// https://docs.aws.amazon.com/robomaker/latest/dg/fleets.html.
func (c *Client) SyncDeploymentJob(ctx context.Context, params *SyncDeploymentJobInput, optFns ...func(*Options)) (*SyncDeploymentJobOutput, error) {
	if params == nil {
		params = &SyncDeploymentJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SyncDeploymentJob", params, optFns, c.addOperationSyncDeploymentJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SyncDeploymentJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SyncDeploymentJobInput struct {

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	//
	// This member is required.
	ClientRequestToken *string

	// The target fleet for the synchronization.
	//
	// This member is required.
	Fleet *string

	noSmithyDocumentSerde
}

type SyncDeploymentJobOutput struct {

	// The Amazon Resource Name (ARN) of the synchronization request.
	Arn *string

	// The time, in milliseconds since the epoch, when the fleet was created.
	CreatedAt *time.Time

	// Information about the deployment application configurations.
	DeploymentApplicationConfigs []types.DeploymentApplicationConfig

	// Information about the deployment configuration.
	DeploymentConfig *types.DeploymentConfig

	// The failure code if the job fails:
	//
	// InternalServiceError Internal service error.
	//
	// RobotApplicationCrash Robot application exited abnormally.
	//
	// SimulationApplicationCrash  Simulation application exited abnormally.
	//
	// BadPermissionsRobotApplication Robot application bundle could not be downloaded.
	//
	// BadPermissionsSimulationApplication Simulation application bundle could not be
	// downloaded.
	//
	// BadPermissionsS3Output Unable to publish outputs to customer-provided S3 bucket.
	//
	// BadPermissionsCloudwatchLogs Unable to publish logs to customer-provided
	// CloudWatch Logs resource.
	//
	// SubnetIpLimitExceeded Subnet IP limit exceeded.
	//
	// ENILimitExceeded ENI limit exceeded.
	//
	// BadPermissionsUserCredentials Unable to use the Role provided.
	//
	// InvalidBundleRobotApplication Robot bundle cannot be extracted (invalid format,
	// bundling error, or other issue).
	//
	// InvalidBundleSimulationApplication Simulation bundle cannot be extracted
	// (invalid format, bundling error, or other issue).
	//
	// RobotApplicationVersionMismatchedEtag Etag for RobotApplication does not match
	// value during version creation.
	//
	// SimulationApplicationVersionMismatchedEtag Etag for SimulationApplication does
	// not match value during version creation.
	FailureCode types.DeploymentJobErrorCode

	// The failure reason if the job fails.
	FailureReason *string

	// The Amazon Resource Name (ARN) of the fleet.
	Fleet *string

	// The status of the synchronization job.
	Status types.DeploymentStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSyncDeploymentJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSyncDeploymentJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSyncDeploymentJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SyncDeploymentJob"); err != nil {
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
	if err = addIdempotencyToken_opSyncDeploymentJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpSyncDeploymentJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSyncDeploymentJob(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpSyncDeploymentJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpSyncDeploymentJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpSyncDeploymentJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*SyncDeploymentJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *SyncDeploymentJobInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opSyncDeploymentJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpSyncDeploymentJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opSyncDeploymentJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SyncDeploymentJob",
	}
}
