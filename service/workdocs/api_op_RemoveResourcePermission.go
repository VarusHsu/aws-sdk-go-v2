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

// Removes the permission for the specified principal from the specified resource.
func (c *Client) RemoveResourcePermission(ctx context.Context, params *RemoveResourcePermissionInput, optFns ...func(*Options)) (*RemoveResourcePermissionOutput, error) {
	stack := middleware.NewStack("RemoveResourcePermission", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpRemoveResourcePermissionMiddlewares(stack)
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
	addOpRemoveResourcePermissionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveResourcePermission(options.Region), middleware.Before)
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
			OperationName: "RemoveResourcePermission",
			Err:           err,
		}
	}
	out := result.(*RemoveResourcePermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveResourcePermissionInput struct {

	// The principal type of the resource.
	PrincipalType types.PrincipalType

	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string

	// The principal ID of the resource.
	//
	// This member is required.
	PrincipalId *string

	// The ID of the resource.
	//
	// This member is required.
	ResourceId *string
}

type RemoveResourcePermissionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpRemoveResourcePermissionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpRemoveResourcePermission{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpRemoveResourcePermission{}, middleware.After)
}

func newServiceMetadataMiddleware_opRemoveResourcePermission(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workdocs",
		OperationName: "RemoveResourcePermission",
	}
}
