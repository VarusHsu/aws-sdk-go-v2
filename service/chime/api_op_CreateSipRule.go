// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a SIP rule which can be used to run a SIP media application as a target
// for a specific trigger type.
//
// This API is is no longer supported and will not be updated. We recommend using
// the latest version, [CreateSipRule], in the Amazon Chime SDK.
//
// Using the latest version requires migrating to a dedicated namespace. For more
// information, refer to [Migrating from the Amazon Chime namespace]in the Amazon Chime SDK Developer Guide.
//
// Deprecated: Replaced by CreateSipRule in the Amazon Chime SDK Voice Namespace
//
// [Migrating from the Amazon Chime namespace]: https://docs.aws.amazon.com/chime-sdk/latest/dg/migrate-from-chm-namespace.html
// [CreateSipRule]: https://docs.aws.amazon.com/chime-sdk/latest/APIReference/API_voice-chime_CreateSipRule.html
func (c *Client) CreateSipRule(ctx context.Context, params *CreateSipRuleInput, optFns ...func(*Options)) (*CreateSipRuleOutput, error) {
	if params == nil {
		params = &CreateSipRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSipRule", params, optFns, c.addOperationCreateSipRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSipRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSipRuleInput struct {

	// The name of the SIP rule.
	//
	// This member is required.
	Name *string

	// List of SIP media applications with priority and AWS Region. Only one SIP
	// application per AWS Region can be used.
	//
	// This member is required.
	TargetApplications []types.SipRuleTargetApplication

	// The type of trigger assigned to the SIP rule in TriggerValue , currently
	// RequestUriHostname or ToPhoneNumber .
	//
	// This member is required.
	TriggerType types.SipRuleTriggerType

	// If TriggerType is RequestUriHostname , the value can be the outbound host name
	// of an Amazon Chime Voice Connector. If TriggerType is ToPhoneNumber , the value
	// can be a customer-owned phone number in the E164 format. The SipMediaApplication
	// specified in the SipRule is triggered if the request URI in an incoming SIP
	// request matches the RequestUriHostname , or if the To header in the incoming
	// SIP request matches the ToPhoneNumber value.
	//
	// This member is required.
	TriggerValue *string

	// Enables or disables a rule. You must disable rules before you can delete them.
	Disabled *bool

	noSmithyDocumentSerde
}

type CreateSipRuleOutput struct {

	// Returns the SIP rule information, including the rule ID, triggers, and target
	// applications.
	SipRule *types.SipRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSipRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSipRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSipRule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateSipRule"); err != nil {
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
	if err = addOpCreateSipRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSipRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSipRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateSipRule",
	}
}
