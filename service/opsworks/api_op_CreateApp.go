// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an app for a specified stack. For more information, see Creating Apps
// (https://docs.aws.amazon.com/opsworks/latest/userguide/workingapps-creating.html).
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) CreateApp(ctx context.Context, params *CreateAppInput, optFns ...func(*Options)) (*CreateAppOutput, error) {
	stack := middleware.NewStack("CreateApp", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateAppMiddlewares(stack)
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
	addOpCreateAppValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateApp(options.Region), middleware.Before)
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
			OperationName: "CreateApp",
			Err:           err,
		}
	}
	out := result.(*CreateAppOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAppInput struct {

	// An SslConfiguration object with the SSL configuration.
	SslConfiguration *types.SslConfiguration

	// The app's short name.
	Shortname *string

	// An array of EnvironmentVariable objects that specify environment variables to be
	// associated with the app. After you deploy the app, these variables are defined
	// on the associated app server instance. For more information, see  Environment
	// Variables
	// (https://docs.aws.amazon.com/opsworks/latest/userguide/workingapps-creating.html#workingapps-creating-environment).
	// There is no specific limit on the number of environment variables. However, the
	// size of the associated data structure - which includes the variables' names,
	// values, and protected flag values - cannot exceed 20 KB. This limit should
	// accommodate most if not all use cases. Exceeding it will cause an exception with
	// the message, "Environment: is too large (maximum is 20KB)." If you have
	// specified one or more environment variables, you cannot modify the stack's Chef
	// version.
	Environment []*types.EnvironmentVariable

	// One or more user-defined key/value pairs to be added to the stack attributes.
	Attributes map[string]*string

	// A Source object that specifies the app repository.
	AppSource *types.Source

	// A description of the app.
	Description *string

	// The app's data source.
	DataSources []*types.DataSource

	// The app type. Each supported type is associated with a particular layer. For
	// example, PHP applications are associated with a PHP layer. AWS OpsWorks Stacks
	// deploys an application to those instances that are members of the corresponding
	// layer. If your app isn't one of the standard types, or you prefer to implement
	// your own Deploy recipes, specify other.
	//
	// This member is required.
	Type types.AppType

	// The app virtual host settings, with multiple domains separated by commas. For
	// example: 'www.example.com, example.com'
	Domains []*string

	// The stack ID.
	//
	// This member is required.
	StackId *string

	// The app name.
	//
	// This member is required.
	Name *string

	// Whether to enable SSL for the app.
	EnableSsl *bool
}

// Contains the response to a CreateApp request.
type CreateAppOutput struct {

	// The app ID.
	AppId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateAppMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateApp{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateApp{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateApp(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "CreateApp",
	}
}
