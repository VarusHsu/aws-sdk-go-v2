// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/fms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Firewall Manager policy.
//
// A Firewall Manager policy is specific to the individual policy type. If you
// want to enforce multiple policy types across accounts, you can create multiple
// policies. You can create more than one policy for each type.
//
// If you add a new account to an organization that you created with
// Organizations, Firewall Manager automatically applies the policy to the
// resources in that account that are within scope of the policy.
//
// Firewall Manager provides the following types of policies:
//
//   - WAF policy - This policy applies WAF web ACL protections to specified
//     accounts and resources.
//
//   - Shield Advanced policy - This policy applies Shield Advanced protection to
//     specified accounts and resources.
//
//   - Security Groups policy - This type of policy gives you control over
//     security groups that are in use throughout your organization in Organizations
//     and lets you enforce a baseline set of rules across your organization.
//
//   - Network ACL policy - This type of policy gives you control over the network
//     ACLs that are in use throughout your organization in Organizations and lets you
//     enforce a baseline set of first and last network ACL rules across your
//     organization.
//
//   - Network Firewall policy - This policy applies Network Firewall protection
//     to your organization's VPCs.
//
//   - DNS Firewall policy - This policy applies Amazon Route 53 Resolver DNS
//     Firewall protections to your organization's VPCs.
//
//   - Third-party firewall policy - This policy applies third-party firewall
//     protections. Third-party firewalls are available by subscription through the
//     Amazon Web Services Marketplace console at [Amazon Web Services Marketplace].
//
//   - Palo Alto Networks Cloud NGFW policy - This policy applies Palo Alto
//     Networks Cloud Next Generation Firewall (NGFW) protections and Palo Alto
//     Networks Cloud NGFW rulestacks to your organization's VPCs.
//
//   - Fortigate CNF policy - This policy applies Fortigate Cloud Native Firewall
//     (CNF) protections. Fortigate CNF is a cloud-centered solution that blocks
//     Zero-Day threats and secures cloud infrastructures with industry-leading
//     advanced threat prevention, smart web application firewalls (WAF), and API
//     protection.
//
// [Amazon Web Services Marketplace]: http://aws.amazon.com/marketplace
func (c *Client) PutPolicy(ctx context.Context, params *PutPolicyInput, optFns ...func(*Options)) (*PutPolicyOutput, error) {
	if params == nil {
		params = &PutPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutPolicy", params, optFns, c.addOperationPutPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutPolicyInput struct {

	// The details of the Firewall Manager policy to be created.
	//
	// This member is required.
	Policy *types.Policy

	// The tags to add to the Amazon Web Services resource.
	TagList []types.Tag

	noSmithyDocumentSerde
}

type PutPolicyOutput struct {

	// The details of the Firewall Manager policy.
	Policy *types.Policy

	// The Amazon Resource Name (ARN) of the policy.
	PolicyArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutPolicy"); err != nil {
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
	if err = addOpPutPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutPolicy",
	}
}
