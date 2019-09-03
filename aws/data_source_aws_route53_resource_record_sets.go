package aws

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAwsRoute53ResourceRecordSets() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAwsRoute53ResourceRecordSetsRead,

		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_record_sets": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ttl": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_records": {
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed: true,
						},
					},
				},
				Computed: true,
			},
		},
	}
}

func dataSourceAwsRoute53ResourceRecordSetsRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).r53conn

	zoneID := d.Get("zone_id").(string)

	req := &route53.ListResourceRecordSetsInput{
		HostedZoneId: aws.String(zoneID),
	}

	log.Printf("[DEBUG] Reading Route53 resource record sets: %s", req)

	resp, err := conn.ListResourceRecordSets(req)
	if err != nil {
		return fmt.Errorf("Failed getting Route53 resource record sets: %s Zone ID: %q", err, zoneID)
	}

	resourceRecordSets := []map[string]interface{}{}
	for _, respResourceRecordSet := range resp.ResourceRecordSets {
		respResourceRecords := respResourceRecordSet.ResourceRecords

		resourceRecords := make([]string, len(respResourceRecords))

		for i, respResourceRecord := range respResourceRecords {
			resourceRecords[i] = *respResourceRecord.Value
		}

		resourceRecordSets = append(resourceRecordSets, map[string]interface{}{
			"name":             *respResourceRecordSet.Name,
			"ttl":              *respResourceRecordSet.TTL,
			"type":             *respResourceRecordSet.Type,
			"resource_records": resourceRecords,
		})
	}

	d.SetId(zoneID)

	if err := d.Set("resource_record_sets", resourceRecordSets); err != nil {
		return fmt.Errorf("error setting resource_record_sets: %s", err)
	}

	return nil
}
