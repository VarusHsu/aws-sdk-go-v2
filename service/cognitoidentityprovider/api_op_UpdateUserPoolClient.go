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

// Updates the specified user pool app client with the specified attributes. You
// can get a list of the current user pool app client settings with . If you don't
// provide a value for an attribute, it will be set to the default value.
func (c *Client) UpdateUserPoolClient(ctx context.Context, params *UpdateUserPoolClientInput, optFns ...func(*Options)) (*UpdateUserPoolClientOutput, error) {
	stack := middleware.NewStack("UpdateUserPoolClient", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateUserPoolClientMiddlewares(stack)
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
	addOpUpdateUserPoolClientValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateUserPoolClient(options.Region), middleware.Before)
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
			OperationName: "UpdateUserPoolClient",
			Err:           err,
		}
	}
	out := result.(*UpdateUserPoolClientOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to update the user pool client.
type UpdateUserPoolClientInput struct {

	// A list of provider names for the identity providers that are supported on this
	// client.
	SupportedIdentityProviders []*string

	// A list of allowed redirect (callback) URLs for the identity providers. A
	// redirect URI must:
	//
	//     * Be an absolute URI.
	//
	//     * Be registered with the
	// authorization server.
	//
	//     * Not include a fragment component.
	//
	// See OAuth 2.0 -
	// Redirection Endpoint (https://tools.ietf.org/html/rfc6749#section-3.1.2). Amazon
	// Cognito requires HTTPS over HTTP except for http://localhost for testing
	// purposes only. App callback URLs such as myapp://example are also supported.
	CallbackURLs []*string

	// The allowed OAuth scopes. Possible values provided by OAuth are: phone, email,
	// openid, and profile. Possible values provided by AWS are:
	// aws.cognito.signin.user.admin. Custom scopes created in Resource Servers are
	// also supported.
	AllowedOAuthScopes []*string

	// The ID of the client associated with the user pool.
	//
	// This member is required.
	ClientId *string

	// A list of allowed logout URLs for the identity providers.
	LogoutURLs []*string

	// The time limit, in days, after which the refresh token is no longer valid and
	// cannot be used.
	RefreshTokenValidity *int32

	// Use this setting to choose which errors and responses are returned by Cognito
	// APIs during authentication, account confirmation, and password recovery when the
	// user does not exist in the user pool. When set to ENABLED and the user does not
	// exist, authentication returns an error indicating either the username or
	// password was incorrect, and account confirmation and password recovery return a
	// response indicating a code was sent to a simulated destination. When set to
	// LEGACY, those APIs will return a UserNotFoundException exception if the user
	// does not exist in the user pool. Valid values include:
	//
	//     * ENABLED - This
	// prevents user existence-related errors.
	//
	//     * LEGACY - This represents the old
	// behavior of Cognito where user existence related errors are not prevented.
	//
	// This
	// setting affects the behavior of following APIs:
	//
	//     * AdminInitiateAuth ()
	//
	//
	// * AdminRespondToAuthChallenge ()
	//
	//     * InitiateAuth ()
	//
	//     *
	// RespondToAuthChallenge ()
	//
	//     * ForgotPassword ()
	//
	//     * ConfirmForgotPassword
	// ()
	//
	//     * ConfirmSignUp ()
	//
	//     * ResendConfirmationCode ()
	//
	// After February 15th
	// 2020, the value of PreventUserExistenceErrors will default to ENABLED for newly
	// created user pool clients if no value is provided.
	PreventUserExistenceErrors types.PreventUserExistenceErrorTypes

	// The writeable attributes of the user pool.
	WriteAttributes []*string

	// The Amazon Pinpoint analytics configuration for collecting metrics for this user
	// pool. Cognito User Pools only supports sending events to Amazon Pinpoint
	// projects in the US East (N. Virginia) us-east-1 Region, regardless of the region
	// in which the user pool resides.
	AnalyticsConfiguration *types.AnalyticsConfigurationType

	// The user pool ID for the user pool where you want to update the user pool
	// client.
	//
	// This member is required.
	UserPoolId *string

	// The allowed OAuth flows. Set to code to initiate a code grant flow, which
	// provides an authorization code as the response. This code can be exchanged for
	// access tokens with the token endpoint. Set to implicit to specify that the
	// client should get the access token (and, optionally, ID token, based on scopes)
	// directly. Set to client_credentials to specify that the client should get the
	// access token (and, optionally, ID token, based on scopes) from the token
	// endpoint using a combination of client and client_secret.
	AllowedOAuthFlows []types.OAuthFlowType

	// The client name from the update user pool client request.
	ClientName *string

	// The authentication flows that are supported by the user pool clients. Flow names
	// without the ALLOW_ prefix are deprecated in favor of new names with the ALLOW_
	// prefix. Note that values with ALLOW_ prefix cannot be used along with values
	// without ALLOW_ prefix. Valid values include:
	//
	//     *
	// ALLOW_ADMIN_USER_PASSWORD_AUTH: Enable admin based user password authentication
	// flow ADMIN_USER_PASSWORD_AUTH. This setting replaces the ADMIN_NO_SRP_AUTH
	// setting. With this authentication flow, Cognito receives the password in the
	// request instead of using the SRP (Secure Remote Password protocol) protocol to
	// verify passwords.
	//
	//     * ALLOW_CUSTOM_AUTH: Enable Lambda trigger based
	// authentication.
	//
	//     * ALLOW_USER_PASSWORD_AUTH: Enable user password-based
	// authentication. In this flow, Cognito receives the password in the request
	// instead of using the SRP protocol to verify passwords.
	//
	//     *
	// ALLOW_USER_SRP_AUTH: Enable SRP based authentication.
	//
	//     *
	// ALLOW_REFRESH_TOKEN_AUTH: Enable authflow to refresh tokens.
	ExplicitAuthFlows []types.ExplicitAuthFlowsType

	// Set to true if the client is allowed to follow the OAuth protocol when
	// interacting with Cognito user pools.
	AllowedOAuthFlowsUserPoolClient *bool

	// The default redirect URI. Must be in the CallbackURLs list. A redirect URI
	// must:
	//
	//     * Be an absolute URI.
	//
	//     * Be registered with the authorization
	// server.
	//
	//     * Not include a fragment component.
	//
	// See OAuth 2.0 - Redirection
	// Endpoint (https://tools.ietf.org/html/rfc6749#section-3.1.2). Amazon Cognito
	// requires HTTPS over HTTP except for http://localhost for testing purposes only.
	// App callback URLs such as myapp://example are also supported.
	DefaultRedirectURI *string

	// The read-only attributes of the user pool.
	ReadAttributes []*string
}

// Represents the response from the server to the request to update the user pool
// client.
type UpdateUserPoolClientOutput struct {

	// The user pool client value from the response from the server when an update user
	// pool client request is made.
	UserPoolClient *types.UserPoolClientType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateUserPoolClientMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateUserPoolClient{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateUserPoolClient{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateUserPoolClient(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "UpdateUserPoolClient",
	}
}
