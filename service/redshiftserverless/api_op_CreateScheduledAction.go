// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftserverless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/redshiftserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a scheduled action. A scheduled action contains a schedule and an
// Amazon Redshift API action. For example, you can create a schedule of when to
// run the CreateSnapshot API operation.
func (c *Client) CreateScheduledAction(ctx context.Context, params *CreateScheduledActionInput, optFns ...func(*Options)) (*CreateScheduledActionOutput, error) {
	if params == nil {
		params = &CreateScheduledActionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateScheduledAction", params, optFns, c.addOperationCreateScheduledActionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateScheduledActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateScheduledActionInput struct {

	// The name of the namespace for which to create a scheduled action.
	//
	// This member is required.
	NamespaceName *string

	// The ARN of the IAM role to assume to run the scheduled action. This IAM role
	// must have permission to run the Amazon Redshift Serverless API operation in the
	// scheduled action. This IAM role must allow the Amazon Redshift scheduler to
	// schedule creating snapshots. (Principal scheduler.redshift.amazonaws.com) to
	// assume permissions on your behalf. For more information about the IAM role to
	// use with the Amazon Redshift scheduler, see [Using Identity-Based Policies for Amazon Redshift]in the Amazon Redshift Management
	// Guide
	//
	// [Using Identity-Based Policies for Amazon Redshift]: https://docs.aws.amazon.com/redshift/latest/mgmt/redshift-iam-access-control-identity-based.html
	//
	// This member is required.
	RoleArn *string

	// The schedule for a one-time (at timestamp format) or recurring (cron format)
	// scheduled action. Schedule invocations must be separated by at least one hour.
	// Times are in UTC.
	//
	//   - Format of at timestamp is yyyy-mm-ddThh:mm:ss . For example,
	//   2016-03-04T17:27:00 .
	//
	//   - Format of cron expression is (Minutes Hours Day-of-month Month Day-of-week
	//   Year) . For example, "(0 10 ? * MON *)" . For more information, see [Cron Expressions]in the
	//   Amazon CloudWatch Events User Guide.
	//
	// [Cron Expressions]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html#CronExpressions
	//
	// This member is required.
	Schedule types.Schedule

	// The name of the scheduled action.
	//
	// This member is required.
	ScheduledActionName *string

	// A JSON format string of the Amazon Redshift Serverless API operation with input
	// parameters. The following is an example of a target action.
	//
	//     "{"CreateSnapshot": {"NamespaceName": "sampleNamespace","SnapshotName":
	//     "sampleSnapshot", "retentionPeriod": "1"}}"
	//
	// This member is required.
	TargetAction types.TargetAction

	// Indicates whether the schedule is enabled. If false, the scheduled action does
	// not trigger. For more information about state of the scheduled action, see [ScheduledAction].
	//
	// [ScheduledAction]: https://docs.aws.amazon.com/redshift-serverless/latest/APIReference/API_ScheduledAction.html
	Enabled *bool

	// The end time in UTC when the schedule is no longer active. After this time, the
	// scheduled action does not trigger.
	EndTime *time.Time

	// The description of the scheduled action.
	ScheduledActionDescription *string

	// The start time in UTC when the schedule is active. Before this time, the
	// scheduled action does not trigger.
	StartTime *time.Time

	noSmithyDocumentSerde
}

type CreateScheduledActionOutput struct {

	// The returned ScheduledAction object that describes the properties of a
	// scheduled action.
	ScheduledAction *types.ScheduledActionResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateScheduledActionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateScheduledAction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateScheduledAction{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateScheduledAction"); err != nil {
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
	if err = addOpCreateScheduledActionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateScheduledAction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateScheduledAction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateScheduledAction",
	}
}
