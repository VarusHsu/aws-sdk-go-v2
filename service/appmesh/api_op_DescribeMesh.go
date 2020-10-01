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

// Describes an existing service mesh.
func (c *Client) DescribeMesh(ctx context.Context, params *DescribeMeshInput, optFns ...func(*Options)) (*DescribeMeshOutput, error) {
	stack := middleware.NewStack("DescribeMesh", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeMeshMiddlewares(stack)
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
	addOpDescribeMeshValidationMiddleware(stack)
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
			OperationName: "DescribeMesh",
			Err:           err,
		}
	}
	out := result.(*DescribeMeshOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeMeshInput struct {

	// The name of the service mesh to describe.
	//
	// This member is required.
	MeshName *string

	// The AWS IAM account ID of the service mesh owner. If the account ID is not your
	// own, then it's the ID of the account that shared the mesh with your account. For
	// more information about mesh sharing, see Working with shared meshes
	// (https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).
	MeshOwner *string
}

//
type DescribeMeshOutput struct {

	// The full description of your service mesh.
	//
	// This member is required.
	Mesh *types.MeshData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeMeshMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeMesh{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeMesh{}, middleware.After)
}
