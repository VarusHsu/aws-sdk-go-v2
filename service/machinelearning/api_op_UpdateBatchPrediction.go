// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the BatchPredictionName of a BatchPrediction. You can use the
// GetBatchPrediction operation to view the contents of the updated data element.
func (c *Client) UpdateBatchPrediction(ctx context.Context, params *UpdateBatchPredictionInput, optFns ...func(*Options)) (*UpdateBatchPredictionOutput, error) {
	stack := middleware.NewStack("UpdateBatchPrediction", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateBatchPredictionMiddlewares(stack)
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
	addOpUpdateBatchPredictionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateBatchPrediction(options.Region), middleware.Before)
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
			OperationName: "UpdateBatchPrediction",
			Err:           err,
		}
	}
	out := result.(*UpdateBatchPredictionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateBatchPredictionInput struct {

	// A new user-supplied name or description of the BatchPrediction.
	//
	// This member is required.
	BatchPredictionName *string

	// The ID assigned to the BatchPrediction during creation.
	//
	// This member is required.
	BatchPredictionId *string
}

// Represents the output of an UpdateBatchPrediction operation. You can see the
// updated content by using the GetBatchPrediction operation.
type UpdateBatchPredictionOutput struct {

	// The ID assigned to the BatchPrediction during creation. This value should be
	// identical to the value of the BatchPredictionId in the request.
	BatchPredictionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateBatchPredictionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateBatchPrediction{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateBatchPrediction{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateBatchPrediction(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "machinelearning",
		OperationName: "UpdateBatchPrediction",
	}
}
