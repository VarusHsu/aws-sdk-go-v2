// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Returns an array of one or more targets associated with a deployment. This
//
// method works with all compute types and should be used instead of the deprecated
// BatchGetDeploymentInstances . The maximum number of targets that can be returned
// is 25.
//
// The type of targets returned depends on the deployment's compute platform or
// deployment method:
//
//   - EC2/On-premises: Information about Amazon EC2 instance targets.
//
//   - Lambda: Information about Lambda functions targets.
//
//   - Amazon ECS: Information about Amazon ECS service targets.
//
//   - CloudFormation: Information about targets of blue/green deployments
//     initiated by a CloudFormation stack update.
func (c *Client) BatchGetDeploymentTargets(ctx context.Context, params *BatchGetDeploymentTargetsInput, optFns ...func(*Options)) (*BatchGetDeploymentTargetsOutput, error) {
	if params == nil {
		params = &BatchGetDeploymentTargetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetDeploymentTargets", params, optFns, c.addOperationBatchGetDeploymentTargetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetDeploymentTargetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetDeploymentTargetsInput struct {

	//  The unique ID of a deployment.
	//
	// This member is required.
	DeploymentId *string

	//  The unique IDs of the deployment targets. The compute platform of the
	// deployment determines the type of the targets and their formats. The maximum
	// number of deployment target IDs you can specify is 25.
	//
	//   - For deployments that use the EC2/On-premises compute platform, the target
	//   IDs are Amazon EC2 or on-premises instances IDs, and their target type is
	//   instanceTarget .
	//
	//   - For deployments that use the Lambda compute platform, the target IDs are
	//   the names of Lambda functions, and their target type is instanceTarget .
	//
	//   - For deployments that use the Amazon ECS compute platform, the target IDs
	//   are pairs of Amazon ECS clusters and services specified using the format : .
	//   Their target type is ecsTarget .
	//
	//   - For deployments that are deployed with CloudFormation, the target IDs are
	//   CloudFormation stack IDs. Their target type is cloudFormationTarget .
	//
	// This member is required.
	TargetIds []string

	noSmithyDocumentSerde
}

type BatchGetDeploymentTargetsOutput struct {

	//  A list of target objects for a deployment. Each target object contains details
	// about the target, such as its status and lifecycle events. The type of the
	// target objects depends on the deployment' compute platform.
	//
	//   - EC2/On-premises: Each target object is an Amazon EC2 or on-premises
	//   instance.
	//
	//   - Lambda: The target object is a specific version of an Lambda function.
	//
	//   - Amazon ECS: The target object is an Amazon ECS service.
	//
	//   - CloudFormation: The target object is an CloudFormation blue/green
	//   deployment.
	DeploymentTargets []types.DeploymentTarget

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetDeploymentTargetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetDeploymentTargets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetDeploymentTargets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchGetDeploymentTargets"); err != nil {
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
	if err = addOpBatchGetDeploymentTargetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetDeploymentTargets(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetDeploymentTargets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchGetDeploymentTargets",
	}
}
