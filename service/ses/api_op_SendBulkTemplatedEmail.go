// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Composes an email message to multiple destinations. The message body is created
// using an email template.
//
// To send email using this operation, your call must meet the following
// requirements:
//
//   - The call must refer to an existing email template. You can create email
//     templates using CreateTemplate.
//
//   - The message must be sent from a verified email address or domain.
//
//   - If your account is still in the Amazon SES sandbox, you may send only to
//     verified addresses or domains, or to email addresses associated with the Amazon
//     SES Mailbox Simulator. For more information, see [Verifying Email Addresses and Domains]in the Amazon SES Developer
//     Guide.
//
//   - The maximum message size is 10 MB.
//
//   - Each Destination parameter must include at least one recipient email
//     address. The recipient address can be a To: address, a CC: address, or a BCC:
//     address. If a recipient email address is invalid (that is, it is not in the
//     format UserName@[SubDomain.]Domain.TopLevelDomain), the entire message is
//     rejected, even if the message contains other recipients that are valid.
//
//   - The message may not include more than 50 recipients, across the To:, CC:
//     and BCC: fields. If you need to send an email message to a larger audience, you
//     can divide your recipient list into groups of 50 or fewer, and then call the
//     SendBulkTemplatedEmail operation several times to send the message to each
//     group.
//
//   - The number of destinations you can contact in a single call can be limited
//     by your account's maximum sending rate.
//
// [Verifying Email Addresses and Domains]: https://docs.aws.amazon.com/ses/latest/dg/verify-addresses-and-domains.html
func (c *Client) SendBulkTemplatedEmail(ctx context.Context, params *SendBulkTemplatedEmailInput, optFns ...func(*Options)) (*SendBulkTemplatedEmailOutput, error) {
	if params == nil {
		params = &SendBulkTemplatedEmailInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendBulkTemplatedEmail", params, optFns, c.addOperationSendBulkTemplatedEmailMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendBulkTemplatedEmailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to send a templated email to multiple destinations using
// Amazon SES. For more information, see the [Amazon SES Developer Guide].
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/send-personalized-email-api.html
type SendBulkTemplatedEmailInput struct {

	// One or more Destination objects. All of the recipients in a Destination receive
	// the same version of the email. You can specify up to 50 Destination objects
	// within a Destinations array.
	//
	// This member is required.
	Destinations []types.BulkEmailDestination

	// The email address that is sending the email. This email address must be either
	// individually verified with Amazon SES, or from a domain that has been verified
	// with Amazon SES. For information about verifying identities, see the [Amazon SES Developer Guide].
	//
	// If you are sending on behalf of another user and have been permitted to do so
	// by a sending authorization policy, then you must also specify the SourceArn
	// parameter. For more information about sending authorization, see the [Amazon SES Developer Guide].
	//
	// Amazon SES does not support the SMTPUTF8 extension, as described in [RFC6531]. For this
	// reason, the email address string must be 7-bit ASCII. If you want to send to or
	// from email addresses that contain Unicode characters in the domain part of an
	// address, you must encode the domain using Punycode. Punycode is not permitted in
	// the local part of the email address (the part before the @ sign) nor in the
	// "friendly from" name. If you want to use Unicode characters in the "friendly
	// from" name, you must encode the "friendly from" name using MIME encoded-word
	// syntax, as described in [Sending raw email using the Amazon SES API]. For more information about Punycode, see [RFC 3492].
	//
	// [RFC6531]: https://tools.ietf.org/html/rfc6531
	// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/sending-authorization.html
	// [Sending raw email using the Amazon SES API]: https://docs.aws.amazon.com/ses/latest/dg/send-email-raw.html
	// [RFC 3492]: http://tools.ietf.org/html/rfc3492
	//
	// This member is required.
	Source *string

	// The template to use when sending this email.
	//
	// This member is required.
	Template *string

	// The name of the configuration set to use when you send an email using
	// SendBulkTemplatedEmail .
	ConfigurationSetName *string

	// A list of tags, in the form of name/value pairs, to apply to an email that you
	// send to a destination using SendBulkTemplatedEmail .
	DefaultTags []types.MessageTag

	// A list of replacement values to apply to the template when replacement data is
	// not specified in a Destination object. These values act as a default or fallback
	// option when no other data is available.
	//
	// The template data is a JSON object, typically consisting of key-value pairs in
	// which the keys correspond to replacement tags in the email template.
	DefaultTemplateData *string

	// The reply-to email address(es) for the message. If the recipient replies to the
	// message, each reply-to address receives the reply.
	ReplyToAddresses []string

	// The email address that bounces and complaints are forwarded to when feedback
	// forwarding is enabled. If the message cannot be delivered to the recipient, then
	// an error message is returned from the recipient's ISP; this message is forwarded
	// to the email address specified by the ReturnPath parameter. The ReturnPath
	// parameter is never overwritten. This email address must be either individually
	// verified with Amazon SES, or from a domain that has been verified with Amazon
	// SES.
	ReturnPath *string

	// This parameter is used only for sending authorization. It is the ARN of the
	// identity that is associated with the sending authorization policy that permits
	// you to use the email address specified in the ReturnPath parameter.
	//
	// For example, if the owner of example.com (which has ARN
	// arn:aws:ses:us-east-1:123456789012:identity/example.com ) attaches a policy to
	// it that authorizes you to use feedback@example.com , then you would specify the
	// ReturnPathArn to be arn:aws:ses:us-east-1:123456789012:identity/example.com ,
	// and the ReturnPath to be feedback@example.com .
	//
	// For more information about sending authorization, see the [Amazon SES Developer Guide].
	//
	// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/sending-authorization.html
	ReturnPathArn *string

	// This parameter is used only for sending authorization. It is the ARN of the
	// identity that is associated with the sending authorization policy that permits
	// you to send for the email address specified in the Source parameter.
	//
	// For example, if the owner of example.com (which has ARN
	// arn:aws:ses:us-east-1:123456789012:identity/example.com ) attaches a policy to
	// it that authorizes you to send from user@example.com , then you would specify
	// the SourceArn to be arn:aws:ses:us-east-1:123456789012:identity/example.com ,
	// and the Source to be user@example.com .
	//
	// For more information about sending authorization, see the [Amazon SES Developer Guide].
	//
	// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/sending-authorization.html
	SourceArn *string

	// The ARN of the template to use when sending this email.
	TemplateArn *string

	noSmithyDocumentSerde
}

type SendBulkTemplatedEmailOutput struct {

	// One object per intended recipient. Check each response object and retry any
	// messages with a failure status. (Note that order of responses will be respective
	// to order of destinations in the request.)Receipt rules enable you to specify
	// which actions
	//
	// This member is required.
	Status []types.BulkEmailDestinationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendBulkTemplatedEmailMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSendBulkTemplatedEmail{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSendBulkTemplatedEmail{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SendBulkTemplatedEmail"); err != nil {
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
	if err = addOpSendBulkTemplatedEmailValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendBulkTemplatedEmail(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSendBulkTemplatedEmail(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SendBulkTemplatedEmail",
	}
}
