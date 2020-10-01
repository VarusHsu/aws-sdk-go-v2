// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates an existing virtual node in a specified service mesh.
func (c *Client) UpdateVirtualNode(ctx context.Context, params *UpdateVirtualNodeInput, optFns ...func(*Options)) (*UpdateVirtualNodeOutput, error) {
	stack := middleware.NewStack("UpdateVirtualNode", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateVirtualNodeMiddlewares(stack)
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
	addOpUpdateVirtualNodeValidationMiddleware(stack)
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
			OperationName: "UpdateVirtualNode",
			Err:           err,
		}
	}
	out := result.(*UpdateVirtualNodeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type UpdateVirtualNodeInput struct {

	// The name of the virtual node to update.
	//
	// This member is required.
	VirtualNodeName *string

	// The name of the service mesh that the virtual node resides in.
	//
	// This member is required.
	MeshName *string

	// The new virtual node specification to apply. This overwrites the existing data.
	//
	// This member is required.
	Spec *types.VirtualNodeSpec

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. Up to 36 letters, numbers, hyphens, and underscores are allowed.
	ClientToken *string

	// The AWS IAM account ID of the service mesh owner. If the account ID is not your
	// own, then it's the ID of the account that shared the mesh with your account. For
	// more information about mesh sharing, see Working with shared meshes
	// (https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).
	MeshOwner *string
}

//
type UpdateVirtualNodeOutput struct {

	// A full description of the virtual node that was updated.
	//
	// This member is required.
	VirtualNode *types.VirtualNodeData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateVirtualNodeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateVirtualNode{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateVirtualNode{}, middleware.After)
}
