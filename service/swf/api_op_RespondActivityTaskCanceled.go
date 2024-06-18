// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Used by workers to tell the service that the ActivityTask identified by the taskToken was
// successfully canceled. Additional details can be provided using the details
// argument.
//
// These details (if provided) appear in the ActivityTaskCanceled event added to
// the workflow history.
//
// Only use this operation if the canceled flag of a RecordActivityTaskHeartbeat request returns true and if
// the activity can be safely undone or abandoned.
//
// A task is considered open from the time that it is scheduled until it is
// closed. Therefore a task is reported as open while a worker is processing it. A
// task is closed after it has been specified in a call to RespondActivityTaskCompleted,
// RespondActivityTaskCanceled, RespondActivityTaskFailed, or the task has [timed out].
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
//   - Use a Resource element with the domain name to limit the action to only
//     specified domains.
//
//   - Use an Action element to allow or deny permission to call this action.
//
//   - You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [timed out]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dg-basic.html#swf-dev-timeout-types
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func (c *Client) RespondActivityTaskCanceled(ctx context.Context, params *RespondActivityTaskCanceledInput, optFns ...func(*Options)) (*RespondActivityTaskCanceledOutput, error) {
	if params == nil {
		params = &RespondActivityTaskCanceledInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RespondActivityTaskCanceled", params, optFns, c.addOperationRespondActivityTaskCanceledMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RespondActivityTaskCanceledOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RespondActivityTaskCanceledInput struct {

	// The taskToken of the ActivityTask.
	//
	// taskToken is generated by the service and should be treated as an opaque value.
	// If the task is passed to another process, its taskToken must also be passed.
	// This enables it to provide its progress and respond with results.
	//
	// This member is required.
	TaskToken *string

	//  Information about the cancellation.
	Details *string

	noSmithyDocumentSerde
}

type RespondActivityTaskCanceledOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRespondActivityTaskCanceledMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpRespondActivityTaskCanceled{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpRespondActivityTaskCanceled{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RespondActivityTaskCanceled"); err != nil {
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
	if err = addOpRespondActivityTaskCanceledValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRespondActivityTaskCanceled(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRespondActivityTaskCanceled(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RespondActivityTaskCanceled",
	}
}
