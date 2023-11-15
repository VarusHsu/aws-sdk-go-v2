// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmeetings

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmeetings/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new attendee for an active Amazon Chime SDK meeting. For more
// information about the Amazon Chime SDK, see Using the Amazon Chime SDK (https://docs.aws.amazon.com/chime/latest/dg/meetings-sdk.html)
// in the Amazon Chime Developer Guide.
func (c *Client) CreateAttendee(ctx context.Context, params *CreateAttendeeInput, optFns ...func(*Options)) (*CreateAttendeeOutput, error) {
	if params == nil {
		params = &CreateAttendeeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAttendee", params, optFns, c.addOperationCreateAttendeeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAttendeeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAttendeeInput struct {

	// The Amazon Chime SDK external user ID. An idempotency token. Links the attendee
	// to an identity managed by a builder application. Pattern:
	// [-_&@+=,(){}\[\]\/«».:|'"#a-zA-Z0-9À-ÿ\s]* Values that begin with aws: are
	// reserved. You can't configure a value that uses this prefix.
	//
	// This member is required.
	ExternalUserId *string

	// The unique ID of the meeting.
	//
	// This member is required.
	MeetingId *string

	// The capabilities ( audio , video , or content ) that you want to grant an
	// attendee. If you don't specify capabilities, all users have send and receive
	// capabilities on all media channels by default. You use the capabilities with a
	// set of values that control what the capabilities can do, such as SendReceive
	// data. For more information about those values, see . When using capabilities, be
	// aware of these corner cases:
	//   - You can't set content capabilities to SendReceive or Receive unless you also
	//   set video capabilities to SendReceive or Receive . If you don't set the video
	//   capability to receive, the response will contain an HTTP 400 Bad Request status
	//   code. However, you can set your video capability to receive and you set your
	//   content capability to not receive.
	//   - When you change an audio capability from None or Receive to Send or
	//   SendReceive , and if the attendee left their microphone unmuted, audio will
	//   flow from the attendee to the other meeting participants.
	//   - When you change a video or content capability from None or Receive to Send
	//   or SendReceive , and if the attendee turned on their video or content streams,
	//   remote attendees can receive those streams, but only after media renegotiation
	//   between the client and the Amazon Chime back-end server.
	Capabilities *types.AttendeeCapabilities

	noSmithyDocumentSerde
}

type CreateAttendeeOutput struct {

	// The attendee information, including attendee ID and join token.
	Attendee *types.Attendee

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAttendeeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAttendee{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAttendee{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAttendee"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpCreateAttendeeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAttendee(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opCreateAttendee(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAttendee",
	}
}
