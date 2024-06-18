// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakergeospatial

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemakergeospatial/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves details of a Vector Enrichment Job for a given job Amazon Resource
// Name (ARN).
func (c *Client) GetVectorEnrichmentJob(ctx context.Context, params *GetVectorEnrichmentJobInput, optFns ...func(*Options)) (*GetVectorEnrichmentJobOutput, error) {
	if params == nil {
		params = &GetVectorEnrichmentJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetVectorEnrichmentJob", params, optFns, c.addOperationGetVectorEnrichmentJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetVectorEnrichmentJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetVectorEnrichmentJobInput struct {

	// The Amazon Resource Name (ARN) of the Vector Enrichment job.
	//
	// This member is required.
	Arn *string

	noSmithyDocumentSerde
}

type GetVectorEnrichmentJobOutput struct {

	// The Amazon Resource Name (ARN) of the Vector Enrichment job.
	//
	// This member is required.
	Arn *string

	// The creation time.
	//
	// This member is required.
	CreationTime *time.Time

	// The duration of the Vector Enrichment job, in seconds.
	//
	// This member is required.
	DurationInSeconds *int32

	// The Amazon Resource Name (ARN) of the IAM role that you specified for the job.
	//
	// This member is required.
	ExecutionRoleArn *string

	// Input configuration information for the Vector Enrichment job.
	//
	// This member is required.
	InputConfig *types.VectorEnrichmentJobInputConfig

	// An object containing information about the job configuration.
	//
	// This member is required.
	JobConfig types.VectorEnrichmentJobConfig

	// The name of the Vector Enrichment job.
	//
	// This member is required.
	Name *string

	// The status of the initiated Vector Enrichment job.
	//
	// This member is required.
	Status types.VectorEnrichmentJobStatus

	// The type of the Vector Enrichment job being initiated.
	//
	// This member is required.
	Type types.VectorEnrichmentJobType

	// Details about the errors generated during the Vector Enrichment job.
	ErrorDetails *types.VectorEnrichmentJobErrorDetails

	// Details about the errors generated during the ExportVectorEnrichmentJob.
	ExportErrorDetails *types.VectorEnrichmentJobExportErrorDetails

	// The export status of the Vector Enrichment job being initiated.
	ExportStatus types.VectorEnrichmentJobExportStatus

	// The Key Management Service key ID for server-side encryption.
	KmsKeyId *string

	// Each tag consists of a key and a value.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetVectorEnrichmentJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetVectorEnrichmentJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetVectorEnrichmentJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetVectorEnrichmentJob"); err != nil {
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
	if err = addOpGetVectorEnrichmentJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetVectorEnrichmentJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetVectorEnrichmentJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetVectorEnrichmentJob",
	}
}
