// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconvert

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a new transcoding job. For information about jobs and job settings, see
// the User Guide at http://docs.aws.amazon.com/mediaconvert/latest/ug/what-is.html
func (c *Client) CreateJob(ctx context.Context, params *CreateJobInput, optFns ...func(*Options)) (*CreateJobOutput, error) {
	if params == nil {
		params = &CreateJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateJob", params, optFns, c.addOperationCreateJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateJobInput struct {

	// Required. The IAM role you use for creating this job. For details about
	// permissions, see the User Guide topic at the User Guide at
	// https://docs.aws.amazon.com/mediaconvert/latest/ug/iam-role.html.
	//
	// This member is required.
	Role *string

	// JobSettings contains all the transcode settings for a job.
	//
	// This member is required.
	Settings *types.JobSettings

	// Optional. Accelerated transcoding can significantly speed up jobs with long,
	// visually complex content. Outputs that use this feature incur pro-tier pricing.
	// For information about feature limitations, see the AWS Elemental MediaConvert
	// User Guide.
	AccelerationSettings *types.AccelerationSettings

	// Optional. Choose a tag type that AWS Billing and Cost Management will use to
	// sort your AWS Elemental MediaConvert costs on any billing report that you set
	// up. Any transcoding outputs that don't have an associated tag will appear in
	// your billing report unsorted. If you don't choose a valid value for this field,
	// your job outputs will appear on the billing report unsorted.
	BillingTagsSource types.BillingTagsSource

	// Prevent duplicate jobs from being created and ensure idempotency for your
	// requests. A client request token can be any string that includes up to 64 ASCII
	// characters. If you reuse a client request token within one minute of a
	// successful request, the API returns the job details of the original request
	// instead. For more information see
	// https://docs.aws.amazon.com/mediaconvert/latest/apireference/idempotency.html.
	ClientRequestToken *string

	// Optional. Use queue hopping to avoid overly long waits in the backlog of the
	// queue that you submit your job to. Specify an alternate queue and the maximum
	// time that your job will wait in the initial queue before hopping. For more
	// information about this feature, see the AWS Elemental MediaConvert User Guide.
	HopDestinations []types.HopDestination

	// Optional. When you create a job, you can either specify a job template or
	// specify the transcoding settings individually.
	JobTemplate *string

	// Optional. Specify the relative priority for this job. In any given queue, the
	// service begins processing the job with the highest value first. When more than
	// one job has the same priority, the service begins processing the job that you
	// submitted first. If you don't specify a priority, the service uses the default
	// value 0.
	Priority *int32

	// Optional. When you create a job, you can specify a queue to send it to. If you
	// don't specify, the job will go to the default queue. For more about queues, see
	// the User Guide topic at
	// https://docs.aws.amazon.com/mediaconvert/latest/ug/what-is.html.
	Queue *string

	// Optional. Enable this setting when you run a test job to estimate how many
	// reserved transcoding slots (RTS) you need. When this is enabled, MediaConvert
	// runs your job from an on-demand queue with similar performance to what you will
	// see with one RTS in a reserved queue. This setting is disabled by default.
	SimulateReservedQueue types.SimulateReservedQueue

	// Optional. Specify how often MediaConvert sends STATUS_UPDATE events to Amazon
	// CloudWatch Events. Set the interval, in seconds, between status updates.
	// MediaConvert sends an update at this interval from the time the service begins
	// processing your job to the time it completes the transcode or encounters an
	// error.
	StatusUpdateInterval types.StatusUpdateInterval

	// Optional. The tags that you want to add to the resource. You can tag resources
	// with a key-value pair or with only a key. Use standard AWS tags on your job for
	// automatic integration with AWS services and for custom integrations and
	// workflows.
	Tags map[string]string

	// Optional. User-defined metadata that you want to associate with an MediaConvert
	// job. You specify metadata in key/value pairs. Use only for existing integrations
	// or workflows that rely on job metadata tags. Otherwise, we recommend that you
	// use standard AWS tags.
	UserMetadata map[string]string

	noSmithyDocumentSerde
}

type CreateJobOutput struct {

	// Each job converts an input file into an output file or files. For more
	// information, see the User Guide at
	// https://docs.aws.amazon.com/mediaconvert/latest/ug/what-is.html
	Job *types.Job

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateJob"); err != nil {
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
	if err = addIdempotencyToken_opCreateJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateJob(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateJobInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateJob",
	}
}
