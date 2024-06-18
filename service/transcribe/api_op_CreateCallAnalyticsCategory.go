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

// Creates a new Call Analytics category.
//
// All categories are automatically applied to your Call Analytics transcriptions.
// Note that in order to apply categories to your transcriptions, you must create
// them before submitting your transcription request, as categories cannot be
// applied retroactively.
//
// When creating a new category, you can use the InputType parameter to label the
// category as a POST_CALL or a REAL_TIME category. POST_CALL categories can only
// be applied to post-call transcriptions and REAL_TIME categories can only be
// applied to real-time transcriptions. If you do not include InputType , your
// category is created as a POST_CALL category by default.
//
// Call Analytics categories are composed of rules. For each category, you must
// create between 1 and 20 rules. Rules can include these parameters: , , , and .
//
// To update an existing category, see .
//
// To learn more about Call Analytics categories, see [Creating categories for post-call transcriptions] and [Creating categories for real-time transcriptions].
//
// [Creating categories for post-call transcriptions]: https://docs.aws.amazon.com/transcribe/latest/dg/tca-categories-batch.html
// [Creating categories for real-time transcriptions]: https://docs.aws.amazon.com/transcribe/latest/dg/tca-categories-stream.html
func (c *Client) CreateCallAnalyticsCategory(ctx context.Context, params *CreateCallAnalyticsCategoryInput, optFns ...func(*Options)) (*CreateCallAnalyticsCategoryOutput, error) {
	if params == nil {
		params = &CreateCallAnalyticsCategoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCallAnalyticsCategory", params, optFns, c.addOperationCreateCallAnalyticsCategoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCallAnalyticsCategoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCallAnalyticsCategoryInput struct {

	// A unique name, chosen by you, for your Call Analytics category. It's helpful to
	// use a detailed naming system that will make sense to you in the future. For
	// example, it's better to use sentiment-positive-last30seconds for a category
	// over a generic name like test-category .
	//
	// Category names are case sensitive.
	//
	// This member is required.
	CategoryName *string

	// Rules define a Call Analytics category. When creating a new category, you must
	// create between 1 and 20 rules for that category. For each rule, you specify a
	// filter you want applied to the attributes of a call. For example, you can choose
	// a sentiment filter that detects if a customer's sentiment was positive during
	// the last 30 seconds of the call.
	//
	// This member is required.
	Rules []types.Rule

	// Choose whether you want to create a real-time or a post-call category for your
	// Call Analytics transcription.
	//
	// Specifying POST_CALL assigns your category to post-call transcriptions;
	// categories with this input type cannot be applied to streaming (real-time)
	// transcriptions.
	//
	// Specifying REAL_TIME assigns your category to streaming transcriptions;
	// categories with this input type cannot be applied to post-call transcriptions.
	//
	// If you do not include InputType , your category is created as a post-call
	// category by default.
	InputType types.InputType

	noSmithyDocumentSerde
}

type CreateCallAnalyticsCategoryOutput struct {

	// Provides you with the properties of your new category, including its associated
	// rules.
	CategoryProperties *types.CategoryProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCallAnalyticsCategoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCallAnalyticsCategory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCallAnalyticsCategory{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateCallAnalyticsCategory"); err != nil {
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
	if err = addOpCreateCallAnalyticsCategoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCallAnalyticsCategory(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCallAnalyticsCategory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateCallAnalyticsCategory",
	}
}
