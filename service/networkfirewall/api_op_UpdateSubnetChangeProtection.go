// Code generated by smithy-go-codegen DO NOT EDIT.

package networkfirewall

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

func (c *Client) UpdateSubnetChangeProtection(ctx context.Context, params *UpdateSubnetChangeProtectionInput, optFns ...func(*Options)) (*UpdateSubnetChangeProtectionOutput, error) {
	if params == nil {
		params = &UpdateSubnetChangeProtectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSubnetChangeProtection", params, optFns, c.addOperationUpdateSubnetChangeProtectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSubnetChangeProtectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSubnetChangeProtectionInput struct {

	// A setting indicating whether the firewall is protected against changes to the
	// subnet associations. Use this setting to protect against accidentally modifying
	// the subnet associations for a firewall that is in use. When you create a
	// firewall, the operation initializes this setting to TRUE .
	//
	// This member is required.
	SubnetChangeProtection bool

	// The Amazon Resource Name (ARN) of the firewall.
	//
	// You must specify the ARN or the name, and you can specify both.
	FirewallArn *string

	// The descriptive name of the firewall. You can't change the name of a firewall
	// after you create it.
	//
	// You must specify the ARN or the name, and you can specify both.
	FirewallName *string

	// An optional token that you can use for optimistic locking. Network Firewall
	// returns a token to your requests that access the firewall. The token marks the
	// state of the firewall resource at the time of the request.
	//
	// To make an unconditional change to the firewall, omit the token in your update
	// request. Without the token, Network Firewall performs your updates regardless of
	// whether the firewall has changed since you last retrieved it.
	//
	// To make a conditional change to the firewall, provide the token in your update
	// request. Network Firewall uses the token to ensure that the firewall hasn't
	// changed since you last retrieved it. If it has changed, the operation fails with
	// an InvalidTokenException . If this happens, retrieve the firewall again to get a
	// current copy of it with a new token. Reapply your changes as needed, then try
	// the operation again using the new token.
	UpdateToken *string

	noSmithyDocumentSerde
}

type UpdateSubnetChangeProtectionOutput struct {

	// The Amazon Resource Name (ARN) of the firewall.
	FirewallArn *string

	// The descriptive name of the firewall. You can't change the name of a firewall
	// after you create it.
	FirewallName *string

	// A setting indicating whether the firewall is protected against changes to the
	// subnet associations. Use this setting to protect against accidentally modifying
	// the subnet associations for a firewall that is in use. When you create a
	// firewall, the operation initializes this setting to TRUE .
	SubnetChangeProtection bool

	// An optional token that you can use for optimistic locking. Network Firewall
	// returns a token to your requests that access the firewall. The token marks the
	// state of the firewall resource at the time of the request.
	//
	// To make an unconditional change to the firewall, omit the token in your update
	// request. Without the token, Network Firewall performs your updates regardless of
	// whether the firewall has changed since you last retrieved it.
	//
	// To make a conditional change to the firewall, provide the token in your update
	// request. Network Firewall uses the token to ensure that the firewall hasn't
	// changed since you last retrieved it. If it has changed, the operation fails with
	// an InvalidTokenException . If this happens, retrieve the firewall again to get a
	// current copy of it with a new token. Reapply your changes as needed, then try
	// the operation again using the new token.
	UpdateToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSubnetChangeProtectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateSubnetChangeProtection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateSubnetChangeProtection{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateSubnetChangeProtection"); err != nil {
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
	if err = addOpUpdateSubnetChangeProtectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSubnetChangeProtection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateSubnetChangeProtection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateSubnetChangeProtection",
	}
}
