// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Replicates a multi-Region key into the specified Region. This operation creates
// a multi-Region replica key based on a multi-Region primary key in a different
// Region of the same AWS partition. You can create multiple replicas of a primary
// key, but each must be in a different Region. To create a multi-Region primary
// key, use the CreateKey operation. This operation supports multi-Region keys, an
// AWS KMS feature that lets you create multiple interoperable CMKs in different
// AWS Regions. Because these CMKs have the same key ID, key material, and other
// metadata, you can use them to encrypt data in one AWS Region and decrypt it in a
// different AWS Region without making a cross-Region call or exposing the
// plaintext data. For more information about multi-Region keys, see Using
// multi-Region keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html)
// in the AWS Key Management Service Developer Guide. A replica key is a
// fully-functional CMK that can be used independently of its primary and peer
// replica keys. A primary key and its replica keys share properties that make them
// interoperable. They have the same key ID
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-id)
// and key material. They also have the same key spec
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-spec),
// key usage
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-usage),
// key material origin
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-origin),
// and automatic key rotation status
// (https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html). AWS
// KMS automatically synchronizes these shared properties among related
// multi-Region keys. All other properties of a replica key can differ, including
// its key policy
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html), tags
// (https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html),
// aliases (https://docs.aws.amazon.com/kms/latest/developerguide/kms-alias.html),
// and key state
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html). AWS KMS
// pricing and quotas for CMKs apply to each primary key and replica key. When this
// operation completes, the new replica key has a transient key state of Creating.
// This key state changes to Enabled (or PendingImport) after a few seconds when
// the process of creating the new replica key is complete. While the key state is
// Creating, you can manage key, but you cannot yet use it in cryptographic
// operations. If you are creating and using the replica key programmatically,
// retry on KMSInvalidStateException or call DescribeKey to check its KeyState
// value before using it. For details about the Creating key state, see Key state:
// Effect on your CMK in the AWS Key Management Service Developer Guide. The AWS
// CloudTrail log of a ReplicateKey operation records a ReplicateKey operation in
// the primary key's Region and a CreateKey operation in the replica key's Region.
// If you replicate a multi-Region primary key with imported key material, the
// replica key is created with no key material. You must import the same key
// material that you imported into the primary key. For details, see Importing key
// material into multi-Region keys in the AWS Key Management Service Developer
// Guide. To convert a replica key to a primary key, use the UpdatePrimaryRegion
// operation. ReplicateKey uses different default values for the KeyPolicy and Tags
// parameters than those used in the AWS KMS console. For details, see the
// parameter descriptions. Cross-account use: No. You cannot use this operation to
// create a CMK in a different AWS account. Required permissions:
//
// *
// kms:ReplicateKey on the primary CMK (in the primary CMK's Region). Include this
// permission in the primary CMK's key policy.
//
// * kms:CreateKey in an IAM policy in
// the replica Region.
//
// * To use the Tags parameter, kms:TagResource in an IAM
// policy in the replica Region.
//
// Related operations
//
// * CreateKey
//
// *
// UpdatePrimaryRegion
func (c *Client) ReplicateKey(ctx context.Context, params *ReplicateKeyInput, optFns ...func(*Options)) (*ReplicateKeyOutput, error) {
	if params == nil {
		params = &ReplicateKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ReplicateKey", params, optFns, c.addOperationReplicateKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ReplicateKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ReplicateKeyInput struct {

	// Identifies the multi-Region primary key that is being replicated. To determine
	// whether a CMK is a multi-Region primary key, use the DescribeKey operation to
	// check the value of the MultiRegionKeyType property. Specify the key ID or key
	// ARN of a multi-Region primary key. For example:
	//
	// * Key ID:
	// mrk-1234abcd12ab34cd56ef1234567890ab
	//
	// * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/mrk-1234abcd12ab34cd56ef1234567890ab
	//
	// To
	// get the key ID and key ARN for a CMK, use ListKeys or DescribeKey.
	//
	// This member is required.
	KeyId *string

	// The Region ID of the AWS Region for this replica key. Enter the Region ID, such
	// as us-east-1 or ap-southeast-2. For a list of AWS Regions in which AWS KMS is
	// supported, see AWS KMS service endpoints
	// (https://docs.aws.amazon.com/general/latest/gr/kms.html#kms_region) in the
	// Amazon Web Services General Reference. The replica must be in a different AWS
	// Region than its primary key and other replicas of that primary key, but in the
	// same AWS partition. AWS KMS must be available in the replica Region. If the
	// Region is not enabled by default, the AWS account must be enabled in the Region.
	// For information about AWS partitions, see Amazon Resource Names (ARNs) in the
	// Amazon Web Services General Reference.
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) For
	// information about enabling and disabling Regions, see Enabling a Region
	// (https://docs.aws.amazon.com/general/latest/gr/rande-manage.html#rande-manage-enable)
	// and Disabling a Region
	// (https://docs.aws.amazon.com/general/latest/gr/rande-manage.html#rande-manage-disable)
	// in the Amazon Web Services General Reference.
	//
	// This member is required.
	ReplicaRegion *string

	// A flag to indicate whether to bypass the key policy lockout safety check.
	// Setting this value to true increases the risk that the CMK becomes unmanageable.
	// Do not set this value to true indiscriminately. For more information, refer to
	// the scenario in the Default Key Policy
	// (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam)
	// section in the AWS Key Management Service Developer Guide. Use this parameter
	// only when you intend to prevent the principal that is making the request from
	// making a subsequent PutKeyPolicy request on the CMK. The default value is false.
	BypassPolicyLockoutSafetyCheck bool

	// A description of the CMK. Use a description that helps you decide whether the
	// CMK is appropriate for a task. The default value is an empty string (no
	// description). The description is not a shared property of multi-Region keys. You
	// can specify the same description or a different description for each key in a
	// set of related multi-Region keys. AWS KMS does not synchronize this property.
	Description *string

	// The key policy to attach to the CMK. This parameter is optional. If you do not
	// provide a key policy, AWS KMS attaches the default key policy
	// (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default)
	// to the CMK. The key policy is not a shared property of multi-Region keys. You
	// can specify the same key policy or a different key policy for each key in a set
	// of related multi-Region keys. AWS KMS does not synchronize this property. If you
	// provide a key policy, it must meet the following criteria:
	//
	// * If you don't set
	// BypassPolicyLockoutSafetyCheck to true, the key policy must give the caller
	// kms:PutKeyPolicy permission on the replica CMK. This reduces the risk that the
	// CMK becomes unmanageable. For more information, refer to the scenario in the
	// Default Key Policy
	// (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam)
	// section of the AWS Key Management Service Developer Guide .
	//
	// * Each statement in
	// the key policy must contain one or more principals. The principals in the key
	// policy must exist and be visible to AWS KMS. When you create a new AWS principal
	// (for example, an IAM user or role), you might need to enforce a delay before
	// including the new principal in a key policy because the new principal might not
	// be immediately visible to AWS KMS. For more information, see Changes that I make
	// are not always immediately visible
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/troubleshoot_general.html#troubleshoot_general_eventual-consistency)
	// in the AWS Identity and Access Management User Guide.
	//
	// * The key policy size
	// quota is 32 kilobytes (32768 bytes).
	Policy *string

	// Assigns one or more tags to the replica key. Use this parameter to tag the CMK
	// when it is created. To tag an existing CMK, use the TagResource operation.
	// Tagging or untagging a CMK can allow or deny permission to the CMK. For details,
	// see Using ABAC in AWS KMS
	// (https://docs.aws.amazon.com/kms/latest/developerguide/abac.html) in the AWS Key
	// Management Service Developer Guide. To use this parameter, you must have
	// kms:TagResource
	// (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
	// permission in an IAM policy. Tags are not a shared property of multi-Region
	// keys. You can specify the same tags or different tags for each key in a set of
	// related multi-Region keys. AWS KMS does not synchronize this property. Each tag
	// consists of a tag key and a tag value. Both the tag key and the tag value are
	// required, but the tag value can be an empty (null) string. You cannot have more
	// than one tag on a CMK with the same tag key. If you specify an existing tag key
	// with a different tag value, AWS KMS replaces the current tag value with the
	// specified one. When you assign tags to an AWS resource, AWS generates a cost
	// allocation report with usage and costs aggregated by tags. Tags can also be used
	// to control access to a CMK. For details, see Tagging Keys
	// (https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html).
	Tags []types.Tag

	noSmithyDocumentSerde
}

type ReplicateKeyOutput struct {

	// Displays details about the new replica CMK, including its Amazon Resource Name
	// (key ARN
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-ARN))
	// and key state
	// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html). It also
	// includes the ARN and AWS Region of its primary key and other replica keys.
	ReplicaKeyMetadata *types.KeyMetadata

	// The key policy of the new replica key. The value is a key policy document in
	// JSON format.
	ReplicaPolicy *string

	// The tags on the new replica key. The value is a list of tag key and tag value
	// pairs.
	ReplicaTags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationReplicateKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpReplicateKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpReplicateKey{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpReplicateKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opReplicateKey(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opReplicateKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "ReplicateKey",
	}
}
