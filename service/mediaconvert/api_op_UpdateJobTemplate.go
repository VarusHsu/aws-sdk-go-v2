// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconvert

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modify one of your existing job templates.
func (c *Client) UpdateJobTemplate(ctx context.Context, params *UpdateJobTemplateInput, optFns ...func(*Options)) (*UpdateJobTemplateOutput, error) {
	if params == nil {
		params = &UpdateJobTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateJobTemplate", params, optFns, c.addOperationUpdateJobTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateJobTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateJobTemplateInput struct {

	// The name of the job template you are modifying
	//
	// This member is required.
	Name *string

	// Accelerated transcoding can significantly speed up jobs with long, visually
	// complex content. Outputs that use this feature incur pro-tier pricing. For
	// information about feature limitations, see the AWS Elemental MediaConvert User
	// Guide.
	AccelerationSettings *types.AccelerationSettings

	// The new category for the job template, if you are changing it.
	Category *string

	// The new description for the job template, if you are changing it.
	Description *string

	// Optional list of hop destinations.
	HopDestinations []types.HopDestination

	// Specify the relative priority for this job. In any given queue, the service
	// begins processing the job with the highest value first. When more than one job
	// has the same priority, the service begins processing the job that you submitted
	// first. If you don't specify a priority, the service uses the default value 0.
	Priority *int32

	// The new queue for the job template, if you are changing it.
	Queue *string

	// JobTemplateSettings contains all the transcode settings saved in the template
	// that will be applied to jobs created from it.
	Settings *types.JobTemplateSettings

	// Specify how often MediaConvert sends STATUS_UPDATE events to Amazon CloudWatch
	// Events. Set the interval, in seconds, between status updates. MediaConvert sends
	// an update at this interval from the time the service begins processing your job
	// to the time it completes the transcode or encounters an error.
	StatusUpdateInterval types.StatusUpdateInterval

	noSmithyDocumentSerde
}

type UpdateJobTemplateOutput struct {

	// A job template is a pre-made set of encoding instructions that you can use to
	// quickly create a job.
	JobTemplate *types.JobTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateJobTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateJobTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateJobTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateJobTemplate"); err != nil {
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
	if err = addOpUpdateJobTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateJobTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateJobTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateJobTemplate",
	}
}
