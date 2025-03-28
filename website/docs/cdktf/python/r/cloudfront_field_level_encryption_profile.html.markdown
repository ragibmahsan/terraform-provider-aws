---
subcategory: "CloudFront"
layout: "aws"
page_title: "AWS: aws_cloudfront_field_level_encryption_profile"
description: |-
  Provides a CloudFront Field-level Encryption Profile resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_cloudfront_field_level_encryption_profile

Provides a CloudFront Field-level Encryption Profile resource.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.cloudfront_field_level_encryption_profile import CloudfrontFieldLevelEncryptionProfile
from imports.aws.cloudfront_public_key import CloudfrontPublicKey
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = CloudfrontPublicKey(self, "example",
            comment="test public key",
            encoded_key=Token.as_string(Fn.file("public_key.pem")),
            name="test_key"
        )
        CloudfrontFieldLevelEncryptionProfile(self, "test",
            comment="test comment",
            encryption_entities=CloudfrontFieldLevelEncryptionProfileEncryptionEntities(
                items=[CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItems(
                    field_patterns=CloudfrontFieldLevelEncryptionProfileEncryptionEntitiesItemsFieldPatterns(
                        items=["DateOfBirth"]
                    ),
                    provider_id="test provider",
                    public_key_id=example.id
                )
                ]
            ),
            name="test profile"
        )
```

## Argument Reference

This resource supports the following arguments:

* `name` - (Required) The name of the Field Level Encryption Profile.
* `comment` - (Optional) An optional comment about the Field Level Encryption Profile.
* `encryption_entities` - (Required) The [encryption entities](#encryption-entities) config block for field-level encryption profiles that contains an attribute `items` which includes the encryption key and field pattern specifications.

### Encryption Entities

* `public_key_id` - (Required) The public key associated with a set of field-level encryption patterns, to be used when encrypting the fields that match the patterns.
* `provider_id` - (Required) The provider associated with the public key being used for encryption.
* `field_patterns` - (Required) Object that contains an attribute `items` that contains the list of field patterns in a field-level encryption content type profile specify the fields that you want to be encrypted.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - The Field Level Encryption Profile ARN.
* `caller_reference` - Internal value used by CloudFront to allow future updates to the Field Level Encryption Profile.
* `etag` - The current version of the Field Level Encryption Profile. For example: `E2QWRUHAPOMQZL`.
* `id` - The identifier for the Field Level Encryption Profile. For example: `K3D5EWEUDCCXON`.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Cloudfront Field Level Encryption Profile using the `id`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.cloudfront_field_level_encryption_profile import CloudfrontFieldLevelEncryptionProfile
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        CloudfrontFieldLevelEncryptionProfile.generate_config_for_import(self, "profile", "K3D5EWEUDCCXON")
```

Using `terraform import`, import Cloudfront Field Level Encryption Profile using the `id`. For example:

```console
% terraform import aws_cloudfront_field_level_encryption_profile.profile K3D5EWEUDCCXON
```

<!-- cache-key: cdktf-0.20.8 input-e5e8e37e15fcd2e27bfd23e7acee3fed0fe068c9c90a71926b1ac34eb1b81810 -->