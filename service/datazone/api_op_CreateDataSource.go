// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates an Amazon DataZone data source.
func (c *Client) CreateDataSource(ctx context.Context, params *CreateDataSourceInput, optFns ...func(*Options)) (*CreateDataSourceOutput, error) {
	if params == nil {
		params = &CreateDataSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDataSource", params, optFns, c.addOperationCreateDataSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDataSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDataSourceInput struct {

	// The ID of the Amazon DataZone domain where the data source is created.
	//
	// This member is required.
	DomainIdentifier *string

	// The unique identifier of the Amazon DataZone environment to which the data
	// source publishes assets.
	//
	// This member is required.
	EnvironmentIdentifier *string

	// The name of the data source.
	//
	// This member is required.
	Name *string

	// The identifier of the Amazon DataZone project in which you want to add this
	// data source.
	//
	// This member is required.
	ProjectIdentifier *string

	// The type of the data source.
	//
	// This member is required.
	Type *string

	// The metadata forms that are to be attached to the assets that this data source
	// works with.
	AssetFormsInput []types.FormInput

	// A unique, case-sensitive identifier that is provided to ensure the idempotency
	// of the request.
	ClientToken *string

	// Specifies the configuration of the data source. It can be set to either
	// glueRunConfiguration or redshiftRunConfiguration .
	Configuration types.DataSourceConfigurationInput

	// The description of the data source.
	Description *string

	// Specifies whether the data source is enabled.
	EnableSetting types.EnableSetting

	// Specifies whether the assets that this data source creates in the inventory are
	// to be also automatically published to the catalog.
	PublishOnImport *bool

	// Specifies whether the business name generation is to be enabled for this data
	// source.
	Recommendation *types.RecommendationConfiguration

	// The schedule of the data source runs.
	Schedule *types.ScheduleConfiguration

	noSmithyDocumentSerde
}

type CreateDataSourceOutput struct {

	// The ID of the Amazon DataZone domain in which the data source is created.
	//
	// This member is required.
	DomainId *string

	// The unique identifier of the Amazon DataZone environment to which the data
	// source publishes assets.
	//
	// This member is required.
	EnvironmentId *string

	// The unique identifier of the data source.
	//
	// This member is required.
	Id *string

	// The name of the data source.
	//
	// This member is required.
	Name *string

	// The ID of the Amazon DataZone project to which the data source is added.
	//
	// This member is required.
	ProjectId *string

	// The metadata forms attached to the assets that this data source creates.
	AssetFormsOutput []types.FormOutput

	// Specifies the configuration of the data source. It can be set to either
	// glueRunConfiguration or redshiftRunConfiguration .
	Configuration types.DataSourceConfigurationOutput

	// The timestamp of when the data source was created.
	CreatedAt *time.Time

	// The description of the data source.
	Description *string

	// Specifies whether the data source is enabled.
	EnableSetting types.EnableSetting

	// Specifies the error message that is returned if the operation cannot be
	// successfully completed.
	ErrorMessage *types.DataSourceErrorMessage

	// The timestamp that specifies when the data source was last run.
	LastRunAt *time.Time

	// Specifies the error message that is returned if the operation cannot be
	// successfully completed.
	LastRunErrorMessage *types.DataSourceErrorMessage

	// The status of the last run of this data source.
	LastRunStatus types.DataSourceRunStatus

	// Specifies whether the assets that this data source creates in the inventory are
	// to be also automatically published to the catalog.
	PublishOnImport *bool

	// Specifies whether the business name generation is to be enabled for this data
	// source.
	Recommendation *types.RecommendationConfiguration

	// The schedule of the data source runs.
	Schedule *types.ScheduleConfiguration

	// The status of the data source.
	Status types.DataSourceStatus

	// The type of the data source.
	Type *string

	// The timestamp of when the data source was updated.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDataSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDataSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDataSource{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDataSource"); err != nil {
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
	if err = addIdempotencyToken_opCreateDataSourceMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateDataSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataSource(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateDataSource struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateDataSource) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateDataSource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateDataSourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateDataSourceInput ")
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
func addIdempotencyToken_opCreateDataSourceMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateDataSource{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateDataSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDataSource",
	}
}
