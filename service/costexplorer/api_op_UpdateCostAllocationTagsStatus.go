// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates status for cost allocation tags in bulk, with maximum batch size of 20.
// If the tag status that's updated is the same as the existing tag status, the
// request doesn't fail. Instead, it doesn't have any effect on the tag status (for
// example, activating the active tag).
func (c *Client) UpdateCostAllocationTagsStatus(ctx context.Context, params *UpdateCostAllocationTagsStatusInput, optFns ...func(*Options)) (*UpdateCostAllocationTagsStatusOutput, error) {
	if params == nil {
		params = &UpdateCostAllocationTagsStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCostAllocationTagsStatus", params, optFns, c.addOperationUpdateCostAllocationTagsStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCostAllocationTagsStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCostAllocationTagsStatusInput struct {

	// The list of CostAllocationTagStatusEntry objects that are used to update cost
	// allocation tags status for this request.
	//
	// This member is required.
	CostAllocationTagsStatus []types.CostAllocationTagStatusEntry

	noSmithyDocumentSerde
}

type UpdateCostAllocationTagsStatusOutput struct {

	// A list of UpdateCostAllocationTagsStatusError objects with error details about
	// each cost allocation tag that can't be updated. If there's no failure, an empty
	// array returns.
	Errors []types.UpdateCostAllocationTagsStatusError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateCostAllocationTagsStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateCostAllocationTagsStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateCostAllocationTagsStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateCostAllocationTagsStatus"); err != nil {
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
	if err = addOpUpdateCostAllocationTagsStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCostAllocationTagsStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateCostAllocationTagsStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateCostAllocationTagsStatus",
	}
}
