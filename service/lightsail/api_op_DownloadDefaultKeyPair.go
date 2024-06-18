// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Downloads the regional Amazon Lightsail default key pair.
//
// This action also creates a Lightsail default key pair if a default key pair
// does not currently exist in the Amazon Web Services Region.
func (c *Client) DownloadDefaultKeyPair(ctx context.Context, params *DownloadDefaultKeyPairInput, optFns ...func(*Options)) (*DownloadDefaultKeyPairOutput, error) {
	if params == nil {
		params = &DownloadDefaultKeyPairInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DownloadDefaultKeyPair", params, optFns, c.addOperationDownloadDefaultKeyPairMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DownloadDefaultKeyPairOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DownloadDefaultKeyPairInput struct {
	noSmithyDocumentSerde
}

type DownloadDefaultKeyPairOutput struct {

	// The timestamp when the default key pair was created.
	CreatedAt *time.Time

	// A base64-encoded RSA private key.
	PrivateKeyBase64 *string

	// A base64-encoded public key of the ssh-rsa type.
	PublicKeyBase64 *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDownloadDefaultKeyPairMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDownloadDefaultKeyPair{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDownloadDefaultKeyPair{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DownloadDefaultKeyPair"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDownloadDefaultKeyPair(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDownloadDefaultKeyPair(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DownloadDefaultKeyPair",
	}
}
