// Code generated by smithy-go-codegen DO NOT EDIT.

package codegurusecurity

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codegurusecurity/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns details about a scan, including whether or not a scan has completed.
func (c *Client) GetScan(ctx context.Context, params *GetScanInput, optFns ...func(*Options)) (*GetScanOutput, error) {
	if params == nil {
		params = &GetScanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetScan", params, optFns, c.addOperationGetScanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetScanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetScanInput struct {

	// The name of the scan you want to view details about.
	//
	// This member is required.
	ScanName *string

	// UUID that identifies the individual scan run you want to view details about.
	// You retrieve this when you call the CreateScan operation. Defaults to the
	// latest scan run if missing.
	RunId *string

	noSmithyDocumentSerde
}

type GetScanOutput struct {

	// The type of analysis CodeGuru Security performed in the scan, either Security
	// or All . The Security type only generates findings related to security. The All
	// type generates both security findings and quality findings.
	//
	// This member is required.
	AnalysisType types.AnalysisType

	// The time the scan was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// UUID that identifies the individual scan run.
	//
	// This member is required.
	RunId *string

	// The name of the scan.
	//
	// This member is required.
	ScanName *string

	// The current state of the scan. Pass either InProgress , Successful , or Failed .
	//
	// This member is required.
	ScanState types.ScanState

	// The number of times a scan has been re-run on a revised resource.
	NumberOfRevisions *int64

	// The ARN for the scan name.
	ScanNameArn *string

	// The time when the scan was last updated. Only available for STANDARD scan types.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetScanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetScan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetScan{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetScan"); err != nil {
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
	if err = addOpGetScanValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetScan(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetScan(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetScan",
	}
}
