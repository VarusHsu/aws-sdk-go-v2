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

// Updates the specified user's attributes, including developer attributes, as an
// administrator. Works on any user. For custom attributes, you must prepend the
// custom: prefix to the attribute name. In addition to updating user attributes,
// this API can also be used to mark phone and email as verified. Calling this
// action requires developer credentials.
func (c *Client) AdminUpdateUserAttributes(ctx context.Context, params *AdminUpdateUserAttributesInput, optFns ...func(*Options)) (*AdminUpdateUserAttributesOutput, error) {
	stack := middleware.NewStack("AdminUpdateUserAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAdminUpdateUserAttributesMiddlewares(stack)
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
	addOpAdminUpdateUserAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAdminUpdateUserAttributes(options.Region), middleware.Before)
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
			OperationName: "AdminUpdateUserAttributes",
			Err:           err,
		}
	}
	out := result.(*AdminUpdateUserAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to update the user's attributes as an administrator.
type AdminUpdateUserAttributesInput struct {

	// A map of custom key-value pairs that you can provide as input for any custom
	// workflows that this action triggers. You create custom workflows by assigning
	// AWS Lambda functions to user pool triggers. When you use the
	// AdminUpdateUserAttributes API action, Amazon Cognito invokes the function that
	// is assigned to the custom message trigger. When Amazon Cognito invokes this
	// function, it passes a JSON payload, which the function receives as input. This
	// payload contains a clientMetadata attribute, which provides the data that you
	// assigned to the ClientMetadata parameter in your AdminUpdateUserAttributes
	// request. In your function code in AWS Lambda, you can process the clientMetadata
	// value to enhance your workflow for your specific needs. For more information,
	// see Customizing User Pool Workflows with Lambda Triggers
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

	// An array of name-value pairs representing user attributes. For custom
	// attributes, you must prepend the custom: prefix to the attribute name.
	//
	// This member is required.
	UserAttributes []*types.AttributeType

	// The user pool ID for the user pool where you want to update user attributes.
	//
	// This member is required.
	UserPoolId *string

	// The user name of the user for whom you want to update user attributes.
	//
	// This member is required.
	Username *string
}

// Represents the response from the server for the request to update user
// attributes as an administrator.
type AdminUpdateUserAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAdminUpdateUserAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAdminUpdateUserAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminUpdateUserAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opAdminUpdateUserAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminUpdateUserAttributes",
	}
}
