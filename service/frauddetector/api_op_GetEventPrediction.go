// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Evaluates an event against a detector version. If a version ID is not provided,
// the detector’s ( ACTIVE ) version is used.
func (c *Client) GetEventPrediction(ctx context.Context, params *GetEventPredictionInput, optFns ...func(*Options)) (*GetEventPredictionOutput, error) {
	if params == nil {
		params = &GetEventPredictionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEventPrediction", params, optFns, c.addOperationGetEventPredictionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEventPredictionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEventPredictionInput struct {

	// The detector ID.
	//
	// This member is required.
	DetectorId *string

	// The entity type (associated with the detector's event type) and specific entity
	// ID representing who performed the event. If an entity id is not available, use
	// "UNKNOWN."
	//
	// This member is required.
	Entities []types.Entity

	// The unique ID used to identify the event.
	//
	// This member is required.
	EventId *string

	// Timestamp that defines when the event under evaluation occurred. The timestamp
	// must be specified using ISO 8601 standard in UTC.
	//
	// This member is required.
	EventTimestamp *string

	// The event type associated with the detector specified for the prediction.
	//
	// This member is required.
	EventTypeName *string

	// Names of the event type's variables you defined in Amazon Fraud Detector to
	// represent data elements and their corresponding values for the event you are
	// sending for evaluation. You must provide at least one eventVariable To ensure
	// most accurate fraud prediction and to simplify your data preparation, Amazon
	// Fraud Detector will replace all missing variables or values as follows: For
	// Amazon Fraud Detector trained models: If a null value is provided explicitly for
	// a variable or if a variable is missing, model will replace the null value or the
	// missing variable (no variable name in the eventVariables map) with calculated
	// default mean/medians for numeric variables and with special values for
	// categorical variables. For imported SageMaker models: If a null value is
	// provided explicitly for a variable, the model and rules will use “null” as the
	// value. If a variable is not provided (no variable name in the eventVariables
	// map), model and rules will use the default value that is provided for the
	// variable.
	//
	// This member is required.
	EventVariables map[string]string

	// The detector version ID.
	DetectorVersionId *string

	// The Amazon SageMaker model endpoint input data blobs.
	ExternalModelEndpointDataBlobs map[string]types.ModelEndpointDataBlob

	noSmithyDocumentSerde
}

type GetEventPredictionOutput struct {

	// The model scores for Amazon SageMaker models.
	ExternalModelOutputs []types.ExternalModelOutputs

	// The model scores. Amazon Fraud Detector generates model scores between 0 and
	// 1000, where 0 is low fraud risk and 1000 is high fraud risk. Model scores are
	// directly related to the false positive rate (FPR). For example, a score of 600
	// corresponds to an estimated 10% false positive rate whereas a score of 900
	// corresponds to an estimated 2% false positive rate.
	ModelScores []types.ModelScores

	// The results from the rules.
	RuleResults []types.RuleResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEventPredictionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetEventPrediction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetEventPrediction{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetEventPrediction"); err != nil {
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
	if err = addOpGetEventPredictionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEventPrediction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetEventPrediction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetEventPrediction",
	}
}
