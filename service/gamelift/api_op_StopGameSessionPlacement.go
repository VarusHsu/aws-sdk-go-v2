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

// Cancels a game session placement that is in PENDING status. To stop a placement,
// provide the placement ID values. If successful, the placement is moved to
// CANCELLED status.
//
//     * CreateGameSession ()
//
//     * DescribeGameSessions ()
//
//
// * DescribeGameSessionDetails ()
//
//     * SearchGameSessions ()
//
//     *
// UpdateGameSession ()
//
//     * GetGameSessionLogUrl ()
//
//     * Game session
// placements
//
//         * StartGameSessionPlacement ()
//
//         *
// DescribeGameSessionPlacement ()
//
//         * StopGameSessionPlacement ()
func (c *Client) StopGameSessionPlacement(ctx context.Context, params *StopGameSessionPlacementInput, optFns ...func(*Options)) (*StopGameSessionPlacementOutput, error) {
	stack := middleware.NewStack("StopGameSessionPlacement", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStopGameSessionPlacementMiddlewares(stack)
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
	addOpStopGameSessionPlacementValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopGameSessionPlacement(options.Region), middleware.Before)
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
			OperationName: "StopGameSessionPlacement",
			Err:           err,
		}
	}
	out := result.(*StopGameSessionPlacementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type StopGameSessionPlacementInput struct {

	// A unique identifier for a game session placement to cancel.
	//
	// This member is required.
	PlacementId *string
}

// Represents the returned data in response to a request action.
type StopGameSessionPlacementOutput struct {

	// Object that describes the canceled game session placement, with CANCELLED status
	// and an end time stamp.
	GameSessionPlacement *types.GameSessionPlacement

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStopGameSessionPlacementMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStopGameSessionPlacement{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopGameSessionPlacement{}, middleware.After)
}

func newServiceMetadataMiddleware_opStopGameSessionPlacement(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "StopGameSessionPlacement",
	}
}
