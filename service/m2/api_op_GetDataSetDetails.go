// Code generated by smithy-go-codegen DO NOT EDIT.

package m2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/m2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets the details of a specific data set.
func (c *Client) GetDataSetDetails(ctx context.Context, params *GetDataSetDetailsInput, optFns ...func(*Options)) (*GetDataSetDetailsOutput, error) {
	if params == nil {
		params = &GetDataSetDetailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDataSetDetails", params, optFns, c.addOperationGetDataSetDetailsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDataSetDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDataSetDetailsInput struct {

	// The unique identifier of the application that this data set is associated with.
	//
	// This member is required.
	ApplicationId *string

	// The name of the data set.
	//
	// This member is required.
	DataSetName *string

	noSmithyDocumentSerde
}

type GetDataSetDetailsOutput struct {

	// The name of the data set.
	//
	// This member is required.
	DataSetName *string

	// The size of the block on disk.
	Blocksize *int32

	// The timestamp when the data set was created.
	CreationTime *time.Time

	// The type of data set. The only supported value is VSAM.
	DataSetOrg types.DatasetDetailOrgAttributes

	// File size of the dataset.
	FileSize *int64

	// The last time the data set was referenced.
	LastReferencedTime *time.Time

	// The last time the data set was updated.
	LastUpdatedTime *time.Time

	// The location where the data set is stored.
	Location *string

	// The length of records in the data set.
	RecordLength *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDataSetDetailsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDataSetDetails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDataSetDetails{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetDataSetDetails"); err != nil {
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
	if err = addOpGetDataSetDetailsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDataSetDetails(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDataSetDetails(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetDataSetDetails",
	}
}
