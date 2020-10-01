// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the parameters of a cluster parameter group to the default value. To
// reset specific parameters, submit a list of the following: ParameterName and
// ApplyMethod. To reset the entire cluster parameter group, specify the
// DBClusterParameterGroupName and ResetAllParameters parameters. When you reset
// the entire group, dynamic parameters are updated immediately and static
// parameters are set to pending-reboot to take effect on the next DB instance
// reboot.
func (c *Client) ResetDBClusterParameterGroup(ctx context.Context, params *ResetDBClusterParameterGroupInput, optFns ...func(*Options)) (*ResetDBClusterParameterGroupOutput, error) {
	stack := middleware.NewStack("ResetDBClusterParameterGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpResetDBClusterParameterGroupMiddlewares(stack)
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
	addOpResetDBClusterParameterGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opResetDBClusterParameterGroup(options.Region), middleware.Before)
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
			OperationName: "ResetDBClusterParameterGroup",
			Err:           err,
		}
	}
	out := result.(*ResetDBClusterParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to ResetDBClusterParameterGroup ().
type ResetDBClusterParameterGroupInput struct {

	// A value that is set to true to reset all parameters in the cluster parameter
	// group to their default values, and false otherwise. You can't use this parameter
	// if there is a list of parameter names specified for the Parameters parameter.
	ResetAllParameters *bool

	// The name of the cluster parameter group to reset.
	//
	// This member is required.
	DBClusterParameterGroupName *string

	// A list of parameter names in the cluster parameter group to reset to the default
	// values. You can't use this parameter if the ResetAllParameters parameter is set
	// to true.
	Parameters []*types.Parameter
}

// Contains the name of a cluster parameter group.
type ResetDBClusterParameterGroupOutput struct {

	// The name of a cluster parameter group. Constraints:
	//
	//     * Must be from 1 to 255
	// letters or numbers.
	//
	//     * The first character must be a letter.
	//
	//     * Cannot
	// end with a hyphen or contain two consecutive hyphens.
	//
	// This value is stored as a
	// lowercase string.
	DBClusterParameterGroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpResetDBClusterParameterGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpResetDBClusterParameterGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpResetDBClusterParameterGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opResetDBClusterParameterGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ResetDBClusterParameterGroup",
	}
}
