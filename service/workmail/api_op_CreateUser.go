// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a user who can be used in Amazon WorkMail by calling the
// RegisterToWorkMail () operation.
func (c *Client) CreateUser(ctx context.Context, params *CreateUserInput, optFns ...func(*Options)) (*CreateUserOutput, error) {
	stack := middleware.NewStack("CreateUser", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateUserMiddlewares(stack)
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
	addOpCreateUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUser(options.Region), middleware.Before)
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
			OperationName: "CreateUser",
			Err:           err,
		}
	}
	out := result.(*CreateUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateUserInput struct {

	// The password for the new user.
	//
	// This member is required.
	Password *string

	// The name for the new user. Simple AD or AD Connector user names have a maximum
	// length of 20. All others have a maximum length of 64.
	//
	// This member is required.
	Name *string

	// The display name for the new user.
	//
	// This member is required.
	DisplayName *string

	// The identifier of the organization for which the user is created.
	//
	// This member is required.
	OrganizationId *string
}

type CreateUserOutput struct {

	// The identifier for the new user.
	UserId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateUserMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateUser{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateUser{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "CreateUser",
	}
}
