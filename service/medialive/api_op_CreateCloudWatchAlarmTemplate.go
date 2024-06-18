// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a cloudwatch alarm template to dynamically generate cloudwatch metric
// alarms on targeted resource types.
func (c *Client) CreateCloudWatchAlarmTemplate(ctx context.Context, params *CreateCloudWatchAlarmTemplateInput, optFns ...func(*Options)) (*CreateCloudWatchAlarmTemplateOutput, error) {
	if params == nil {
		params = &CreateCloudWatchAlarmTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCloudWatchAlarmTemplate", params, optFns, c.addOperationCreateCloudWatchAlarmTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCloudWatchAlarmTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for CreateCloudWatchAlarmTemplateRequest
type CreateCloudWatchAlarmTemplateInput struct {

	// The comparison operator used to compare the specified statistic and the
	// threshold.
	//
	// This member is required.
	ComparisonOperator types.CloudWatchAlarmTemplateComparisonOperator

	// The number of periods over which data is compared to the specified threshold.
	//
	// This member is required.
	EvaluationPeriods *int32

	// A cloudwatch alarm template group's identifier. Can be either be its id or
	// current name.
	//
	// This member is required.
	GroupIdentifier *string

	// The name of the metric associated with the alarm. Must be compatible with
	// targetResourceType.
	//
	// This member is required.
	MetricName *string

	// A resource's name. Names must be unique within the scope of a resource type in
	// a specific region.
	//
	// This member is required.
	Name *string

	// The period, in seconds, over which the specified statistic is applied.
	//
	// This member is required.
	Period *int32

	// The statistic to apply to the alarm's metric data.
	//
	// This member is required.
	Statistic types.CloudWatchAlarmTemplateStatistic

	// The resource type this template should dynamically generate cloudwatch metric
	// alarms for.
	//
	// This member is required.
	TargetResourceType types.CloudWatchAlarmTemplateTargetResourceType

	// The threshold value to compare with the specified statistic.
	//
	// This member is required.
	Threshold *float64

	// Specifies how missing data points are treated when evaluating the alarm's
	// condition.
	//
	// This member is required.
	TreatMissingData types.CloudWatchAlarmTemplateTreatMissingData

	// The number of datapoints within the evaluation period that must be breaching to
	// trigger the alarm.
	DatapointsToAlarm *int32

	// A resource's optional description.
	Description *string

	// Represents the tags associated with a resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Placeholder documentation for CreateCloudWatchAlarmTemplateResponse
type CreateCloudWatchAlarmTemplateOutput struct {

	// A cloudwatch alarm template's ARN (Amazon Resource Name)
	Arn *string

	// The comparison operator used to compare the specified statistic and the
	// threshold.
	ComparisonOperator types.CloudWatchAlarmTemplateComparisonOperator

	// Placeholder documentation for __timestampIso8601
	CreatedAt *time.Time

	// The number of datapoints within the evaluation period that must be breaching to
	// trigger the alarm.
	DatapointsToAlarm *int32

	// A resource's optional description.
	Description *string

	// The number of periods over which data is compared to the specified threshold.
	EvaluationPeriods *int32

	// A cloudwatch alarm template group's id. AWS provided template groups have ids
	// that start with aws-
	GroupId *string

	// A cloudwatch alarm template's id. AWS provided templates have ids that start
	// with aws-
	Id *string

	// The name of the metric associated with the alarm. Must be compatible with
	// targetResourceType.
	MetricName *string

	// Placeholder documentation for __timestampIso8601
	ModifiedAt *time.Time

	// A resource's name. Names must be unique within the scope of a resource type in
	// a specific region.
	Name *string

	// The period, in seconds, over which the specified statistic is applied.
	Period *int32

	// The statistic to apply to the alarm's metric data.
	Statistic types.CloudWatchAlarmTemplateStatistic

	// Represents the tags associated with a resource.
	Tags map[string]string

	// The resource type this template should dynamically generate cloudwatch metric
	// alarms for.
	TargetResourceType types.CloudWatchAlarmTemplateTargetResourceType

	// The threshold value to compare with the specified statistic.
	Threshold *float64

	// Specifies how missing data points are treated when evaluating the alarm's
	// condition.
	TreatMissingData types.CloudWatchAlarmTemplateTreatMissingData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCloudWatchAlarmTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateCloudWatchAlarmTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateCloudWatchAlarmTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateCloudWatchAlarmTemplate"); err != nil {
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
	if err = addOpCreateCloudWatchAlarmTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCloudWatchAlarmTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCloudWatchAlarmTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateCloudWatchAlarmTemplate",
	}
}
