// Code generated by smithy-go-codegen DO NOT EDIT.

package mq

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns the specified configuration revision for the specified configuration.
func (c *Client) DescribeConfigurationRevision(ctx context.Context, params *DescribeConfigurationRevisionInput, optFns ...func(*Options)) (*DescribeConfigurationRevisionOutput, error) {
	stack := middleware.NewStack("DescribeConfigurationRevision", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeConfigurationRevisionMiddlewares(stack)
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
	addOpDescribeConfigurationRevisionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConfigurationRevision(options.Region), middleware.Before)
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
			OperationName: "DescribeConfigurationRevision",
			Err:           err,
		}
	}
	out := result.(*DescribeConfigurationRevisionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConfigurationRevisionInput struct {

	// The revision of the configuration.
	//
	// This member is required.
	ConfigurationRevision *string

	// The unique ID that Amazon MQ generates for the configuration.
	//
	// This member is required.
	ConfigurationId *string
}

type DescribeConfigurationRevisionOutput struct {

	// Required. The unique ID that Amazon MQ generates for the configuration.
	ConfigurationId *string

	// Required. The base64-encoded XML configuration.
	Data *string

	// The description of the configuration.
	Description *string

	// Required. The date and time of the configuration.
	Created *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeConfigurationRevisionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeConfigurationRevision{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeConfigurationRevision{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeConfigurationRevision(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mq",
		OperationName: "DescribeConfigurationRevision",
	}
}
