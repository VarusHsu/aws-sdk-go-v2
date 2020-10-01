// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the state of the AWS Systems Manager Change Calendar at an optional,
// specified time. If you specify a time, GetCalendarState returns the state of the
// calendar at a specific time, and returns the next time that the Change Calendar
// state will transition. If you do not specify a time, GetCalendarState assumes
// the current time. Change Calendar entries have two possible states: OPEN or
// CLOSED. For more information about Systems Manager Change Calendar, see AWS
// Systems Manager Change Calendar
// (https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-change-calendar.html)
// in the AWS Systems Manager User Guide.
func (c *Client) GetCalendarState(ctx context.Context, params *GetCalendarStateInput, optFns ...func(*Options)) (*GetCalendarStateOutput, error) {
	stack := middleware.NewStack("GetCalendarState", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetCalendarStateMiddlewares(stack)
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
	addOpGetCalendarStateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCalendarState(options.Region), middleware.Before)
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
			OperationName: "GetCalendarState",
			Err:           err,
		}
	}
	out := result.(*GetCalendarStateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCalendarStateInput struct {

	// (Optional) The specific time for which you want to get calendar state
	// information, in ISO 8601 (https://en.wikipedia.org/wiki/ISO_8601) format. If you
	// do not add AtTime, the current time is assumed.
	AtTime *string

	// The names or Amazon Resource Names (ARNs) of the Systems Manager documents that
	// represent the calendar entries for which you want to get the state.
	//
	// This member is required.
	CalendarNames []*string
}

type GetCalendarStateOutput struct {

	// The state of the calendar. An OPEN calendar indicates that actions are allowed
	// to proceed, and a CLOSED calendar indicates that actions are not allowed to
	// proceed.
	State types.CalendarState

	// The time, as an ISO 8601 (https://en.wikipedia.org/wiki/ISO_8601) string, that
	// you specified in your command. If you did not specify a time, GetCalendarState
	// uses the current time.
	AtTime *string

	// The time, as an ISO 8601 (https://en.wikipedia.org/wiki/ISO_8601) string, that
	// the calendar state will change. If the current calendar state is OPEN,
	// NextTransitionTime indicates when the calendar state changes to CLOSED, and
	// vice-versa.
	NextTransitionTime *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetCalendarStateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetCalendarState{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCalendarState{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetCalendarState(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "GetCalendarState",
	}
}
