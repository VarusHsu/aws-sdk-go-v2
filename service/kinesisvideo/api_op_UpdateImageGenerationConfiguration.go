// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the StreamInfo and ImageProcessingConfiguration fields.
func (c *Client) UpdateImageGenerationConfiguration(ctx context.Context, params *UpdateImageGenerationConfigurationInput, optFns ...func(*Options)) (*UpdateImageGenerationConfigurationOutput, error) {
	if params == nil {
		params = &UpdateImageGenerationConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateImageGenerationConfiguration", params, optFns, c.addOperationUpdateImageGenerationConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateImageGenerationConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateImageGenerationConfigurationInput struct {

	// The structure that contains the information required for the KVS images
	// delivery. If the structure is null, the configuration will be deleted from the
	// stream.
	ImageGenerationConfiguration *types.ImageGenerationConfiguration

	// The Amazon Resource Name (ARN) of the Kinesis video stream from where you want
	// to update the image generation configuration. You must specify either the
	// StreamName or the StreamARN .
	StreamARN *string

	// The name of the stream from which to update the image generation configuration.
	// You must specify either the StreamName or the StreamARN .
	StreamName *string

	noSmithyDocumentSerde
}

type UpdateImageGenerationConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateImageGenerationConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateImageGenerationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateImageGenerationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateImageGenerationConfiguration"); err != nil {
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
	if err = addOpUpdateImageGenerationConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateImageGenerationConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateImageGenerationConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateImageGenerationConfiguration",
	}
}
