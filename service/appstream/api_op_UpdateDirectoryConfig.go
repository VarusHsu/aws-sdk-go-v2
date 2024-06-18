// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified Directory Config object in AppStream 2.0. This object
// includes the configuration information required to join fleets and image
// builders to Microsoft Active Directory domains.
func (c *Client) UpdateDirectoryConfig(ctx context.Context, params *UpdateDirectoryConfigInput, optFns ...func(*Options)) (*UpdateDirectoryConfigOutput, error) {
	if params == nil {
		params = &UpdateDirectoryConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDirectoryConfig", params, optFns, c.addOperationUpdateDirectoryConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDirectoryConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDirectoryConfigInput struct {

	// The name of the Directory Config object.
	//
	// This member is required.
	DirectoryName *string

	// The certificate-based authentication properties used to authenticate SAML 2.0
	// Identity Provider (IdP) user identities to Active Directory domain-joined
	// streaming instances. Fallback is turned on by default when certificate-based
	// authentication is Enabled . Fallback allows users to log in using their AD
	// domain password if certificate-based authentication is unsuccessful, or to
	// unlock a desktop lock screen. Enabled_no_directory_login_fallback enables
	// certificate-based authentication, but does not allow users to log in using their
	// AD domain password. Users will be disconnected to re-authenticate using
	// certificates.
	CertificateBasedAuthProperties *types.CertificateBasedAuthProperties

	// The distinguished names of the organizational units for computer accounts.
	OrganizationalUnitDistinguishedNames []string

	// The credentials for the service account used by the fleet or image builder to
	// connect to the directory.
	ServiceAccountCredentials *types.ServiceAccountCredentials

	noSmithyDocumentSerde
}

type UpdateDirectoryConfigOutput struct {

	// Information about the Directory Config object.
	DirectoryConfig *types.DirectoryConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDirectoryConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDirectoryConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDirectoryConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateDirectoryConfig"); err != nil {
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
	if err = addOpUpdateDirectoryConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDirectoryConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDirectoryConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateDirectoryConfig",
	}
}
