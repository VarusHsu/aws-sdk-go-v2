// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets details about the specified Git repository.
func (c *Client) DescribeCodeRepository(ctx context.Context, params *DescribeCodeRepositoryInput, optFns ...func(*Options)) (*DescribeCodeRepositoryOutput, error) {
	if params == nil {
		params = &DescribeCodeRepositoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCodeRepository", params, optFns, c.addOperationDescribeCodeRepositoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCodeRepositoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCodeRepositoryInput struct {

	// The name of the Git repository to describe.
	//
	// This member is required.
	CodeRepositoryName *string

	noSmithyDocumentSerde
}

type DescribeCodeRepositoryOutput struct {

	// The Amazon Resource Name (ARN) of the Git repository.
	//
	// This member is required.
	CodeRepositoryArn *string

	// The name of the Git repository.
	//
	// This member is required.
	CodeRepositoryName *string

	// The date and time that the repository was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The date and time that the repository was last changed.
	//
	// This member is required.
	LastModifiedTime *time.Time

	// Configuration details about the repository, including the URL where the
	// repository is located, the default branch, and the Amazon Resource Name (ARN) of
	// the Amazon Web Services Secrets Manager secret that contains the credentials
	// used to access the repository.
	GitConfig *types.GitConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeCodeRepositoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCodeRepository{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCodeRepository{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeCodeRepository"); err != nil {
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
	if err = addOpDescribeCodeRepositoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCodeRepository(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeCodeRepository(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeCodeRepository",
	}
}
