// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appconfig/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieve information about a configuration profile.
func (c *Client) GetConfigurationProfile(ctx context.Context, params *GetConfigurationProfileInput, optFns ...func(*Options)) (*GetConfigurationProfileOutput, error) {
	stack := middleware.NewStack("GetConfigurationProfile", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetConfigurationProfileMiddlewares(stack)
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
	addOpGetConfigurationProfileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetConfigurationProfile(options.Region), middleware.Before)
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
			OperationName: "GetConfigurationProfile",
			Err:           err,
		}
	}
	out := result.(*GetConfigurationProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetConfigurationProfileInput struct {

	// The ID of the configuration profile you want to get.
	//
	// This member is required.
	ConfigurationProfileId *string

	// The ID of the application that includes the configuration profile you want to
	// get.
	//
	// This member is required.
	ApplicationId *string
}

type GetConfigurationProfileOutput struct {

	// The application ID.
	ApplicationId *string

	// The name of the configuration profile.
	Name *string

	// A list of methods for validating the configuration.
	Validators []*types.Validator

	// The ARN of an IAM role with permission to access the configuration at the
	// specified LocationUri.
	RetrievalRoleArn *string

	// The URI location of the configuration.
	LocationUri *string

	// The configuration profile ID.
	Id *string

	// The configuration profile description.
	Description *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetConfigurationProfileMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetConfigurationProfile{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetConfigurationProfile{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetConfigurationProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "GetConfigurationProfile",
	}
}
