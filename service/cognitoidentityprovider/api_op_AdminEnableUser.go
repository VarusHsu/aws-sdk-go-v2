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

// Enables the specified user as an administrator. Works on any user. Calling this
// action requires developer credentials.
func (c *Client) AdminEnableUser(ctx context.Context, params *AdminEnableUserInput, optFns ...func(*Options)) (*AdminEnableUserOutput, error) {
	stack := middleware.NewStack("AdminEnableUser", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAdminEnableUserMiddlewares(stack)
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
	addOpAdminEnableUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAdminEnableUser(options.Region), middleware.Before)
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
			OperationName: "AdminEnableUser",
			Err:           err,
		}
	}
	out := result.(*AdminEnableUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request that enables the user as an administrator.
type AdminEnableUserInput struct {

	// The user name of the user you wish to enable.
	//
	// This member is required.
	Username *string

	// The user pool ID for the user pool where you want to enable the user.
	//
	// This member is required.
	UserPoolId *string
}

// Represents the response from the server for the request to enable a user as an
// administrator.
type AdminEnableUserOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAdminEnableUserMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAdminEnableUser{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminEnableUser{}, middleware.After)
}

func newServiceMetadataMiddleware_opAdminEnableUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminEnableUser",
	}
}
