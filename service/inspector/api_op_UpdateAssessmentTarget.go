// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the assessment target that is specified by the ARN of the assessment
// target. If resourceGroupArn is not specified, all EC2 instances in the current
// AWS account and region are included in the assessment target.
func (c *Client) UpdateAssessmentTarget(ctx context.Context, params *UpdateAssessmentTargetInput, optFns ...func(*Options)) (*UpdateAssessmentTargetOutput, error) {
	stack := middleware.NewStack("UpdateAssessmentTarget", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateAssessmentTargetMiddlewares(stack)
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
	addOpUpdateAssessmentTargetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAssessmentTarget(options.Region), middleware.Before)
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
			OperationName: "UpdateAssessmentTarget",
			Err:           err,
		}
	}
	out := result.(*UpdateAssessmentTargetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAssessmentTargetInput struct {

	// The ARN of the assessment target that you want to update.
	//
	// This member is required.
	AssessmentTargetArn *string

	// The ARN of the resource group that is used to specify the new resource group to
	// associate with the assessment target.
	ResourceGroupArn *string

	// The name of the assessment target that you want to update.
	//
	// This member is required.
	AssessmentTargetName *string
}

type UpdateAssessmentTargetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateAssessmentTargetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateAssessmentTarget{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateAssessmentTarget{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateAssessmentTarget(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "UpdateAssessmentTarget",
	}
}
