// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use this operation to suppress anomaly detection for a specified anomaly or
// pattern. If you suppress an anomaly, CloudWatch Logs won’t report new
// occurrences of that anomaly and won't update that anomaly with new data. If you
// suppress a pattern, CloudWatch Logs won’t report any anomalies related to that
// pattern.
//
// You must specify either anomalyId or patternId , but you can't specify both
// parameters in the same operation.
//
// If you have previously used this operation to suppress detection of a pattern
// or anomaly, you can use it again to cause CloudWatch Logs to end the
// suppression. To do this, use this operation and specify the anomaly or pattern
// to stop suppressing, and omit the suppressionType and suppressionPeriod
// parameters.
func (c *Client) UpdateAnomaly(ctx context.Context, params *UpdateAnomalyInput, optFns ...func(*Options)) (*UpdateAnomalyOutput, error) {
	if params == nil {
		params = &UpdateAnomalyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAnomaly", params, optFns, c.addOperationUpdateAnomalyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAnomalyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAnomalyInput struct {

	// The ARN of the anomaly detector that this operation is to act on.
	//
	// This member is required.
	AnomalyDetectorArn *string

	// If you are suppressing or unsuppressing an anomaly, specify its unique ID here.
	// You can find anomaly IDs by using the [ListAnomalies]operation.
	//
	// [ListAnomalies]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_ListAnomalies.html
	AnomalyId *string

	// If you are suppressing or unsuppressing an pattern, specify its unique ID here.
	// You can find pattern IDs by using the [ListAnomalies]operation.
	//
	// [ListAnomalies]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_ListAnomalies.html
	PatternId *string

	// If you are temporarily suppressing an anomaly or pattern, use this structure to
	// specify how long the suppression is to last.
	SuppressionPeriod *types.SuppressionPeriod

	// Use this to specify whether the suppression to be temporary or infinite. If you
	// specify LIMITED , you must also specify a suppressionPeriod . If you specify
	// INFINITE , any value for suppressionPeriod is ignored.
	SuppressionType types.SuppressionType

	noSmithyDocumentSerde
}

type UpdateAnomalyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAnomalyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateAnomaly{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateAnomaly{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateAnomaly"); err != nil {
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
	if err = addOpUpdateAnomalyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAnomaly(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAnomaly(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateAnomaly",
	}
}
