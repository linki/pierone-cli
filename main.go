package main

import (
	"fmt"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/linki/pierone-cli/client"
	"github.com/linki/pierone-cli/client/docker_registry_api_v2"
)

func main() {
	transport := httptransport.New("registry.opensource.zalan.do", "/", []string{"https"})
	transport.Debug = true

	c := client.New(transport, strfmt.Default)

	r, err := c.DockerRegistryAPIV2.OrgZalandoStupsPieroneAPIV2ListTags(
		docker_registry_api_v2.NewOrgZalandoStupsPieroneAPIV2ListTagsParams().
			WithTeam("teapot").
			WithArtifact("gerry"),
		httptransport.BearerToken("123"),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(*r.Payload.Name)
	for _, t := range r.Payload.Tags {
		fmt.Println("- ", t)
	}
}
