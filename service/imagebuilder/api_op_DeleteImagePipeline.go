// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an image pipeline.
func (c *Client) DeleteImagePipeline(ctx context.Context, params *DeleteImagePipelineInput, optFns ...func(*Options)) (*DeleteImagePipelineOutput, error) {
	stack := middleware.NewStack("DeleteImagePipeline", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteImagePipelineMiddlewares(stack)
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
	addOpDeleteImagePipelineValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteImagePipeline(options.Region), middleware.Before)
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
			OperationName: "DeleteImagePipeline",
			Err:           err,
		}
	}
	out := result.(*DeleteImagePipelineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteImagePipelineInput struct {

	// The Amazon Resource Name (ARN) of the image pipeline to delete.
	//
	// This member is required.
	ImagePipelineArn *string
}

type DeleteImagePipelineOutput struct {

	// The Amazon Resource Name (ARN) of the image pipeline that was deleted.
	ImagePipelineArn *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteImagePipelineMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteImagePipeline{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteImagePipeline{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteImagePipeline(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "DeleteImagePipeline",
	}
}
