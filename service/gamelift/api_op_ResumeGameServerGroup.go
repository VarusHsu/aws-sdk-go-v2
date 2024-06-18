// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	This operation is used with the Amazon GameLift FleetIQ solution and game
//
// server groups.
//
// Reinstates activity on a game server group after it has been suspended. A game
// server group might be suspended by the SuspendGameServerGroupoperation, or it might be suspended
// involuntarily due to a configuration problem. In the second case, you can
// manually resume activity on the group once the configuration problem has been
// resolved. Refer to the game server group status and status reason for more
// information on why group activity is suspended.
//
// To resume activity, specify a game server group ARN and the type of activity to
// be resumed. If successful, a GameServerGroup object is returned showing that
// the resumed activity is no longer listed in SuspendedActions .
//
// # Learn more
//
// [Amazon GameLift FleetIQ Guide]
//
// [Amazon GameLift FleetIQ Guide]: https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/gsg-intro.html
func (c *Client) ResumeGameServerGroup(ctx context.Context, params *ResumeGameServerGroupInput, optFns ...func(*Options)) (*ResumeGameServerGroupOutput, error) {
	if params == nil {
		params = &ResumeGameServerGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResumeGameServerGroup", params, optFns, c.addOperationResumeGameServerGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResumeGameServerGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResumeGameServerGroupInput struct {

	// A unique identifier for the game server group. Use either the name or ARN value.
	//
	// This member is required.
	GameServerGroupName *string

	// The activity to resume for this game server group.
	//
	// This member is required.
	ResumeActions []types.GameServerGroupAction

	noSmithyDocumentSerde
}

type ResumeGameServerGroupOutput struct {

	// An object that describes the game server group resource, with the
	// SuspendedActions property updated to reflect the resumed activity.
	GameServerGroup *types.GameServerGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationResumeGameServerGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpResumeGameServerGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpResumeGameServerGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ResumeGameServerGroup"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpResumeGameServerGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opResumeGameServerGroup(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opResumeGameServerGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ResumeGameServerGroup",
	}
}
