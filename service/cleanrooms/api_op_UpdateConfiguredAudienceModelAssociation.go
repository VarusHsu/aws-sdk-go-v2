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

// Provides the details necessary to update a configured audience model
// association.
func (c *Client) UpdateConfiguredAudienceModelAssociation(ctx context.Context, params *UpdateConfiguredAudienceModelAssociationInput, optFns ...func(*Options)) (*UpdateConfiguredAudienceModelAssociationOutput, error) {
	if params == nil {
		params = &UpdateConfiguredAudienceModelAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateConfiguredAudienceModelAssociation", params, optFns, c.addOperationUpdateConfiguredAudienceModelAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateConfiguredAudienceModelAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateConfiguredAudienceModelAssociationInput struct {

	// A unique identifier for the configured audience model association that you want
	// to update.
	//
	// This member is required.
	ConfiguredAudienceModelAssociationIdentifier *string

	// A unique identifier of the membership that contains the configured audience
	// model association that you want to update.
	//
	// This member is required.
	MembershipIdentifier *string

	// A new description for the configured audience model association.
	Description *string

	// A new name for the configured audience model association.
	Name *string

	noSmithyDocumentSerde
}

type UpdateConfiguredAudienceModelAssociationOutput struct {

	// Details about the configured audience model association that you updated.
	//
	// This member is required.
	ConfiguredAudienceModelAssociation *types.ConfiguredAudienceModelAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateConfiguredAudienceModelAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateConfiguredAudienceModelAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateConfiguredAudienceModelAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateConfiguredAudienceModelAssociation"); err != nil {
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
	if err = addOpUpdateConfiguredAudienceModelAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateConfiguredAudienceModelAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateConfiguredAudienceModelAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateConfiguredAudienceModelAssociation",
	}
}
