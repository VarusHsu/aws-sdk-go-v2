// Code generated by smithy-go-codegen DO NOT EDIT.

package personalizeruntime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalizeruntime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of recommended items. The required input depends on the recipe
// type used to create the solution backing the campaign, as follows:
//
//     *
// RELATED_ITEMS - itemId required, userId not used
//
//     * USER_PERSONALIZATION -
// itemId optional, userId required
//
// Campaigns that are backed by a solution
// created using a recipe of type PERSONALIZED_RANKING use the API.
func (c *Client) GetRecommendations(ctx context.Context, params *GetRecommendationsInput, optFns ...func(*Options)) (*GetRecommendationsOutput, error) {
	stack := middleware.NewStack("GetRecommendations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetRecommendationsMiddlewares(stack)
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
	addOpGetRecommendationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRecommendations(options.Region), middleware.Before)
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
			OperationName: "GetRecommendations",
			Err:           err,
		}
	}
	out := result.(*GetRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRecommendationsInput struct {

	// The contextual metadata to use when getting recommendations. Contextual metadata
	// includes any interaction information that might be relevant when getting a
	// user's recommendations, such as the user's current location or device type.
	Context map[string]*string

	// The user ID to provide recommendations for. Required for USER_PERSONALIZATION
	// recipe type.
	UserId *string

	// The Amazon Resource Name (ARN) of the campaign to use for getting
	// recommendations.
	//
	// This member is required.
	CampaignArn *string

	// The ARN of the filter to apply to the returned recommendations. For more
	// information, see Using Filters with Amazon Personalize.
	FilterArn *string

	// The item ID to provide recommendations for. Required for RELATED_ITEMS recipe
	// type.
	ItemId *string

	// The number of results to return. The default is 25. The maximum is 500.
	NumResults *int32
}

type GetRecommendationsOutput struct {

	// A list of recommendations sorted in ascending order by prediction score. There
	// can be a maximum of 500 items in the list.
	ItemList []*types.PredictedItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetRecommendationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetRecommendations{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRecommendations{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetRecommendations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "GetRecommendations",
	}
}
