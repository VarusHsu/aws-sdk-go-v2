// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunedata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists Neptune ML model-training jobs. See [Model training using the modeltraining command]modeltraining .
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:neptune-db:ListMLModelTrainingJobs]IAM action in that cluster.
//
// [Model training using the modeltraining command]: https://docs.aws.amazon.com/neptune/latest/userguide/machine-learning-api-modeltraining.html
// [neptune-db:neptune-db:ListMLModelTrainingJobs]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#neptune-db:listmlmodeltrainingjobs
func (c *Client) ListMLModelTrainingJobs(ctx context.Context, params *ListMLModelTrainingJobsInput, optFns ...func(*Options)) (*ListMLModelTrainingJobsOutput, error) {
	if params == nil {
		params = &ListMLModelTrainingJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMLModelTrainingJobs", params, optFns, c.addOperationListMLModelTrainingJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMLModelTrainingJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMLModelTrainingJobsInput struct {

	// The maximum number of items to return (from 1 to 1024; the default is 10).
	MaxItems *int32

	// The ARN of an IAM role that provides Neptune access to SageMaker and Amazon S3
	// resources. This must be listed in your DB cluster parameter group or an error
	// will occur.
	NeptuneIamRoleArn *string

	noSmithyDocumentSerde
}

type ListMLModelTrainingJobsOutput struct {

	// A page of the list of model training job IDs.
	Ids []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMLModelTrainingJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListMLModelTrainingJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListMLModelTrainingJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListMLModelTrainingJobs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMLModelTrainingJobs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListMLModelTrainingJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListMLModelTrainingJobs",
	}
}
