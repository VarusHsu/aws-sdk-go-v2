// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a URL for a specified UserProfile in a Domain. When accessed in a web
// browser, the user will be automatically signed in to Amazon SageMaker Studio,
// and granted access to all of the Apps and files associated with the Domain's
// Amazon Elastic File System (EFS) volume. This operation can only be called when
// the authentication mode equals IAM. The IAM role or user passed to this API
// defines the permissions to access the app. Once the presigned URL is created, no
// additional permission is required to access this URL. IAM authorization policies
// for this API are also enforced for every HTTP request and WebSocket frame that
// attempts to connect to the app. You can restrict access to this API and to the
// URL that it returns to a list of IP addresses, Amazon VPCs or Amazon VPC
// Endpoints that you specify. For more information, see Connect to SageMaker
// Studio Through an Interface VPC Endpoint (https://docs.aws.amazon.com/sagemaker/latest/dg/studio-interface-endpoint.html)
// . The URL that you get from a call to CreatePresignedDomainUrl has a default
// timeout of 5 minutes. You can configure this value using ExpiresInSeconds . If
// you try to use the URL after the timeout limit expires, you are directed to the
// Amazon Web Services console sign-in page.
func (c *Client) CreatePresignedDomainUrl(ctx context.Context, params *CreatePresignedDomainUrlInput, optFns ...func(*Options)) (*CreatePresignedDomainUrlOutput, error) {
	if params == nil {
		params = &CreatePresignedDomainUrlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePresignedDomainUrl", params, optFns, c.addOperationCreatePresignedDomainUrlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePresignedDomainUrlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePresignedDomainUrlInput struct {

	// The domain ID.
	//
	// This member is required.
	DomainId *string

	// The name of the UserProfile to sign-in as.
	//
	// This member is required.
	UserProfileName *string

	// The number of seconds until the pre-signed URL expires. This value defaults to
	// 300.
	ExpiresInSeconds *int32

	// The session expiration duration in seconds. This value defaults to 43200.
	SessionExpirationDurationInSeconds *int32

	// The name of the space.
	SpaceName *string

	noSmithyDocumentSerde
}

type CreatePresignedDomainUrlOutput struct {

	// The presigned URL.
	AuthorizedUrl *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePresignedDomainUrlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePresignedDomainUrl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePresignedDomainUrl{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreatePresignedDomainUrl"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpCreatePresignedDomainUrlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePresignedDomainUrl(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opCreatePresignedDomainUrl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreatePresignedDomainUrl",
	}
}
