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
	"time"
)

// Gets information about an audit mitigation task that is used to apply mitigation
// actions to a set of audit findings. Properties include the actions being
// applied, the audit checks to which they're being applied, the task status, and
// aggregated task statistics.
func (c *Client) DescribeAuditMitigationActionsTask(ctx context.Context, params *DescribeAuditMitigationActionsTaskInput, optFns ...func(*Options)) (*DescribeAuditMitigationActionsTaskOutput, error) {
	stack := middleware.NewStack("DescribeAuditMitigationActionsTask", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeAuditMitigationActionsTaskMiddlewares(stack)
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
	addOpDescribeAuditMitigationActionsTaskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAuditMitigationActionsTask(options.Region), middleware.Before)
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
			OperationName: "DescribeAuditMitigationActionsTask",
			Err:           err,
		}
	}
	out := result.(*DescribeAuditMitigationActionsTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAuditMitigationActionsTaskInput struct {

	// The unique identifier for the audit mitigation task.
	//
	// This member is required.
	TaskId *string
}

type DescribeAuditMitigationActionsTaskOutput struct {

	// Specifies the mitigation actions that should be applied to specific audit
	// checks.
	AuditCheckToActionsMapping map[string][]*string

	// The date and time when the task was completed or canceled.
	EndTime *time.Time

	// Identifies the findings to which the mitigation actions are applied. This can be
	// by audit checks, by audit task, or a set of findings.
	Target *types.AuditMitigationActionsTaskTarget

	// Aggregate counts of the results when the mitigation tasks were applied to the
	// findings for this audit mitigation actions task.
	TaskStatistics map[string]*types.TaskStatisticsForAuditCheck

	// The current status of the task.
	TaskStatus types.AuditMitigationActionsTaskStatus

	// Specifies the mitigation actions and their parameters that are applied as part
	// of this task.
	ActionsDefinition []*types.MitigationAction

	// The date and time when the task was started.
	StartTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeAuditMitigationActionsTaskMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAuditMitigationActionsTask{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAuditMitigationActionsTask{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAuditMitigationActionsTask(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeAuditMitigationActionsTask",
	}
}
