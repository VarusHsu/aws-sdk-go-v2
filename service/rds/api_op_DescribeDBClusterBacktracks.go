// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about backtracks for a DB cluster. For more information on
// Amazon Aurora, see  What Is Amazon Aurora?
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide. This action only applies to Aurora MySQL DB
// clusters.
func (c *Client) DescribeDBClusterBacktracks(ctx context.Context, params *DescribeDBClusterBacktracksInput, optFns ...func(*Options)) (*DescribeDBClusterBacktracksOutput, error) {
	stack := middleware.NewStack("DescribeDBClusterBacktracks", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeDBClusterBacktracksMiddlewares(stack)
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
	addOpDescribeDBClusterBacktracksValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBClusterBacktracks(options.Region), middleware.Before)
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
			OperationName: "DescribeDBClusterBacktracks",
			Err:           err,
		}
	}
	out := result.(*DescribeDBClusterBacktracksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeDBClusterBacktracksInput struct {

	// A filter that specifies one or more DB clusters to describe. Supported filters
	// include the following:
	//
	//     * db-cluster-backtrack-id - Accepts backtrack
	// identifiers. The results list includes information about only the backtracks
	// identified by these identifiers.
	//
	//     * db-cluster-backtrack-status - Accepts
	// any of the following backtrack status values:
	//
	//         * applying
	//
	//         *
	// completed
	//
	//         * failed
	//
	//         * pending
	//
	//     The results list includes
	// information about only the backtracks identified by these values.
	Filters []*types.Filter

	// If specified, this value is the backtrack identifier of the backtrack to be
	// described. Constraints:
	//
	//     * Must contain a valid universally unique
	// identifier (UUID). For more information about UUIDs, see A Universally Unique
	// Identifier (UUID) URN Namespace (http://www.ietf.org/rfc/rfc4122.txt).
	//
	// Example:
	// 123e4567-e89b-12d3-a456-426655440000
	BacktrackIdentifier *string

	// The DB cluster identifier of the DB cluster to be described. This parameter is
	// stored as a lowercase string. Constraints:
	//
	//     * Must contain from 1 to 63
	// alphanumeric characters or hyphens.
	//
	//     * First character must be a letter.
	//
	//
	// * Can't end with a hyphen or contain two consecutive hyphens.
	//
	// Example:
	// my-cluster1
	//
	// This member is required.
	DBClusterIdentifier *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	// An optional pagination token provided by a previous DescribeDBClusterBacktracks
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string
}

// Contains the result of a successful invocation of the
// DescribeDBClusterBacktracks action.
type DescribeDBClusterBacktracksOutput struct {

	// A pagination token that can be used in a later DescribeDBClusterBacktracks
	// request.
	Marker *string

	// Contains a list of backtracks for the user.
	DBClusterBacktracks []*types.DBClusterBacktrack

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeDBClusterBacktracksMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBClusterBacktracks{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBClusterBacktracks{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeDBClusterBacktracks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBClusterBacktracks",
	}
}
