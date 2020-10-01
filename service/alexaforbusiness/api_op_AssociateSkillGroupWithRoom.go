// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates a skill group with a given room. This enables all skills in the
// associated skill group on all devices in the room.
func (c *Client) AssociateSkillGroupWithRoom(ctx context.Context, params *AssociateSkillGroupWithRoomInput, optFns ...func(*Options)) (*AssociateSkillGroupWithRoomOutput, error) {
	stack := middleware.NewStack("AssociateSkillGroupWithRoom", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAssociateSkillGroupWithRoomMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateSkillGroupWithRoom(options.Region), middleware.Before)
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
			OperationName: "AssociateSkillGroupWithRoom",
			Err:           err,
		}
	}
	out := result.(*AssociateSkillGroupWithRoomOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateSkillGroupWithRoomInput struct {

	// The ARN of the skill group to associate with a room. Required.
	SkillGroupArn *string

	// The ARN of the room with which to associate the skill group. Required.
	RoomArn *string
}

type AssociateSkillGroupWithRoomOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAssociateSkillGroupWithRoomMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateSkillGroupWithRoom{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateSkillGroupWithRoom{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssociateSkillGroupWithRoom(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "AssociateSkillGroupWithRoom",
	}
}
