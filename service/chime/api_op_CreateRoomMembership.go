// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds a member to a chat room in an Amazon Chime Enterprise account. A member can
// be either a user or a bot. The member role designates whether the member is a
// chat room administrator or a general chat room member.
func (c *Client) CreateRoomMembership(ctx context.Context, params *CreateRoomMembershipInput, optFns ...func(*Options)) (*CreateRoomMembershipOutput, error) {
	stack := middleware.NewStack("CreateRoomMembership", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateRoomMembershipMiddlewares(stack)
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
	addOpCreateRoomMembershipValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRoomMembership(options.Region), middleware.Before)
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
			OperationName: "CreateRoomMembership",
			Err:           err,
		}
	}
	out := result.(*CreateRoomMembershipOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRoomMembershipInput struct {

	// The role of the member.
	Role types.RoomMembershipRole

	// The Amazon Chime member ID (user ID or bot ID).
	//
	// This member is required.
	MemberId *string

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string

	// The room ID.
	//
	// This member is required.
	RoomId *string
}

type CreateRoomMembershipOutput struct {

	// The room membership details.
	RoomMembership *types.RoomMembership

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateRoomMembershipMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateRoomMembership{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRoomMembership{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateRoomMembership(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "CreateRoomMembership",
	}
}
