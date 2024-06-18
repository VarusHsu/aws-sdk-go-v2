// Code generated by smithy-go-codegen DO NOT EDIT.

package mediatailor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a channel. For information about MediaTailor channels, see [Working with channels] in the
// MediaTailor User Guide.
//
// [Working with channels]: https://docs.aws.amazon.com/mediatailor/latest/ug/channel-assembly-channels.html
func (c *Client) CreateChannel(ctx context.Context, params *CreateChannelInput, optFns ...func(*Options)) (*CreateChannelOutput, error) {
	if params == nil {
		params = &CreateChannelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateChannel", params, optFns, c.addOperationCreateChannelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateChannelInput struct {

	// The name of the channel.
	//
	// This member is required.
	ChannelName *string

	// The channel's output properties.
	//
	// This member is required.
	Outputs []types.RequestOutputItem

	// The type of playback mode to use for this channel.
	//
	// LINEAR - The programs in the schedule play once back-to-back in the schedule.
	//
	// LOOP - The programs in the schedule play back-to-back in an endless loop. When
	// the last program in the schedule stops playing, playback loops back to the first
	// program in the schedule.
	//
	// This member is required.
	PlaybackMode types.PlaybackMode

	// The list of audiences defined in channel.
	Audiences []string

	// The slate used to fill gaps between programs in the schedule. You must
	// configure filler slate if your channel uses the LINEAR PlaybackMode .
	// MediaTailor doesn't support filler slate for channels using the LOOP PlaybackMode
	// .
	FillerSlate *types.SlateSource

	// The tags to assign to the channel. Tags are key-value pairs that you can
	// associate with Amazon resources to help with organization, access control, and
	// cost tracking. For more information, see [Tagging AWS Elemental MediaTailor Resources].
	//
	// [Tagging AWS Elemental MediaTailor Resources]: https://docs.aws.amazon.com/mediatailor/latest/ug/tagging.html
	Tags map[string]string

	// The tier of the channel.
	Tier types.Tier

	//  The time-shifted viewing configuration you want to associate to the channel.
	TimeShiftConfiguration *types.TimeShiftConfiguration

	noSmithyDocumentSerde
}

type CreateChannelOutput struct {

	// The Amazon Resource Name (ARN) to assign to the channel.
	Arn *string

	// The list of audiences defined in channel.
	Audiences []string

	// The name to assign to the channel.
	ChannelName *string

	// Indicates whether the channel is in a running state or not.
	ChannelState types.ChannelState

	// The timestamp of when the channel was created.
	CreationTime *time.Time

	// Contains information about the slate used to fill gaps between programs in the
	// schedule.
	FillerSlate *types.SlateSource

	// The timestamp of when the channel was last modified.
	LastModifiedTime *time.Time

	// The output properties to assign to the channel.
	Outputs []types.ResponseOutputItem

	// The playback mode to assign to the channel.
	PlaybackMode *string

	// The tags to assign to the channel. Tags are key-value pairs that you can
	// associate with Amazon resources to help with organization, access control, and
	// cost tracking. For more information, see [Tagging AWS Elemental MediaTailor Resources].
	//
	// [Tagging AWS Elemental MediaTailor Resources]: https://docs.aws.amazon.com/mediatailor/latest/ug/tagging.html
	Tags map[string]string

	// The tier of the channel.
	Tier *string

	//  The time-shifted viewing configuration assigned to the channel.
	TimeShiftConfiguration *types.TimeShiftConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateChannelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateChannel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateChannel{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateChannel"); err != nil {
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
	if err = addOpCreateChannelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateChannel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateChannel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateChannel",
	}
}
