// Code generated by smithy-go-codegen DO NOT EDIT.

package textract

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/textract/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets configuration information for an adapter specified by an AdapterId,
// returning information on AdapterName, Description, CreationTime, AutoUpdate
// status, and FeatureTypes.
func (c *Client) GetAdapter(ctx context.Context, params *GetAdapterInput, optFns ...func(*Options)) (*GetAdapterOutput, error) {
	if params == nil {
		params = &GetAdapterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAdapter", params, optFns, c.addOperationGetAdapterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAdapterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAdapterInput struct {

	// A string containing a unique ID for the adapter.
	//
	// This member is required.
	AdapterId *string

	noSmithyDocumentSerde
}

type GetAdapterOutput struct {

	// A string identifying the adapter that information has been retrieved for.
	AdapterId *string

	// The name of the requested adapter.
	AdapterName *string

	// Binary value indicating if the adapter is being automatically updated or not.
	AutoUpdate types.AutoUpdate

	// The date and time the requested adapter was created at.
	CreationTime *time.Time

	// The description for the requested adapter.
	Description *string

	// List of the targeted feature types for the requested adapter.
	FeatureTypes []types.FeatureType

	// A set of tags (key-value pairs) associated with the adapter that has been
	// retrieved.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAdapterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetAdapter{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAdapter{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAdapter"); err != nil {
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
	if err = addOpGetAdapterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAdapter(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAdapter(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAdapter",
	}
}
