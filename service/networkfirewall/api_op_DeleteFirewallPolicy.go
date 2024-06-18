// Code generated by smithy-go-codegen DO NOT EDIT.

package networkfirewall

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified FirewallPolicy.
func (c *Client) DeleteFirewallPolicy(ctx context.Context, params *DeleteFirewallPolicyInput, optFns ...func(*Options)) (*DeleteFirewallPolicyOutput, error) {
	if params == nil {
		params = &DeleteFirewallPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteFirewallPolicy", params, optFns, c.addOperationDeleteFirewallPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteFirewallPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteFirewallPolicyInput struct {

	// The Amazon Resource Name (ARN) of the firewall policy.
	//
	// You must specify the ARN or the name, and you can specify both.
	FirewallPolicyArn *string

	// The descriptive name of the firewall policy. You can't change the name of a
	// firewall policy after you create it.
	//
	// You must specify the ARN or the name, and you can specify both.
	FirewallPolicyName *string

	noSmithyDocumentSerde
}

type DeleteFirewallPolicyOutput struct {

	// The object containing the definition of the FirewallPolicyResponse that you asked to delete.
	//
	// This member is required.
	FirewallPolicyResponse *types.FirewallPolicyResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteFirewallPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteFirewallPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteFirewallPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteFirewallPolicy"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteFirewallPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteFirewallPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteFirewallPolicy",
	}
}
