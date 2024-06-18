// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update a review template answer.
func (c *Client) UpdateReviewTemplateAnswer(ctx context.Context, params *UpdateReviewTemplateAnswerInput, optFns ...func(*Options)) (*UpdateReviewTemplateAnswerOutput, error) {
	if params == nil {
		params = &UpdateReviewTemplateAnswerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateReviewTemplateAnswer", params, optFns, c.addOperationUpdateReviewTemplateAnswerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateReviewTemplateAnswerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateReviewTemplateAnswerInput struct {

	// The alias of the lens.
	//
	// For Amazon Web Services official lenses, this is either the lens alias, such as
	// serverless , or the lens ARN, such as
	// arn:aws:wellarchitected:us-east-1::lens/serverless . Note that some operations
	// (such as ExportLens and CreateLensShare) are not permitted on Amazon Web
	// Services official lenses.
	//
	// For custom lenses, this is the lens ARN, such as
	// arn:aws:wellarchitected:us-west-2:123456789012:lens/0123456789abcdef01234567890abcdef
	// .
	//
	// Each lens is identified by its LensSummary$LensAlias.
	//
	// This member is required.
	LensAlias *string

	// The ID of the question.
	//
	// This member is required.
	QuestionId *string

	// The review template ARN.
	//
	// This member is required.
	TemplateArn *string

	// A list of choices to be updated.
	ChoiceUpdates map[string]types.ChoiceUpdate

	// Defines whether this question is applicable to a lens review.
	IsApplicable *bool

	// The notes associated with the workload.
	//
	// For a review template, these are the notes that will be associated with the
	// workload when the template is applied.
	Notes *string

	// The update reason.
	Reason types.AnswerReason

	// List of selected choice IDs in a question answer.
	//
	// The values entered replace the previously selected choices.
	SelectedChoices []string

	noSmithyDocumentSerde
}

type UpdateReviewTemplateAnswerOutput struct {

	// An answer of the question.
	Answer *types.ReviewTemplateAnswer

	// The alias of the lens.
	//
	// For Amazon Web Services official lenses, this is either the lens alias, such as
	// serverless , or the lens ARN, such as
	// arn:aws:wellarchitected:us-east-1::lens/serverless . Note that some operations
	// (such as ExportLens and CreateLensShare) are not permitted on Amazon Web
	// Services official lenses.
	//
	// For custom lenses, this is the lens ARN, such as
	// arn:aws:wellarchitected:us-west-2:123456789012:lens/0123456789abcdef01234567890abcdef
	// .
	//
	// Each lens is identified by its LensSummary$LensAlias.
	LensAlias *string

	// The review template ARN.
	TemplateArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateReviewTemplateAnswerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateReviewTemplateAnswer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateReviewTemplateAnswer{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateReviewTemplateAnswer"); err != nil {
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
	if err = addOpUpdateReviewTemplateAnswerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateReviewTemplateAnswer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateReviewTemplateAnswer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateReviewTemplateAnswer",
	}
}
