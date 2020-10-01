// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackage

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediapackage/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new HarvestJob record.
func (c *Client) CreateHarvestJob(ctx context.Context, params *CreateHarvestJobInput, optFns ...func(*Options)) (*CreateHarvestJobOutput, error) {
	stack := middleware.NewStack("CreateHarvestJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateHarvestJobMiddlewares(stack)
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
	addOpCreateHarvestJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHarvestJob(options.Region), middleware.Before)
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
			OperationName: "CreateHarvestJob",
			Err:           err,
		}
	}
	out := result.(*CreateHarvestJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Configuration parameters used to create a new HarvestJob.
type CreateHarvestJobInput struct {

	// The end of the time-window which will be harvested
	//
	// This member is required.
	EndTime *string

	// The ID of the OriginEndpoint that the HarvestJob will harvest from. This cannot
	// be changed after the HarvestJob is submitted.
	//
	// This member is required.
	OriginEndpointId *string

	// The start of the time-window which will be harvested
	//
	// This member is required.
	StartTime *string

	// The ID of the HarvestJob. The ID must be unique within the region and it cannot
	// be changed after the HarvestJob is submitted
	//
	// This member is required.
	Id *string

	// Configuration parameters for where in an S3 bucket to place the harvested
	// content
	//
	// This member is required.
	S3Destination *types.S3Destination
}

type CreateHarvestJobOutput struct {

	// The ID of the HarvestJob. The ID must be unique within the region and it cannot
	// be changed after the HarvestJob is submitted.
	Id *string

	// The start of the time-window which will be harvested.
	StartTime *string

	// The end of the time-window which will be harvested.
	EndTime *string

	// The Amazon Resource Name (ARN) assigned to the HarvestJob.
	Arn *string

	// The time the HarvestJob was submitted
	CreatedAt *string

	// The ID of the OriginEndpoint that the HarvestJob will harvest from. This cannot
	// be changed after the HarvestJob is submitted.
	OriginEndpointId *string

	// Configuration parameters for where in an S3 bucket to place the harvested
	// content
	S3Destination *types.S3Destination

	// The current status of the HarvestJob. Consider setting up a CloudWatch Event to
	// listen for HarvestJobs as they succeed or fail. In the event of failure, the
	// CloudWatch Event will include an explanation of why the HarvestJob failed.
	Status types.Status

	// The ID of the Channel that the HarvestJob will harvest from.
	ChannelId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateHarvestJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateHarvestJob{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateHarvestJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateHarvestJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediapackage",
		OperationName: "CreateHarvestJob",
	}
}
