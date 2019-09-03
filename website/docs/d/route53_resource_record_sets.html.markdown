---
layout: "aws"
page_title: "AWS: aws_route53_resource_record_sets"
sidebar_current: "docs-aws-datasource-route53-resource-record-sets"
description: |-
    Lists resource record sets of a Route53 zone
---

# Data Source: aws_route53_resource_record_sets

`aws_route53_resource_record_sets` lists the resource record sets of a Route53 zone.

## Example Usage

The following example shows how to get the resource record sets of a zone from its id.

```hcl
data "aws_route53_resource_record_sets" "resource_record_sets" {
  zone_id = "MQWGHCBFAKEID"
}
```

## Argument Reference

* `zone_id` - (Required) The Hosted Zone ID which to retrieve the resource record sets from.

The following attribute is additionally exported:

* `resource_record_sets` - The list of resource record set from the zone.
  * `resource_record_sets.#.name` - The name of the resource record set.
  * `resource_record_sets.#.ttl` - The TTL of the resource record set.
  * `resource_record_sets.#.type` - The type of the resource record set.
  * `resource_record_sets.#.resource_records` - The resource records of the resource record set.
