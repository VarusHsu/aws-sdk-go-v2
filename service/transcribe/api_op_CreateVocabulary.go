// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a new custom vocabulary.
//
// When creating a new custom vocabulary, you can either upload a text file that
// contains your new entries, phrases, and terms into an Amazon S3 bucket and
// include the URI in your request. Or you can include a list of terms directly in
// your request using the Phrases flag.
//
// Each language has a character set that contains all allowed characters for that
// specific language. If you use unsupported characters, your custom vocabulary
// request fails. Refer to [Character Sets for Custom Vocabularies]to get the character set for your language.
//
// For more information, see [Custom vocabularies].
//
// [Custom vocabularies]: https://docs.aws.amazon.com/transcribe/latest/dg/custom-vocabulary.html
// [Character Sets for Custom Vocabularies]: https://docs.aws.amazon.com/transcribe/latest/dg/charsets.html
func (c *Client) CreateVocabulary(ctx context.Context, params *CreateVocabularyInput, optFns ...func(*Options)) (*CreateVocabularyOutput, error) {
	if params == nil {
		params = &CreateVocabularyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVocabulary", params, optFns, c.addOperationCreateVocabularyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVocabularyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVocabularyInput struct {

	// The language code that represents the language of the entries in your custom
	// vocabulary. Each custom vocabulary must contain terms in only one language.
	//
	// A custom vocabulary can only be used to transcribe files in the same language
	// as the custom vocabulary. For example, if you create a custom vocabulary using
	// US English ( en-US ), you can only apply this custom vocabulary to files that
	// contain English audio.
	//
	// For a list of supported languages and their associated language codes, refer to
	// the [Supported languages]table.
	//
	// [Supported languages]: https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// A unique name, chosen by you, for your new custom vocabulary.
	//
	// This name is case sensitive, cannot contain spaces, and must be unique within
	// an Amazon Web Services account. If you try to create a new custom vocabulary
	// with the same name as an existing custom vocabulary, you get a ConflictException
	// error.
	//
	// This member is required.
	VocabularyName *string

	// The Amazon Resource Name (ARN) of an IAM role that has permissions to access
	// the Amazon S3 bucket that contains your input files (in this case, your custom
	// vocabulary). If the role that you specify doesn’t have the appropriate
	// permissions to access the specified Amazon S3 location, your request fails.
	//
	// IAM role ARNs have the format
	// arn:partition:iam::account:role/role-name-with-path . For example:
	// arn:aws:iam::111122223333:role/Admin .
	//
	// For more information, see [IAM ARNs].
	//
	// [IAM ARNs]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-arns
	DataAccessRoleArn *string

	// Use this parameter if you want to create your custom vocabulary by including
	// all desired terms, as comma-separated values, within your request. The other
	// option for creating your custom vocabulary is to save your entries in a text
	// file and upload them to an Amazon S3 bucket, then specify the location of your
	// file using the VocabularyFileUri parameter.
	//
	// Note that if you include Phrases in your request, you cannot use
	// VocabularyFileUri ; you must choose one or the other.
	//
	// Each language has a character set that contains all allowed characters for that
	// specific language. If you use unsupported characters, your custom vocabulary
	// filter request fails. Refer to [Character Sets for Custom Vocabularies]to get the character set for your language.
	//
	// [Character Sets for Custom Vocabularies]: https://docs.aws.amazon.com/transcribe/latest/dg/charsets.html
	Phrases []string

	// Adds one or more custom tags, each in the form of a key:value pair, to a new
	// custom vocabulary at the time you create this new custom vocabulary.
	//
	// To learn more about using tags with Amazon Transcribe, refer to [Tagging resources].
	//
	// [Tagging resources]: https://docs.aws.amazon.com/transcribe/latest/dg/tagging.html
	Tags []types.Tag

	// The Amazon S3 location of the text file that contains your custom vocabulary.
	// The URI must be located in the same Amazon Web Services Region as the resource
	// you're calling.
	//
	// Here's an example URI path: s3://DOC-EXAMPLE-BUCKET/my-vocab-file.txt
	//
	// Note that if you include VocabularyFileUri in your request, you cannot use the
	// Phrases flag; you must choose one or the other.
	VocabularyFileUri *string

	noSmithyDocumentSerde
}

type CreateVocabularyOutput struct {

	// If VocabularyState is FAILED , FailureReason contains information about why the
	// custom vocabulary request failed. See also: [Common Errors].
	//
	// [Common Errors]: https://docs.aws.amazon.com/transcribe/latest/APIReference/CommonErrors.html
	FailureReason *string

	// The language code you selected for your custom vocabulary.
	LanguageCode types.LanguageCode

	// The date and time you created your custom vocabulary.
	//
	// Timestamps are in the format YYYY-MM-DD'T'HH:MM:SS.SSSSSS-UTC . For example,
	// 2022-05-04T12:32:58.761000-07:00 represents 12:32 PM UTC-7 on May 4, 2022.
	LastModifiedTime *time.Time

	// The name you chose for your custom vocabulary.
	VocabularyName *string

	// The processing state of your custom vocabulary. If the state is READY , you can
	// use the custom vocabulary in a StartTranscriptionJob request.
	VocabularyState types.VocabularyState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateVocabularyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateVocabulary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateVocabulary{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateVocabulary"); err != nil {
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
	if err = addOpCreateVocabularyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVocabulary(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateVocabulary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateVocabulary",
	}
}
