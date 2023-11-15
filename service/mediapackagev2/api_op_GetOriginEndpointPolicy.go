// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackagev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the specified origin endpoint policy that's configured in AWS
// Elemental MediaPackage.
func (c *Client) GetOriginEndpointPolicy(ctx context.Context, params *GetOriginEndpointPolicyInput, optFns ...func(*Options)) (*GetOriginEndpointPolicyOutput, error) {
	if params == nil {
		params = &GetOriginEndpointPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetOriginEndpointPolicy", params, optFns, c.addOperationGetOriginEndpointPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetOriginEndpointPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetOriginEndpointPolicyInput struct {

	// The name that describes the channel group. The name is the primary identifier
	// for the channel group, and must be unique for your account in the AWS Region.
	//
	// This member is required.
	ChannelGroupName *string

	// The name that describes the channel. The name is the primary identifier for the
	// channel, and must be unique for your account in the AWS Region and channel
	// group.
	//
	// This member is required.
	ChannelName *string

	// The name that describes the origin endpoint. The name is the primary identifier
	// for the origin endpoint, and and must be unique for your account in the AWS
	// Region and channel.
	//
	// This member is required.
	OriginEndpointName *string

	noSmithyDocumentSerde
}

type GetOriginEndpointPolicyOutput struct {

	// The name that describes the channel group. The name is the primary identifier
	// for the channel group, and must be unique for your account in the AWS Region.
	//
	// This member is required.
	ChannelGroupName *string

	// The name that describes the channel. The name is the primary identifier for the
	// channel, and must be unique for your account in the AWS Region and channel
	// group.
	//
	// This member is required.
	ChannelName *string

	// The name that describes the origin endpoint. The name is the primary identifier
	// for the origin endpoint, and and must be unique for your account in the AWS
	// Region and channel.
	//
	// This member is required.
	OriginEndpointName *string

	// The policy assigned to the origin endpoint.
	//
	// This member is required.
	Policy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetOriginEndpointPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetOriginEndpointPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetOriginEndpointPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetOriginEndpointPolicy"); err != nil {
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
	if err = addOpGetOriginEndpointPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetOriginEndpointPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetOriginEndpointPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetOriginEndpointPolicy",
	}
}
