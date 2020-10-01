// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new fleet to run your game servers. whether they are custom game
// builds or Realtime Servers with game-specific script. A fleet is a set of Amazon
// Elastic Compute Cloud (Amazon EC2) instances, each of which can host multiple
// game sessions. When creating a fleet, you choose the hardware specifications,
// set some configuration options, and specify the game server to deploy on the new
// fleet. To create a new fleet, provide the following: (1) a fleet name, (2) an
// EC2 instance type and fleet type (spot or on-demand), (3) the build ID for your
// game build or script ID if using Realtime Servers, and (4) a runtime
// configuration, which determines how game servers will run on each instance in
// the fleet.  <p>If the <code>CreateFleet</code> call is successful, Amazon
// GameLift performs the following tasks. You can track the process of a fleet by
// checking the fleet status or by monitoring fleet creation events:</p> <ul> <li>
// <p>Creates a fleet resource. Status: <code>NEW</code>.</p> </li> <li> <p>Begins
// writing events to the fleet event log, which can be accessed in the Amazon
// GameLift console.</p> </li> <li> <p>Sets the fleet's target capacity to 1
// (desired instances), which triggers Amazon GameLift to start one new EC2
// instance.</p> </li> <li> <p>Downloads the game build or Realtime script to the
// new instance and installs it. Statuses: <code>DOWNLOADING</code>,
// <code>VALIDATING</code>, <code>BUILDING</code>. </p> </li> <li> <p>Starts
// launching server processes on the instance. If the fleet is configured to run
// multiple server processes per instance, Amazon GameLift staggers each process
// launch by a few seconds. Status: <code>ACTIVATING</code>.</p> </li> <li> <p>Sets
// the fleet's status to <code>ACTIVE</code> as soon as one server process is ready
// to host a game session.</p> </li> </ul> <p> <b>Learn more</b> </p> <p> <a
// href="https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html">Setting
// Up Fleets</a> </p> <p> <a
// href="https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-creating-debug.html#fleets-creating-debug-creation">Debug
// Fleet Creation Issues</a> </p> <p> <b>Related operations</b> </p> <ul> <li> <p>
// <a>CreateFleet</a> </p> </li> <li> <p> <a>ListFleets</a> </p> </li> <li> <p>
// <a>DeleteFleet</a> </p> </li> <li> <p> <a>DescribeFleetAttributes</a> </p> </li>
// <li> <p> <a>UpdateFleetAttributes</a> </p> </li> <li> <p>
// <a>StartFleetActions</a> or <a>StopFleetActions</a> </p> </li> </ul>
func (c *Client) CreateFleet(ctx context.Context, params *CreateFleetInput, optFns ...func(*Options)) (*CreateFleetOutput, error) {
	stack := middleware.NewStack("CreateFleet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateFleetMiddlewares(stack)
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
	addOpCreateFleetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFleet(options.Region), middleware.Before)
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
			OperationName: "CreateFleet",
			Err:           err,
		}
	}
	out := result.(*CreateFleetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type CreateFleetInput struct {

	// A unique identifier for the AWS account with the VPC that you want to peer your
	// Amazon GameLift fleet with. You can find your account ID in the AWS Management
	// Console under account settings.
	PeerVpcAwsAccountId *string

	// Instructions for launching server processes on each instance in the fleet.
	// Server processes run either a custom game build executable or a Realtime script.
	// The runtime configuration defines the server executables or launch script file,
	// launch parameters, and the number of processes to run concurrently on each
	// instance. When creating a fleet, the runtime configuration must have at least
	// one server process configuration; otherwise the request fails with an invalid
	// request exception. (This parameter replaces the parameters ServerLaunchPath and
	// ServerLaunchParameters, although requests that contain values for these
	// parameters instead of a runtime configuration will continue to work.) This
	// parameter is required unless the parameters ServerLaunchPath and
	// ServerLaunchParameters are defined. Runtime configuration replaced these
	// parameters, but fleets that use them will continue to work.
	RuntimeConfiguration *types.RuntimeConfiguration

	// A unique identifier for an AWS IAM role that manages access to your AWS
	// services. With an instance role ARN set, any application that runs on an
	// instance in this fleet can assume the role, including install scripts, server
	// processes, and daemons (background processes). Create a role or look up a role's
	// ARN from the IAM dashboard (https://console.aws.amazon.com/iam/) in the AWS
	// Management Console. Learn more about using on-box credentials for your game
	// servers at  Access external resources from a game server
	// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-resources.html).
	InstanceRoleArn *string

	// A list of labels to assign to the new fleet resource. Tags are developer-defined
	// key-value pairs. Tagging AWS resources are useful for resource management,
	// access management and cost allocation. For more information, see  Tagging AWS
	// Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in
	// the AWS General Reference. Once the resource is created, you can use TagResource
	// (), UntagResource (), and ListTagsForResource () to add, remove, and view tags.
	// The maximum tag limit may be lower than stated. See the AWS General Reference
	// for actual tagging limits.
	Tags []*types.Tag

	// The name of an Amazon CloudWatch metric group to add this fleet to. A metric
	// group aggregates the metrics for all fleets in the group. Specify an existing
	// metric group name, or provide a new name to create a new metric group. A fleet
	// can only be included in one metric group at a time.
	MetricGroups []*string

	// A policy that limits the number of game sessions an individual player can create
	// over a span of time for this fleet.
	ResourceCreationLimitPolicy *types.ResourceCreationLimitPolicy

	// A unique identifier for a build to be deployed on the new fleet. You can use
	// either the build ID or ARN value. The custom game server build must have been
	// successfully uploaded to Amazon GameLift and be in a READY status. This fleet
	// setting cannot be changed once the fleet is created.
	BuildId *string

	// A game session protection policy to apply to all instances in this fleet. If
	// this parameter is not set, instances in this fleet default to no protection. You
	// can change a fleet's protection policy using UpdateFleetAttributes (), but this
	// change will only affect sessions created after the policy change. You can also
	// set protection for individual instances using UpdateGameSession ().
	//
	//     *
	// NoProtection - The game session can be terminated during a scale-down event.
	//
	//
	// * FullProtection - If the game session is in an ACTIVE status, it cannot be
	// terminated during a scale-down event.
	NewGameSessionProtectionPolicy types.ProtectionPolicy

	// A unique identifier for a Realtime script to be deployed on the new fleet. You
	// can use either the script ID or ARN value. The Realtime script must have been
	// successfully uploaded to Amazon GameLift. This fleet setting cannot be changed
	// once the fleet is created.
	ScriptId *string

	// Range of IP addresses and port settings that permit inbound traffic to access
	// game sessions that are running on the fleet. For fleets using a custom game
	// build, this parameter is required before game sessions running on the fleet can
	// accept connections. For Realtime Servers fleets, Amazon GameLift automatically
	// sets TCP and UDP ranges for use by the Realtime servers. You can specify
	// multiple permission settings or add more by updating the fleet.
	EC2InboundPermissions []*types.IpPermission

	// The name of an EC2 instance type that is supported in Amazon GameLift. A fleet
	// instance type determines the computing resources of each instance in the fleet,
	// including CPU, memory, storage, and networking capacity. Amazon GameLift
	// supports the following EC2 instance types. See Amazon EC2 Instance Types
	// (http://aws.amazon.com/ec2/instance-types/) for detailed descriptions.
	//
	// This member is required.
	EC2InstanceType types.EC2InstanceType

	// A human-readable description of a fleet.
	Description *string

	// This parameter is no longer used. Instead, specify server launch parameters in
	// the RuntimeConfiguration parameter. (Requests that specify a server launch path
	// and launch parameters instead of a runtime configuration will continue to work.)
	ServerLaunchParameters *string

	// Indicates whether to use On-Demand instances or Spot instances for this fleet.
	// If empty, the default is ON_DEMAND. Both categories of instances use identical
	// hardware and configurations based on the instance type selected for this fleet.
	// Learn more about  On-Demand versus Spot Instances
	// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-ec2-instances.html#gamelift-ec2-instances-spot).
	FleetType types.FleetType

	// Indicates whether to generate a TLS/SSL certificate for the new fleet. TLS
	// certificates are used for encrypting traffic between game clients and game
	// servers running on GameLift. If this parameter is not specified, the default
	// value, DISABLED, is used. This fleet setting cannot be changed once the fleet is
	// created. Learn more at Securing Client/Server Communication
	// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-howitworks.html#gamelift-howitworks-security).
	// Note: This feature requires the AWS Certificate Manager (ACM) service, which is
	// available in the AWS global partition but not in all other partitions. When
	// working in a partition that does not support this feature, a request for a new
	// fleet with certificate generation results fails with a 4xx unsupported Region
	// error. Valid values include:
	//
	//     * GENERATED - Generate a TLS/SSL certificate
	// for this fleet.
	//
	//     * DISABLED - (default) Do not generate a TLS/SSL
	// certificate for this fleet.
	CertificateConfiguration *types.CertificateConfiguration

	// This parameter is no longer used. Instead, specify a server launch path using
	// the RuntimeConfiguration parameter. Requests that specify a server launch path
	// and launch parameters instead of a runtime configuration will continue to work.
	ServerLaunchPath *string

	// A descriptive label that is associated with a fleet. Fleet names do not need to
	// be unique.
	//
	// This member is required.
	Name *string

	// This parameter is no longer used. Instead, to specify where Amazon GameLift
	// should store log files once a server process shuts down, use the Amazon GameLift
	// server API ProcessReady() and specify one or more directory paths in
	// logParameters. See more information in the Server API Reference
	// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api-ref.html#gamelift-sdk-server-api-ref-dataypes-process).
	LogPaths []*string

	// A unique identifier for a VPC with resources to be accessed by your Amazon
	// GameLift fleet. The VPC must be in the same Region as your fleet. To look up a
	// VPC ID, use the VPC Dashboard (https://console.aws.amazon.com/vpc/) in the AWS
	// Management Console. Learn more about VPC peering in VPC Peering with Amazon
	// GameLift Fleets
	// (https://docs.aws.amazon.com/gamelift/latest/developerguide/vpc-peering.html).
	PeerVpcId *string
}

// Represents the returned data in response to a request action.
type CreateFleetOutput struct {

	// Properties for the newly created fleet.
	FleetAttributes *types.FleetAttributes

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateFleetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateFleet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateFleet{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateFleet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "CreateFleet",
	}
}
