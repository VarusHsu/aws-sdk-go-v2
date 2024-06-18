// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This action might generate an SMS text message. Starting June 1, 2021, US
// telecom carriers require you to register an origination phone number before you
// can send SMS messages to US phone numbers. If you use SMS text messages in
// Amazon Cognito, you must register a phone number with [Amazon Pinpoint]. Amazon Cognito uses the
// registered number automatically. Otherwise, Amazon Cognito users who must
// receive SMS messages might not be able to sign up, activate their accounts, or
// sign in.
//
// If you have never used SMS text messages with Amazon Cognito or any other
// Amazon Web Service, Amazon Simple Notification Service might place your account
// in the SMS sandbox. In [sandbox mode], you can send messages only to verified phone numbers.
// After you test your app while in the sandbox environment, you can move out of
// the sandbox and into production. For more information, see [SMS message settings for Amazon Cognito user pools]in the Amazon
// Cognito Developer Guide.
//
// Creates a new Amazon Cognito user pool and sets the password policy for the
// pool.
//
// If you don't provide a value for an attribute, Amazon Cognito sets it to its
// default value.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [SMS message settings for Amazon Cognito user pools]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [sandbox mode]: https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
// [Amazon Pinpoint]: https://console.aws.amazon.com/pinpoint/home/
func (c *Client) CreateUserPool(ctx context.Context, params *CreateUserPoolInput, optFns ...func(*Options)) (*CreateUserPoolOutput, error) {
	if params == nil {
		params = &CreateUserPoolInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateUserPool", params, optFns, c.addOperationCreateUserPoolMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateUserPoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to create a user pool.
type CreateUserPoolInput struct {

	// A string used to name the user pool.
	//
	// This member is required.
	PoolName *string

	// The available verified method a user can use to recover their password when
	// they call ForgotPassword . You can use this setting to define a preferred method
	// when a user has more than one method available. With this setting, SMS doesn't
	// qualify for a valid password recovery mechanism if the user also has SMS
	// multi-factor authentication (MFA) activated. In the absence of this setting,
	// Amazon Cognito uses the legacy behavior to determine the recovery method where
	// SMS is preferred through email.
	AccountRecoverySetting *types.AccountRecoverySettingType

	// The configuration for AdminCreateUser requests.
	AdminCreateUserConfig *types.AdminCreateUserConfigType

	// Attributes supported as an alias for this user pool. Possible values:
	// phone_number, email, or preferred_username.
	AliasAttributes []types.AliasAttributeType

	// The attributes to be auto-verified. Possible values: email, phone_number.
	AutoVerifiedAttributes []types.VerifiedAttributeType

	// When active, DeletionProtection prevents accidental deletion of your user pool.
	// Before you can delete a user pool that you have protected against deletion, you
	// must deactivate this feature.
	//
	// When you try to delete a protected user pool in a DeleteUserPool API request,
	// Amazon Cognito returns an InvalidParameterException error. To delete a
	// protected user pool, send a new DeleteUserPool request after you deactivate
	// deletion protection in an UpdateUserPool API request.
	DeletionProtection types.DeletionProtectionType

	// The device-remembering configuration for a user pool. A null value indicates
	// that you have deactivated device remembering in your user pool.
	//
	// When you provide a value for any DeviceConfiguration field, you activate the
	// Amazon Cognito device-remembering feature.
	DeviceConfiguration *types.DeviceConfigurationType

	// The email configuration of your user pool. The email configuration type sets
	// your preferred sending method, Amazon Web Services Region, and sender for
	// messages from your user pool.
	EmailConfiguration *types.EmailConfigurationType

	// This parameter is no longer used. See [VerificationMessageTemplateType].
	//
	// [VerificationMessageTemplateType]: https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_VerificationMessageTemplateType.html
	EmailVerificationMessage *string

	// This parameter is no longer used. See [VerificationMessageTemplateType].
	//
	// [VerificationMessageTemplateType]: https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_VerificationMessageTemplateType.html
	EmailVerificationSubject *string

	// The Lambda trigger configuration information for the new user pool.
	//
	// In a push model, event sources (such as Amazon S3 and custom applications) need
	// permission to invoke a function. So you must make an extra call to add
	// permission for these event sources to invoke your Lambda function.
	//
	// For more information on using the Lambda API to add permission, see[AddPermission] .
	//
	// For adding permission using the CLI, see[add-permission] .
	//
	// [AddPermission]: https://docs.aws.amazon.com/lambda/latest/dg/API_AddPermission.html
	// [add-permission]: https://docs.aws.amazon.com/cli/latest/reference/lambda/add-permission.html
	LambdaConfig *types.LambdaConfigType

	// Specifies MFA configuration details.
	MfaConfiguration types.UserPoolMfaType

	// The policies associated with the new user pool.
	Policies *types.UserPoolPolicyType

	// An array of schema attributes for the new user pool. These attributes can be
	// standard or custom attributes.
	Schema []types.SchemaAttributeType

	// A string representing the SMS authentication message.
	SmsAuthenticationMessage *string

	// The SMS configuration with the settings that your Amazon Cognito user pool must
	// use to send an SMS message from your Amazon Web Services account through Amazon
	// Simple Notification Service. To send SMS messages with Amazon SNS in the Amazon
	// Web Services Region that you want, the Amazon Cognito user pool uses an Identity
	// and Access Management (IAM) role in your Amazon Web Services account.
	SmsConfiguration *types.SmsConfigurationType

	// This parameter is no longer used. See [VerificationMessageTemplateType].
	//
	// [VerificationMessageTemplateType]: https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_VerificationMessageTemplateType.html
	SmsVerificationMessage *string

	// The settings for updates to user attributes. These settings include the
	// property AttributesRequireVerificationBeforeUpdate , a user-pool setting that
	// tells Amazon Cognito how to handle changes to the value of your users' email
	// address and phone number attributes. For more information, see [Verifying updates to email addresses and phone numbers].
	//
	// [Verifying updates to email addresses and phone numbers]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-email-phone-verification.html#user-pool-settings-verifications-verify-attribute-updates
	UserAttributeUpdateSettings *types.UserAttributeUpdateSettingsType

	// User pool add-ons. Contains settings for activation of advanced security
	// features. To log user security information but take no action, set to AUDIT . To
	// configure automatic security responses to risky traffic to your user pool, set
	// to ENFORCED .
	//
	// For more information, see [Adding advanced security to a user pool].
	//
	// [Adding advanced security to a user pool]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-advanced-security.html
	UserPoolAddOns *types.UserPoolAddOnsType

	// The tag keys and values to assign to the user pool. A tag is a label that you
	// can use to categorize and manage user pools in different ways, such as by
	// purpose, owner, environment, or other criteria.
	UserPoolTags map[string]string

	// Specifies whether a user can use an email address or phone number as a username
	// when they sign up.
	UsernameAttributes []types.UsernameAttributeType

	// Case sensitivity on the username input for the selected sign-in option. When
	// case sensitivity is set to False (case insensitive), users can sign in with any
	// combination of capital and lowercase letters. For example, username , USERNAME ,
	// or UserName , or for email, email@example.com or EMaiL@eXamplE.Com . For most
	// use cases, set case sensitivity to False (case insensitive) as a best practice.
	// When usernames and email addresses are case insensitive, Amazon Cognito treats
	// any variation in case as the same user, and prevents a case variation from being
	// assigned to the same attribute for a different user.
	//
	// This configuration is immutable after you set it. For more information, see [UsernameConfigurationType].
	//
	// [UsernameConfigurationType]: https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UsernameConfigurationType.html
	UsernameConfiguration *types.UsernameConfigurationType

	// The template for the verification message that the user sees when the app
	// requests permission to access the user's information.
	VerificationMessageTemplate *types.VerificationMessageTemplateType

	noSmithyDocumentSerde
}

// Represents the response from the server for the request to create a user pool.
type CreateUserPoolOutput struct {

	// A container for the user pool details.
	UserPool *types.UserPoolType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateUserPoolMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateUserPool{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateUserPool{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateUserPool"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpCreateUserPoolValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUserPool(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateUserPool(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateUserPool",
	}
}
