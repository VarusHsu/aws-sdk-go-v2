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

// Deletes the assessment run that is specified by the ARN of the assessment run.
func (c *Client) DeleteAssessmentRun(ctx context.Context, params *DeleteAssessmentRunInput, optFns ...func(*Options)) (*DeleteAssessmentRunOutput, error) {
	stack := middleware.NewStack("DeleteAssessmentRun", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteAssessmentRunMiddlewares(stack)
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
	addOpDeleteAssessmentRunValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAssessmentRun(options.Region), middleware.Before)
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
			OperationName: "DeleteAssessmentRun",
			Err:           err,
		}
	}
	out := result.(*DeleteAssessmentRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAssessmentRunInput struct {

	// The ARN that specifies the assessment run that you want to delete.
	//
	// This member is required.
	AssessmentRunArn *string
}

type DeleteAssessmentRunOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteAssessmentRunMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteAssessmentRun{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteAssessmentRun{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteAssessmentRun(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "DeleteAssessmentRun",
	}
}
