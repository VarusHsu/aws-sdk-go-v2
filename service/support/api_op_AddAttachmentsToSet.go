// Code generated by smithy-go-codegen DO NOT EDIT.

package support

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds one or more attachments to an attachment set. An attachment set is a
// temporary container for attachments that you add to a case or case
// communication. The set is available for 1 hour after it's created. The
// expiryTime returned in the response is when the set expires.
//
//     * You must
// have a Business or Enterprise support plan to use the AWS Support API.
//
//     * If
// you call the AWS Support API from an account that does not have a Business or
// Enterprise support plan, the SubscriptionRequiredException error message
// appears. For information about changing your support plan, see AWS Support
// (http://aws.amazon.com/premiumsupport/).
func (c *Client) AddAttachmentsToSet(ctx context.Context, params *AddAttachmentsToSetInput, optFns ...func(*Options)) (*AddAttachmentsToSetOutput, error) {
	stack := middleware.NewStack("AddAttachmentsToSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAddAttachmentsToSetMiddlewares(stack)
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
	addOpAddAttachmentsToSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddAttachmentsToSet(options.Region), middleware.Before)
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
			OperationName: "AddAttachmentsToSet",
			Err:           err,
		}
	}
	out := result.(*AddAttachmentsToSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddAttachmentsToSetInput struct {

	// One or more attachments to add to the set. You can add up to three attachments
	// per set. The size limit is 5 MB per attachment. In the Attachment object, use
	// the data parameter to specify the contents of the attachment file. In the
	// previous request syntax, the value for data appear as blob, which is represented
	// as a base64-encoded string. The value for fileName is the name of the
	// attachment, such as troubleshoot-screenshot.png.
	//
	// This member is required.
	Attachments []*types.Attachment

	// The ID of the attachment set. If an attachmentSetId is not specified, a new
	// attachment set is created, and the ID of the set is returned in the response. If
	// an attachmentSetId is specified, the attachments are added to the specified set,
	// if it exists.
	AttachmentSetId *string
}

// The ID and expiry time of the attachment set returned by the AddAttachmentsToSet
// () operation.
type AddAttachmentsToSetOutput struct {

	// The ID of the attachment set. If an attachmentSetId was not specified, a new
	// attachment set is created, and the ID of the set is returned in the response. If
	// an attachmentSetId was specified, the attachments are added to the specified
	// set, if it exists.
	AttachmentSetId *string

	// The time and date when the attachment set expires.
	ExpiryTime *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAddAttachmentsToSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAddAttachmentsToSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddAttachmentsToSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opAddAttachmentsToSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "support",
		OperationName: "AddAttachmentsToSet",
	}
}
