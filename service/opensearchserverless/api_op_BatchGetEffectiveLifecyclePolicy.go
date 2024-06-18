// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearchserverless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/opensearchserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of successful and failed retrievals for the OpenSearch
// Serverless indexes. For more information, see [Viewing data lifecycle policies].
//
// [Viewing data lifecycle policies]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-lifecycle.html#serverless-lifecycle-list
func (c *Client) BatchGetEffectiveLifecyclePolicy(ctx context.Context, params *BatchGetEffectiveLifecyclePolicyInput, optFns ...func(*Options)) (*BatchGetEffectiveLifecyclePolicyOutput, error) {
	if params == nil {
		params = &BatchGetEffectiveLifecyclePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetEffectiveLifecyclePolicy", params, optFns, c.addOperationBatchGetEffectiveLifecyclePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetEffectiveLifecyclePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetEffectiveLifecyclePolicyInput struct {

	// The unique identifiers of policy types and resource names.
	//
	// This member is required.
	ResourceIdentifiers []types.LifecyclePolicyResourceIdentifier

	noSmithyDocumentSerde
}

type BatchGetEffectiveLifecyclePolicyOutput struct {

	// A list of lifecycle policies applied to the OpenSearch Serverless indexes.
	EffectiveLifecyclePolicyDetails []types.EffectiveLifecyclePolicyDetail

	// A list of resources for which retrieval failed.
	EffectiveLifecyclePolicyErrorDetails []types.EffectiveLifecyclePolicyErrorDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetEffectiveLifecyclePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpBatchGetEffectiveLifecyclePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpBatchGetEffectiveLifecyclePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchGetEffectiveLifecyclePolicy"); err != nil {
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
	if err = addOpBatchGetEffectiveLifecyclePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetEffectiveLifecyclePolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetEffectiveLifecyclePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchGetEffectiveLifecyclePolicy",
	}
}
