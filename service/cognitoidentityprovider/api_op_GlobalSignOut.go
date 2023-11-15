// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Signs out a user from all devices. GlobalSignOut invalidates all identity,
// access and refresh tokens that Amazon Cognito has issued to a user. A user can
// still use a hosted UI cookie to retrieve new tokens for the duration of the
// 1-hour cookie validity period. Your app isn't aware that a user's access token
// is revoked unless it attempts to authorize a user pools API request with an
// access token that contains the scope aws.cognito.signin.user.admin . Your app
// might otherwise accept access tokens until they expire. Amazon Cognito doesn't
// evaluate Identity and Access Management (IAM) policies in requests for this API
// operation. For this operation, you can't use IAM credentials to authorize
// requests, and you can't grant IAM permissions in policies. For more information
// about authorization models in Amazon Cognito, see Using the Amazon Cognito
// native and OIDC APIs (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html)
// .
func (c *Client) GlobalSignOut(ctx context.Context, params *GlobalSignOutInput, optFns ...func(*Options)) (*GlobalSignOutOutput, error) {
	if params == nil {
		params = &GlobalSignOutInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GlobalSignOut", params, optFns, c.addOperationGlobalSignOutMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GlobalSignOutOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to sign out all devices.
type GlobalSignOutInput struct {

	// A valid access token that Amazon Cognito issued to the user who you want to
	// sign out.
	//
	// This member is required.
	AccessToken *string

	noSmithyDocumentSerde
}

// The response to the request to sign out all devices.
type GlobalSignOutOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGlobalSignOutMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGlobalSignOut{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGlobalSignOut{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GlobalSignOut"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpGlobalSignOutValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGlobalSignOut(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opGlobalSignOut(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GlobalSignOut",
	}
}
