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

// Gets the specified identity provider.
func (c *Client) GetIdentityProviderByIdentifier(ctx context.Context, params *GetIdentityProviderByIdentifierInput, optFns ...func(*Options)) (*GetIdentityProviderByIdentifierOutput, error) {
	stack := middleware.NewStack("GetIdentityProviderByIdentifier", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetIdentityProviderByIdentifierMiddlewares(stack)
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
	addOpGetIdentityProviderByIdentifierValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetIdentityProviderByIdentifier(options.Region), middleware.Before)
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
			OperationName: "GetIdentityProviderByIdentifier",
			Err:           err,
		}
	}
	out := result.(*GetIdentityProviderByIdentifierOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetIdentityProviderByIdentifierInput struct {

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string

	// The identity provider ID.
	//
	// This member is required.
	IdpIdentifier *string
}

type GetIdentityProviderByIdentifierOutput struct {

	// The identity provider object.
	//
	// This member is required.
	IdentityProvider *types.IdentityProviderType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetIdentityProviderByIdentifierMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetIdentityProviderByIdentifier{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetIdentityProviderByIdentifier{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetIdentityProviderByIdentifier(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "GetIdentityProviderByIdentifier",
	}
}
