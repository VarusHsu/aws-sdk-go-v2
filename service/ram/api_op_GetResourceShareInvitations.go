// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the invitations for resource sharing that you've received.
func (c *Client) GetResourceShareInvitations(ctx context.Context, params *GetResourceShareInvitationsInput, optFns ...func(*Options)) (*GetResourceShareInvitationsOutput, error) {
	stack := middleware.NewStack("GetResourceShareInvitations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetResourceShareInvitationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetResourceShareInvitations(options.Region), middleware.Before)
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
			OperationName: "GetResourceShareInvitations",
			Err:           err,
		}
	}
	out := result.(*GetResourceShareInvitationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResourceShareInvitationsInput struct {

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The Amazon Resource Names (ARN) of the invitations.
	ResourceShareInvitationArns []*string

	// The token for the next page of results.
	NextToken *string

	// The Amazon Resource Names (ARN) of the resource shares.
	ResourceShareArns []*string
}

type GetResourceShareInvitationsOutput struct {

	// Information about the invitations.
	ResourceShareInvitations []*types.ResourceShareInvitation

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetResourceShareInvitationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetResourceShareInvitations{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetResourceShareInvitations{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetResourceShareInvitations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "GetResourceShareInvitations",
	}
}
