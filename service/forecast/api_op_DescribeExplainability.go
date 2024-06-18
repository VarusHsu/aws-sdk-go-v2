// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes an Explainability resource created using the CreateExplainability operation.
func (c *Client) DescribeExplainability(ctx context.Context, params *DescribeExplainabilityInput, optFns ...func(*Options)) (*DescribeExplainabilityOutput, error) {
	if params == nil {
		params = &DescribeExplainabilityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeExplainability", params, optFns, c.addOperationDescribeExplainabilityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeExplainabilityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeExplainabilityInput struct {

	// The Amazon Resource Name (ARN) of the Explaianability to describe.
	//
	// This member is required.
	ExplainabilityArn *string

	noSmithyDocumentSerde
}

type DescribeExplainabilityOutput struct {

	// When the Explainability resource was created.
	CreationTime *time.Time

	// The source of your data, an Identity and Access Management (IAM) role that
	// allows Amazon Forecast to access the data and, optionally, an Key Management
	// Service (KMS) key.
	DataSource *types.DataSource

	// Whether the visualization was enabled for the Explainability resource.
	EnableVisualization *bool

	// If TimePointGranularity is set to SPECIFIC , the last time point in the
	// Explainability.
	EndDateTime *string

	// The estimated time remaining in minutes for the CreateExplainability job to complete.
	EstimatedTimeRemainingInMinutes *int64

	// The Amazon Resource Name (ARN) of the Explainability.
	ExplainabilityArn *string

	// The configuration settings that define the granularity of time series and time
	// points for the Explainability.
	ExplainabilityConfig *types.ExplainabilityConfig

	// The name of the Explainability.
	ExplainabilityName *string

	// The last time the resource was modified. The timestamp depends on the status of
	// the job:
	//
	//   - CREATE_PENDING - The CreationTime .
	//
	//   - CREATE_IN_PROGRESS - The current timestamp.
	//
	//   - CREATE_STOPPING - The current timestamp.
	//
	//   - CREATE_STOPPED - When the job stopped.
	//
	//   - ACTIVE or CREATE_FAILED - When the job finished or failed.
	LastModificationTime *time.Time

	// If an error occurred, a message about the error.
	Message *string

	// The Amazon Resource Name (ARN) of the Predictor or Forecast used to create the
	// Explainability resource.
	ResourceArn *string

	// Defines the fields of a dataset.
	Schema *types.Schema

	// If TimePointGranularity is set to SPECIFIC , the first time point in the
	// Explainability.
	StartDateTime *string

	// The status of the Explainability resource. States include:
	//
	//   - ACTIVE
	//
	//   - CREATE_PENDING , CREATE_IN_PROGRESS , CREATE_FAILED
	//
	//   - CREATE_STOPPING , CREATE_STOPPED
	//
	//   - DELETE_PENDING , DELETE_IN_PROGRESS , DELETE_FAILED
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeExplainabilityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeExplainability{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeExplainability{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeExplainability"); err != nil {
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
	if err = addOpDescribeExplainabilityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeExplainability(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeExplainability(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeExplainability",
	}
}
