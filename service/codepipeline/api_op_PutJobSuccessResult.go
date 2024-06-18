// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Represents the success of a job as returned to the pipeline by a job worker.
// Used for custom actions only.
func (c *Client) PutJobSuccessResult(ctx context.Context, params *PutJobSuccessResultInput, optFns ...func(*Options)) (*PutJobSuccessResultOutput, error) {
	if params == nil {
		params = &PutJobSuccessResultInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutJobSuccessResult", params, optFns, c.addOperationPutJobSuccessResultMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutJobSuccessResultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a PutJobSuccessResult action.
type PutJobSuccessResultInput struct {

	// The unique system-generated ID of the job that succeeded. This is the same ID
	// returned from PollForJobs .
	//
	// This member is required.
	JobId *string

	// A token generated by a job worker, such as a CodeDeploy deployment ID, that a
	// successful job provides to identify a custom action in progress. Future jobs use
	// this token to identify the running instance of the action. It can be reused to
	// return more information about the progress of the custom action. When the action
	// is complete, no continuation token should be supplied.
	ContinuationToken *string

	// The ID of the current revision of the artifact successfully worked on by the
	// job.
	CurrentRevision *types.CurrentRevision

	// The execution details of the successful job, such as the actions taken by the
	// job worker.
	ExecutionDetails *types.ExecutionDetails

	// Key-value pairs produced as output by a job worker that can be made available
	// to a downstream action configuration. outputVariables can be included only when
	// there is no continuation token on the request.
	OutputVariables map[string]string

	noSmithyDocumentSerde
}

type PutJobSuccessResultOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutJobSuccessResultMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutJobSuccessResult{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutJobSuccessResult{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutJobSuccessResult"); err != nil {
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
	if err = addOpPutJobSuccessResultValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutJobSuccessResult(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutJobSuccessResult(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutJobSuccessResult",
	}
}
