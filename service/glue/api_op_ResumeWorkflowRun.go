// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Restarts any completed nodes in a workflow run and resumes the run execution.
func (c *Client) ResumeWorkflowRun(ctx context.Context, params *ResumeWorkflowRunInput, optFns ...func(*Options)) (*ResumeWorkflowRunOutput, error) {
	stack := middleware.NewStack("ResumeWorkflowRun", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpResumeWorkflowRunMiddlewares(stack)
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
	addOpResumeWorkflowRunValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opResumeWorkflowRun(options.Region), middleware.Before)
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
			OperationName: "ResumeWorkflowRun",
			Err:           err,
		}
	}
	out := result.(*ResumeWorkflowRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResumeWorkflowRunInput struct {

	// The name of the workflow to resume.
	//
	// This member is required.
	Name *string

	// The ID of the workflow run to resume.
	//
	// This member is required.
	RunId *string

	// A list of the node IDs for the nodes you want to restart. The nodes that are to
	// be restarted must have an execution attempt in the original run.
	//
	// This member is required.
	NodeIds []*string
}

type ResumeWorkflowRunOutput struct {

	// A list of the node IDs for the nodes that were actually restarted.
	NodeIds []*string

	// The new ID assigned to the resumed workflow run. Each resume of a workflow run
	// will have a new run ID.
	RunId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpResumeWorkflowRunMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpResumeWorkflowRun{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpResumeWorkflowRun{}, middleware.After)
}

func newServiceMetadataMiddleware_opResumeWorkflowRun(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "ResumeWorkflowRun",
	}
}
