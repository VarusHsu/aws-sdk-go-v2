// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmessaging

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a specified number of users and bots to a channel.
func (c *Client) BatchCreateChannelMembership(ctx context.Context, params *BatchCreateChannelMembershipInput, optFns ...func(*Options)) (*BatchCreateChannelMembershipOutput, error) {
	if params == nil {
		params = &BatchCreateChannelMembershipInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchCreateChannelMembership", params, optFns, c.addOperationBatchCreateChannelMembershipMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchCreateChannelMembershipOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchCreateChannelMembershipInput struct {

	// The ARN of the channel to which you're adding users or bots.
	//
	// This member is required.
	ChannelArn *string

	// The ARN of the AppInstanceUser or AppInstanceBot that makes the API call.
	//
	// This member is required.
	ChimeBearer *string

	// The ARNs of the members you want to add to the channel. Only AppInstanceUsers
	// and AppInstanceBots can be added as a channel member.
	//
	// This member is required.
	MemberArns []string

	// The ID of the SubChannel in the request.
	//
	// Only required when creating membership in a SubChannel for a moderator in an
	// elastic channel.
	SubChannelId *string

	// The membership type of a user, DEFAULT or HIDDEN . Default members are always
	// returned as part of ListChannelMemberships . Hidden members are only returned if
	// the type filter in ListChannelMemberships equals HIDDEN . Otherwise hidden
	// members are not returned. This is only supported by moderators.
	Type types.ChannelMembershipType

	noSmithyDocumentSerde
}

type BatchCreateChannelMembershipOutput struct {

	// The list of channel memberships in the response.
	BatchChannelMemberships *types.BatchChannelMemberships

	// If the action fails for one or more of the memberships in the request, a list
	// of the memberships is returned, along with error codes and error messages.
	Errors []types.BatchCreateChannelMembershipError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchCreateChannelMembershipMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchCreateChannelMembership{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchCreateChannelMembership{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchCreateChannelMembership"); err != nil {
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
	if err = addOpBatchCreateChannelMembershipValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchCreateChannelMembership(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchCreateChannelMembership(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchCreateChannelMembership",
	}
}
