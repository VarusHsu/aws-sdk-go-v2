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

// Copies the specified DB parameter group.
func (c *Client) CopyDBParameterGroup(ctx context.Context, params *CopyDBParameterGroupInput, optFns ...func(*Options)) (*CopyDBParameterGroupOutput, error) {
	stack := middleware.NewStack("CopyDBParameterGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCopyDBParameterGroupMiddlewares(stack)
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
	addOpCopyDBParameterGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCopyDBParameterGroup(options.Region), middleware.Before)
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
			OperationName: "CopyDBParameterGroup",
			Err:           err,
		}
	}
	out := result.(*CopyDBParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopyDBParameterGroupInput struct {

	// A description for the copied DB parameter group.
	//
	// This member is required.
	TargetDBParameterGroupDescription *string

	// The identifier for the copied DB parameter group. Constraints:
	//
	//     * Cannot be
	// null, empty, or blank.
	//
	//     * Must contain from 1 to 255 letters, numbers, or
	// hyphens.
	//
	//     * First character must be a letter.
	//
	//     * Cannot end with a
	// hyphen or contain two consecutive hyphens.
	//
	//     <p>Example:
	// <code>my-db-parameter-group</code> </p>
	//
	// This member is required.
	TargetDBParameterGroupIdentifier *string

	// The tags to be assigned to the copied DB parameter group.
	Tags []*types.Tag

	// The identifier or ARN for the source DB parameter group. For information about
	// creating an ARN, see  Constructing an Amazon Resource Name (ARN)
	// (https://docs.aws.amazon.com/neptune/latest/UserGuide/tagging.ARN.html#tagging.ARN.Constructing).
	// <p>Constraints:</p> <ul> <li> <p>Must specify a valid DB parameter group.</p>
	// </li> <li> <p>Must specify a valid DB parameter group identifier, for example
	// <code>my-db-param-group</code>, or a valid ARN.</p> </li> </ul>
	//
	// This member is required.
	SourceDBParameterGroupIdentifier *string
}

type CopyDBParameterGroupOutput struct {

	// Contains the details of an Amazon Neptune DB parameter group. This data type is
	// used as a response element in the DescribeDBParameterGroups () action.
	DBParameterGroup *types.DBParameterGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCopyDBParameterGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCopyDBParameterGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCopyDBParameterGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opCopyDBParameterGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CopyDBParameterGroup",
	}
}
