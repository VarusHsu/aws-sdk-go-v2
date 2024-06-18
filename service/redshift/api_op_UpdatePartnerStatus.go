// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the status of a partner integration.
func (c *Client) UpdatePartnerStatus(ctx context.Context, params *UpdatePartnerStatusInput, optFns ...func(*Options)) (*UpdatePartnerStatusOutput, error) {
	if params == nil {
		params = &UpdatePartnerStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePartnerStatus", params, optFns, c.addOperationUpdatePartnerStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePartnerStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePartnerStatusInput struct {

	// The Amazon Web Services account ID that owns the cluster.
	//
	// This member is required.
	AccountId *string

	// The cluster identifier of the cluster whose partner integration status is being
	// updated.
	//
	// This member is required.
	ClusterIdentifier *string

	// The name of the database whose partner integration status is being updated.
	//
	// This member is required.
	DatabaseName *string

	// The name of the partner whose integration status is being updated.
	//
	// This member is required.
	PartnerName *string

	// The value of the updated status.
	//
	// This member is required.
	Status types.PartnerIntegrationStatus

	// The status message provided by the partner.
	StatusMessage *string

	noSmithyDocumentSerde
}

type UpdatePartnerStatusOutput struct {

	// The name of the database that receives data from the partner.
	DatabaseName *string

	// The name of the partner that is authorized to send data.
	PartnerName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePartnerStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUpdatePartnerStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdatePartnerStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdatePartnerStatus"); err != nil {
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
	if err = addOpUpdatePartnerStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePartnerStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePartnerStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdatePartnerStatus",
	}
}
