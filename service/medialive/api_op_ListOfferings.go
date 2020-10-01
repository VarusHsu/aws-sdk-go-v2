// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List offerings available for purchase.
func (c *Client) ListOfferings(ctx context.Context, params *ListOfferingsInput, optFns ...func(*Options)) (*ListOfferingsOutput, error) {
	stack := middleware.NewStack("ListOfferings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListOfferingsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListOfferings(options.Region), middleware.Before)
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
			OperationName: "ListOfferings",
			Err:           err,
		}
	}
	out := result.(*ListOfferingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for ListOfferingsRequest
type ListOfferingsInput struct {

	// Filter by codec, 'AVC', 'HEVC', 'MPEG2', or 'AUDIO'
	Codec *string

	// Filter by channel class, 'STANDARD' or 'SINGLE_PIPELINE'
	ChannelClass *string

	// Filter by offering duration, e.g. '12'
	Duration *string

	// Filter by resolution, 'SD', 'HD', 'FHD', or 'UHD'
	Resolution *string

	// Filter by framerate, 'MAX_30_FPS' or 'MAX_60_FPS'
	MaximumFramerate *string

	// Filter to offerings that match the configuration of an existing channel, e.g.
	// '2345678' (a channel ID)
	ChannelConfiguration *string

	// Filter by resource type, 'INPUT', 'OUTPUT', 'MULTIPLEX', or 'CHANNEL'
	ResourceType *string

	// Filter by video quality, 'STANDARD', 'ENHANCED', or 'PREMIUM'
	VideoQuality *string

	// Placeholder documentation for __string
	NextToken *string

	// Placeholder documentation for MaxResults
	MaxResults *int32

	// Filter by bitrate, 'MAX_10_MBPS', 'MAX_20_MBPS', or 'MAX_50_MBPS'
	MaximumBitrate *string

	// Filter by special feature, 'ADVANCED_AUDIO' or 'AUDIO_NORMALIZATION'
	SpecialFeature *string
}

// Placeholder documentation for ListOfferingsResponse
type ListOfferingsOutput struct {

	// List of offerings
	Offerings []*types.Offering

	// Token to retrieve the next page of results
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListOfferingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListOfferings{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListOfferings{}, middleware.After)
}

func newServiceMetadataMiddleware_opListOfferings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "ListOfferings",
	}
}
