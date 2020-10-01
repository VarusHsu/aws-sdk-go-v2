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

// Returns a list of existing routes in a service mesh.
func (c *Client) ListRoutes(ctx context.Context, params *ListRoutesInput, optFns ...func(*Options)) (*ListRoutesOutput, error) {
	stack := middleware.NewStack("ListRoutes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListRoutesMiddlewares(stack)
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
	addOpListRoutesValidationMiddleware(stack)
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
			OperationName: "ListRoutes",
			Err:           err,
		}
	}
	out := result.(*ListRoutesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ListRoutesInput struct {

	// The name of the service mesh to list routes in.
	//
	// This member is required.
	MeshName *string

	// The name of the virtual router to list routes in.
	//
	// This member is required.
	VirtualRouterName *string

	// The nextToken value returned from a previous paginated ListRoutes request where
	// limit was used and the results exceeded the value of that parameter. Pagination
	// continues from the end of the previous results that returned the nextToken
	// value.
	NextToken *string

	// The maximum number of results returned by ListRoutes in paginated output. When
	// you use this parameter, ListRoutes returns only limit results in a single page
	// along with a nextToken response element. You can see the remaining results of
	// the initial request by sending another ListRoutes request with the returned
	// nextToken value. This value can be between 1 and 100. If you don't use this
	// parameter, ListRoutes returns up to 100 results and a nextToken value if
	// applicable.
	Limit *int32

	// The AWS IAM account ID of the service mesh owner. If the account ID is not your
	// own, then it's the ID of the account that shared the mesh with your account. For
	// more information about mesh sharing, see Working with shared meshes
	// (https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).
	MeshOwner *string
}

//
type ListRoutesOutput struct {

	// The list of existing routes for the specified service mesh and virtual router.
	//
	// This member is required.
	Routes []*types.RouteRef

	// The nextToken value to include in a future ListRoutes request. When the results
	// of a ListRoutes request exceed limit, you can use this value to retrieve the
	// next page of results. This value is null when there are no more results to
	// return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListRoutesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListRoutes{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListRoutes{}, middleware.After)
}
