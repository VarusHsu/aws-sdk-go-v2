// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Receive information about a configuration. AWS AppConfig uses the value of the
// ClientConfigurationVersion parameter to identify the configuration version on
// your clients. If you don’t send ClientConfigurationVersion with each call to
// GetConfiguration, your clients receive the current configuration. You are
// charged each time your clients receive a configuration. To avoid excess charges,
// we recommend that you include the ClientConfigurationVersion value with every
// call to GetConfiguration. This value must be saved on your client. Subsequent
// calls to GetConfiguration must pass this value by using the
// ClientConfigurationVersion parameter.
func (c *Client) GetConfiguration(ctx context.Context, params *GetConfigurationInput, optFns ...func(*Options)) (*GetConfigurationOutput, error) {
	stack := middleware.NewStack("GetConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetConfigurationMiddlewares(stack)
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
	addOpGetConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetConfiguration(options.Region), middleware.Before)
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
			OperationName: "GetConfiguration",
			Err:           err,
		}
	}
	out := result.(*GetConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetConfigurationInput struct {

	// The environment to get. Specify either the environment name or the environment
	// ID.
	//
	// This member is required.
	Environment *string

	// The configuration to get. Specify either the configuration name or the
	// configuration ID.
	//
	// This member is required.
	Configuration *string

	// A unique ID to identify the client for the configuration. This ID enables
	// AppConfig to deploy the configuration in intervals, as defined in the deployment
	// strategy.
	//
	// This member is required.
	ClientId *string

	// The application to get. Specify either the application name or the application
	// ID.
	//
	// This member is required.
	Application *string

	// The configuration version returned in the most recent GetConfiguration response.
	// AWS AppConfig uses the value of the ClientConfigurationVersion parameter to
	// identify the configuration version on your clients. If you don’t send
	// ClientConfigurationVersion with each call to GetConfiguration, your clients
	// receive the current configuration. You are charged each time your clients
	// receive a configuration. To avoid excess charges, we recommend that you include
	// the ClientConfigurationVersion value with every call to GetConfiguration. This
	// value must be saved on your client. Subsequent calls to GetConfiguration must
	// pass this value by using the ClientConfigurationVersion parameter. For more
	// information about working with configurations, see Retrieving the Configuration
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/appconfig-retrieving-the-configuration.html)
	// in the AWS AppConfig User Guide.
	ClientConfigurationVersion *string
}

type GetConfigurationOutput struct {

	// The content of the configuration or the configuration data.
	Content []byte

	// A standard MIME type describing the format of the configuration content. For
	// more information, see Content-Type
	// (http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17).
	ContentType *string

	// The configuration version.
	ConfigurationVersion *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "GetConfiguration",
	}
}
