// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchain

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchain/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Casts a vote for a specified ProposalId on behalf of a member. The member to
// vote as, specified by VoterMemberId, must be in the same AWS account as the
// principal that calls the action.
func (c *Client) VoteOnProposal(ctx context.Context, params *VoteOnProposalInput, optFns ...func(*Options)) (*VoteOnProposalOutput, error) {
	stack := middleware.NewStack("VoteOnProposal", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpVoteOnProposalMiddlewares(stack)
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
	addOpVoteOnProposalValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opVoteOnProposal(options.Region), middleware.Before)
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
			OperationName: "VoteOnProposal",
			Err:           err,
		}
	}
	out := result.(*VoteOnProposalOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type VoteOnProposalInput struct {

	// The unique identifier of the network.
	//
	// This member is required.
	NetworkId *string

	// The value of the vote.
	//
	// This member is required.
	Vote types.VoteValue

	// The unique identifier of the member casting the vote.
	//
	// This member is required.
	VoterMemberId *string

	// The unique identifier of the proposal.
	//
	// This member is required.
	ProposalId *string
}

type VoteOnProposalOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpVoteOnProposalMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpVoteOnProposal{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpVoteOnProposal{}, middleware.After)
}

func newServiceMetadataMiddleware_opVoteOnProposal(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "managedblockchain",
		OperationName: "VoteOnProposal",
	}
}
