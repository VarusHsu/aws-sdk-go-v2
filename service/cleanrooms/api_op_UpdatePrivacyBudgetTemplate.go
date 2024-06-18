// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanrooms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cleanrooms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the privacy budget template for the specified membership.
func (c *Client) UpdatePrivacyBudgetTemplate(ctx context.Context, params *UpdatePrivacyBudgetTemplateInput, optFns ...func(*Options)) (*UpdatePrivacyBudgetTemplateOutput, error) {
	if params == nil {
		params = &UpdatePrivacyBudgetTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePrivacyBudgetTemplate", params, optFns, c.addOperationUpdatePrivacyBudgetTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePrivacyBudgetTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePrivacyBudgetTemplateInput struct {

	// A unique identifier for one of your memberships for a collaboration. The
	// privacy budget template is updated in the collaboration that this membership
	// belongs to. Accepts a membership ID.
	//
	// This member is required.
	MembershipIdentifier *string

	// A unique identifier for your privacy budget template that you want to update.
	//
	// This member is required.
	PrivacyBudgetTemplateIdentifier *string

	// Specifies the type of the privacy budget template.
	//
	// This member is required.
	PrivacyBudgetType types.PrivacyBudgetType

	// Specifies the epsilon and noise parameters for the privacy budget template.
	Parameters types.PrivacyBudgetTemplateUpdateParameters

	noSmithyDocumentSerde
}

type UpdatePrivacyBudgetTemplateOutput struct {

	// Summary of the privacy budget template.
	//
	// This member is required.
	PrivacyBudgetTemplate *types.PrivacyBudgetTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePrivacyBudgetTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdatePrivacyBudgetTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdatePrivacyBudgetTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdatePrivacyBudgetTemplate"); err != nil {
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
	if err = addOpUpdatePrivacyBudgetTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePrivacyBudgetTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePrivacyBudgetTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdatePrivacyBudgetTemplate",
	}
}
