// Code generated by smithy-go-codegen DO NOT EDIT.

package codegurureviewer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codegurureviewer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use to create a code review with a [CodeReviewType] of RepositoryAnalysis . This type of code
// review analyzes all code under a specified branch in an associated repository.
// PullRequest code reviews are automatically triggered by a pull request.
//
// [CodeReviewType]: https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_CodeReviewType.html
func (c *Client) CreateCodeReview(ctx context.Context, params *CreateCodeReviewInput, optFns ...func(*Options)) (*CreateCodeReviewOutput, error) {
	if params == nil {
		params = &CreateCodeReviewInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCodeReview", params, optFns, c.addOperationCreateCodeReviewMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCodeReviewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCodeReviewInput struct {

	// The name of the code review. The name of each code review in your Amazon Web
	// Services account must be unique.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of the [RepositoryAssociation] object. You can retrieve this ARN by
	// calling [ListRepositoryAssociations].
	//
	// A code review can only be created on an associated repository. This is the ARN
	// of the associated repository.
	//
	// [ListRepositoryAssociations]: https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_ListRepositoryAssociations.html
	// [RepositoryAssociation]: https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociation.html
	//
	// This member is required.
	RepositoryAssociationArn *string

	// The type of code review to create. This is specified using a [CodeReviewType] object. You can
	// create a code review only of type RepositoryAnalysis .
	//
	// [CodeReviewType]: https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_CodeReviewType.html
	//
	// This member is required.
	Type *types.CodeReviewType

	// Amazon CodeGuru Reviewer uses this value to prevent the accidental creation of
	// duplicate code reviews if there are failures and retries.
	ClientRequestToken *string

	noSmithyDocumentSerde
}

type CreateCodeReviewOutput struct {

	// Information about a code review. A code review belongs to the associated
	// repository that contains the reviewed code.
	CodeReview *types.CodeReview

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCodeReviewMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateCodeReview{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateCodeReview{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateCodeReview"); err != nil {
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
	if err = addIdempotencyToken_opCreateCodeReviewMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateCodeReviewValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCodeReview(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateCodeReview struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateCodeReview) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateCodeReview) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateCodeReviewInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateCodeReviewInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateCodeReviewMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateCodeReview{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateCodeReview(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateCodeReview",
	}
}
