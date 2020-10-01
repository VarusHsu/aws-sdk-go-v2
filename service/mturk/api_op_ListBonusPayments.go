// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The ListBonusPayments operation retrieves the amounts of bonuses you have paid
// to Workers for a given HIT or assignment.
func (c *Client) ListBonusPayments(ctx context.Context, params *ListBonusPaymentsInput, optFns ...func(*Options)) (*ListBonusPaymentsOutput, error) {
	stack := middleware.NewStack("ListBonusPayments", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListBonusPaymentsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListBonusPayments(options.Region), middleware.Before)
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
			OperationName: "ListBonusPayments",
			Err:           err,
		}
	}
	out := result.(*ListBonusPaymentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBonusPaymentsInput struct {

	// Pagination token
	NextToken *string

	// The ID of the HIT associated with the bonus payments to retrieve. If not
	// specified, all bonus payments for all assignments for the given HIT are
	// returned. Either the HITId parameter or the AssignmentId parameter must be
	// specified
	HITId *string

	MaxResults *int32

	// The ID of the assignment associated with the bonus payments to retrieve. If
	// specified, only bonus payments for the given assignment are returned. Either the
	// HITId parameter or the AssignmentId parameter must be specified
	AssignmentId *string
}

type ListBonusPaymentsOutput struct {

	// A successful request to the ListBonusPayments operation returns a list of
	// BonusPayment objects.
	BonusPayments []*types.BonusPayment

	// The number of bonus payments on this page in the filtered results list,
	// equivalent to the number of bonus payments being returned by this call.
	NumResults *int32

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Mechanical Turk returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListBonusPaymentsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListBonusPayments{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListBonusPayments{}, middleware.After)
}

func newServiceMetadataMiddleware_opListBonusPayments(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "ListBonusPayments",
	}
}
