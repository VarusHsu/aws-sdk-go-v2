// Code generated by smithy-go-codegen DO NOT EDIT.

package workdocs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves details of the current user for whom the authentication token was
// generated. This is not a valid action for SigV4 (administrative API) clients.
// This action requires an authentication token. To get an authentication token,
// register an application with Amazon WorkDocs. For more information, see
// Authentication and Access Control for User Applications
// (https://docs.aws.amazon.com/workdocs/latest/developerguide/wd-auth-user.html)
// in the Amazon WorkDocs Developer Guide.
func (c *Client) GetCurrentUser(ctx context.Context, params *GetCurrentUserInput, optFns ...func(*Options)) (*GetCurrentUserOutput, error) {
	stack := middleware.NewStack("GetCurrentUser", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetCurrentUserMiddlewares(stack)
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
	addOpGetCurrentUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCurrentUser(options.Region), middleware.Before)
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
			OperationName: "GetCurrentUser",
			Err:           err,
		}
	}
	out := result.(*GetCurrentUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCurrentUserInput struct {

	// Amazon WorkDocs authentication token.
	//
	// This member is required.
	AuthenticationToken *string
}

type GetCurrentUserOutput struct {

	// Metadata of the user.
	User *types.User

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetCurrentUserMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetCurrentUser{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCurrentUser{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetCurrentUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workdocs",
		OperationName: "GetCurrentUser",
	}
}
