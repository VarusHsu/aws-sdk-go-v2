// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sends a diagnostic interrupt to the specified Amazon EC2 instance to trigger a
// kernel panic (on Linux instances), or a blue screen/stop error (on Windows
// instances). For instances based on Intel and AMD processors, the interrupt is
// received as a non-maskable interrupt (NMI).  <p>In general, the operating system
// crashes and reboots when a kernel panic or stop error is triggered. The
// operating system can also be configured to perform diagnostic tasks, such as
// generating a memory dump file, loading a secondary kernel, or obtaining a call
// trace.</p> <p>Before sending a diagnostic interrupt to your instance, ensure
// that its operating system is configured to perform the required diagnostic
// tasks.</p> <p>For more information about configuring your operating system to
// generate a crash dump when a kernel panic or stop error occurs, see <a
// href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/diagnostic-interrupt.html">Send
// a diagnostic interrupt</a> (Linux instances) or <a
// href="https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/diagnostic-interrupt.html">Send
// a Diagnostic Interrupt</a> (Windows instances).</p>
func (c *Client) SendDiagnosticInterrupt(ctx context.Context, params *SendDiagnosticInterruptInput, optFns ...func(*Options)) (*SendDiagnosticInterruptOutput, error) {
	stack := middleware.NewStack("SendDiagnosticInterrupt", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpSendDiagnosticInterruptMiddlewares(stack)
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
	addOpSendDiagnosticInterruptValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSendDiagnosticInterrupt(options.Region), middleware.Before)
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
			OperationName: "SendDiagnosticInterrupt",
			Err:           err,
		}
	}
	out := result.(*SendDiagnosticInterruptOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendDiagnosticInterruptInput struct {

	// The ID of the instance.
	//
	// This member is required.
	InstanceId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type SendDiagnosticInterruptOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpSendDiagnosticInterruptMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpSendDiagnosticInterrupt{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpSendDiagnosticInterrupt{}, middleware.After)
}

func newServiceMetadataMiddleware_opSendDiagnosticInterrupt(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "SendDiagnosticInterrupt",
	}
}
