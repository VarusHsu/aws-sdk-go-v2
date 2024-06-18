// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a definition for a job that monitors data quality and drift. For
// information about model monitor, see [Amazon SageMaker Model Monitor].
//
// [Amazon SageMaker Model Monitor]: https://docs.aws.amazon.com/sagemaker/latest/dg/model-monitor.html
func (c *Client) CreateDataQualityJobDefinition(ctx context.Context, params *CreateDataQualityJobDefinitionInput, optFns ...func(*Options)) (*CreateDataQualityJobDefinitionOutput, error) {
	if params == nil {
		params = &CreateDataQualityJobDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDataQualityJobDefinition", params, optFns, c.addOperationCreateDataQualityJobDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDataQualityJobDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDataQualityJobDefinitionInput struct {

	// Specifies the container that runs the monitoring job.
	//
	// This member is required.
	DataQualityAppSpecification *types.DataQualityAppSpecification

	// A list of inputs for the monitoring job. Currently endpoints are supported as
	// monitoring inputs.
	//
	// This member is required.
	DataQualityJobInput *types.DataQualityJobInput

	// The output configuration for monitoring jobs.
	//
	// This member is required.
	DataQualityJobOutputConfig *types.MonitoringOutputConfig

	// The name for the monitoring job definition.
	//
	// This member is required.
	JobDefinitionName *string

	// Identifies the resources to deploy for a monitoring job.
	//
	// This member is required.
	JobResources *types.MonitoringResources

	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume
	// to perform tasks on your behalf.
	//
	// This member is required.
	RoleArn *string

	// Configures the constraints and baselines for the monitoring job.
	DataQualityBaselineConfig *types.DataQualityBaselineConfig

	// Specifies networking configuration for the monitoring job.
	NetworkConfig *types.MonitoringNetworkConfig

	// A time limit for how long the monitoring job is allowed to run before stopping.
	StoppingCondition *types.MonitoringStoppingCondition

	// (Optional) An array of key-value pairs. For more information, see [Using Cost Allocation Tags] in the
	// Amazon Web Services Billing and Cost Management User Guide.
	//
	// [Using Cost Allocation Tags]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-whatURL
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateDataQualityJobDefinitionOutput struct {

	// The Amazon Resource Name (ARN) of the job definition.
	//
	// This member is required.
	JobDefinitionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDataQualityJobDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDataQualityJobDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDataQualityJobDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDataQualityJobDefinition"); err != nil {
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
	if err = addOpCreateDataQualityJobDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataQualityJobDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDataQualityJobDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDataQualityJobDefinition",
	}
}
