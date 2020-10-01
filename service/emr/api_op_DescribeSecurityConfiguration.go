// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

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

// Provides the details of a security configuration by returning the configuration
// JSON.
func (c *Client) DescribeSecurityConfiguration(ctx context.Context, params *DescribeSecurityConfigurationInput, optFns ...func(*Options)) (*DescribeSecurityConfigurationOutput, error) {
	stack := middleware.NewStack("DescribeSecurityConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeSecurityConfigurationMiddlewares(stack)
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
	addOpDescribeSecurityConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSecurityConfiguration(options.Region), middleware.Before)
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
			OperationName: "DescribeSecurityConfiguration",
			Err:           err,
		}
	}
	out := result.(*DescribeSecurityConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSecurityConfigurationInput struct {

	// The name of the security configuration.
	//
	// This member is required.
	Name *string
}

type DescribeSecurityConfigurationOutput struct {

	// The name of the security configuration.
	Name *string

	// The security configuration details in JSON format.
	SecurityConfiguration *string

	// The date and time the security configuration was created
	CreationDateTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeSecurityConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeSecurityConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeSecurityConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeSecurityConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "DescribeSecurityConfiguration",
	}
}
