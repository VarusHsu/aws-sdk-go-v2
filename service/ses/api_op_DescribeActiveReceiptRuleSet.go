// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the metadata and receipt rules for the receipt rule set that is
// currently active. For information about setting up receipt rule sets, see the
// Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-receipt-rule-set.html).
// You can execute this operation no more than once per second.
func (c *Client) DescribeActiveReceiptRuleSet(ctx context.Context, params *DescribeActiveReceiptRuleSetInput, optFns ...func(*Options)) (*DescribeActiveReceiptRuleSetOutput, error) {
	stack := middleware.NewStack("DescribeActiveReceiptRuleSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeActiveReceiptRuleSetMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeActiveReceiptRuleSet(options.Region), middleware.Before)
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
			OperationName: "DescribeActiveReceiptRuleSet",
			Err:           err,
		}
	}
	out := result.(*DescribeActiveReceiptRuleSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to return the metadata and receipt rules for the receipt
// rule set that is currently active. You use receipt rule sets to receive email
// with Amazon SES. For more information, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-concepts.html).
type DescribeActiveReceiptRuleSetInput struct {
}

// Represents the metadata and receipt rules for the receipt rule set that is
// currently active.
type DescribeActiveReceiptRuleSetOutput struct {

	// The metadata for the currently active receipt rule set. The metadata consists of
	// the rule set name and a timestamp of when the rule set was created.
	Metadata *types.ReceiptRuleSetMetadata

	// The receipt rules that belong to the active rule set.
	Rules []*types.ReceiptRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeActiveReceiptRuleSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeActiveReceiptRuleSet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeActiveReceiptRuleSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeActiveReceiptRuleSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "DescribeActiveReceiptRuleSet",
	}
}
