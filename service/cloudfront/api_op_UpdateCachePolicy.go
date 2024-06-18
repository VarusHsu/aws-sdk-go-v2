// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a cache policy configuration.
//
// When you update a cache policy configuration, all the fields are updated with
// the values provided in the request. You cannot update some fields independent of
// others. To update a cache policy configuration:
//
//   - Use GetCachePolicyConfig to get the current configuration.
//
//   - Locally modify the fields in the cache policy configuration that you want
//     to update.
//
//   - Call UpdateCachePolicy by providing the entire cache policy configuration,
//     including the fields that you modified and those that you didn't.
func (c *Client) UpdateCachePolicy(ctx context.Context, params *UpdateCachePolicyInput, optFns ...func(*Options)) (*UpdateCachePolicyOutput, error) {
	if params == nil {
		params = &UpdateCachePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCachePolicy", params, optFns, c.addOperationUpdateCachePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCachePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCachePolicyInput struct {

	// A cache policy configuration.
	//
	// This member is required.
	CachePolicyConfig *types.CachePolicyConfig

	// The unique identifier for the cache policy that you are updating. The
	// identifier is returned in a cache behavior's CachePolicyId field in the
	// response to GetDistributionConfig .
	//
	// This member is required.
	Id *string

	// The version of the cache policy that you are updating. The version is returned
	// in the cache policy's ETag field in the response to GetCachePolicyConfig .
	IfMatch *string

	noSmithyDocumentSerde
}

type UpdateCachePolicyOutput struct {

	// A cache policy.
	CachePolicy *types.CachePolicy

	// The current version of the cache policy.
	ETag *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateCachePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpUpdateCachePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpUpdateCachePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateCachePolicy"); err != nil {
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
	if err = addOpUpdateCachePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCachePolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateCachePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateCachePolicy",
	}
}
