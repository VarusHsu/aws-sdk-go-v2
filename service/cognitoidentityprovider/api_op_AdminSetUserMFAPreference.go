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

// The user's multi-factor authentication (MFA) preference, including which MFA
// options are activated, and if any are preferred. Only one factor can be set as
// preferred. The preferred MFA factor will be used to authenticate a user if
// multiple factors are activated. If multiple options are activated and no
// preference is set, a challenge to choose an MFA option will be returned during
// sign-in.
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
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func (c *Client) AdminSetUserMFAPreference(ctx context.Context, params *AdminSetUserMFAPreferenceInput, optFns ...func(*Options)) (*AdminSetUserMFAPreferenceOutput, error) {
	if params == nil {
		params = &AdminSetUserMFAPreferenceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdminSetUserMFAPreference", params, optFns, c.addOperationAdminSetUserMFAPreferenceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdminSetUserMFAPreferenceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AdminSetUserMFAPreferenceInput struct {

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string

	// The username of the user that you want to query or modify. The value of this
	// parameter is typically your user's username, but it can be any of their alias
	// attributes. If username isn't an alias attribute in your user pool, this value
	// must be the sub of a local user or the username of a user from a third-party
	// IdP.
	//
	// This member is required.
	Username *string

	// The SMS text message MFA settings.
	SMSMfaSettings *types.SMSMfaSettingsType

	// The time-based one-time password software token MFA settings.
	SoftwareTokenMfaSettings *types.SoftwareTokenMfaSettingsType

	noSmithyDocumentSerde
}

type AdminSetUserMFAPreferenceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAdminSetUserMFAPreferenceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdminSetUserMFAPreference{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminSetUserMFAPreference{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AdminSetUserMFAPreference"); err != nil {
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
	if err = addOpAdminSetUserMFAPreferenceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAdminSetUserMFAPreference(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAdminSetUserMFAPreference(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AdminSetUserMFAPreference",
	}
}
