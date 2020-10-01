// Code generated by smithy-go-codegen DO NOT EDIT.

package elastictranscoder

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The TestRole operation tests the IAM role used to create the pipeline. The
// TestRole action lets you determine whether the IAM role you are using has
// sufficient permissions to let Elastic Transcoder perform tasks associated with
// the transcoding process. The action attempts to assume the specified IAM role,
// checks read access to the input and output buckets, and tries to send a test
// notification to Amazon SNS topics that you specify.
func (c *Client) TestRole(ctx context.Context, params *TestRoleInput, optFns ...func(*Options)) (*TestRoleOutput, error) {
	stack := middleware.NewStack("TestRole", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpTestRoleMiddlewares(stack)
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
	addOpTestRoleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTestRole(options.Region), middleware.Before)
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
			OperationName: "TestRole",
			Err:           err,
		}
	}
	out := result.(*TestRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The TestRoleRequest structure.
type TestRoleInput struct {

	// The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder
	// to test.
	//
	// This member is required.
	Role *string

	// The Amazon S3 bucket that Elastic Transcoder writes transcoded media files to.
	// The action attempts to read from this bucket.
	//
	// This member is required.
	OutputBucket *string

	// The ARNs of one or more Amazon Simple Notification Service (Amazon SNS) topics
	// that you want the action to send a test notification to.
	//
	// This member is required.
	Topics []*string

	// The Amazon S3 bucket that contains media files to be transcoded. The action
	// attempts to read from this bucket.
	//
	// This member is required.
	InputBucket *string
}

// The TestRoleResponse structure.
type TestRoleOutput struct {

	// If the operation is successful, this value is true; otherwise, the value is
	// false.
	Success *string

	// If the Success element contains false, this value is an array of one or more
	// error messages that were generated during the test process.
	Messages []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpTestRoleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpTestRole{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpTestRole{}, middleware.After)
}

func newServiceMetadataMiddleware_opTestRole(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elastictranscoder",
		OperationName: "TestRole",
	}
}
