// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the history of the specified workflow execution. The results may be
// split into multiple pages. To retrieve subsequent pages, make the call again
// using the nextPageToken returned by the initial call. This operation is
// eventually consistent. The results are best effort and may not exactly reflect
// recent updates and changes. Access Control You can use IAM policies to control
// this action's access to Amazon SWF resources as follows:
//
//     * Use a Resource
// element with the domain name to limit the action to only specified domains.
//
//
// * Use an Action element to allow or deny permission to call this action.
//
//     *
// You cannot use an IAM policy to constrain this action's parameters.
//
// If the
// caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED.
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows
// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) GetWorkflowExecutionHistory(ctx context.Context, params *GetWorkflowExecutionHistoryInput, optFns ...func(*Options)) (*GetWorkflowExecutionHistoryOutput, error) {
	stack := middleware.NewStack("GetWorkflowExecutionHistory", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpGetWorkflowExecutionHistoryMiddlewares(stack)
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
	addOpGetWorkflowExecutionHistoryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetWorkflowExecutionHistory(options.Region), middleware.Before)
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
			OperationName: "GetWorkflowExecutionHistory",
			Err:           err,
		}
	}
	out := result.(*GetWorkflowExecutionHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWorkflowExecutionHistoryInput struct {

	// When set to true, returns the events in reverse order. By default the results
	// are returned in ascending order of the eventTimeStamp of the events.
	ReverseOrder *bool

	// The name of the domain containing the workflow execution.
	//
	// This member is required.
	Domain *string

	// If NextPageToken is returned there are more results available. The value of
	// NextPageToken is a unique pagination token for each page. Make the call again
	// using the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 60 seconds. Using an expired
	// pagination token will return a 400 error: "Specified token has exceeded its
	// maximum lifetime".  <p>The configured <code>maximumPageSize</code> determines
	// how many results can be returned in a single call. </p>
	NextPageToken *string

	// Specifies the workflow execution for which to return the history.
	//
	// This member is required.
	Execution *types.WorkflowExecution

	// The maximum number of results that are returned per call. Use nextPageToken to
	// obtain further pages of results.
	MaximumPageSize *int32
}

// Paginated representation of a workflow history for a workflow execution. This is
// the up to date, complete and authoritative record of the events related to all
// tasks and events in the life of the workflow execution.
type GetWorkflowExecutionHistoryOutput struct {

	// The list of history events.
	//
	// This member is required.
	Events []*types.HistoryEvent

	// If a NextPageToken was returned by a previous call, there are more results
	// available. To retrieve the next page of results, make the call again using the
	// returned token in nextPageToken. Keep all other arguments unchanged. The
	// configured maximumPageSize determines how many results can be returned in a
	// single call.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpGetWorkflowExecutionHistoryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpGetWorkflowExecutionHistory{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetWorkflowExecutionHistory{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetWorkflowExecutionHistory(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "GetWorkflowExecutionHistory",
	}
}
