// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers a CA certificate with Amazon Web Services IoT Core. There is no limit
// to the number of CA certificates you can register in your Amazon Web Services
// account. You can register up to 10 CA certificates with the same CA subject
// field per Amazon Web Services account.
//
// Requires permission to access the [RegisterCACertificate] action.
//
// [RegisterCACertificate]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func (c *Client) RegisterCACertificate(ctx context.Context, params *RegisterCACertificateInput, optFns ...func(*Options)) (*RegisterCACertificateOutput, error) {
	if params == nil {
		params = &RegisterCACertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterCACertificate", params, optFns, c.addOperationRegisterCACertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterCACertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input to the RegisterCACertificate operation.
type RegisterCACertificateInput struct {

	// The CA certificate.
	//
	// This member is required.
	CaCertificate *string

	// Allows this CA certificate to be used for auto registration of device
	// certificates.
	AllowAutoRegistration bool

	// Describes the certificate mode in which the Certificate Authority (CA) will be
	// registered. If the verificationCertificate field is not provided, set
	// certificateMode to be SNI_ONLY . If the verificationCertificate field is
	// provided, set certificateMode to be DEFAULT . When certificateMode is not
	// provided, it defaults to DEFAULT . All the device certificates that are
	// registered using this CA will be registered in the same certificate mode as the
	// CA. For more information about certificate mode for device certificates, see [certificate mode].
	//
	// [certificate mode]: https://docs.aws.amazon.com/iot/latest/apireference/API_CertificateDescription.html#iot-Type-CertificateDescription-certificateMode
	CertificateMode types.CertificateMode

	// Information about the registration configuration.
	RegistrationConfig *types.RegistrationConfig

	// A boolean value that specifies if the CA certificate is set to active.
	//
	// Valid values: ACTIVE | INACTIVE
	SetAsActive bool

	// Metadata which can be used to manage the CA certificate.
	//
	// For URI Request parameters use format: ...key1=value1&key2=value2...
	//
	// For the CLI command-line parameter use format: &&tags
	// "key1=value1&key2=value2..."
	//
	// For the cli-input-json file use format: "tags": "key1=value1&key2=value2..."
	Tags []types.Tag

	// The private key verification certificate. If certificateMode is SNI_ONLY , the
	// verificationCertificate field must be empty. If certificateMode is DEFAULT or
	// not provided, the verificationCertificate field must not be empty.
	VerificationCertificate *string

	noSmithyDocumentSerde
}

// The output from the RegisterCACertificateResponse operation.
type RegisterCACertificateOutput struct {

	// The CA certificate ARN.
	CertificateArn *string

	// The CA certificate identifier.
	CertificateId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterCACertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRegisterCACertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRegisterCACertificate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RegisterCACertificate"); err != nil {
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
	if err = addOpRegisterCACertificateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterCACertificate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterCACertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RegisterCACertificate",
	}
}
