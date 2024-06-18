// Code generated by smithy-go-codegen DO NOT EDIT.

package paymentcryptography

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/paymentcryptography/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon Web Services Payment Cryptography key, a logical
// representation of a cryptographic key, that is unique in your account and Amazon
// Web Services Region. You use keys for cryptographic functions such as encryption
// and decryption.
//
// In addition to the key material used in cryptographic operations, an Amazon Web
// Services Payment Cryptography key includes metadata such as the key ARN, key
// usage, key origin, creation date, description, and key state.
//
// When you create a key, you specify both immutable and mutable data about the
// key. The immutable data contains key attributes that define the scope and
// cryptographic operations that you can perform using the key, for example key
// class (example: SYMMETRIC_KEY ), key algorithm (example: TDES_2KEY ), key usage
// (example: TR31_P0_PIN_ENCRYPTION_KEY ) and key modes of use (example: Encrypt ).
// For information about valid combinations of key attributes, see [Understanding key attributes]in the Amazon
// Web Services Payment Cryptography User Guide. The mutable data contained within
// a key includes usage timestamp and key deletion timestamp and can be modified
// after creation.
//
// Amazon Web Services Payment Cryptography binds key attributes to keys using key
// blocks when you store or export them. Amazon Web Services Payment Cryptography
// stores the key contents wrapped and never stores or transmits them in the clear.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [DeleteKey]
//
// [GetKey]
//
// [ListKeys]
//
// [DeleteKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_DeleteKey.html
// [GetKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetKey.html
// [ListKeys]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ListKeys.html
// [Understanding key attributes]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-validattributes.html
func (c *Client) CreateKey(ctx context.Context, params *CreateKeyInput, optFns ...func(*Options)) (*CreateKeyOutput, error) {
	if params == nil {
		params = &CreateKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateKey", params, optFns, c.addOperationCreateKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateKeyInput struct {

	// Specifies whether the key is exportable from the service.
	//
	// This member is required.
	Exportable *bool

	// The role of the key, the algorithm it supports, and the cryptographic
	// operations allowed with the key. This data is immutable after the key is
	// created.
	//
	// This member is required.
	KeyAttributes *types.KeyAttributes

	// Specifies whether to enable the key. If the key is enabled, it is activated for
	// use within the service. If the key is not enabled, then it is created but not
	// activated. The default value is enabled.
	Enabled *bool

	// The algorithm that Amazon Web Services Payment Cryptography uses to calculate
	// the key check value (KCV). It is used to validate the key integrity.
	//
	// For TDES keys, the KCV is computed by encrypting 8 bytes, each with value of
	// zero, with the key to be checked and retaining the 3 highest order bytes of the
	// encrypted result. For AES keys, the KCV is computed using a CMAC algorithm where
	// the input data is 16 bytes of zero and retaining the 3 highest order bytes of
	// the encrypted result.
	KeyCheckValueAlgorithm types.KeyCheckValueAlgorithm

	// Assigns one or more tags to the Amazon Web Services Payment Cryptography key.
	// Use this parameter to tag a key when it is created. To tag an existing Amazon
	// Web Services Payment Cryptography key, use the [TagResource]operation.
	//
	// Each tag consists of a tag key and a tag value. Both the tag key and the tag
	// value are required, but the tag value can be an empty (null) string. You can't
	// have more than one tag on an Amazon Web Services Payment Cryptography key with
	// the same tag key.
	//
	// Don't include personal, confidential or sensitive information in this field.
	// This field may be displayed in plaintext in CloudTrail logs and other output.
	//
	// Tagging or untagging an Amazon Web Services Payment Cryptography key can allow
	// or deny permission to the key.
	//
	// [TagResource]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_TagResource.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateKeyOutput struct {

	// The key material that contains all the key attributes.
	//
	// This member is required.
	Key *types.Key

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateKey{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateKey"); err != nil {
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
	if err = addOpCreateKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateKey(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateKey",
	}
}
