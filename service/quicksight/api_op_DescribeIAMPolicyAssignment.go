// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes an existing IAM policy assignment, as specified by the assignment
// name.
func (c *Client) DescribeIAMPolicyAssignment(ctx context.Context, params *DescribeIAMPolicyAssignmentInput, optFns ...func(*Options)) (*DescribeIAMPolicyAssignmentOutput, error) {
	stack := middleware.NewStack("DescribeIAMPolicyAssignment", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeIAMPolicyAssignmentMiddlewares(stack)
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
	addOpDescribeIAMPolicyAssignmentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeIAMPolicyAssignment(options.Region), middleware.Before)
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
			OperationName: "DescribeIAMPolicyAssignment",
			Err:           err,
		}
	}
	out := result.(*DescribeIAMPolicyAssignmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeIAMPolicyAssignmentInput struct {

	// The ID of the AWS account that contains the assignment that you want to
	// describe.
	//
	// This member is required.
	AwsAccountId *string

	// The namespace that contains the assignment.
	//
	// This member is required.
	Namespace *string

	// The name of the assignment.
	//
	// This member is required.
	AssignmentName *string
}

type DescribeIAMPolicyAssignmentOutput struct {

	// The AWS request ID for this operation.
	RequestId *string

	// Information describing the IAM policy assignment.
	IAMPolicyAssignment *types.IAMPolicyAssignment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeIAMPolicyAssignmentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeIAMPolicyAssignment{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeIAMPolicyAssignment{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeIAMPolicyAssignment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DescribeIAMPolicyAssignment",
	}
}
