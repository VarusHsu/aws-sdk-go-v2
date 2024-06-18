// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes a user from a channel's ban list.
//
// The x-amz-chime-bearer request header is mandatory. Use the AppInstanceUserArn
// of the user that makes the API call as the value in the header.
//
// This API is is no longer supported and will not be updated. We recommend using
// the latest version, [DeleteChannelBan], in the Amazon Chime SDK.
//
// Using the latest version requires migrating to a dedicated namespace. For more
// information, refer to [Migrating from the Amazon Chime namespace]in the Amazon Chime SDK Developer Guide.
//
// Deprecated: Replaced by DeleteChannelBan in the Amazon Chime SDK Messaging
// Namespace
//
// [DeleteChannelBan]: https://docs.aws.amazon.com/chime-sdk/latest/APIReference/API_messaging-chime_DeleteChannelBan.html
// [Migrating from the Amazon Chime namespace]: https://docs.aws.amazon.com/chime-sdk/latest/dg/migrate-from-chm-namespace.html
func (c *Client) DeleteChannelBan(ctx context.Context, params *DeleteChannelBanInput, optFns ...func(*Options)) (*DeleteChannelBanOutput, error) {
	if params == nil {
		params = &DeleteChannelBanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteChannelBan", params, optFns, c.addOperationDeleteChannelBanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteChannelBanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteChannelBanInput struct {

	// The ARN of the channel from which the AppInstanceUser was banned.
	//
	// This member is required.
	ChannelArn *string

	// The ARN of the AppInstanceUser that you want to reinstate.
	//
	// This member is required.
	MemberArn *string

	// The AppInstanceUserArn of the user that makes the API call.
	ChimeBearer *string

	noSmithyDocumentSerde
}

type DeleteChannelBanOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteChannelBanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteChannelBan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteChannelBan{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteChannelBan"); err != nil {
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
	if err = addEndpointPrefix_opDeleteChannelBanMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDeleteChannelBanValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteChannelBan(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opDeleteChannelBanMiddleware struct {
}

func (*endpointPrefix_opDeleteChannelBanMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDeleteChannelBanMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "messaging-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opDeleteChannelBanMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opDeleteChannelBanMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opDeleteChannelBan(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteChannelBan",
	}
}
