// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the settings for the Insights event selectors that you configured for
// your trail. GetInsightSelectors shows if CloudTrail Insights event logging is
// enabled on the trail, and if it is, which insight types are enabled. If you run
// GetInsightSelectors on a trail that does not have Insights events enabled, the
// operation throws the exception InsightNotEnabledException For more information,
// see Logging CloudTrail Insights Events for Trails
// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-insights-events-with-cloudtrail.html)
// in the AWS CloudTrail User Guide.
func (c *Client) GetInsightSelectors(ctx context.Context, params *GetInsightSelectorsInput, optFns ...func(*Options)) (*GetInsightSelectorsOutput, error) {
	stack := middleware.NewStack("GetInsightSelectors", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetInsightSelectorsMiddlewares(stack)
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
	addOpGetInsightSelectorsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetInsightSelectors(options.Region), middleware.Before)
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
			OperationName: "GetInsightSelectors",
			Err:           err,
		}
	}
	out := result.(*GetInsightSelectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInsightSelectorsInput struct {

	// Specifies the name of the trail or trail ARN. If you specify a trail name, the
	// string must meet the following requirements:
	//
	//     * Contain only ASCII letters
	// (a-z, A-Z), numbers (0-9), periods (.), underscores (_), or dashes (-)
	//
	//     *
	// Start with a letter or number, and end with a letter or number
	//
	//     * Be between
	// 3 and 128 characters
	//
	//     * Have no adjacent periods, underscores or dashes.
	// Names like my-_namespace and my--namespace are not valid.
	//
	//     * Not be in IP
	// address format (for example, 192.168.5.4)
	//
	// If you specify a trail ARN, it must
	// be in the format: arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail
	//
	// This member is required.
	TrailName *string
}

type GetInsightSelectorsOutput struct {

	// A JSON string that contains the insight types you want to log on a trail. In
	// this release, only ApiCallRateInsight is supported as an insight type.
	InsightSelectors []*types.InsightSelector

	// The Amazon Resource Name (ARN) of a trail for which you want to get Insights
	// selectors.
	TrailARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetInsightSelectorsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetInsightSelectors{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetInsightSelectors{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetInsightSelectors(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudtrail",
		OperationName: "GetInsightSelectors",
	}
}
