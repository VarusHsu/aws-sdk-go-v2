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

// Finds new players to fill open slots in an existing game session. This operation
// can be used to add players to matched games that start with fewer than the
// maximum number of players or to replace players when they drop out. By
// backfilling with the same matchmaker used to create the original match, you
// ensure that new players meet the match criteria and maintain a consistent
// experience throughout the game session. You can backfill a match anytime after a
// game session has been created. To request a match backfill, specify a unique
// ticket ID, the existing game session's ARN, a matchmaking configuration, and a
// set of data that describes all current players in the game session. If
// successful, a match backfill ticket is created and returned with status set to
// QUEUED. The ticket is placed in the matchmaker's ticket pool and processed.
// Track the status of the ticket to respond as needed. The process of finding
// backfill matches is essentially identical to the initial matchmaking process.
// The matchmaker searches the pool and groups tickets together to form potential
// matches, allowing only one backfill ticket per potential match. Once the a match
// is formed, the matchmaker creates player sessions for the new players. All
// tickets in the match are updated with the game session's connection information,
// and the GameSession () object is updated to include matchmaker data on the new
// players. For more detail on how match backfill requests are processed, see  How
// Amazon GameLift FlexMatch Works
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-match.html).
// Learn more  Backfill Existing Games with FlexMatch
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-backfill.html)
// How GameLift FlexMatch Works
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-match.html)
// Related operations
//
//     * StartMatchmaking ()
//
//     * DescribeMatchmaking ()
//
//
// * StopMatchmaking ()
//
//     * AcceptMatch ()
//
//     * StartMatchBackfill ()
func (c *Client) StartMatchBackfill(ctx context.Context, params *StartMatchBackfillInput, optFns ...func(*Options)) (*StartMatchBackfillOutput, error) {
	stack := middleware.NewStack("StartMatchBackfill", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartMatchBackfillMiddlewares(stack)
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
	addOpStartMatchBackfillValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartMatchBackfill(options.Region), middleware.Before)
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
			OperationName: "StartMatchBackfill",
			Err:           err,
		}
	}
	out := result.(*StartMatchBackfillOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type StartMatchBackfillInput struct {

	// Match information on all players that are currently assigned to the game
	// session. This information is used by the matchmaker to find new players and add
	// them to the existing game.
	//
	//     * PlayerID, PlayerAttributes, Team -\\- This
	// information is maintained in the GameSession () object, MatchmakerData property,
	// for all players who are currently assigned to the game session. The matchmaker
	// data is in JSON syntax, formatted as a string. For more details, see  Match Data
	// (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-server.html#match-server-data).
	//
	//
	// * LatencyInMs -\\- If the matchmaker uses player latency, include a latency
	// value, in milliseconds, for the Region that the game session is currently in. Do
	// not include latency values for any other Region.
	//
	// This member is required.
	Players []*types.Player

	// Name of the matchmaker to use for this request. You can use either the
	// configuration name or ARN value. The ARN of the matchmaker that was used with
	// the original game session is listed in the GameSession () object, MatchmakerData
	// property.
	//
	// This member is required.
	ConfigurationName *string

	// Amazon Resource Name (ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html))
	// that is assigned to a game session and uniquely identifies it. This is the same
	// as the game session ID.
	//
	// This member is required.
	GameSessionArn *string

	// A unique identifier for a matchmaking ticket. If no ticket ID is specified here,
	// Amazon GameLift will generate one in the form of a UUID. Use this identifier to
	// track the match backfill ticket status and retrieve match results.
	TicketId *string
}

// Represents the returned data in response to a request action.
type StartMatchBackfillOutput struct {

	// Ticket representing the backfill matchmaking request. This object includes the
	// information in the request, ticket status, and match results as generated during
	// the matchmaking process.
	MatchmakingTicket *types.MatchmakingTicket

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartMatchBackfillMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartMatchBackfill{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartMatchBackfill{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartMatchBackfill(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "StartMatchBackfill",
	}
}
