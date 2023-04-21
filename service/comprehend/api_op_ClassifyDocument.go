// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new document classification request to analyze a single document in
// real-time, using a previously created and trained custom model and an endpoint.
// You can input plain text or you can upload a single-page input document (text,
// PDF, Word, or image). If the system detects errors while processing a page in
// the input document, the API response includes an entry in Errors that describes
// the errors. If the system detects a document-level error in your input document,
// the API returns an InvalidRequestException error response. For details about
// this exception, see Errors in semi-structured documents (https://docs.aws.amazon.com/comprehend/latest/dg/idp-inputs-sync-err.html)
// in the Comprehend Developer Guide.
func (c *Client) ClassifyDocument(ctx context.Context, params *ClassifyDocumentInput, optFns ...func(*Options)) (*ClassifyDocumentOutput, error) {
	if params == nil {
		params = &ClassifyDocumentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ClassifyDocument", params, optFns, c.addOperationClassifyDocumentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ClassifyDocumentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ClassifyDocumentInput struct {

	// The Amazon Resource Number (ARN) of the endpoint. For information about
	// endpoints, see Managing endpoints (https://docs.aws.amazon.com/comprehend/latest/dg/manage-endpoints.html)
	// .
	//
	// This member is required.
	EndpointArn *string

	// Use the Bytes parameter to input a text, PDF, Word or image file. You can also
	// use the Bytes parameter to input an Amazon Textract DetectDocumentText or
	// AnalyzeDocument output file. Provide the input document as a sequence of
	// base64-encoded bytes. If your code uses an Amazon Web Services SDK to classify
	// documents, the SDK may encode the document file bytes for you. The maximum
	// length of this field depends on the input document type. For details, see
	// Inputs for real-time custom analysis (https://docs.aws.amazon.com/comprehend/latest/dg/idp-inputs-sync.html)
	// in the Comprehend Developer Guide. If you use the Bytes parameter, do not use
	// the Text parameter.
	Bytes []byte

	// Provides configuration parameters to override the default actions for
	// extracting text from PDF documents and image files.
	DocumentReaderConfig *types.DocumentReaderConfig

	// The document text to be analyzed. If you enter text using this parameter, do
	// not use the Bytes parameter.
	Text *string

	noSmithyDocumentSerde
}

type ClassifyDocumentOutput struct {

	// The classes used by the document being analyzed. These are used for multi-class
	// trained models. Individual classes are mutually exclusive and each document is
	// expected to have only a single class assigned to it. For example, an animal can
	// be a dog or a cat, but not both at the same time.
	Classes []types.DocumentClass

	// Extraction information about the document. This field is present in the
	// response only if your request includes the Byte parameter.
	DocumentMetadata *types.DocumentMetadata

	// The document type for each page in the input document. This field is present in
	// the response only if your request includes the Byte parameter.
	DocumentType []types.DocumentTypeListItem

	// Page-level errors that the system detected while processing the input document.
	// The field is empty if the system encountered no errors.
	Errors []types.ErrorsListItem

	// The labels used the document being analyzed. These are used for multi-label
	// trained models. Individual labels represent different categories that are
	// related in some manner and are not mutually exclusive. For example, a movie can
	// be just an action movie, or it can be an action movie, a science fiction movie,
	// and a comedy, all at the same time.
	Labels []types.DocumentLabel

	// Warnings detected while processing the input document. The response includes a
	// warning if there is a mismatch between the input document type and the model
	// type associated with the endpoint that you specified. The response can also
	// include warnings for individual pages that have a mismatch. The field is empty
	// if the system generated no warnings.
	Warnings []types.WarningsListItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationClassifyDocumentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpClassifyDocument{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpClassifyDocument{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpClassifyDocumentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opClassifyDocument(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opClassifyDocument(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "ClassifyDocument",
	}
}
