// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the effects of an organization's access control rules as they apply to a
// specified IPv4 address, access protocol action, and user ID or impersonation
// role ID. You must provide either the user ID or impersonation role ID.
// Impersonation role ID can only be used with Action EWS.
func (c *Client) GetAccessControlEffect(ctx context.Context, params *GetAccessControlEffectInput, optFns ...func(*Options)) (*GetAccessControlEffectOutput, error) {
	if params == nil {
		params = &GetAccessControlEffectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAccessControlEffect", params, optFns, c.addOperationGetAccessControlEffectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAccessControlEffectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAccessControlEffectInput struct {

	// The access protocol action. Valid values include ActiveSync , AutoDiscover , EWS
	// , IMAP , SMTP , WindowsOutlook , and WebMail .
	//
	// This member is required.
	Action *string

	// The IPv4 address.
	//
	// This member is required.
	IpAddress *string

	// The identifier for the organization.
	//
	// This member is required.
	OrganizationId *string

	// The impersonation role ID.
	ImpersonationRoleId *string

	// The user ID.
	UserId *string

	noSmithyDocumentSerde
}

type GetAccessControlEffectOutput struct {

	// The rule effect.
	Effect types.AccessControlRuleEffect

	// The rules that match the given parameters, resulting in an effect.
	MatchedRules []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAccessControlEffectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetAccessControlEffect{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAccessControlEffect{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAccessControlEffect"); err != nil {
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
	if err = addOpGetAccessControlEffectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAccessControlEffect(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAccessControlEffect(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAccessControlEffect",
	}
}
