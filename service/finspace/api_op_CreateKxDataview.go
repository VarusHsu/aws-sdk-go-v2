// Code generated by smithy-go-codegen DO NOT EDIT.

package finspace

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/finspace/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

//	Creates a snapshot of kdb database with tiered storage capabilities and a
//
// pre-warmed cache, ready for mounting on kdb clusters. Dataviews are only
// available for clusters running on a scaling group. They are not supported on
// dedicated clusters.
func (c *Client) CreateKxDataview(ctx context.Context, params *CreateKxDataviewInput, optFns ...func(*Options)) (*CreateKxDataviewOutput, error) {
	if params == nil {
		params = &CreateKxDataviewInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateKxDataview", params, optFns, c.addOperationCreateKxDataviewMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateKxDataviewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateKxDataviewInput struct {

	// The number of availability zones you want to assign per volume. Currently,
	// FinSpace only supports SINGLE for volumes. This places dataview in a single AZ.
	//
	// This member is required.
	AzMode types.KxAzMode

	// A token that ensures idempotency. This token expires in 10 minutes.
	//
	// This member is required.
	ClientToken *string

	//  The name of the database where you want to create a dataview.
	//
	// This member is required.
	DatabaseName *string

	// A unique identifier for the dataview.
	//
	// This member is required.
	DataviewName *string

	// A unique identifier for the kdb environment, where you want to create the
	// dataview.
	//
	// This member is required.
	EnvironmentId *string

	// The option to specify whether you want to apply all the future additions and
	// corrections automatically to the dataview, when you ingest new changesets. The
	// default value is false.
	AutoUpdate bool

	//  The identifier of the availability zones.
	AvailabilityZoneId *string

	//  A unique identifier of the changeset that you want to use to ingest data.
	ChangesetId *string

	// A description of the dataview.
	Description *string

	//  The option to specify whether you want to make the dataview writable to
	// perform database maintenance. The following are some considerations related to
	// writable dataviews.
	//
	//   - You cannot create partial writable dataviews. When you create writeable
	//   dataviews you must provide the entire database path.
	//
	//   - You cannot perform updates on a writeable dataview. Hence, autoUpdate must
	//   be set as False if readWrite is True for a dataview.
	//
	//   - You must also use a unique volume for creating a writeable dataview. So, if
	//   you choose a volume that is already in use by another dataview, the dataview
	//   creation fails.
	//
	//   - Once you create a dataview as writeable, you cannot change it to read-only.
	//   So, you cannot update the readWrite parameter later.
	ReadWrite bool

	//  The configuration that contains the database path of the data that you want to
	// place on each selected volume. Each segment must have a unique database path for
	// each volume. If you do not explicitly specify any database path for a volume,
	// they are accessible from the cluster through the default S3/object store
	// segment.
	SegmentConfigurations []types.KxDataviewSegmentConfiguration

	//  A list of key-value pairs to label the dataview. You can add up to 50 tags to
	// a dataview.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateKxDataviewOutput struct {

	// The option to select whether you want to apply all the future additions and
	// corrections automatically to the dataview when you ingest new changesets. The
	// default value is false.
	AutoUpdate bool

	//  The identifier of the availability zones.
	AvailabilityZoneId *string

	// The number of availability zones you want to assign per volume. Currently,
	// FinSpace only supports SINGLE for volumes. This places dataview in a single AZ.
	AzMode types.KxAzMode

	// A unique identifier for the changeset.
	ChangesetId *string

	//  The timestamp at which the dataview was created in FinSpace. The value is
	// determined as epoch time in milliseconds. For example, the value for Monday,
	// November 1, 2021 12:00:00 PM UTC is specified as 1635768000000.
	CreatedTimestamp *time.Time

	// The name of the database where you want to create a dataview.
	DatabaseName *string

	// A unique identifier for the dataview.
	DataviewName *string

	// A description of the dataview.
	Description *string

	// A unique identifier for the kdb environment, where you want to create the
	// dataview.
	EnvironmentId *string

	//  The last time that the dataview was updated in FinSpace. The value is
	// determined as epoch time in milliseconds. For example, the value for Monday,
	// November 1, 2021 12:00:00 PM UTC is specified as 1635768000000.
	LastModifiedTimestamp *time.Time

	// Returns True if the dataview is created as writeable and False otherwise.
	ReadWrite bool

	//  The configuration that contains the database path of the data that you want to
	// place on each selected volume. Each segment must have a unique database path for
	// each volume. If you do not explicitly specify any database path for a volume,
	// they are accessible from the cluster through the default S3/object store
	// segment.
	SegmentConfigurations []types.KxDataviewSegmentConfiguration

	//  The status of dataview creation.
	//
	//   - CREATING – The dataview creation is in progress.
	//
	//   - UPDATING – The dataview is in the process of being updated.
	//
	//   - ACTIVE – The dataview is active.
	Status types.KxDataviewStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateKxDataviewMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateKxDataview{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateKxDataview{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateKxDataview"); err != nil {
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
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
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
	if err = addIdempotencyToken_opCreateKxDataviewMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateKxDataviewValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateKxDataview(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateKxDataview struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateKxDataview) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateKxDataview) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateKxDataviewInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateKxDataviewInput ")
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
func addIdempotencyToken_opCreateKxDataviewMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateKxDataview{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateKxDataview(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateKxDataview",
	}
}
