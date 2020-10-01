// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Signs out users from all devices, as an administrator. It also invalidates all
// refresh tokens issued to a user. The user's current access and Id tokens remain
// valid until their expiry. Access and Id tokens expire one hour after they are
// issued. Calling this action requires developer credentials.
func (c *Client) AdminUserGlobalSignOut(ctx context.Context, params *AdminUserGlobalSignOutInput, optFns ...func(*Options)) (*AdminUserGlobalSignOutOutput, error) {
	stack := middleware.NewStack("AdminUserGlobalSignOut", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAdminUserGlobalSignOutMiddlewares(stack)
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
	addOpAdminUserGlobalSignOutValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAdminUserGlobalSignOut(options.Region), middleware.Before)
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
			OperationName: "AdminUserGlobalSignOut",
			Err:           err,
		}
	}
	out := result.(*AdminUserGlobalSignOutOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to sign out of all devices, as an administrator.
type AdminUserGlobalSignOutInput struct {

	// The user name.
	//
	// This member is required.
	Username *string

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string
}

// The global sign-out response, as an administrator.
type AdminUserGlobalSignOutOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAdminUserGlobalSignOutMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAdminUserGlobalSignOut{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminUserGlobalSignOut{}, middleware.After)
}

func newServiceMetadataMiddleware_opAdminUserGlobalSignOut(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminUserGlobalSignOut",
	}
}
