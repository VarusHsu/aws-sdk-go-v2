// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds one or more tags to an IAM user. If a tag with the same key name already
// exists, then that tag is overwritten with the new value.  <p>A tag consists of a
// key name and an associated value. By assigning tags to your resources, you can
// do the following:</p> <ul> <li> <p> <b>Administrative grouping and discovery</b>
// - Attach tags to resources to aid in organization and search. For example, you
// could search for all resources with the key name <i>Project</i> and the value
// <i>MyImportantProject</i>. Or search for all resources with the key name <i>Cost
// Center</i> and the value <i>41200</i>. </p> </li> <li> <p> <b>Access control</b>
// - Reference tags in IAM user-based and resource-based policies. You can use tags
// to restrict access to only an IAM requesting user or to a role that has a
// specified tag attached. You can also restrict access to only those resources
// that have a certain tag attached. For examples of policies that show how to use
// tags to control access, see <a
// href="https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html">Control
// Access Using IAM Tags</a> in the <i>IAM User Guide</i>.</p> </li> <li> <p>
// <b>Cost allocation</b> - Use tags to help track which individuals and teams are
// using which AWS resources.</p> </li> </ul> <note> <ul> <li> <p>Make sure that
// you have no invalid tags and that you do not exceed the allowed number of tags
// per role. In either case, the entire request fails and <i>no</i> tags are added
// to the role.</p> </li> <li> <p>AWS always interprets the tag <code>Value</code>
// as a single string. If you need to store an array, you can store comma-separated
// values in the string. However, you must interpret the value in your code.</p>
// </li> </ul> </note> <p>For more information about tagging, see <a
// href="https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html">Tagging IAM
// Identities</a> in the <i>IAM User Guide</i>.</p>
func (c *Client) TagUser(ctx context.Context, params *TagUserInput, optFns ...func(*Options)) (*TagUserOutput, error) {
	stack := middleware.NewStack("TagUser", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpTagUserMiddlewares(stack)
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
	addOpTagUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTagUser(options.Region), middleware.Before)
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
			OperationName: "TagUser",
			Err:           err,
		}
	}
	out := result.(*TagUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TagUserInput struct {

	// The list of tags that you want to attach to the user. Each tag consists of a key
	// name and an associated value.
	//
	// This member is required.
	Tags []*types.Tag

	// The name of the user that you want to add tags to. This parameter accepts
	// (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters that consist of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: =,.@-
	//
	// This member is required.
	UserName *string
}

type TagUserOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpTagUserMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpTagUser{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpTagUser{}, middleware.After)
}

func newServiceMetadataMiddleware_opTagUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "TagUser",
	}
}
