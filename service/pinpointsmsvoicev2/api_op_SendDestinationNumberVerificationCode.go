// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Before you can send test messages to a verified destination phone number you
// need to opt-in the verified destination phone number. Creates a new text message
// with a verification code and send it to a verified destination phone number.
// Once you have the verification code use VerifyDestinationNumberto opt-in the verified destination
// phone number to receive messages.
func (c *Client) SendDestinationNumberVerificationCode(ctx context.Context, params *SendDestinationNumberVerificationCodeInput, optFns ...func(*Options)) (*SendDestinationNumberVerificationCodeOutput, error) {
	if params == nil {
		params = &SendDestinationNumberVerificationCodeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendDestinationNumberVerificationCode", params, optFns, c.addOperationSendDestinationNumberVerificationCodeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendDestinationNumberVerificationCodeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendDestinationNumberVerificationCodeInput struct {

	// Choose to send the verification code as an SMS or voice message.
	//
	// This member is required.
	VerificationChannel types.VerificationChannel

	// The unique identifier for the verified destination phone number.
	//
	// This member is required.
	VerifiedDestinationNumberId *string

	// The name of the configuration set to use. This can be either the
	// ConfigurationSetName or ConfigurationSetArn.
	ConfigurationSetName *string

	// You can specify custom data in this field. If you do, that data is logged to
	// the event destination.
	Context map[string]string

	// This field is used for any country-specific registration requirements.
	// Currently, this setting is only used when you send messages to recipients in
	// India using a sender ID. For more information see [Special requirements for sending SMS messages to recipients in India].
	//
	// [Special requirements for sending SMS messages to recipients in India]: https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-sms-senderid-india.html
	DestinationCountryParameters map[string]string

	// Choose the language to use for the message.
	LanguageCode types.LanguageCode

	// The origination identity of the message. This can be either the PhoneNumber,
	// PhoneNumberId, PhoneNumberArn, SenderId, SenderIdArn, PoolId, or PoolArn.
	OriginationIdentity *string

	noSmithyDocumentSerde
}

type SendDestinationNumberVerificationCodeOutput struct {

	// The unique identifier for the message.
	//
	// This member is required.
	MessageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendDestinationNumberVerificationCodeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpSendDestinationNumberVerificationCode{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpSendDestinationNumberVerificationCode{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SendDestinationNumberVerificationCode"); err != nil {
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
	if err = addOpSendDestinationNumberVerificationCodeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendDestinationNumberVerificationCode(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSendDestinationNumberVerificationCode(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SendDestinationNumberVerificationCode",
	}
}
