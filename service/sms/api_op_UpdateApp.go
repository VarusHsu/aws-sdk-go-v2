// Code generated by smithy-go-codegen DO NOT EDIT.

package sms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sms/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates an application.
func (c *Client) UpdateApp(ctx context.Context, params *UpdateAppInput, optFns ...func(*Options)) (*UpdateAppOutput, error) {
	stack := middleware.NewStack("UpdateApp", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateAppMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateApp(options.Region), middleware.Before)
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
			OperationName: "UpdateApp",
			Err:           err,
		}
	}
	out := result.(*UpdateAppOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAppInput struct {

	// Name of the service role in the customer's account used by AWS SMS.
	RoleName *string

	// New name of the application.
	Name *string

	// ID of the application to update.
	AppId *string

	// List of server groups in the application to update.
	ServerGroups []*types.ServerGroup

	// List of tags to associate with the application.
	Tags []*types.Tag

	// New description of the application.
	Description *string
}

type UpdateAppOutput struct {

	// Summary description of the application.
	AppSummary *types.AppSummary

	// List of tags associated with the application.
	Tags []*types.Tag

	// List of updated server groups in the application.
	ServerGroups []*types.ServerGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateAppMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateApp{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateApp{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateApp(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms",
		OperationName: "UpdateApp",
	}
}
