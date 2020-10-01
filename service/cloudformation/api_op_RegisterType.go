// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Registers a type with the CloudFormation service. Registering a type makes it
// available for use in CloudFormation templates in your AWS account, and
// includes:
//
//     * Validating the resource schema
//
//     * Determining which
// handlers have been specified for the resource
//
//     * Making the resource type
// available for use in your account
//
// For more information on how to develop types
// and ready them for registeration, see Creating Resource Providers
// (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-types.html)
// in the CloudFormation CLI User Guide. You can have a maximum of 50 resource type
// versions registered at a time. This maximum is per account and per region. Use
// DeregisterType () to deregister specific resource type versions if necessary.
// Once you have initiated a registration request using RegisterType (), you can
// use DescribeTypeRegistration () to monitor the progress of the registration
// request.
func (c *Client) RegisterType(ctx context.Context, params *RegisterTypeInput, optFns ...func(*Options)) (*RegisterTypeOutput, error) {
	stack := middleware.NewStack("RegisterType", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpRegisterTypeMiddlewares(stack)
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
	addOpRegisterTypeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterType(options.Region), middleware.Before)
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
			OperationName: "RegisterType",
			Err:           err,
		}
	}
	out := result.(*RegisterTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterTypeInput struct {

	// The Amazon Resource Name (ARN) of the IAM role for CloudFormation to assume when
	// invoking the resource provider. If your resource type calls AWS APIs in any of
	// its handlers, you must create an IAM execution role
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html) that includes
	// the necessary permissions to call those AWS APIs, and provision that execution
	// role in your account. When CloudFormation needs to invoke the resource provider
	// handler, CloudFormation assumes this execution role to create a temporary
	// session token, which it then passes to the resource provider handler, thereby
	// supplying your resource provider with the appropriate credentials.
	ExecutionRoleArn *string

	// A unique identifier that acts as an idempotency key for this registration
	// request. Specifying a client request token prevents CloudFormation from
	// generating more than one version of a type from the same registeration request,
	// even if the request is submitted multiple times.
	ClientRequestToken *string

	// The name of the type being registered. We recommend that type names adhere to
	// the following pattern: company_or_organization::service::type. The following
	// organization namespaces are reserved and cannot be used in your resource type
	// names:
	//
	//     * Alexa
	//
	//     * AMZN
	//
	//     * Amazon
	//
	//     * AWS
	//
	//     * Custom
	//
	//     *
	// Dev
	//
	// This member is required.
	TypeName *string

	// Specifies logging configuration information for a type.
	LoggingConfig *types.LoggingConfig

	// The kind of type. Currently, the only valid value is RESOURCE.
	Type types.RegistryType

	// A url to the S3 bucket containing the schema handler package that contains the
	// schema, event handlers, and associated files for the type you want to register.
	// For information on generating a schema handler package for the type you want to
	// register, see submit
	// (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-cli-submit.html)
	// in the CloudFormation CLI User Guide. As part of registering a resource provider
	// type, CloudFormation must be able to access the S3 bucket which contains the
	// schema handler package for that resource provider. For more information, see IAM
	// Permissions for Registering a Resource Provider
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry.html#registry-register-permissions)
	// in the AWS CloudFormation User Guide.
	//
	// This member is required.
	SchemaHandlerPackage *string
}

type RegisterTypeOutput struct {

	// The identifier for this registration request. Use this registration token when
	// calling DescribeTypeRegistration (), which returns information about the status
	// and IDs of the type registration.
	RegistrationToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpRegisterTypeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpRegisterType{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpRegisterType{}, middleware.After)
}

func newServiceMetadataMiddleware_opRegisterType(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "RegisterType",
	}
}
