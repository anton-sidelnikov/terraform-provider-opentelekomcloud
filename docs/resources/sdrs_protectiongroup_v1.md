---
subcategory: "Storage Disaster Recovery Service (SDRS)"
---

# opentelekomcloud_sdrs_protectiongroup_v1

Manages a SDRS protection group resource within OpenTelekomCloud.

## Example Usage

```hcl
data "opentelekomcloud_sdrs_domain_v1" "dom_1" {}

resource "opentelekomcloud_sdrs_protectiongroup_v1" "group_1" {
  name        = "group_1"
  description = "test description"

  source_availability_zone = "eu-de-01"
  target_availability_zone = "eu-de-02"

  domain_id     = data.opentelekomcloud_sdrs_domain_v1.dom_1.id
  source_vpc_id = var.vpc_id
  dr_type       = "migration"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of a protection group.

* `description` - (Optional) The description of a protection group. Changing this creates a new group.

* `source_availability_zone` - (Required) Specifies the source AZ of a protection group. Changing this creates a new group.

* `target_availability_zone` - (Required) Specifies the target AZ of a protection group. Changing this creates a new group.

* `domain_id` - (Required) Specifies the ID of an ``active-active domain``. Changing this creates a new group.
  An ``active-active domain`` id can be extracted from ``data/opentelekomcloud_sdrs_domain_v1`` and shouldn't be confused
  with tenant ``domain``.

* `source_vpc_id` - (Required) Specifies the ID of the source VPC. Changing this creates a new group.

* `dr_type` - (Optional) Specifies the deployment model. The default value is migration indicating migration within a VPC.
  Changing this creates a new group.


## Attributes Reference

The following attributes are exported:

* `id` -  ID of the protection group.

## Import

Protection groups can be imported using the `id`, e.g.

```sh
terraform import opentelekomcloud_sdrs_protectiongroup_v1.group_1 7117d38e-4c8f-4624-a505-bd96b97d024c
```
