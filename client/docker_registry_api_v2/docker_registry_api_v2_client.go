package docker_registry_api_v2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new docker registry api v2 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for docker registry api v2 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
OrgZalandoStupsPieroneAPIV2GetBlob gets layer
*/
func (a *Client) OrgZalandoStupsPieroneAPIV2GetBlob(params *OrgZalandoStupsPieroneAPIV2GetBlobParams, authInfo runtime.ClientAuthInfoWriter) (*OrgZalandoStupsPieroneAPIV2GetBlobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrgZalandoStupsPieroneAPIV2GetBlobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "org.zalando.stups.pierone.api-v2/get-blob",
		Method:             "GET",
		PathPattern:        "/v2/{team}/{artifact}/blobs/{digest}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrgZalandoStupsPieroneAPIV2GetBlobReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OrgZalandoStupsPieroneAPIV2GetBlobOK), nil
}

/*
OrgZalandoStupsPieroneAPIV2GetManifest gets manifest
*/
func (a *Client) OrgZalandoStupsPieroneAPIV2GetManifest(params *OrgZalandoStupsPieroneAPIV2GetManifestParams, authInfo runtime.ClientAuthInfoWriter) (*OrgZalandoStupsPieroneAPIV2GetManifestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrgZalandoStupsPieroneAPIV2GetManifestParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "org.zalando.stups.pierone.api-v2/get-manifest",
		Method:             "GET",
		PathPattern:        "/v2/{team}/{artifact}/manifests/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrgZalandoStupsPieroneAPIV2GetManifestReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OrgZalandoStupsPieroneAPIV2GetManifestOK), nil
}

/*
OrgZalandoStupsPieroneAPIV2HeadBlob checks layer
*/
func (a *Client) OrgZalandoStupsPieroneAPIV2HeadBlob(params *OrgZalandoStupsPieroneAPIV2HeadBlobParams, authInfo runtime.ClientAuthInfoWriter) (*OrgZalandoStupsPieroneAPIV2HeadBlobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrgZalandoStupsPieroneAPIV2HeadBlobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "org.zalando.stups.pierone.api-v2/head-blob",
		Method:             "HEAD",
		PathPattern:        "/v2/{team}/{artifact}/blobs/{digest}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrgZalandoStupsPieroneAPIV2HeadBlobReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OrgZalandoStupsPieroneAPIV2HeadBlobOK), nil
}

/*
OrgZalandoStupsPieroneAPIV2ListRepositories lists browse all repositories
*/
func (a *Client) OrgZalandoStupsPieroneAPIV2ListRepositories(params *OrgZalandoStupsPieroneAPIV2ListRepositoriesParams, authInfo runtime.ClientAuthInfoWriter) (*OrgZalandoStupsPieroneAPIV2ListRepositoriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrgZalandoStupsPieroneAPIV2ListRepositoriesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "org.zalando.stups.pierone.api-v2/list-repositories",
		Method:             "GET",
		PathPattern:        "/v2/_catalog",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrgZalandoStupsPieroneAPIV2ListRepositoriesReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OrgZalandoStupsPieroneAPIV2ListRepositoriesOK), nil
}

/*
OrgZalandoStupsPieroneAPIV2ListTags lists tags
*/
func (a *Client) OrgZalandoStupsPieroneAPIV2ListTags(params *OrgZalandoStupsPieroneAPIV2ListTagsParams, authInfo runtime.ClientAuthInfoWriter) (*OrgZalandoStupsPieroneAPIV2ListTagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrgZalandoStupsPieroneAPIV2ListTagsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "org.zalando.stups.pierone.api-v2/list-tags",
		Method:             "GET",
		PathPattern:        "/v2/{team}/{artifact}/tags/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrgZalandoStupsPieroneAPIV2ListTagsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OrgZalandoStupsPieroneAPIV2ListTagsOK), nil
}

/*
OrgZalandoStupsPieroneAPIV2PatchUpload uploads f s layer blob
*/
func (a *Client) OrgZalandoStupsPieroneAPIV2PatchUpload(params *OrgZalandoStupsPieroneAPIV2PatchUploadParams, authInfo runtime.ClientAuthInfoWriter) (*OrgZalandoStupsPieroneAPIV2PatchUploadAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrgZalandoStupsPieroneAPIV2PatchUploadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "org.zalando.stups.pierone.api-v2/patch-upload",
		Method:             "PATCH",
		PathPattern:        "/v2/{team}/{artifact}/blobs/uploads/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/octet-stream"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrgZalandoStupsPieroneAPIV2PatchUploadReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OrgZalandoStupsPieroneAPIV2PatchUploadAccepted), nil
}

/*
OrgZalandoStupsPieroneAPIV2Ping checks compatibility

Checks for compatibility with Docker Registry v2.
*/
func (a *Client) OrgZalandoStupsPieroneAPIV2Ping(params *OrgZalandoStupsPieroneAPIV2PingParams) (*OrgZalandoStupsPieroneAPIV2PingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrgZalandoStupsPieroneAPIV2PingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "org.zalando.stups.pierone.api-v2/ping",
		Method:             "GET",
		PathPattern:        "/v2/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrgZalandoStupsPieroneAPIV2PingReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*OrgZalandoStupsPieroneAPIV2PingOK), nil
}

/*
OrgZalandoStupsPieroneAPIV2PostUpload starts upload
*/
func (a *Client) OrgZalandoStupsPieroneAPIV2PostUpload(params *OrgZalandoStupsPieroneAPIV2PostUploadParams, authInfo runtime.ClientAuthInfoWriter) (*OrgZalandoStupsPieroneAPIV2PostUploadAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrgZalandoStupsPieroneAPIV2PostUploadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "org.zalando.stups.pierone.api-v2/post-upload",
		Method:             "POST",
		PathPattern:        "/v2/{team}/{artifact}/blobs/uploads/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrgZalandoStupsPieroneAPIV2PostUploadReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OrgZalandoStupsPieroneAPIV2PostUploadAccepted), nil
}

/*
OrgZalandoStupsPieroneAPIV2PutManifest puts manifest
*/
func (a *Client) OrgZalandoStupsPieroneAPIV2PutManifest(params *OrgZalandoStupsPieroneAPIV2PutManifestParams, authInfo runtime.ClientAuthInfoWriter) (*OrgZalandoStupsPieroneAPIV2PutManifestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrgZalandoStupsPieroneAPIV2PutManifestParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "org.zalando.stups.pierone.api-v2/put-manifest",
		Method:             "PUT",
		PathPattern:        "/v2/{team}/{artifact}/manifests/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/octet-stream", "application/vnd.docker.distribution.manifest.v1+prettyjws", "application/vnd.docker.distribution.manifest.v2+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrgZalandoStupsPieroneAPIV2PutManifestReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OrgZalandoStupsPieroneAPIV2PutManifestOK), nil
}

/*
OrgZalandoStupsPieroneAPIV2PutUpload uploads complete
*/
func (a *Client) OrgZalandoStupsPieroneAPIV2PutUpload(params *OrgZalandoStupsPieroneAPIV2PutUploadParams, authInfo runtime.ClientAuthInfoWriter) (*OrgZalandoStupsPieroneAPIV2PutUploadCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrgZalandoStupsPieroneAPIV2PutUploadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "org.zalando.stups.pierone.api-v2/put-upload",
		Method:             "PUT",
		PathPattern:        "/v2/{team}/{artifact}/blobs/uploads/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/octet-stream"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrgZalandoStupsPieroneAPIV2PutUploadReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OrgZalandoStupsPieroneAPIV2PutUploadCreated), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}