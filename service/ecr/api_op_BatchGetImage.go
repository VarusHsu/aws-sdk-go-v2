// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets detailed information for an image. Images are specified with either an
// imageTag or imageDigest. When an image is pulled, the BatchGetImage API is
// called once to retrieve the image manifest.
func (c *Client) BatchGetImage(ctx context.Context, params *BatchGetImageInput, optFns ...func(*Options)) (*BatchGetImageOutput, error) {
	stack := middleware.NewStack("BatchGetImage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpBatchGetImageMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpBatchGetImageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetImage(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "BatchGetImage",
			Err:           err,
		}
	}
	out := result.(*BatchGetImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetImageInput struct {

	// A list of image ID references that correspond to images to describe. The format
	// of the imageIds reference is imageTag=tag or imageDigest=digest.
	//
	// This member is required.
	ImageIds []*types.ImageIdentifier

	// The repository that contains the images to describe.
	//
	// This member is required.
	RepositoryName *string

	// The AWS account ID associated with the registry that contains the images to
	// describe. If you do not specify a registry, the default registry is assumed.
	RegistryId *string

	// The accepted media types for the request. Valid values:
	// application/vnd.docker.distribution.manifest.v1+json |
	// application/vnd.docker.distribution.manifest.v2+json |
	// application/vnd.oci.image.manifest.v1+json
	AcceptedMediaTypes []*string
}

type BatchGetImageOutput struct {

	// Any failures associated with the call.
	Failures []*types.ImageFailure

	// A list of image objects corresponding to the image references in the request.
	Images []*types.Image

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpBatchGetImageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetImage{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetImage{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchGetImage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr",
		OperationName: "BatchGetImage",
	}
}
