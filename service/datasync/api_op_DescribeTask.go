// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides information about a task, which defines where and how DataSync
// transfers your data.
func (c *Client) DescribeTask(ctx context.Context, params *DescribeTaskInput, optFns ...func(*Options)) (*DescribeTaskOutput, error) {
	if params == nil {
		params = &DescribeTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTask", params, optFns, c.addOperationDescribeTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeTaskRequest
type DescribeTaskInput struct {

	// Specifies the Amazon Resource Name (ARN) of the transfer task that you want
	// information about.
	//
	// This member is required.
	TaskArn *string

	noSmithyDocumentSerde
}

// DescribeTaskResponse
type DescribeTaskOutput struct {

	// The Amazon Resource Name (ARN) of an Amazon CloudWatch log group for monitoring
	// your task. For more information, see Monitoring DataSync with Amazon CloudWatch (https://docs.aws.amazon.com/datasync/latest/userguide/monitor-datasync.html)
	// .
	CloudWatchLogGroupArn *string

	// The time that the task was created.
	CreationTime *time.Time

	// The ARN of the most recent task execution.
	CurrentTaskExecutionArn *string

	// The ARN of your transfer's destination location.
	DestinationLocationArn *string

	// The ARNs of the network interfaces (https://docs.aws.amazon.com/datasync/latest/userguide/datasync-network.html#required-network-interfaces)
	// that DataSync created for your destination location.
	DestinationNetworkInterfaceArns []string

	// If there's an issue with your task, you can use the error code to help you
	// troubleshoot the problem. For more information, see Troubleshooting issues with
	// DataSync transfers (https://docs.aws.amazon.com/datasync/latest/userguide/troubleshooting-datasync-locations-tasks.html)
	// .
	ErrorCode *string

	// If there's an issue with your task, you can use the error details to help you
	// troubleshoot the problem. For more information, see Troubleshooting issues with
	// DataSync transfers (https://docs.aws.amazon.com/datasync/latest/userguide/troubleshooting-datasync-locations-tasks.html)
	// .
	ErrorDetail *string

	// The exclude filters that define the files, objects, and folders in your source
	// location that you don't want DataSync to transfer. For more information and
	// examples, see Specifying what DataSync transfers by using filters (https://docs.aws.amazon.com/datasync/latest/userguide/filtering.html)
	// .
	Excludes []types.FilterRule

	// The include filters that define the files, objects, and folders in your source
	// location that you want DataSync to transfer. For more information and examples,
	// see Specifying what DataSync transfers by using filters (https://docs.aws.amazon.com/datasync/latest/userguide/filtering.html)
	// .
	Includes []types.FilterRule

	// The configuration of the manifest that lists the files or objects that you want
	// DataSync to transfer. For more information, see Specifying what DataSync
	// transfers by using a manifest (https://docs.aws.amazon.com/datasync/latest/userguide/transferring-with-manifest.html)
	// .
	ManifestConfig *types.ManifestConfig

	// The name of your task.
	Name *string

	// The task's settings. For example, what file metadata gets preserved, how data
	// integrity gets verified at the end of your transfer, bandwidth limits, among
	// other options.
	Options *types.Options

	// The schedule for when you want your task to run. For more information, see
	// Scheduling your task (https://docs.aws.amazon.com/datasync/latest/userguide/task-scheduling.html)
	// .
	Schedule *types.TaskSchedule

	// The details about your task schedule (https://docs.aws.amazon.com/datasync/latest/userguide/task-scheduling.html)
	// .
	ScheduleDetails *types.TaskScheduleDetails

	// The ARN of your transfer's source location.
	SourceLocationArn *string

	// The ARNs of the network interfaces (https://docs.aws.amazon.com/datasync/latest/userguide/datasync-network.html#required-network-interfaces)
	// that DataSync created for your source location.
	SourceNetworkInterfaceArns []string

	// The status of your task. For information about what each status means, see Task
	// statuses (https://docs.aws.amazon.com/datasync/latest/userguide/understand-task-statuses.html#understand-task-creation-statuses)
	// .
	Status types.TaskStatus

	// The ARN of your task.
	TaskArn *string

	// The configuration of your task report, which provides detailed information
	// about your DataSync transfer. For more information, see Monitoring your
	// DataSync transfers with task reports (https://docs.aws.amazon.com/datasync/latest/userguide/task-reports.html)
	// .
	TaskReportConfig *types.TaskReportConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTask{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeTask"); err != nil {
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
	if err = addOpDescribeTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeTask",
	}
}
