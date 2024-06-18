// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoveryreadiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets recommendations about architecture designs for improving resiliency for an
// application, based on a recovery group.
func (c *Client) GetArchitectureRecommendations(ctx context.Context, params *GetArchitectureRecommendationsInput, optFns ...func(*Options)) (*GetArchitectureRecommendationsOutput, error) {
	if params == nil {
		params = &GetArchitectureRecommendationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetArchitectureRecommendations", params, optFns, c.addOperationGetArchitectureRecommendationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetArchitectureRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetArchitectureRecommendationsInput struct {

	// The name of a recovery group.
	//
	// This member is required.
	RecoveryGroupName *string

	// The number of objects that you want to return with this call.
	MaxResults *int32

	// The token that identifies which batch of results you want to see.
	NextToken *string

	noSmithyDocumentSerde
}

type GetArchitectureRecommendationsOutput struct {

	// The time that a recovery group was last assessed for recommendations, in UTC
	// ISO-8601 format.
	LastAuditTimestamp *time.Time

	// The token that identifies which batch of results you want to see.
	NextToken *string

	// A list of the recommendations for the customer's application.
	Recommendations []types.Recommendation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetArchitectureRecommendationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetArchitectureRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetArchitectureRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetArchitectureRecommendations"); err != nil {
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
	if err = addOpGetArchitectureRecommendationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetArchitectureRecommendations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetArchitectureRecommendations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetArchitectureRecommendations",
	}
}
