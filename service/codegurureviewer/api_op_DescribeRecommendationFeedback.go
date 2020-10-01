// Code generated by smithy-go-codegen DO NOT EDIT.

package codegurureviewer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codegurureviewer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the customer feedback for a CodeGuru Reviewer recommendation.
func (c *Client) DescribeRecommendationFeedback(ctx context.Context, params *DescribeRecommendationFeedbackInput, optFns ...func(*Options)) (*DescribeRecommendationFeedbackOutput, error) {
	stack := middleware.NewStack("DescribeRecommendationFeedback", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeRecommendationFeedbackMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDescribeRecommendationFeedbackValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRecommendationFeedback(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "DescribeRecommendationFeedback",
			Err:           err,
		}
	}
	out := result.(*DescribeRecommendationFeedbackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRecommendationFeedbackInput struct {

	// The recommendation ID that can be used to track the provided recommendations and
	// then to collect the feedback.
	//
	// This member is required.
	RecommendationId *string

	// The Amazon Resource Name (ARN) of the CodeReview
	// (https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_CodeReview.html)
	// object.
	//
	// This member is required.
	CodeReviewArn *string

	// Optional parameter to describe the feedback for a given user. If this is not
	// supplied, it defaults to the user making the request. The UserId is an IAM
	// principal that can be specified as an AWS account ID or an Amazon Resource Name
	// (ARN). For more information, see  Specifying a Principal
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html#Principal_specifying)
	// in the AWS Identity and Access Management User Guide.
	UserId *string
}

type DescribeRecommendationFeedbackOutput struct {

	// The recommendation feedback given by the user.
	RecommendationFeedback *types.RecommendationFeedback

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeRecommendationFeedbackMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeRecommendationFeedback{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeRecommendationFeedback{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeRecommendationFeedback(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-reviewer",
		OperationName: "DescribeRecommendationFeedback",
	}
}
