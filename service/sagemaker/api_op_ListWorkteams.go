// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of work teams that you have defined in a region. The list may be
// empty if no work team satisfies the filter specified in the NameContains
// parameter.
func (c *Client) ListWorkteams(ctx context.Context, params *ListWorkteamsInput, optFns ...func(*Options)) (*ListWorkteamsOutput, error) {
	stack := middleware.NewStack("ListWorkteams", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListWorkteamsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListWorkteams(options.Region), middleware.Before)
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
			OperationName: "ListWorkteams",
			Err:           err,
		}
	}
	out := result.(*ListWorkteamsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorkteamsInput struct {

	// If the result of the previous ListWorkteams request was truncated, the response
	// includes a NextToken. To retrieve the next set of labeling jobs, use the token
	// in the next request.
	NextToken *string

	// The maximum number of work teams to return in each page of the response.
	MaxResults *int32

	// The field to sort results by. The default is CreationTime.
	SortBy types.ListWorkteamsSortByOptions

	// The sort order for results. The default is Ascending.
	SortOrder types.SortOrder

	// A string in the work team's name. This filter returns only work teams whose name
	// contains the specified string.
	NameContains *string
}

type ListWorkteamsOutput struct {

	// An array of Workteam objects, each describing a work team.
	//
	// This member is required.
	Workteams []*types.Workteam

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of work teams, use it in the subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListWorkteamsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListWorkteams{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListWorkteams{}, middleware.After)
}

func newServiceMetadataMiddleware_opListWorkteams(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListWorkteams",
	}
}
