package grafanastack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceStack() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the resource, also acts as it's unique ID",
				ForceNew:    true,
			},
			"accesskey": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the resource, also acts as it's unique ID",
				ForceNew:    true,
			},
			"stackname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the resource, also acts as it's unique ID",
				ForceNew:    true,
			},
			"slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A description of an item",
			},
			"region": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "An optional list of tags, represented as a key, value pair",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
		Create: resourceCreateStack,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

type CreateStackRequest struct {
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Region string `json:"region"`
}

func resourceCreateStack(d *schema.ResourceData, m interface{}) error {
	client := &http.Client{}
	request := CreateStackRequest{Name: d.Get("name").(string), Slug: d.Get("slug").(string), Region: d.Get("region").(string)}
	data, err := json.Marshal(request)

	fmt.Println(string(data))

	req, err := http.NewRequest("POST", d.Get("url").(string), bytes.NewBuffer(data))
	req.Header.Set("Authorization", "Bearer "+d.Get("accessKey").(string))
	req.Header.Add("Content-Type", "application/json")

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(data))
	}
	return nil
}
