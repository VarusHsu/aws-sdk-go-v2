// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the Certificate expiry time for a group.
func (c *Client) UpdateGroupCertificateConfiguration(ctx context.Context, params *UpdateGroupCertificateConfigurationInput, optFns ...func(*Options)) (*UpdateGroupCertificateConfigurationOutput, error) {
	if params == nil {
		params = &UpdateGroupCertificateConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateGroupCertificateConfiguration", params, optFns, c.addOperationUpdateGroupCertificateConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateGroupCertificateConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGroupCertificateConfigurationInput struct {

	// The ID of the Greengrass group.
	//
	// This member is required.
	GroupId *string

	// The amount of time remaining before the certificate expires, in milliseconds.
	CertificateExpiryInMilliseconds *string

	noSmithyDocumentSerde
}

type UpdateGroupCertificateConfigurationOutput struct {

	// The amount of time remaining before the certificate authority expires, in
	// milliseconds.
	CertificateAuthorityExpiryInMilliseconds *string

	// The amount of time remaining before the certificate expires, in milliseconds.
	CertificateExpiryInMilliseconds *string

	// The ID of the group certificate configuration.
	GroupId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateGroupCertificateConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateGroupCertificateConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateGroupCertificateConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateGroupCertificateConfiguration"); err != nil {
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
	if err = addOpUpdateGroupCertificateConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGroupCertificateConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateGroupCertificateConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateGroupCertificateConfiguration",
	}
}
