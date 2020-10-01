// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the parameters of a DB parameter group to the engine/system default
// value. To reset specific parameters, provide a list of the following:
// ParameterName and ApplyMethod. To reset the entire DB parameter group, specify
// the DBParameterGroup name and ResetAllParameters parameters. When resetting the
// entire group, dynamic parameters are updated immediately and static parameters
// are set to pending-reboot to take effect on the next DB instance restart or
// RebootDBInstance request.
func (c *Client) ResetDBParameterGroup(ctx context.Context, params *ResetDBParameterGroupInput, optFns ...func(*Options)) (*ResetDBParameterGroupOutput, error) {
	stack := middleware.NewStack("ResetDBParameterGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpResetDBParameterGroupMiddlewares(stack)
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
	addOpResetDBParameterGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opResetDBParameterGroup(options.Region), middleware.Before)
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
			OperationName: "ResetDBParameterGroup",
			Err:           err,
		}
	}
	out := result.(*ResetDBParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResetDBParameterGroupInput struct {

	// The name of the DB parameter group. Constraints:
	//
	//     * Must match the name of
	// an existing DBParameterGroup.
	//
	// This member is required.
	DBParameterGroupName *string

	// To reset the entire DB parameter group, specify the DBParameterGroup name and
	// ResetAllParameters parameters. To reset specific parameters, provide a list of
	// the following: ParameterName and ApplyMethod. A maximum of 20 parameters can be
	// modified in a single request. Valid Values (for Apply method): pending-reboot
	Parameters []*types.Parameter

	// Specifies whether (true) or not (false) to reset all parameters in the DB
	// parameter group to default values. Default: true
	ResetAllParameters *bool
}

type ResetDBParameterGroupOutput struct {

	// Provides the name of the DB parameter group.
	DBParameterGroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpResetDBParameterGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpResetDBParameterGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpResetDBParameterGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opResetDBParameterGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ResetDBParameterGroup",
	}
}
