// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates and defines the settings for a classification job.
func (c *Client) CreateClassificationJob(ctx context.Context, params *CreateClassificationJobInput, optFns ...func(*Options)) (*CreateClassificationJobOutput, error) {
	if params == nil {
		params = &CreateClassificationJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateClassificationJob", params, optFns, c.addOperationCreateClassificationJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateClassificationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateClassificationJobInput struct {

	// A unique, case-sensitive token that you provide to ensure the idempotency of
	// the request.
	//
	// This member is required.
	ClientToken *string

	// The schedule for running the job. Valid values are:
	//   - ONE_TIME - Run the job only once. If you specify this value, don't specify
	//   a value for the scheduleFrequency property.
	//   - SCHEDULED - Run the job on a daily, weekly, or monthly basis. If you
	//   specify this value, use the scheduleFrequency property to define the recurrence
	//   pattern for the job.
	//
	// This member is required.
	JobType types.JobType

	// A custom name for the job. The name can contain as many as 500 characters.
	//
	// This member is required.
	Name *string

	// The S3 buckets that contain the objects to analyze, and the scope of that
	// analysis.
	//
	// This member is required.
	S3JobDefinition *types.S3JobDefinition

	// An array of unique identifiers, one for each allow list for the job to use when
	// it analyzes data.
	AllowListIds []string

	// An array of unique identifiers, one for each custom data identifier for the job
	// to use when it analyzes data. To use only managed data identifiers, don't
	// specify a value for this property and specify a value other than NONE for the
	// managedDataIdentifierSelector property.
	CustomDataIdentifierIds []string

	// A custom description of the job. The description can contain as many as 200
	// characters.
	Description *string

	// For a recurring job, specifies whether to analyze all existing, eligible
	// objects immediately after the job is created (true). To analyze only those
	// objects that are created or changed after you create the job and before the
	// job's first scheduled run, set this value to false. If you configure the job to
	// run only once, don't specify a value for this property.
	InitialRun *bool

	// An array of unique identifiers, one for each managed data identifier for the
	// job to include (use) or exclude (not use) when it analyzes data. Inclusion or
	// exclusion depends on the managed data identifier selection type that you specify
	// for the job (managedDataIdentifierSelector). To retrieve a list of valid values
	// for this property, use the ListManagedDataIdentifiers operation.
	ManagedDataIdentifierIds []string

	// The selection type to apply when determining which managed data identifiers the
	// job uses to analyze data. Valid values are:
	//   - ALL - Use all managed data identifiers. If you specify this value, don't
	//   specify any values for the managedDataIdentifierIds property.
	//   - EXCLUDE - Use all managed data identifiers except the ones specified by the
	//   managedDataIdentifierIds property.
	//   - INCLUDE - Use only the managed data identifiers specified by the
	//   managedDataIdentifierIds property.
	//   - NONE - Don't use any managed data identifiers. If you specify this value,
	//   specify at least one value for the customDataIdentifierIds property and don't
	//   specify any values for the managedDataIdentifierIds property.
	//   - RECOMMENDED (default) - Use the recommended set of managed data
	//   identifiers. If you specify this value, don't specify any values for the
	//   managedDataIdentifierIds property.
	// If you don't specify a value for this property, the job uses the recommended
	// set of managed data identifiers. If the job is a recurring job and you specify
	// ALL or EXCLUDE, each job run automatically uses new managed data identifiers
	// that are released. If you specify RECOMMENDED for a recurring job, each job run
	// automatically uses all the managed data identifiers that are in the recommended
	// set when the run starts. For information about individual managed data
	// identifiers or to determine which ones are in the recommended set, see Using
	// managed data identifiers (https://docs.aws.amazon.com/macie/latest/user/managed-data-identifiers.html)
	// and Recommended managed data identifiers (https://docs.aws.amazon.com/macie/latest/user/discovery-jobs-mdis-recommended.html)
	// in the Amazon Macie User Guide.
	ManagedDataIdentifierSelector types.ManagedDataIdentifierSelector

	// The sampling depth, as a percentage, for the job to apply when processing
	// objects. This value determines the percentage of eligible objects that the job
	// analyzes. If this value is less than 100, Amazon Macie selects the objects to
	// analyze at random, up to the specified percentage, and analyzes all the data in
	// those objects.
	SamplingPercentage *int32

	// The recurrence pattern for running the job. To run the job only once, don't
	// specify a value for this property and set the value for the jobType property to
	// ONE_TIME.
	ScheduleFrequency *types.JobScheduleFrequency

	// A map of key-value pairs that specifies the tags to associate with the job. A
	// job can have a maximum of 50 tags. Each tag consists of a tag key and an
	// associated tag value. The maximum length of a tag key is 128 characters. The
	// maximum length of a tag value is 256 characters.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateClassificationJobOutput struct {

	// The Amazon Resource Name (ARN) of the job.
	JobArn *string

	// The unique identifier for the job.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateClassificationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateClassificationJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateClassificationJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateClassificationJob"); err != nil {
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
	if err = addIdempotencyToken_opCreateClassificationJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateClassificationJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateClassificationJob(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateClassificationJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateClassificationJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateClassificationJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateClassificationJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateClassificationJobInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateClassificationJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateClassificationJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateClassificationJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateClassificationJob",
	}
}
