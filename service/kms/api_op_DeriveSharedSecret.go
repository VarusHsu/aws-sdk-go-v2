// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Derives a shared secret using a key agreement algorithm.
//
// You must use an asymmetric NIST-recommended elliptic curve (ECC) or SM2 (China
// Regions only) KMS key pair with a KeyUsage value of KEY_AGREEMENT to call
// DeriveSharedSecret.
//
// DeriveSharedSecret uses the [Elliptic Curve Cryptography Cofactor Diffie-Hellman Primitive] (ECDH) to establish a key agreement between two
// peers by deriving a shared secret from their elliptic curve public-private key
// pairs. You can use the raw shared secret that DeriveSharedSecret returns to
// derive a symmetric key that can encrypt and decrypt data that is sent between
// the two peers, or that can generate and verify HMACs. KMS recommends that you
// follow [NIST recommendations for key derivation]when using the raw shared secret to derive a symmetric key.
//
// The following workflow demonstrates how to establish key agreement over an
// insecure communication channel using DeriveSharedSecret.
//
//   - Alice calls CreateKeyto create an asymmetric KMS key pair with a KeyUsage value of
//     KEY_AGREEMENT .
//
// The asymmetric KMS key must use a NIST-recommended elliptic curve (ECC) or SM2
//
//	(China Regions only) key spec.
//
//	- Bob creates an elliptic curve key pair.
//
// Bob can call CreateKeyto create an asymmetric KMS key pair or generate a key pair
//
//	outside of KMS. Bob's key pair must use the same NIST-recommended elliptic curve
//	(ECC) or SM2 (China Regions ony) curve as Alice.
//
//	- Alice and Bob exchange their public keys through an insecure communication
//	channel (like the internet).
//
// Use GetPublicKeyto download the public key of your asymmetric KMS key pair.
//
// KMS strongly recommends verifying that the public key you receive came from the
//
//	expected party before using it to derive a shared secret.
//
//	- Alice calls DeriveSharedSecret.
//
// KMS uses the private key from the KMS key pair generated in Step 1, Bob's
//
//	public key, and the Elliptic Curve Cryptography Cofactor Diffie-Hellman
//	Primitive to derive the shared secret. The private key in your KMS key pair
//	never leaves KMS unencrypted. DeriveSharedSecret returns the raw shared secret.
//
//	- Bob uses the Elliptic Curve Cryptography Cofactor Diffie-Hellman Primitive
//	to calculate the same raw secret using his private key and Alice's public key.
//
// To derive a shared secret you must provide a key agreement algorithm, the
// private key of the caller's asymmetric NIST-recommended elliptic curve or SM2
// (China Regions only) KMS key pair, and the public key from your peer's
// NIST-recommended elliptic curve or SM2 (China Regions only) key pair. The public
// key can be from another asymmetric KMS key pair or from a key pair generated
// outside of KMS, but both key pairs must be on the same elliptic curve.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: Yes. To perform this operation with a KMS key in a different
// Amazon Web Services account, specify the key ARN or alias ARN in the value of
// the KeyId parameter.
//
// Required permissions: [kms:DeriveSharedSecret] (key policy)
//
// Related operations:
//
// # CreateKey
//
// # GetPublicKey
//
// # DescribeKey
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [kms:DeriveSharedSecret]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Elliptic Curve Cryptography Cofactor Diffie-Hellman Primitive]: https://nvlpubs.nist.gov/nistpubs/SpecialPublications/NIST.SP.800-56Ar3.pdf#page=60
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/programming-eventual-consistency.html
// [NIST recommendations for key derivation]: https://nvlpubs.nist.gov/nistpubs/SpecialPublications/NIST.SP.800-56Cr2.pdf
func (c *Client) DeriveSharedSecret(ctx context.Context, params *DeriveSharedSecretInput, optFns ...func(*Options)) (*DeriveSharedSecretOutput, error) {
	if params == nil {
		params = &DeriveSharedSecretInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeriveSharedSecret", params, optFns, c.addOperationDeriveSharedSecretMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeriveSharedSecretOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeriveSharedSecretInput struct {

	// Specifies the key agreement algorithm used to derive the shared secret. The
	// only valid value is ECDH .
	//
	// This member is required.
	KeyAgreementAlgorithm types.KeyAgreementAlgorithmSpec

	// Identifies an asymmetric NIST-recommended ECC or SM2 (China Regions only) KMS
	// key. KMS uses the private key in the specified key pair to derive the shared
	// secret. The key usage of the KMS key must be KEY_AGREEMENT . To find the
	// KeyUsage of a KMS key, use the DescribeKey operation.
	//
	// To specify a KMS key, use its key ID, key ARN, alias name, or alias ARN. When
	// using an alias name, prefix it with "alias/" . To specify a KMS key in a
	// different Amazon Web Services account, you must use the key ARN or alias ARN.
	//
	// For example:
	//
	//   - Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//   - Key ARN:
	//   arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//   - Alias name: alias/ExampleAlias
	//
	//   - Alias ARN: arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	//
	// To get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey. To get the alias name
	// and alias ARN, use ListAliases.
	//
	// This member is required.
	KeyId *string

	// Specifies the public key in your peer's NIST-recommended elliptic curve (ECC)
	// or SM2 (China Regions only) key pair.
	//
	// The public key must be a DER-encoded X.509 public key, also known as
	// SubjectPublicKeyInfo (SPKI), as defined in [RFC 5280].
	//
	// GetPublicKeyreturns the public key of an asymmetric KMS key pair in the required
	// DER-encoded format.
	//
	// If you use [Amazon Web Services CLI version 1], you must provide the DER-encoded X.509 public key in a file.
	// Otherwise, the Amazon Web Services CLI Base64-encodes the public key a second
	// time, resulting in a ValidationException .
	//
	// You can specify the public key as binary data in a file using fileb ( fileb:// )
	// or in-line using a Base64 encoded string.
	//
	// [RFC 5280]: https://tools.ietf.org/html/rfc5280
	// [Amazon Web Services CLI version 1]: https://docs.aws.amazon.com/cli/v1/userguide/cli-chap-welcome.html
	//
	// This member is required.
	PublicKey []byte

	// Checks if your request will succeed. DryRun is an optional parameter.
	//
	// To learn more about how to use this parameter, see [Testing your KMS API calls] in the Key Management
	// Service Developer Guide.
	//
	// [Testing your KMS API calls]: https://docs.aws.amazon.com/kms/latest/developerguide/programming-dryrun.html
	DryRun *bool

	// A list of grant tokens.
	//
	// Use a grant token when your permission to call this operation comes from a new
	// grant that has not yet achieved eventual consistency. For more information, see [Grant token]
	// and [Using a grant token]in the Key Management Service Developer Guide.
	//
	// [Grant token]: https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#grant_token
	// [Using a grant token]: https://docs.aws.amazon.com/kms/latest/developerguide/grant-manage.html#using-grant-token
	GrantTokens []string

	// A signed [attestation document] from an Amazon Web Services Nitro enclave and the encryption
	// algorithm to use with the enclave's public key. The only valid encryption
	// algorithm is RSAES_OAEP_SHA_256 .
	//
	// This parameter only supports attestation documents for Amazon Web Services
	// Nitro Enclaves. To call DeriveSharedSecret for an Amazon Web Services Nitro
	// Enclaves, use the [Amazon Web Services Nitro Enclaves SDK]to generate the attestation document and then use the
	// Recipient parameter from any Amazon Web Services SDK to provide the attestation
	// document for the enclave.
	//
	// When you use this parameter, instead of returning a plaintext copy of the
	// shared secret, KMS encrypts the plaintext shared secret under the public key in
	// the attestation document, and returns the resulting ciphertext in the
	// CiphertextForRecipient field in the response. This ciphertext can be decrypted
	// only with the private key in the enclave. The CiphertextBlob field in the
	// response contains the encrypted shared secret derived from the KMS key specified
	// by the KeyId parameter and public key specified by the PublicKey parameter. The
	// SharedSecret field in the response is null or empty.
	//
	// For information about the interaction between KMS and Amazon Web Services Nitro
	// Enclaves, see [How Amazon Web Services Nitro Enclaves uses KMS]in the Key Management Service Developer Guide.
	//
	// [attestation document]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nitro-enclave-how.html#term-attestdoc
	// [How Amazon Web Services Nitro Enclaves uses KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/services-nitro-enclaves.html
	// [Amazon Web Services Nitro Enclaves SDK]: https://docs.aws.amazon.com/enclaves/latest/user/developing-applications.html#sdk
	Recipient *types.RecipientInfo

	noSmithyDocumentSerde
}

type DeriveSharedSecretOutput struct {

	// The plaintext shared secret encrypted with the public key in the attestation
	// document.
	//
	// This field is included in the response only when the Recipient parameter in the
	// request includes a valid attestation document from an Amazon Web Services Nitro
	// enclave. For information about the interaction between KMS and Amazon Web
	// Services Nitro Enclaves, see [How Amazon Web Services Nitro Enclaves uses KMS]in the Key Management Service Developer Guide.
	//
	// [How Amazon Web Services Nitro Enclaves uses KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/services-nitro-enclaves.html
	CiphertextForRecipient []byte

	// Identifies the key agreement algorithm used to derive the shared secret.
	KeyAgreementAlgorithm types.KeyAgreementAlgorithmSpec

	// Identifies the KMS key used to derive the shared secret.
	KeyId *string

	// The source of the key material for the specified KMS key.
	//
	// When this value is AWS_KMS , KMS created the key material. When this value is
	// EXTERNAL , the key material was imported or the KMS key doesn't have any key
	// material.
	//
	// The only valid values for DeriveSharedSecret are AWS_KMS and EXTERNAL .
	// DeriveSharedSecret does not support KMS keys with a KeyOrigin value of
	// AWS_CLOUDHSM or EXTERNAL_KEY_STORE .
	KeyOrigin types.OriginType

	// The raw secret derived from the specified key agreement algorithm, private key
	// in the asymmetric KMS key, and your peer's public key.
	//
	// If the response includes the CiphertextForRecipient field, the SharedSecret
	// field is null or empty.
	SharedSecret []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeriveSharedSecretMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeriveSharedSecret{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeriveSharedSecret{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeriveSharedSecret"); err != nil {
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
	if err = addOpDeriveSharedSecretValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeriveSharedSecret(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeriveSharedSecret(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeriveSharedSecret",
	}
}
