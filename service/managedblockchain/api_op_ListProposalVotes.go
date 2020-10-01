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

// Returns the listing of votes for a specified proposal, including the value of
// each vote and the unique identifier of the member that cast the vote.
func (c *Client) ListProposalVotes(ctx context.Context, params *ListProposalVotesInput, optFns ...func(*Options)) (*ListProposalVotesOutput, error) {
	stack := middleware.NewStack("ListProposalVotes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListProposalVotesMiddlewares(stack)
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
	addOpListProposalVotesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListProposalVotes(options.Region), middleware.Before)
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
			OperationName: "ListProposalVotes",
			Err:           err,
		}
	}
	out := result.(*ListProposalVotesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProposalVotesInput struct {

	// The unique identifier of the network.
	//
	// This member is required.
	NetworkId *string

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	// The maximum number of votes to return.
	MaxResults *int32

	// The unique identifier of the proposal.
	//
	// This member is required.
	ProposalId *string
}

type ListProposalVotesOutput struct {

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	// The listing of votes.
	ProposalVotes []*types.VoteSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListProposalVotesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListProposalVotes{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListProposalVotes{}, middleware.After)
}

func newServiceMetadataMiddleware_opListProposalVotes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "managedblockchain",
		OperationName: "ListProposalVotes",
	}
}
