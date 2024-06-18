// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new Elasticsearch domain. For more information, see [Creating Elasticsearch Domains] in the Amazon
// Elasticsearch Service Developer Guide.
//
// [Creating Elasticsearch Domains]: http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomains
func (c *Client) CreateElasticsearchDomain(ctx context.Context, params *CreateElasticsearchDomainInput, optFns ...func(*Options)) (*CreateElasticsearchDomainOutput, error) {
	if params == nil {
		params = &CreateElasticsearchDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateElasticsearchDomain", params, optFns, c.addOperationCreateElasticsearchDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateElasticsearchDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateElasticsearchDomainInput struct {

	// The name of the Elasticsearch domain that you are creating. Domain names are
	// unique across the domains owned by an account within an AWS region. Domain names
	// must start with a lowercase letter and can contain the following characters: a-z
	// (lowercase), 0-9, and - (hyphen).
	//
	// This member is required.
	DomainName *string

	//  IAM access policy as a JSON-formatted string.
	AccessPolicies *string

	//  Option to allow references to indices in an HTTP request body. Must be false
	// when configuring access to individual sub-resources. By default, the value is
	// true . See [Configuration Advanced Options] for more information.
	//
	// [Configuration Advanced Options]: http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-advanced-options
	AdvancedOptions map[string]string

	// Specifies advanced security options.
	AdvancedSecurityOptions *types.AdvancedSecurityOptionsInput

	// Specifies Auto-Tune options.
	AutoTuneOptions *types.AutoTuneOptionsInput

	// Options to specify the Cognito user and identity pools for Kibana
	// authentication. For more information, see [Amazon Cognito Authentication for Kibana].
	//
	// [Amazon Cognito Authentication for Kibana]: http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-cognito-auth.html
	CognitoOptions *types.CognitoOptions

	// Options to specify configuration that will be applied to the domain endpoint.
	DomainEndpointOptions *types.DomainEndpointOptions

	// Options to enable, disable and specify the type and size of EBS storage
	// volumes.
	EBSOptions *types.EBSOptions

	// Configuration options for an Elasticsearch domain. Specifies the instance type
	// and number of instances in the domain cluster.
	ElasticsearchClusterConfig *types.ElasticsearchClusterConfig

	// String of format X.Y to specify version for the Elasticsearch domain eg. "1.5"
	// or "2.3". For more information, see [Creating Elasticsearch Domains]in the Amazon Elasticsearch Service
	// Developer Guide.
	//
	// [Creating Elasticsearch Domains]: http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomains
	ElasticsearchVersion *string

	// Specifies the Encryption At Rest Options.
	EncryptionAtRestOptions *types.EncryptionAtRestOptions

	// Map of LogType and LogPublishingOption , each containing options to publish a
	// given type of Elasticsearch log.
	LogPublishingOptions map[string]types.LogPublishingOption

	// Specifies the NodeToNodeEncryptionOptions.
	NodeToNodeEncryptionOptions *types.NodeToNodeEncryptionOptions

	// Option to set time, in UTC format, of the daily automated snapshot. Default
	// value is 0 hours.
	SnapshotOptions *types.SnapshotOptions

	// A list of Tag added during domain creation.
	TagList []types.Tag

	// Options to specify the subnets and security groups for VPC endpoint. For more
	// information, see [Creating a VPC]in VPC Endpoints for Amazon Elasticsearch Service Domains
	//
	// [Creating a VPC]: http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-creating-vpc
	VPCOptions *types.VPCOptions

	noSmithyDocumentSerde
}

// The result of a CreateElasticsearchDomain operation. Contains the status of the
// newly created Elasticsearch domain.
type CreateElasticsearchDomainOutput struct {

	// The status of the newly created Elasticsearch domain.
	DomainStatus *types.ElasticsearchDomainStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateElasticsearchDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateElasticsearchDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateElasticsearchDomain{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateElasticsearchDomain"); err != nil {
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
	if err = addOpCreateElasticsearchDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateElasticsearchDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateElasticsearchDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateElasticsearchDomain",
	}
}
