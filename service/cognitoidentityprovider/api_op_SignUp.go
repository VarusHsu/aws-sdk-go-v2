// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Registers the user in the specified user pool and creates a user name, password,
// and user attributes.
func (c *Client) SignUp(ctx context.Context, params *SignUpInput, optFns ...func(*Options)) (*SignUpOutput, error) {
	stack := middleware.NewStack("SignUp", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpSignUpMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpSignUpValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSignUp(options.Region), middleware.Before)
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
			OperationName: "SignUp",
			Err:           err,
		}
	}
	out := result.(*SignUpOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to register a user.
type SignUpInput struct {

	// The user name of the user you wish to register.
	//
	// This member is required.
	Username *string

	// A map of custom key-value pairs that you can provide as input for any custom
	// workflows that this action triggers. You create custom workflows by assigning
	// AWS Lambda functions to user pool triggers. When you use the SignUp API action,
	// Amazon Cognito invokes any functions that are assigned to the following
	// triggers: pre sign-up, custom message, and post confirmation. When Amazon
	// Cognito invokes any of these functions, it passes a JSON payload, which the
	// function receives as input. This payload contains a clientMetadata attribute,
	// which provides the data that you assigned to the ClientMetadata parameter in
	// your SignUp request. In your function code in AWS Lambda, you can process the
	// clientMetadata value to enhance your workflow for your specific needs. For more
	// information, see Customizing User Pool Workflows with Lambda Triggers
	// (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html)
	// in the Amazon Cognito Developer Guide. Take the following limitations into
	// consideration when you use the ClientMetadata parameter:
	//
	//     * Amazon Cognito
	// does not store the ClientMetadata value. This data is available only to AWS
	// Lambda triggers that are assigned to a user pool to support custom workflows. If
	// your user pool configuration does not include triggers, the ClientMetadata
	// parameter serves no purpose.
	//
	//     * Amazon Cognito does not validate the
	// ClientMetadata value.
	//
	//     * Amazon Cognito does not encrypt the the
	// ClientMetadata value, so don't use it to provide sensitive information.
	ClientMetadata map[string]*string

	// The password of the user you wish to register.
	//
	// This member is required.
	Password *string

	// The Amazon Pinpoint analytics metadata for collecting metrics for SignUp calls.
	AnalyticsMetadata *types.AnalyticsMetadataType

	// An array of name-value pairs representing user attributes. For custom
	// attributes, you must prepend the custom: prefix to the attribute name.
	UserAttributes []*types.AttributeType

	// The ID of the client associated with the user pool.
	//
	// This member is required.
	ClientId *string

	// Contextual data such as the user's device fingerprint, IP address, or location
	// used for evaluating the risk of an unexpected event by Amazon Cognito advanced
	// security.
	UserContextData *types.UserContextDataType

	// A keyed-hash message authentication code (HMAC) calculated using the secret key
	// of a user pool client and username plus the client ID in the message.
	SecretHash *string

	// The validation data in the request to register a user.
	ValidationData []*types.AttributeType
}

// The response from the server for a registration request.
type SignUpOutput struct {

	// A response from the server indicating that a user registration has been
	// confirmed.
	//
	// This member is required.
	UserConfirmed *bool

	// The UUID of the authenticated user. This is not the same as username.
	//
	// This member is required.
	UserSub *string

	// The code delivery details returned by the server response to the user
	// registration request.
	CodeDeliveryDetails *types.CodeDeliveryDetailsType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpSignUpMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpSignUp{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpSignUp{}, middleware.After)
}

func newServiceMetadataMiddleware_opSignUp(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SignUp",
	}
}
