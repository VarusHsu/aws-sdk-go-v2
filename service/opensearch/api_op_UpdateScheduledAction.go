// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Reschedules a planned domain configuration change for a later time. This change
// can be a scheduled service software update
// (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/service-software.html)
// or a blue/green Auto-Tune enhancement
// (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/auto-tune.html#auto-tune-types).
func (c *Client) UpdateScheduledAction(ctx context.Context, params *UpdateScheduledActionInput, optFns ...func(*Options)) (*UpdateScheduledActionOutput, error) {
	if params == nil {
		params = &UpdateScheduledActionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateScheduledAction", params, optFns, c.addOperationUpdateScheduledActionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateScheduledActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateScheduledActionInput struct {

	// The unique identifier of the action to reschedule. To retrieve this ID, send a
	// ListScheduledActions
	// (https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_ListScheduledActions.html)
	// request.
	//
	// This member is required.
	ActionID *string

	// The type of action to reschedule. Can be one of SERVICE_SOFTWARE_UPDATE,
	// JVM_HEAP_SIZE_TUNING, or JVM_YOUNG_GEN_TUNING. To retrieve this value, send a
	// ListScheduledActions
	// (https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_ListScheduledActions.html)
	// request.
	//
	// This member is required.
	ActionType types.ActionType

	// The name of the domain to reschedule an action for.
	//
	// This member is required.
	DomainName *string

	// When to schedule the action.
	//
	// * NOW - Immediately schedules the update to happen
	// in the current hour if there's capacity available.
	//
	// * TIMESTAMP - Lets you
	// specify a custom date and time to apply the update. If you specify this value,
	// you must also provide a value for DesiredStartTime.
	//
	// * OFF_PEAK_WINDOW - Marks
	// the action to be picked up during an upcoming off-peak window. There's no
	// guarantee that the change will be implemented during the next immediate window.
	// Depending on capacity, it might happen in subsequent days.
	//
	// This member is required.
	ScheduleAt types.ScheduleAt

	// The time to implement the change, in Coordinated Universal Time (UTC). Only
	// specify this parameter if you set ScheduleAt to TIMESTAMP.
	DesiredStartTime *int64

	noSmithyDocumentSerde
}

type UpdateScheduledActionOutput struct {

	// Information about the rescheduled action.
	ScheduledAction *types.ScheduledAction

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateScheduledActionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateScheduledAction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateScheduledAction{}, middleware.After)
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
	if err = addOpUpdateScheduledActionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateScheduledAction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateScheduledAction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "UpdateScheduledAction",
	}
}