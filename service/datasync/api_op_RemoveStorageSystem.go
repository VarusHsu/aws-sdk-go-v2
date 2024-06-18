// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Permanently removes a storage system resource from DataSync Discovery,
// including the associated discovery jobs, collected data, and recommendations.
func (c *Client) RemoveStorageSystem(ctx context.Context, params *RemoveStorageSystemInput, optFns ...func(*Options)) (*RemoveStorageSystemOutput, error) {
	if params == nil {
		params = &RemoveStorageSystemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveStorageSystem", params, optFns, c.addOperationRemoveStorageSystemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveStorageSystemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveStorageSystemInput struct {

	// Specifies the Amazon Resource Name (ARN) of the storage system that you want to
	// permanently remove from DataSync Discovery.
	//
	// This member is required.
	StorageSystemArn *string

	noSmithyDocumentSerde
}

type RemoveStorageSystemOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRemoveStorageSystemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRemoveStorageSystem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRemoveStorageSystem{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RemoveStorageSystem"); err != nil {
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
	if err = addEndpointPrefix_opRemoveStorageSystemMiddleware(stack); err != nil {
		return err
	}
	if err = addOpRemoveStorageSystemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveStorageSystem(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opRemoveStorageSystemMiddleware struct {
}

func (*endpointPrefix_opRemoveStorageSystemMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opRemoveStorageSystemMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "discovery-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opRemoveStorageSystemMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opRemoveStorageSystemMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opRemoveStorageSystem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RemoveStorageSystem",
	}
}
