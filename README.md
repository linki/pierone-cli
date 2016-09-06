# pierone-cli

A test app to show case how a client library generated by go-swagger looks.

## About

The only custom thing here is `main.go`, your business login. `pierone-api.yaml` is the unchanged swagger spec of pierone as of commit `58186e7`. The folders `models` and `client` were generated by running `swagger generate client -f pierone-api.yaml` and are unchanged. It supports authentication via OAuth2 out of the box.

The generated method names are cumbersome but can be greatly improved by tagging the actions in the swagger spec. This could result in function calls as clean as

```
client.Tags.ListTags(
  tags.NewListTagsParams().WithTeam("teapot").WithArtifact("gerry")
)
```

Since all api interactions are generated and type-safe there's a high degree of confidence that the client conforms to the spec.

## Usage

Just run `main.go` to fetch all tags of `registry.opensource.zalan.do/teapot/gerry`.

```
$ go run main.go
GET /v2/teapot/gerry/tags/list HTTP/1.1
Host: registry.opensource.zalan.do
User-Agent: Go-http-client/1.1
Accept: application/json
Authorization: Bearer 123
Accept-Encoding: gzip


HTTP/1.1 200 OK
Transfer-Encoding: chunked
Access-Control-Allow-Headers:
Access-Control-Allow-Methods: GET, POST, DELETE, PUT, PATCH, OPTIONS
Access-Control-Allow-Origin: *
Access-Control-Max-Age: 3600
Connection: keep-alive
Content-Type: application/json
Date: Tue, 06 Sep 2016 18:51:04 GMT
Server: Jetty(9.2.z-SNAPSHOT)
Strict-Transport-Security: max-age=10886400

114
{"name":"teapot/gerry","tags":["v0.0.0-SNAPSHOT","v0.0.1","v0.0.2","v0.0.3","v0.0.3-17-gf7d2c39","v0.0.3-SNAPSHOT","v0.0.4","v0.0.4-10-g588f1f4","v0.0.4-2-g9c63d2c","v0.0.4-2-gf7d2c39","v0.0.4-4-g0a8d339","v0.0.5","v0.0.5-11-g9b60177","v0.0.5-3-gbf8938c","v0.0.5-7-g510aaee"]}
0


teapot/gerry
-  v0.0.0-SNAPSHOT
-  v0.0.1
-  v0.0.2
-  v0.0.3
-  v0.0.3-17-gf7d2c39
-  v0.0.3-SNAPSHOT
-  v0.0.4
-  v0.0.4-10-g588f1f4
-  v0.0.4-2-g9c63d2c
-  v0.0.4-2-gf7d2c39
-  v0.0.4-4-g0a8d339
-  v0.0.5
-  v0.0.5-11-g9b60177
-  v0.0.5-3-gbf8938c
-  v0.0.5-7-g510aaee
```
