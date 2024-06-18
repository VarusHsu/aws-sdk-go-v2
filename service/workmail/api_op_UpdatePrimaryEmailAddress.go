// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the primary email for a user, group, or resource. The current email is
// moved into the list of aliases (or swapped between an existing alias and the
// current primary email), and the email provided in the input is promoted as the
// primary.
func (c *Client) UpdatePrimaryEmailAddress(ctx context.Context, params *UpdatePrimaryEmailAddressInput, optFns ...func(*Options)) (*UpdatePrimaryEmailAddressOutput, error) {
	if params == nil {
		params = &UpdatePrimaryEmailAddressInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePrimaryEmailAddress", params, optFns, c.addOperationUpdatePrimaryEmailAddressMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePrimaryEmailAddressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePrimaryEmailAddressInput struct {

	// The value of the email to be updated as primary.
	//
	// This member is required.
	Email *string

	// The user, group, or resource to update.
	//
	// The identifier can accept UseriD, ResourceId, or GroupId, Username,
	// Resourcename, or Groupname, or email. The following identity formats are
	// available:
	//
	//   - Entity ID: 12345678-1234-1234-1234-123456789012,
	//   r-0123456789a0123456789b0123456789, or
	//   S-1-1-12-1234567890-123456789-123456789-1234
	//
	//   - Email address: entity@domain.tld
	//
	//   - Entity name: entity
	//
	// This member is required.
	EntityId *string

	// The organization that contains the user, group, or resource to update.
	//
	// This member is required.
	OrganizationId *string

	noSmithyDocumentSerde
}

type UpdatePrimaryEmailAddressOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePrimaryEmailAddressMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdatePrimaryEmailAddress{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdatePrimaryEmailAddress{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdatePrimaryEmailAddress"); err != nil {
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
	if err = addOpUpdatePrimaryEmailAddressValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePrimaryEmailAddress(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePrimaryEmailAddress(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdatePrimaryEmailAddress",
	}
}
