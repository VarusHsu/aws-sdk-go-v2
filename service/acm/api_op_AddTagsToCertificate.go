// Code generated by smithy-go-codegen DO NOT EDIT.

package acm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds one or more tags to an ACM certificate. Tags are labels that you can use
// to identify and organize your Amazon Web Services resources. Each tag consists
// of a key and an optional value . You specify the certificate on input by its
// Amazon Resource Name (ARN). You specify the tag by using a key-value pair.
//
// You can apply a tag to just one certificate if you want to identify a specific
// characteristic of that certificate, or you can apply the same tag to multiple
// certificates if you want to filter for a common relationship among those
// certificates. Similarly, you can apply the same tag to multiple resources if you
// want to specify a relationship among those resources. For example, you can add
// the same tag to an ACM certificate and an Elastic Load Balancing load balancer
// to indicate that they are both used by the same website. For more information,
// see [Tagging ACM certificates].
//
// To remove one or more tags, use the RemoveTagsFromCertificate action. To view all of the tags that have
// been applied to the certificate, use the ListTagsForCertificateaction.
//
// [Tagging ACM certificates]: https://docs.aws.amazon.com/acm/latest/userguide/tags.html
func (c *Client) AddTagsToCertificate(ctx context.Context, params *AddTagsToCertificateInput, optFns ...func(*Options)) (*AddTagsToCertificateOutput, error) {
	if params == nil {
		params = &AddTagsToCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddTagsToCertificate", params, optFns, c.addOperationAddTagsToCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddTagsToCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddTagsToCertificateInput struct {

	// String that contains the ARN of the ACM certificate to which the tag is to be
	// applied. This must be of the form:
	//
	//     arn:aws:acm:region:123456789012:certificate/12345678-1234-1234-1234-123456789012
	//
	// For more information about ARNs, see [Amazon Resource Names (ARNs)].
	//
	// [Amazon Resource Names (ARNs)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
	//
	// This member is required.
	CertificateArn *string

	// The key-value pair that defines the tag. The tag value is optional.
	//
	// This member is required.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type AddTagsToCertificateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddTagsToCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAddTagsToCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddTagsToCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AddTagsToCertificate"); err != nil {
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
	if err = addOpAddTagsToCertificateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddTagsToCertificate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddTagsToCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AddTagsToCertificate",
	}
}
