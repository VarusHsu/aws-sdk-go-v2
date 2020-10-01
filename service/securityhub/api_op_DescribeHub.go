// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns details about the Hub resource in your account, including the HubArn and
// the time when you enabled Security Hub.
func (c *Client) DescribeHub(ctx context.Context, params *DescribeHubInput, optFns ...func(*Options)) (*DescribeHubOutput, error) {
	stack := middleware.NewStack("DescribeHub", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeHubMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeHub(options.Region), middleware.Before)
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
			OperationName: "DescribeHub",
			Err:           err,
		}
	}
	out := result.(*DescribeHubOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeHubInput struct {

	// The ARN of the Hub resource to retrieve.
	HubArn *string
}

type DescribeHubOutput struct {

	// Whether to automatically enable new controls when they are added to standards
	// that are enabled. If set to true, then new controls for enabled standards are
	// enabled automatically. If set to false, then new controls are not enabled.
	AutoEnableControls *bool

	// The ARN of the Hub resource that was retrieved.
	HubArn *string

	// The date and time when Security Hub was enabled in the account.
	SubscribedAt *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeHubMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeHub{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeHub{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeHub(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "DescribeHub",
	}
}
