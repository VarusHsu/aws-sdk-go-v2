// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new Amazon Chime SDK meeting in the specified media Region, with
// attendees. For more information about specifying media Regions, see Amazon Chime
// SDK Media Regions
// (https://docs.aws.amazon.com/chime/latest/dg/chime-sdk-meetings-regions.html) in
// the Amazon Chime Developer Guide. For more information about the Amazon Chime
// SDK, see Using the Amazon Chime SDK
// (https://docs.aws.amazon.com/chime/latest/dg/meetings-sdk.html) in the Amazon
// Chime Developer Guide.
func (c *Client) CreateMeetingWithAttendees(ctx context.Context, params *CreateMeetingWithAttendeesInput, optFns ...func(*Options)) (*CreateMeetingWithAttendeesOutput, error) {
	stack := middleware.NewStack("CreateMeetingWithAttendees", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateMeetingWithAttendeesMiddlewares(stack)
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
	addIdempotencyToken_opCreateMeetingWithAttendeesMiddleware(stack, options)
	addOpCreateMeetingWithAttendeesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMeetingWithAttendees(options.Region), middleware.Before)
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
			OperationName: "CreateMeetingWithAttendees",
			Err:           err,
		}
	}
	out := result.(*CreateMeetingWithAttendeesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMeetingWithAttendeesInput struct {

	// The configuration for resource targets to receive notifications when Amazon
	// Chime SDK meeting and attendee events occur. The Amazon Chime SDK supports
	// resource targets located in the US East (N. Virginia) AWS Region (us-east-1).
	NotificationsConfiguration *types.MeetingNotificationConfiguration

	// The unique identifier for the client request. Use a different token for
	// different meetings.
	//
	// This member is required.
	ClientRequestToken *string

	// The external meeting ID.
	ExternalMeetingId *string

	// The request containing the attendees to create.
	Attendees []*types.CreateAttendeeRequestItem

	// The Region in which to create the meeting. Default: us-east-1. Available values:
	// ap-northeast-1, ap-southeast-1, ap-southeast-2, ca-central-1, eu-central-1,
	// eu-north-1, eu-west-1, eu-west-2, eu-west-3, sa-east-1, us-east-1, us-east-2,
	// us-west-1, us-west-2.
	MediaRegion *string

	// Reserved.
	MeetingHostId *string

	// The tag key-value pairs.
	Tags []*types.Tag
}

type CreateMeetingWithAttendeesOutput struct {

	// If the action fails for one or more of the attendees in the request, a list of
	// the attendees is returned, along with error codes and error messages.
	Errors []*types.CreateAttendeeError

	// The attendee information, including attendees IDs and join tokens.
	Attendees []*types.Attendee

	// A meeting created using the Amazon Chime SDK.
	Meeting *types.Meeting

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateMeetingWithAttendeesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateMeetingWithAttendees{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMeetingWithAttendees{}, middleware.After)
}

type idempotencyToken_initializeOpCreateMeetingWithAttendees struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateMeetingWithAttendees) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateMeetingWithAttendees) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateMeetingWithAttendeesInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateMeetingWithAttendeesInput ")
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
func addIdempotencyToken_opCreateMeetingWithAttendeesMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateMeetingWithAttendees{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateMeetingWithAttendees(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "CreateMeetingWithAttendees",
	}
}
