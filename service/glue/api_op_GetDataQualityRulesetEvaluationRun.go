// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves a specific run where a ruleset is evaluated against a data source.
func (c *Client) GetDataQualityRulesetEvaluationRun(ctx context.Context, params *GetDataQualityRulesetEvaluationRunInput, optFns ...func(*Options)) (*GetDataQualityRulesetEvaluationRunOutput, error) {
	if params == nil {
		params = &GetDataQualityRulesetEvaluationRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDataQualityRulesetEvaluationRun", params, optFns, c.addOperationGetDataQualityRulesetEvaluationRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDataQualityRulesetEvaluationRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDataQualityRulesetEvaluationRunInput struct {

	// The unique run identifier associated with this run.
	//
	// This member is required.
	RunId *string

	noSmithyDocumentSerde
}

type GetDataQualityRulesetEvaluationRunOutput struct {

	// A map of reference strings to additional data sources you can specify for an
	// evaluation run.
	AdditionalDataSources map[string]types.DataSource

	// Additional run options you can specify for an evaluation run.
	AdditionalRunOptions *types.DataQualityEvaluationRunAdditionalRunOptions

	// The date and time when this run was completed.
	CompletedOn *time.Time

	// The data source (an Glue table) associated with this evaluation run.
	DataSource *types.DataSource

	// The error strings that are associated with the run.
	ErrorString *string

	// The amount of time (in seconds) that the run consumed resources.
	ExecutionTime int32

	// A timestamp. The last point in time when this data quality rule recommendation
	// run was modified.
	LastModifiedOn *time.Time

	// The number of G.1X workers to be used in the run. The default is 5.
	NumberOfWorkers *int32

	// A list of result IDs for the data quality results for the run.
	ResultIds []string

	// An IAM role supplied to encrypt the results of the run.
	Role *string

	// A list of ruleset names for the run. Currently, this parameter takes only one
	// Ruleset name.
	RulesetNames []string

	// The unique run identifier associated with this run.
	RunId *string

	// The date and time when this run started.
	StartedOn *time.Time

	// The status for this run.
	Status types.TaskStatusType

	// The timeout for a run in minutes. This is the maximum time that a run can
	// consume resources before it is terminated and enters TIMEOUT status. The
	// default is 2,880 minutes (48 hours).
	Timeout *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDataQualityRulesetEvaluationRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDataQualityRulesetEvaluationRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDataQualityRulesetEvaluationRun{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetDataQualityRulesetEvaluationRun"); err != nil {
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
	if err = addOpGetDataQualityRulesetEvaluationRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDataQualityRulesetEvaluationRun(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDataQualityRulesetEvaluationRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetDataQualityRulesetEvaluationRun",
	}
}
