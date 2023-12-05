// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the result of a data quality rule evaluation.
func (c *Client) GetDataQualityResult(ctx context.Context, params *GetDataQualityResultInput, optFns ...func(*Options)) (*GetDataQualityResultOutput, error) {
	if params == nil {
		params = &GetDataQualityResultInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDataQualityResult", params, optFns, c.addOperationGetDataQualityResultMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDataQualityResultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDataQualityResultInput struct {

	// A unique result ID for the data quality result.
	//
	// This member is required.
	ResultId *string

	noSmithyDocumentSerde
}

type GetDataQualityResultOutput struct {

	// A list of DataQualityAnalyzerResult objects representing the results for each
	// analyzer.
	AnalyzerResults []types.DataQualityAnalyzerResult

	// The date and time when the run for this data quality result was completed.
	CompletedOn *time.Time

	// The table associated with the data quality result, if any.
	DataSource *types.DataSource

	// In the context of a job in Glue Studio, each node in the canvas is typically
	// assigned some sort of name and data quality nodes will have names. In the case
	// of multiple nodes, the evaluationContext can differentiate the nodes.
	EvaluationContext *string

	// The job name associated with the data quality result, if any.
	JobName *string

	// The job run ID associated with the data quality result, if any.
	JobRunId *string

	// A list of DataQualityObservation objects representing the observations
	// generated after evaluating the rules and analyzers.
	Observations []types.DataQualityObservation

	// A unique result ID for the data quality result.
	ResultId *string

	// A list of DataQualityRuleResult objects representing the results for each rule.
	RuleResults []types.DataQualityRuleResult

	// The unique run ID associated with the ruleset evaluation.
	RulesetEvaluationRunId *string

	// The name of the ruleset associated with the data quality result.
	RulesetName *string

	// An aggregate data quality score. Represents the ratio of rules that passed to
	// the total number of rules.
	Score *float64

	// The date and time when the run for this data quality result started.
	StartedOn *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDataQualityResultMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDataQualityResult{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDataQualityResult{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetDataQualityResult"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpGetDataQualityResultValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDataQualityResult(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opGetDataQualityResult(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetDataQualityResult",
	}
}
