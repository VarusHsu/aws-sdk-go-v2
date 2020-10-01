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

// Returns the default engine and system parameter information for the cluster
// database engine. For more information on Amazon Aurora, see  What Is Amazon
// Aurora?
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide.
func (c *Client) DescribeEngineDefaultClusterParameters(ctx context.Context, params *DescribeEngineDefaultClusterParametersInput, optFns ...func(*Options)) (*DescribeEngineDefaultClusterParametersOutput, error) {
	stack := middleware.NewStack("DescribeEngineDefaultClusterParameters", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeEngineDefaultClusterParametersMiddlewares(stack)
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
	addOpDescribeEngineDefaultClusterParametersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEngineDefaultClusterParameters(options.Region), middleware.Before)
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
			OperationName: "DescribeEngineDefaultClusterParameters",
			Err:           err,
		}
	}
	out := result.(*DescribeEngineDefaultClusterParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeEngineDefaultClusterParametersInput struct {

	// An optional pagination token provided by a previous
	// DescribeEngineDefaultClusterParameters request. If this parameter is specified,
	// the response includes only records beyond the marker, up to the value specified
	// by MaxRecords.
	Marker *string

	// This parameter isn't currently supported.
	Filters []*types.Filter

	// The name of the DB cluster parameter group family to return engine parameter
	// information for.
	//
	// This member is required.
	DBParameterGroupFamily *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
}

type DescribeEngineDefaultClusterParametersOutput struct {

	// Contains the result of a successful invocation of the
	// DescribeEngineDefaultParameters action.
	EngineDefaults *types.EngineDefaults

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeEngineDefaultClusterParametersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeEngineDefaultClusterParameters{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeEngineDefaultClusterParameters{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeEngineDefaultClusterParameters(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeEngineDefaultClusterParameters",
	}
}
