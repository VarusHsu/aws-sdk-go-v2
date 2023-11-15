// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a single DNS Firewall rule in the specified rule group, using the
// specified domain list.
func (c *Client) CreateFirewallRule(ctx context.Context, params *CreateFirewallRuleInput, optFns ...func(*Options)) (*CreateFirewallRuleOutput, error) {
	if params == nil {
		params = &CreateFirewallRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFirewallRule", params, optFns, c.addOperationCreateFirewallRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFirewallRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFirewallRuleInput struct {

	// The action that DNS Firewall should take on a DNS query when it matches one of
	// the domains in the rule's domain list:
	//   - ALLOW - Permit the request to go through.
	//   - ALERT - Permit the request and send metrics and logs to Cloud Watch.
	//   - BLOCK - Disallow the request. This option requires additional details in the
	//   rule's BlockResponse .
	//
	// This member is required.
	Action types.Action

	// A unique string that identifies the request and that allows you to retry failed
	// requests without the risk of running the operation twice. CreatorRequestId can
	// be any unique string, for example, a date/time stamp.
	//
	// This member is required.
	CreatorRequestId *string

	// The ID of the domain list that you want to use in the rule.
	//
	// This member is required.
	FirewallDomainListId *string

	// The unique identifier of the firewall rule group where you want to create the
	// rule.
	//
	// This member is required.
	FirewallRuleGroupId *string

	// A name that lets you identify the rule in the rule group.
	//
	// This member is required.
	Name *string

	// The setting that determines the processing order of the rule in the rule group.
	// DNS Firewall processes the rules in a rule group by order of priority, starting
	// from the lowest setting. You must specify a unique priority for each rule in a
	// rule group. To make it easier to insert rules later, leave space between the
	// numbers, for example, use 100, 200, and so on. You can change the priority
	// setting for the rules in a rule group at any time.
	//
	// This member is required.
	Priority *int32

	// The DNS record's type. This determines the format of the record value that you
	// provided in BlockOverrideDomain . Used for the rule action BLOCK with a
	// BlockResponse setting of OVERRIDE . This setting is required if the
	// BlockResponse setting is OVERRIDE .
	BlockOverrideDnsType types.BlockOverrideDnsType

	// The custom DNS record to send back in response to the query. Used for the rule
	// action BLOCK with a BlockResponse setting of OVERRIDE . This setting is required
	// if the BlockResponse setting is OVERRIDE .
	BlockOverrideDomain *string

	// The recommended amount of time, in seconds, for the DNS resolver or web browser
	// to cache the provided override record. Used for the rule action BLOCK with a
	// BlockResponse setting of OVERRIDE . This setting is required if the
	// BlockResponse setting is OVERRIDE .
	BlockOverrideTtl *int32

	// The way that you want DNS Firewall to block the request, used with the rule
	// action setting BLOCK .
	//   - NODATA - Respond indicating that the query was successful, but no response
	//   is available for it.
	//   - NXDOMAIN - Respond indicating that the domain name that's in the query
	//   doesn't exist.
	//   - OVERRIDE - Provide a custom override in the response. This option requires
	//   custom handling details in the rule's BlockOverride* settings.
	// This setting is required if the rule action setting is BLOCK .
	BlockResponse types.BlockResponse

	noSmithyDocumentSerde
}

type CreateFirewallRuleOutput struct {

	// The firewall rule that you just created.
	FirewallRule *types.FirewallRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFirewallRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateFirewallRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateFirewallRule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateFirewallRule"); err != nil {
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
	if err = addIdempotencyToken_opCreateFirewallRuleMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateFirewallRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFirewallRule(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateFirewallRule struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateFirewallRule) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateFirewallRule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateFirewallRuleInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateFirewallRuleInput ")
	}

	if input.CreatorRequestId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.CreatorRequestId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateFirewallRuleMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateFirewallRule{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateFirewallRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateFirewallRule",
	}
}
