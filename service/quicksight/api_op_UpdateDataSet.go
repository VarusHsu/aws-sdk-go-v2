// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a dataset. This operation doesn't support datasets that include
// uploaded files as a source. Partial updates are not supported by this operation.
func (c *Client) UpdateDataSet(ctx context.Context, params *UpdateDataSetInput, optFns ...func(*Options)) (*UpdateDataSetOutput, error) {
	if params == nil {
		params = &UpdateDataSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDataSet", params, optFns, c.addOperationUpdateDataSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDataSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDataSetInput struct {

	// The Amazon Web Services account ID.
	//
	// This member is required.
	AwsAccountId *string

	// The ID for the dataset that you want to update. This ID is unique per Amazon
	// Web Services Region for each Amazon Web Services account.
	//
	// This member is required.
	DataSetId *string

	// Indicates whether you want to import the data into SPICE.
	//
	// This member is required.
	ImportMode types.DataSetImportMode

	// The display name for the dataset.
	//
	// This member is required.
	Name *string

	// Declares the physical tables that are available in the underlying data sources.
	//
	// This member is required.
	PhysicalTableMap map[string]types.PhysicalTable

	// Groupings of columns that work together in certain Amazon QuickSight features.
	// Currently, only geospatial hierarchy is supported.
	ColumnGroups []types.ColumnGroup

	// A set of one or more definitions of a [ColumnLevelPermissionRule].
	//
	// [ColumnLevelPermissionRule]: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_ColumnLevelPermissionRule.html
	ColumnLevelPermissionRules []types.ColumnLevelPermissionRule

	// The usage configuration to apply to child datasets that reference this dataset
	// as a source.
	DataSetUsageConfiguration *types.DataSetUsageConfiguration

	// The parameter declarations of the dataset.
	DatasetParameters []types.DatasetParameter

	// The folder that contains fields and nested subfolders for your dataset.
	FieldFolders map[string]types.FieldFolder

	// Configures the combination and transformation of the data from the physical
	// tables.
	LogicalTableMap map[string]types.LogicalTable

	// The row-level security configuration for the data you want to create.
	RowLevelPermissionDataSet *types.RowLevelPermissionDataSet

	// The configuration of tags on a dataset to set row-level security. Row-level
	// security tags are currently supported for anonymous embedding only.
	RowLevelPermissionTagConfiguration *types.RowLevelPermissionTagConfiguration

	noSmithyDocumentSerde
}

type UpdateDataSetOutput struct {

	// The Amazon Resource Name (ARN) of the dataset.
	Arn *string

	// The ID for the dataset that you want to create. This ID is unique per Amazon
	// Web Services Region for each Amazon Web Services account.
	DataSetId *string

	// The ARN for the ingestion, which is triggered as a result of dataset creation
	// if the import mode is SPICE.
	IngestionArn *string

	// The ID of the ingestion, which is triggered as a result of dataset creation if
	// the import mode is SPICE.
	IngestionId *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDataSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDataSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDataSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateDataSet"); err != nil {
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
	if err = addOpUpdateDataSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDataSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDataSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateDataSet",
	}
}
