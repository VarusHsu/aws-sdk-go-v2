// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The GetFileUploadURL operation generates and returns a temporary URL. You use
// the temporary URL to retrieve a file uploaded by a Worker as an answer to a
// FileUploadAnswer question for a HIT. The temporary URL is generated the instant
// the GetFileUploadURL operation is called, and is valid for 60 seconds. You can
// get a temporary file upload URL any time until the HIT is disposed. After the
// HIT is disposed, any uploaded files are deleted, and cannot be retrieved.
// Pending Deprecation on December 12, 2017. The Answer Specification structure
// will no longer support the <code>FileUploadAnswer</code> element to be used for
// the QuestionForm data structure. Instead, we recommend that Requesters who want
// to create HITs asking Workers to upload files to use Amazon S3. </p>
func (c *Client) GetFileUploadURL(ctx context.Context, params *GetFileUploadURLInput, optFns ...func(*Options)) (*GetFileUploadURLOutput, error) {
	stack := middleware.NewStack("GetFileUploadURL", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetFileUploadURLMiddlewares(stack)
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
	addOpGetFileUploadURLValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetFileUploadURL(options.Region), middleware.Before)
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
			OperationName: "GetFileUploadURL",
			Err:           err,
		}
	}
	out := result.(*GetFileUploadURLOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFileUploadURLInput struct {

	// The identifier of the question with a FileUploadAnswer, as specified in the
	// QuestionForm of the HIT.
	//
	// This member is required.
	QuestionIdentifier *string

	// The ID of the assignment that contains the question with a FileUploadAnswer.
	//
	// This member is required.
	AssignmentId *string
}

type GetFileUploadURLOutput struct {

	// A temporary URL for the file that the Worker uploaded for the answer.
	FileUploadURL *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetFileUploadURLMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetFileUploadURL{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetFileUploadURL{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetFileUploadURL(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "GetFileUploadURL",
	}
}
