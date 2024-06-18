// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a metadata form type.
func (c *Client) CreateFormType(ctx context.Context, params *CreateFormTypeInput, optFns ...func(*Options)) (*CreateFormTypeOutput, error) {
	if params == nil {
		params = &CreateFormTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFormType", params, optFns, c.addOperationCreateFormTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFormTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFormTypeInput struct {

	// The ID of the Amazon DataZone domain in which this metadata form type is
	// created.
	//
	// This member is required.
	DomainIdentifier *string

	// The model of this Amazon DataZone metadata form type.
	//
	// This member is required.
	Model types.Model

	// The name of this Amazon DataZone metadata form type.
	//
	// This member is required.
	Name *string

	// The ID of the Amazon DataZone project that owns this metadata form type.
	//
	// This member is required.
	OwningProjectIdentifier *string

	// The description of this Amazon DataZone metadata form type.
	Description *string

	// The status of this Amazon DataZone metadata form type.
	Status types.FormTypeStatus

	noSmithyDocumentSerde
}

type CreateFormTypeOutput struct {

	// The ID of the Amazon DataZone domain in which this metadata form type is
	// created.
	//
	// This member is required.
	DomainId *string

	// The name of this Amazon DataZone metadata form type.
	//
	// This member is required.
	Name *string

	// The revision of this Amazon DataZone metadata form type.
	//
	// This member is required.
	Revision *string

	// The description of this Amazon DataZone metadata form type.
	Description *string

	// The ID of the Amazon DataZone domain in which this metadata form type was
	// originally created.
	OriginDomainId *string

	// The ID of the project in which this Amazon DataZone metadata form type was
	// originally created.
	OriginProjectId *string

	// The ID of the project that owns this Amazon DataZone metadata form type.
	OwningProjectId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFormTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateFormType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateFormType{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateFormType"); err != nil {
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
	if err = addOpCreateFormTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFormType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateFormType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateFormType",
	}
}
