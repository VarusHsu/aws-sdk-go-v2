// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// The action to start the generation of test set.
func (c *Client) StartTestSetGeneration(ctx context.Context, params *StartTestSetGenerationInput, optFns ...func(*Options)) (*StartTestSetGenerationOutput, error) {
	if params == nil {
		params = &StartTestSetGenerationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartTestSetGeneration", params, optFns, c.addOperationStartTestSetGenerationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartTestSetGenerationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartTestSetGenerationInput struct {

	// The data source for the test set generation.
	//
	// This member is required.
	GenerationDataSource *types.TestSetGenerationDataSource

	// The roleARN used for any operation in the test set to access resources in the
	// Amazon Web Services account.
	//
	// This member is required.
	RoleArn *string

	// The Amazon S3 storage location for the test set generation.
	//
	// This member is required.
	StorageLocation *types.TestSetStorageLocation

	// The test set name for the test set generation request.
	//
	// This member is required.
	TestSetName *string

	// The test set description for the test set generation request.
	Description *string

	// A list of tags to add to the test set. You can only add tags when you
	// import/generate a new test set. You can't use the UpdateTestSet operation to
	// update tags. To update tags, use the TagResource operation.
	TestSetTags map[string]string

	noSmithyDocumentSerde
}

type StartTestSetGenerationOutput struct {

	//  The creation date and time for the test set generation.
	CreationDateTime *time.Time

	// The description used for the test set generation.
	Description *string

	//  The data source for the test set generation.
	GenerationDataSource *types.TestSetGenerationDataSource

	// The roleARN used for any operation in the test set to access resources in the
	// Amazon Web Services account.
	RoleArn *string

	// The Amazon S3 storage location for the test set generation.
	StorageLocation *types.TestSetStorageLocation

	// The unique identifier of the test set generation to describe.
	TestSetGenerationId *string

	//  The status for the test set generation.
	TestSetGenerationStatus types.TestSetGenerationStatus

	// The test set name used for the test set generation.
	TestSetName *string

	// A list of tags that was used for the test set that is being generated.
	TestSetTags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartTestSetGenerationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartTestSetGeneration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartTestSetGeneration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartTestSetGeneration"); err != nil {
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
	if err = addOpStartTestSetGenerationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartTestSetGeneration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartTestSetGeneration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartTestSetGeneration",
	}
}
