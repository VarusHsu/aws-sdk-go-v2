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

// Client method for returning the configuration information and metadata of the
// specified user pool app client.
func (c *Client) DescribeUserPoolClient(ctx context.Context, params *DescribeUserPoolClientInput, optFns ...func(*Options)) (*DescribeUserPoolClientOutput, error) {
	stack := middleware.NewStack("DescribeUserPoolClient", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeUserPoolClientMiddlewares(stack)
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
	addOpDescribeUserPoolClientValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeUserPoolClient(options.Region), middleware.Before)
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
			OperationName: "DescribeUserPoolClient",
			Err:           err,
		}
	}
	out := result.(*DescribeUserPoolClientOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to describe a user pool client.
type DescribeUserPoolClientInput struct {

	// The app client ID of the app associated with the user pool.
	//
	// This member is required.
	ClientId *string

	// The user pool ID for the user pool you want to describe.
	//
	// This member is required.
	UserPoolId *string
}

// Represents the response from the server from a request to describe the user pool
// client.
type DescribeUserPoolClientOutput struct {

	// The user pool client from a server response to describe the user pool client.
	UserPoolClient *types.UserPoolClientType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeUserPoolClientMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeUserPoolClient{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeUserPoolClient{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeUserPoolClient(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "DescribeUserPoolClient",
	}
}
