// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides information about the specified medical transcription job.
//
// To view the status of the specified medical transcription job, check the
// TranscriptionJobStatus field. If the status is COMPLETED , the job is finished.
// You can find the results at the location specified in TranscriptFileUri . If the
// status is FAILED , FailureReason provides details on why your transcription job
// failed.
//
// To get a list of your medical transcription jobs, use the operation.
func (c *Client) GetMedicalTranscriptionJob(ctx context.Context, params *GetMedicalTranscriptionJobInput, optFns ...func(*Options)) (*GetMedicalTranscriptionJobOutput, error) {
	if params == nil {
		params = &GetMedicalTranscriptionJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMedicalTranscriptionJob", params, optFns, c.addOperationGetMedicalTranscriptionJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMedicalTranscriptionJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMedicalTranscriptionJobInput struct {

	// The name of the medical transcription job you want information about. Job names
	// are case sensitive.
	//
	// This member is required.
	MedicalTranscriptionJobName *string

	noSmithyDocumentSerde
}

type GetMedicalTranscriptionJobOutput struct {

	// Provides detailed information about the specified medical transcription job,
	// including job status and, if applicable, failure reason.
	MedicalTranscriptionJob *types.MedicalTranscriptionJob

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMedicalTranscriptionJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetMedicalTranscriptionJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetMedicalTranscriptionJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetMedicalTranscriptionJob"); err != nil {
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
	if err = addOpGetMedicalTranscriptionJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMedicalTranscriptionJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMedicalTranscriptionJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetMedicalTranscriptionJob",
	}
}
