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
	"time"
)

// Gets the specified user by user name in a user pool as an administrator. Works
// on any user. Calling this action requires developer credentials.
func (c *Client) AdminGetUser(ctx context.Context, params *AdminGetUserInput, optFns ...func(*Options)) (*AdminGetUserOutput, error) {
	stack := middleware.NewStack("AdminGetUser", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAdminGetUserMiddlewares(stack)
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
	addOpAdminGetUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAdminGetUser(options.Region), middleware.Before)
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
			OperationName: "AdminGetUser",
			Err:           err,
		}
	}
	out := result.(*AdminGetUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to get the specified user as an administrator.
type AdminGetUserInput struct {

	// The user pool ID for the user pool where you want to get information about the
	// user.
	//
	// This member is required.
	UserPoolId *string

	// The user name of the user you wish to retrieve.
	//
	// This member is required.
	Username *string
}

// Represents the response from the server from the request to get the specified
// user as an administrator.
type AdminGetUserOutput struct {

	// The user's preferred MFA setting.
	PreferredMfaSetting *string

	// The date the user was created.
	UserCreateDate *time.Time

	// The user name of the user about whom you are receiving information.
	//
	// This member is required.
	Username *string

	// The MFA options that are enabled for the user. The possible values in this list
	// are SMS_MFA and SOFTWARE_TOKEN_MFA.
	UserMFASettingList []*string

	// Indicates that the status is enabled.
	Enabled *bool

	// An array of name-value pairs representing user attributes.
	UserAttributes []*types.AttributeType

	// The user status. Can be one of the following:
	//
	//     * UNCONFIRMED - User has been
	// created but not confirmed.
	//
	//     * CONFIRMED - User has been confirmed.
	//
	//     *
	// ARCHIVED - User is no longer active.
	//
	//     * COMPROMISED - User is disabled due
	// to a potential security threat.
	//
	//     * UNKNOWN - User status is not known.
	//
	//
	// * RESET_REQUIRED - User is confirmed, but the user must request a code and reset
	// his or her password before he or she can sign in.
	//
	//     * FORCE_CHANGE_PASSWORD -
	// The user is confirmed and the user can sign in using a temporary password, but
	// on first sign-in, the user must change his or her password to a new value before
	// doing anything else.
	UserStatus types.UserStatusType

	// This response parameter is no longer supported. It provides information only
	// about SMS MFA configurations. It doesn't provide information about TOTP software
	// token MFA configurations. To look up information about either type of MFA
	// configuration, use the AdminGetUserResponse$UserMFASettingList () response
	// instead.
	MFAOptions []*types.MFAOptionType

	// The date the user was last modified.
	UserLastModifiedDate *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAdminGetUserMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAdminGetUser{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminGetUser{}, middleware.After)
}

func newServiceMetadataMiddleware_opAdminGetUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminGetUser",
	}
}
