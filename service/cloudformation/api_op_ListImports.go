// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all stacks that are importing an exported output value. To modify or
// remove an exported output value, first use this action to see which stacks are
// using it. To see the exported output values in your account, see ListExports ().
// For more information about importing an exported output value, see the
// Fn::ImportValue
// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-importvalue.html)
// function.
func (c *Client) ListImports(ctx context.Context, params *ListImportsInput, optFns ...func(*Options)) (*ListImportsOutput, error) {
	stack := middleware.NewStack("ListImports", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpListImportsMiddlewares(stack)
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
	addOpListImportsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListImports(options.Region), middleware.Before)
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
			OperationName: "ListImports",
			Err:           err,
		}
	}
	out := result.(*ListImportsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListImportsInput struct {

	// The name of the exported output value. AWS CloudFormation returns the stack
	// names that are importing this value.
	//
	// This member is required.
	ExportName *string

	// A string (provided by the ListImports () response output) that identifies the
	// next page of stacks that are importing the specified exported output value.
	NextToken *string
}

type ListImportsOutput struct {

	// A string that identifies the next page of exports. If there is no additional
	// page, this value is null.
	NextToken *string

	// A list of stack names that are importing the specified exported output value.
	Imports []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpListImportsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpListImports{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpListImports{}, middleware.After)
}

func newServiceMetadataMiddleware_opListImports(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "ListImports",
	}
}
