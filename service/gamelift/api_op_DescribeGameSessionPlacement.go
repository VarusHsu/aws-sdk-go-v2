// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves properties and current status of a game session placement request. To
// get game session placement details, specify the placement ID. If successful, a
// GameSessionPlacement () object is returned.
//
//     * CreateGameSession ()
//
//     *
// DescribeGameSessions ()
//
//     * DescribeGameSessionDetails ()
//
//     *
// SearchGameSessions ()
//
//     * UpdateGameSession ()
//
//     * GetGameSessionLogUrl
// ()
//
//     * Game session placements
//
//         * StartGameSessionPlacement ()
//
//
// * DescribeGameSessionPlacement ()
//
//         * StopGameSessionPlacement ()
func (c *Client) DescribeGameSessionPlacement(ctx context.Context, params *DescribeGameSessionPlacementInput, optFns ...func(*Options)) (*DescribeGameSessionPlacementOutput, error) {
	stack := middleware.NewStack("DescribeGameSessionPlacement", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeGameSessionPlacementMiddlewares(stack)
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
	addOpDescribeGameSessionPlacementValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeGameSessionPlacement(options.Region), middleware.Before)
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
			OperationName: "DescribeGameSessionPlacement",
			Err:           err,
		}
	}
	out := result.(*DescribeGameSessionPlacementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type DescribeGameSessionPlacementInput struct {

	// A unique identifier for a game session placement to retrieve.
	//
	// This member is required.
	PlacementId *string
}

// Represents the returned data in response to a request action.
type DescribeGameSessionPlacementOutput struct {

	// Object that describes the requested game session placement.
	GameSessionPlacement *types.GameSessionPlacement

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeGameSessionPlacementMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeGameSessionPlacement{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeGameSessionPlacement{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeGameSessionPlacement(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeGameSessionPlacement",
	}
}
