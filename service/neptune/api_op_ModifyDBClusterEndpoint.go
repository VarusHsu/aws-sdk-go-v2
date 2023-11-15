// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the properties of an endpoint in an Amazon Neptune DB cluster.
func (c *Client) ModifyDBClusterEndpoint(ctx context.Context, params *ModifyDBClusterEndpointInput, optFns ...func(*Options)) (*ModifyDBClusterEndpointOutput, error) {
	if params == nil {
		params = &ModifyDBClusterEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyDBClusterEndpoint", params, optFns, c.addOperationModifyDBClusterEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyDBClusterEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyDBClusterEndpointInput struct {

	// The identifier of the endpoint to modify. This parameter is stored as a
	// lowercase string.
	//
	// This member is required.
	DBClusterEndpointIdentifier *string

	// The type of the endpoint. One of: READER , WRITER , ANY .
	EndpointType *string

	// List of DB instance identifiers that aren't part of the custom endpoint group.
	// All other eligible instances are reachable through the custom endpoint. Only
	// relevant if the list of static members is empty.
	ExcludedMembers []string

	// List of DB instance identifiers that are part of the custom endpoint group.
	StaticMembers []string

	noSmithyDocumentSerde
}

// This data type represents the information you need to connect to an Amazon
// Neptune DB cluster. This data type is used as a response element in the
// following actions:
//   - CreateDBClusterEndpoint
//   - DescribeDBClusterEndpoints
//   - ModifyDBClusterEndpoint
//   - DeleteDBClusterEndpoint
//
// For the data structure that represents Amazon RDS DB instance endpoints, see
// Endpoint .
type ModifyDBClusterEndpointOutput struct {

	// The type associated with a custom endpoint. One of: READER , WRITER , ANY .
	CustomEndpointType *string

	// The Amazon Resource Name (ARN) for the endpoint.
	DBClusterEndpointArn *string

	// The identifier associated with the endpoint. This parameter is stored as a
	// lowercase string.
	DBClusterEndpointIdentifier *string

	// A unique system-generated identifier for an endpoint. It remains the same for
	// the whole life of the endpoint.
	DBClusterEndpointResourceIdentifier *string

	// The DB cluster identifier of the DB cluster associated with the endpoint. This
	// parameter is stored as a lowercase string.
	DBClusterIdentifier *string

	// The DNS address of the endpoint.
	Endpoint *string

	// The type of the endpoint. One of: READER , WRITER , CUSTOM .
	EndpointType *string

	// List of DB instance identifiers that aren't part of the custom endpoint group.
	// All other eligible instances are reachable through the custom endpoint. Only
	// relevant if the list of static members is empty.
	ExcludedMembers []string

	// List of DB instance identifiers that are part of the custom endpoint group.
	StaticMembers []string

	// The current status of the endpoint. One of: creating , available , deleting ,
	// inactive , modifying . The inactive state applies to an endpoint that cannot be
	// used for a certain kind of cluster, such as a writer endpoint for a read-only
	// secondary cluster in a global database.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyDBClusterEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBClusterEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBClusterEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyDBClusterEndpoint"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpModifyDBClusterEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBClusterEndpoint(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opModifyDBClusterEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyDBClusterEndpoint",
	}
}
