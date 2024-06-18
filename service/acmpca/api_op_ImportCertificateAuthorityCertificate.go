// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Imports a signed private CA certificate into Amazon Web Services Private CA.
// This action is used when you are using a chain of trust whose root is located
// outside Amazon Web Services Private CA. Before you can call this action, the
// following preparations must in place:
//
//   - In Amazon Web Services Private CA, call the [CreateCertificateAuthority]action to create the private CA
//     that you plan to back with the imported certificate.
//
//   - Call the [GetCertificateAuthorityCsr]action to generate a certificate signing request (CSR).
//
//   - Sign the CSR using a root or intermediate CA hosted by either an
//     on-premises PKI hierarchy or by a commercial CA.
//
//   - Create a certificate chain and copy the signed certificate and the
//     certificate chain to your working directory.
//
// Amazon Web Services Private CA supports three scenarios for installing a CA
// certificate:
//
//   - Installing a certificate for a root CA hosted by Amazon Web Services
//     Private CA.
//
//   - Installing a subordinate CA certificate whose parent authority is hosted by
//     Amazon Web Services Private CA.
//
//   - Installing a subordinate CA certificate whose parent authority is
//     externally hosted.
//
// The following additional requirements apply when you import a CA certificate.
//
//   - Only a self-signed certificate can be imported as a root CA.
//
//   - A self-signed certificate cannot be imported as a subordinate CA.
//
//   - Your certificate chain must not include the private CA certificate that you
//     are importing.
//
//   - Your root CA must be the last certificate in your chain. The subordinate
//     certificate, if any, that your root CA signed must be next to last. The
//     subordinate certificate signed by the preceding subordinate CA must come next,
//     and so on until your chain is built.
//
//   - The chain must be PEM-encoded.
//
//   - The maximum allowed size of a certificate is 32 KB.
//
//   - The maximum allowed size of a certificate chain is 2 MB.
//
// # Enforcement of Critical Constraints
//
// Amazon Web Services Private CA allows the following extensions to be marked
// critical in the imported CA certificate or chain.
//
//   - Authority key identifier
//
//   - Basic constraints (must be marked critical)
//
//   - Certificate policies
//
//   - Extended key usage
//
//   - Inhibit anyPolicy
//
//   - Issuer alternative name
//
//   - Key usage
//
//   - Name constraints
//
//   - Policy mappings
//
//   - Subject alternative name
//
//   - Subject directory attributes
//
//   - Subject key identifier
//
//   - Subject information access
//
// Amazon Web Services Private CA rejects the following extensions when they are
// marked critical in an imported CA certificate or chain.
//
//   - Authority information access
//
//   - CRL distribution points
//
//   - Freshest CRL
//
//   - Policy constraints
//
// Amazon Web Services Private Certificate Authority will also reject any other
// extension marked as critical not contained on the preceding list of allowed
// extensions.
//
// [GetCertificateAuthorityCsr]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_GetCertificateAuthorityCsr.html
// [CreateCertificateAuthority]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html
func (c *Client) ImportCertificateAuthorityCertificate(ctx context.Context, params *ImportCertificateAuthorityCertificateInput, optFns ...func(*Options)) (*ImportCertificateAuthorityCertificateOutput, error) {
	if params == nil {
		params = &ImportCertificateAuthorityCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportCertificateAuthorityCertificate", params, optFns, c.addOperationImportCertificateAuthorityCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportCertificateAuthorityCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportCertificateAuthorityCertificateInput struct {

	// The PEM-encoded certificate for a private CA. This may be a self-signed
	// certificate in the case of a root CA, or it may be signed by another CA that you
	// control.
	//
	// This member is required.
	Certificate []byte

	// The Amazon Resource Name (ARN) that was returned when you called [CreateCertificateAuthority]. This must be
	// of the form:
	//
	//     arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	//
	// [CreateCertificateAuthority]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html
	//
	// This member is required.
	CertificateAuthorityArn *string

	// A PEM-encoded file that contains all of your certificates, other than the
	// certificate you're importing, chaining up to your root CA. Your Amazon Web
	// Services Private CA-hosted or on-premises root certificate is the last in the
	// chain, and each certificate in the chain signs the one preceding.
	//
	// This parameter must be supplied when you import a subordinate CA. When you
	// import a root CA, there is no chain.
	CertificateChain []byte

	noSmithyDocumentSerde
}

type ImportCertificateAuthorityCertificateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportCertificateAuthorityCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpImportCertificateAuthorityCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpImportCertificateAuthorityCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ImportCertificateAuthorityCertificate"); err != nil {
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
	if err = addOpImportCertificateAuthorityCertificateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportCertificateAuthorityCertificate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opImportCertificateAuthorityCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ImportCertificateAuthorityCertificate",
	}
}
