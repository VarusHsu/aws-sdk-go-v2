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

// Creates a multiplayer game session for players. This action creates a game
// session record and assigns an available server process in the specified fleet to
// host the game session. A fleet must have an ACTIVE status before a game session
// can be created in it. To create a game session, specify either fleet ID or alias
// ID and indicate a maximum number of players to allow in the game session. You
// can also provide a name and game-specific properties for this game session. If
// successful, a GameSession () object is returned containing the game session
// properties and other settings you specified. Idempotency tokens. You can add a
// token that uniquely identifies game session requests. This is useful for
// ensuring that game session requests are idempotent. Multiple requests with the
// same idempotency token are processed only once; subsequent requests return the
// original result. All response values are the same with the exception of game
// session status, which may change. Resource creation limits. If you are creating
// a game session on a fleet with a resource creation limit policy in force, then
// you must specify a creator ID. Without this ID, Amazon GameLift has no way to
// evaluate the policy for this new game session request. Player acceptance policy.
// By default, newly created game sessions are open to new players. You can
// restrict new player access by using UpdateGameSession () to change the game
// session's player session creation policy. Game session logs. Logs are retained
// for all active game sessions for 14 days. To access the logs, call
// GetGameSessionLogUrl () to download the log files. Available in Amazon GameLift
// Local.
//
//     * CreateGameSession ()
//
//     * DescribeGameSessions ()
//
//     *
// DescribeGameSessionDetails ()
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
func (c *Client) CreateGameSession(ctx context.Context, params *CreateGameSessionInput, optFns ...func(*Options)) (*CreateGameSessionOutput, error) {
	stack := middleware.NewStack("CreateGameSession", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateGameSessionMiddlewares(stack)
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
	addOpCreateGameSessionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGameSession(options.Region), middleware.Before)
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
			OperationName: "CreateGameSession",
			Err:           err,
		}
	}
	out := result.(*CreateGameSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type CreateGameSessionInput struct {

	// The maximum number of players that can be connected simultaneously to the game
	// session.
	//
	// This member is required.
	MaximumPlayerSessionCount *int32

	// A unique identifier for an alias associated with the fleet to create a game
	// session in. You can use either the alias ID or ARN value. Each request must
	// reference either a fleet ID or alias ID, but not both.
	AliasId *string

	// A unique identifier for a player or entity creating the game session. This ID is
	// used to enforce a resource protection policy (if one exists) that limits the
	// number of concurrent active game sessions one player can have.
	CreatorId *string

	// A descriptive label that is associated with a game session. Session names do not
	// need to be unique.
	Name *string

	// Set of custom properties for a game session, formatted as key:value pairs. These
	// properties are passed to a game server process in the GameSession () object with
	// a request to start a new game session (see Start a Game Session
	// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession)).
	GameProperties []*types.GameProperty

	// Set of custom game session properties, formatted as a single string value. This
	// data is passed to a game server process in the GameSession () object with a
	// request to start a new game session (see Start a Game Session
	// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession)).
	GameSessionData *string

	// This parameter is no longer preferred. Please use IdempotencyToken instead.
	// Custom string that uniquely identifies a request for a new game session. Maximum
	// token length is 48 characters. If provided, this string is included in the new
	// game session's ID. (A game session ARN has the following format:
	// arn:aws:gamelift:::gamesession//.)
	GameSessionId *string

	// A unique identifier for a fleet to create a game session in. You can use either
	// the fleet ID or ARN value. Each request must reference either a fleet ID or
	// alias ID, but not both.
	FleetId *string

	// Custom string that uniquely identifies a request for a new game session. Maximum
	// token length is 48 characters. If provided, this string is included in the new
	// game session's ID. (A game session ARN has the following format:
	// arn:aws:gamelift:::gamesession//.) Idempotency tokens remain in use for 30 days
	// after a game session has ended; game session objects are retained for this time
	// period and then deleted.
	IdempotencyToken *string
}

// Represents the returned data in response to a request action.
type CreateGameSessionOutput struct {

	// Object that describes the newly created game session record.
	GameSession *types.GameSession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateGameSessionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateGameSession{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateGameSession{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateGameSession(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "CreateGameSession",
	}
}
