// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Stops a sentiment detection job in progress.
//
// If the job state is IN_PROGRESS , the job is marked for termination and put into
// the STOP_REQUESTED state. If the job completes before it can be stopped, it is
// put into the COMPLETED state; otherwise the job is be stopped and put into the
// STOPPED state.
//
// If the job is in the COMPLETED or FAILED state when you call the
// StopDominantLanguageDetectionJob operation, the operation returns a 400 Internal
// Request Exception.
//
// When a job is stopped, any documents already processed are written to the
// output location.
func (c *Client) StopSentimentDetectionJob(ctx context.Context, params *StopSentimentDetectionJobInput, optFns ...func(*Options)) (*StopSentimentDetectionJobOutput, error) {
	if params == nil {
		params = &StopSentimentDetectionJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopSentimentDetectionJob", params, optFns, c.addOperationStopSentimentDetectionJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopSentimentDetectionJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopSentimentDetectionJobInput struct {

	// The identifier of the sentiment detection job to stop.
	//
	// This member is required.
	JobId *string

	noSmithyDocumentSerde
}

type StopSentimentDetectionJobOutput struct {

	// The identifier of the sentiment detection job to stop.
	JobId *string

	// Either STOP_REQUESTED if the job is currently running, or STOPPED if the job
	// was previously stopped with the StopSentimentDetectionJob operation.
	JobStatus types.JobStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStopSentimentDetectionJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopSentimentDetectionJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopSentimentDetectionJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StopSentimentDetectionJob"); err != nil {
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
	if err = addOpStopSentimentDetectionJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopSentimentDetectionJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopSentimentDetectionJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StopSentimentDetectionJob",
	}
}
