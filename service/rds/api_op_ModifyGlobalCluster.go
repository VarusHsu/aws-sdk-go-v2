// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies a setting for an Amazon Aurora global database cluster. You can change
// one or more database configuration parameters by specifying these parameters and
// the new values in the request. For more information on Amazon Aurora, see [What is Amazon Aurora?]in
// the Amazon Aurora User Guide.
//
// This operation only applies to Aurora global database clusters.
//
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
func (c *Client) ModifyGlobalCluster(ctx context.Context, params *ModifyGlobalClusterInput, optFns ...func(*Options)) (*ModifyGlobalClusterOutput, error) {
	if params == nil {
		params = &ModifyGlobalClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyGlobalCluster", params, optFns, c.addOperationModifyGlobalClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyGlobalClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyGlobalClusterInput struct {

	// Specifies whether to allow major version upgrades.
	//
	// Constraints: Must be enabled if you specify a value for the EngineVersion
	// parameter that's a different major version than the global cluster's current
	// version.
	//
	// If you upgrade the major version of a global database, the cluster and DB
	// instance parameter groups are set to the default parameter groups for the new
	// version. Apply any custom parameter groups after completing the upgrade.
	AllowMajorVersionUpgrade *bool

	// Specifies whether to enable deletion protection for the global database
	// cluster. The global database cluster can't be deleted when deletion protection
	// is enabled.
	DeletionProtection *bool

	// The version number of the database engine to which you want to upgrade.
	//
	// To list all of the available engine versions for aurora-mysql (for MySQL-based
	// Aurora global databases), use the following command:
	//
	//     aws rds describe-db-engine-versions --engine aurora-mysql --query
	//     '*[]|[?SupportsGlobalDatabases == `true`].[EngineVersion]'
	//
	// To list all of the available engine versions for aurora-postgresql (for
	// PostgreSQL-based Aurora global databases), use the following command:
	//
	//     aws rds describe-db-engine-versions --engine aurora-postgresql --query
	//     '*[]|[?SupportsGlobalDatabases == `true`].[EngineVersion]'
	EngineVersion *string

	// The cluster identifier for the global cluster to modify. This parameter isn't
	// case-sensitive.
	//
	// Constraints:
	//
	//   - Must match the identifier of an existing global database cluster.
	GlobalClusterIdentifier *string

	// The new cluster identifier for the global database cluster. This value is
	// stored as a lowercase string.
	//
	// Constraints:
	//
	//   - Must contain from 1 to 63 letters, numbers, or hyphens.
	//
	//   - The first character must be a letter.
	//
	//   - Can't end with a hyphen or contain two consecutive hyphens.
	//
	// Example: my-cluster2
	NewGlobalClusterIdentifier *string

	noSmithyDocumentSerde
}

type ModifyGlobalClusterOutput struct {

	// A data type representing an Aurora global database.
	GlobalCluster *types.GlobalCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyGlobalClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyGlobalCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyGlobalCluster{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyGlobalCluster"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyGlobalCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyGlobalCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyGlobalCluster",
	}
}
