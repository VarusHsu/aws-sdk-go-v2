// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets an Amazon DataZone environment.
func (c *Client) GetEnvironment(ctx context.Context, params *GetEnvironmentInput, optFns ...func(*Options)) (*GetEnvironmentOutput, error) {
	if params == nil {
		params = &GetEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEnvironment", params, optFns, c.addOperationGetEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEnvironmentInput struct {

	// The ID of the Amazon DataZone domain where the environment exists.
	//
	// This member is required.
	DomainIdentifier *string

	// The ID of the Amazon DataZone environment.
	//
	// This member is required.
	Identifier *string

	noSmithyDocumentSerde
}

type GetEnvironmentOutput struct {

	// The Amazon DataZone user who created the environment.
	//
	// This member is required.
	CreatedBy *string

	// The ID of the Amazon DataZone domain where the environment exists.
	//
	// This member is required.
	DomainId *string

	// The name of the environment.
	//
	// This member is required.
	Name *string

	// The ID of the Amazon DataZone project in which this environment is created.
	//
	// This member is required.
	ProjectId *string

	// The provider of this Amazon DataZone environment.
	//
	// This member is required.
	Provider *string

	// The ID of the Amazon Web Services account where the environment exists.
	AwsAccountId *string

	// The Amazon Web Services region where the environment exists.
	AwsAccountRegion *string

	// The timestamp of when the environment was created.
	CreatedAt *time.Time

	// The deployment properties of the environment.
	DeploymentProperties *types.DeploymentProperties

	// The description of the environment.
	Description *string

	// The actions of the environment.
	EnvironmentActions []types.ConfigurableEnvironmentAction

	// The blueprint with which the environment is created.
	EnvironmentBlueprintId *string

	// The ID of the environment profile with which the environment is created.
	EnvironmentProfileId *string

	// The business glossary terms that can be used in this environment.
	GlossaryTerms []string

	// The ID of the environment.
	Id *string

	// The details of the last deployment of the environment.
	LastDeployment *types.Deployment

	// The provisioned resources of this Amazon DataZone environment.
	ProvisionedResources []types.Resource

	// The provisioning properties of this Amazon DataZone environment.
	ProvisioningProperties types.ProvisioningProperties

	// The status of this Amazon DataZone environment.
	Status types.EnvironmentStatus

	// The timestamp of when this environment was updated.
	UpdatedAt *time.Time

	// The user parameters of this Amazon DataZone environment.
	UserParameters []types.CustomParameter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetEnvironment"); err != nil {
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
	if err = addOpGetEnvironmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEnvironment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetEnvironment",
	}
}
