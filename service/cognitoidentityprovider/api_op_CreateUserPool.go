// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new Amazon Cognito user pool and sets the password policy for the
// pool.
func (c *Client) CreateUserPool(ctx context.Context, params *CreateUserPoolInput, optFns ...func(*Options)) (*CreateUserPoolOutput, error) {
	stack := middleware.NewStack("CreateUserPool", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateUserPoolMiddlewares(stack)
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
	addOpCreateUserPoolValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUserPool(options.Region), middleware.Before)
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
			OperationName: "CreateUserPool",
			Err:           err,
		}
	}
	out := result.(*CreateUserPoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to create a user pool.
type CreateUserPoolInput struct {

	// The email configuration.
	EmailConfiguration *types.EmailConfigurationType

	// The configuration for AdminCreateUser requests.
	AdminCreateUserConfig *types.AdminCreateUserConfigType

	// A string representing the email verification message.
	EmailVerificationMessage *string

	// The device configuration.
	DeviceConfiguration *types.DeviceConfigurationType

	// The SMS configuration.
	SmsConfiguration *types.SmsConfigurationType

	// The tag keys and values to assign to the user pool. A tag is a label that you
	// can use to categorize and manage user pools in different ways, such as by
	// purpose, owner, environment, or other criteria.
	UserPoolTags map[string]*string

	// A string representing the email verification subject.
	EmailVerificationSubject *string

	// Used to enable advanced security risk detection. Set the key
	// AdvancedSecurityMode to the value "AUDIT".
	UserPoolAddOns *types.UserPoolAddOnsType

	// The policies associated with the new user pool.
	Policies *types.UserPoolPolicyType

	// An array of schema attributes for the new user pool. These attributes can be
	// standard or custom attributes.
	Schema []*types.SchemaAttributeType

	// The attributes to be auto-verified. Possible values: email, phone_number.
	AutoVerifiedAttributes []types.VerifiedAttributeType

	// A string representing the SMS authentication message.
	SmsAuthenticationMessage *string

	// Use this setting to define which verified available method a user can use to
	// recover their password when they call ForgotPassword. It allows you to define a
	// preferred method when a user has more than one method available. With this
	// setting, SMS does not qualify for a valid password recovery mechanism if the
	// user also has SMS MFA enabled. In the absence of this setting, Cognito uses the
	// legacy behavior to determine the recovery method where SMS is preferred over
	// email. Starting February 1, 2020, the value of AccountRecoverySetting will
	// default to verified_email first and verified_phone_number as the second option
	// for newly created user pools if no value is provided.
	AccountRecoverySetting *types.AccountRecoverySettingType

	// Specifies whether email addresses or phone numbers can be specified as usernames
	// when a user signs up.
	UsernameAttributes []types.UsernameAttributeType

	// The Lambda trigger configuration information for the new user pool. In a push
	// model, event sources (such as Amazon S3 and custom applications) need permission
	// to invoke a function. So you will need to make an extra call to add permission
	// for these event sources to invoke your Lambda function. For more information on
	// using the Lambda API to add permission, see  AddPermission
	// (https://docs.aws.amazon.com/lambda/latest/dg/API_AddPermission.html). For
	// adding permission using the AWS CLI, see  add-permission
	// (https://docs.aws.amazon.com/cli/latest/reference/lambda/add-permission.html).
	LambdaConfig *types.LambdaConfigType

	// The template for the verification message that the user sees when the app
	// requests permission to access the user's information.
	VerificationMessageTemplate *types.VerificationMessageTemplateType

	// A string used to name the user pool.
	//
	// This member is required.
	PoolName *string

	// You can choose to set case sensitivity on the username input for the selected
	// sign-in option. For example, when this is set to False, users will be able to
	// sign in using either "username" or "Username". This configuration is immutable
	// once it has been set. For more information, see .
	UsernameConfiguration *types.UsernameConfigurationType

	// Attributes supported as an alias for this user pool. Possible values:
	// phone_number, email, or preferred_username.
	AliasAttributes []types.AliasAttributeType

	// A string representing the SMS verification message.
	SmsVerificationMessage *string

	// Specifies MFA configuration details.
	MfaConfiguration types.UserPoolMfaType
}

// Represents the response from the server for the request to create a user pool.
type CreateUserPoolOutput struct {

	// A container for the user pool details.
	UserPool *types.UserPoolType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateUserPoolMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateUserPool{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateUserPool{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateUserPool(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "CreateUserPool",
	}
}
