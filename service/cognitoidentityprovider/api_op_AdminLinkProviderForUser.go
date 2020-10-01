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

// Links an existing user account in a user pool (DestinationUser) to an identity
// from an external identity provider (SourceUser) based on a specified attribute
// name and value from the external identity provider. This allows you to create a
// link from the existing user account to an external federated user identity that
// has not yet been used to sign in, so that the federated user identity can be
// used to sign in as the existing user account. For example, if there is an
// existing user with a username and password, this API links that user to a
// federated user identity, so that when the federated user identity is used, the
// user signs in as the existing user account. Because this API allows a user with
// an external federated identity to sign in as an existing user in the user pool,
// it is critical that it only be used with external identity providers and
// provider attributes that have been trusted by the application owner. See also .
// This action is enabled only for admin access and requires developer credentials.
func (c *Client) AdminLinkProviderForUser(ctx context.Context, params *AdminLinkProviderForUserInput, optFns ...func(*Options)) (*AdminLinkProviderForUserOutput, error) {
	stack := middleware.NewStack("AdminLinkProviderForUser", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAdminLinkProviderForUserMiddlewares(stack)
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
	addOpAdminLinkProviderForUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAdminLinkProviderForUser(options.Region), middleware.Before)
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
			OperationName: "AdminLinkProviderForUser",
			Err:           err,
		}
	}
	out := result.(*AdminLinkProviderForUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AdminLinkProviderForUserInput struct {

	// The existing user in the user pool to be linked to the external identity
	// provider user account. Can be a native (Username + Password) Cognito User Pools
	// user or a federated user (for example, a SAML or Facebook user). If the user
	// doesn't exist, an exception is thrown. This is the user that is returned when
	// the new user (with the linked identity provider attribute) signs in. For a
	// native username + password user, the ProviderAttributeValue for the
	// DestinationUser should be the username in the user pool. For a federated user,
	// it should be the provider-specific user_id. The ProviderAttributeName of the
	// DestinationUser is ignored. The ProviderName should be set to Cognito for users
	// in Cognito user pools.
	//
	// This member is required.
	DestinationUser *types.ProviderUserIdentifierType

	// The user pool ID for the user pool.
	//
	// This member is required.
	UserPoolId *string

	// An external identity provider account for a user who does not currently exist
	// yet in the user pool. This user must be a federated user (for example, a SAML or
	// Facebook user), not another native user. If the SourceUser is a federated social
	// identity provider user (Facebook, Google, or Login with Amazon), you must set
	// the ProviderAttributeName to Cognito_Subject. For social identity providers, the
	// ProviderName will be Facebook, Google, or LoginWithAmazon, and Cognito will
	// automatically parse the Facebook, Google, and Login with Amazon tokens for id,
	// sub, and user_id, respectively. The ProviderAttributeValue for the user must be
	// the same value as the id, sub, or user_id value found in the social identity
	// provider token. For SAML, the ProviderAttributeName can be any value that
	// matches a claim in the SAML assertion. If you wish to link SAML users based on
	// the subject of the SAML assertion, you should map the subject to a claim through
	// the SAML identity provider and submit that claim name as the
	// ProviderAttributeName. If you set ProviderAttributeName to Cognito_Subject,
	// Cognito will automatically parse the default unique identifier found in the
	// subject from the SAML token.
	//
	// This member is required.
	SourceUser *types.ProviderUserIdentifierType
}

type AdminLinkProviderForUserOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAdminLinkProviderForUserMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAdminLinkProviderForUser{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminLinkProviderForUser{}, middleware.After)
}

func newServiceMetadataMiddleware_opAdminLinkProviderForUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminLinkProviderForUser",
	}
}
