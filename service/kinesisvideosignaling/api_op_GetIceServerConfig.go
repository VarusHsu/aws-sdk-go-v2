// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideosignaling

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideosignaling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the Interactive Connectivity Establishment (ICE) server configuration
// information, including URIs, username, and password which can be used to
// configure the WebRTC connection. The ICE component uses this configuration
// information to setup the WebRTC connection, including authenticating with the
// Traversal Using Relays around NAT (TURN) relay server. TURN is a protocol that
// is used to improve the connectivity of peer-to-peer applications. By providing a
// cloud-based relay service, TURN ensures that a connection can be established
// even when one or more peers are incapable of a direct peer-to-peer connection.
// For more information, see A REST API For Access To TURN Services (https://tools.ietf.org/html/draft-uberti-rtcweb-turn-rest-00)
// . You can invoke this API to establish a fallback mechanism in case either of
// the peers is unable to establish a direct peer-to-peer connection over a
// signaling channel. You must specify either a signaling channel ARN or the client
// ID in order to invoke this API.
func (c *Client) GetIceServerConfig(ctx context.Context, params *GetIceServerConfigInput, optFns ...func(*Options)) (*GetIceServerConfigOutput, error) {
	if params == nil {
		params = &GetIceServerConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIceServerConfig", params, optFns, c.addOperationGetIceServerConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIceServerConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetIceServerConfigInput struct {

	// The ARN of the signaling channel to be used for the peer-to-peer connection
	// between configured peers.
	//
	// This member is required.
	ChannelARN *string

	// Unique identifier for the viewer. Must be unique within the signaling channel.
	ClientId *string

	// Specifies the desired service. Currently, TURN is the only valid value.
	Service types.Service

	// An optional user ID to be associated with the credentials.
	Username *string

	noSmithyDocumentSerde
}

type GetIceServerConfigOutput struct {

	// The list of ICE server information objects.
	IceServerList []types.IceServer

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetIceServerConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetIceServerConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetIceServerConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetIceServerConfig"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpGetIceServerConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetIceServerConfig(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opGetIceServerConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetIceServerConfig",
	}
}
