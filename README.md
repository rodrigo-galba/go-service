# Go microservice template

## Context

- Go (1.17)
- Gin framework (1.7.7) for web implementation [Github repo](https://github.com/gin-gonic/gin)
- Go Swagger (OpenAPI 2.0) for API specification generation
- Docker configuration to build a container image

## Project structure

- `cmd`: Main applications for this project.
- `internal`: Private application and library code. To avoid others importing in their applications or libraries.
- `pkg`: Library code that's ok to use by external applications.
- `config`: Configuration file templates or default configs.
- `api`: OpenAPI/Swagger specs, JSON schema files, protocol definition files.

## Project Setup

To setup the project locally on your development environment, install the following tools:

- Go (1.17)
- Go Swagger (latest)

After with all in place, run the following from the project's root folder:
```shell
go build cmd/go-service
```

To run the app natively:
```shell
go build cmd/go-service
./go-service.exe
```

## Docker build

To build its image locally, run:
```shell
docker build . -t rodrigo-galba/go-service
```
To run the app from the local image:
```shell
docker run -p 8080:8080 rodrigo-galba/go-service:latest
curl --request GET \
  --url http://localhost:8080/recipes
```

## AWS Elastic beanstalk

To deploy project as a Docker application on Elastic Beanstalk service, use its `eb` CLI.  
Initialize as a Docker project:
```shell
eb init -p docker go-service
```

Create and deploy an environment:
```shell
$ eb create development --profile guru -r us-east-1 # set aws profile (guru) and region
Environment details for: development
  Application name: go-service
  Region: us-east-1
  Deployed Version: app-aea4-220220_133835
  Environment ID: e-g4dfgzqi3m
  Platform: arn:aws:elasticbeanstalk:us-east-1::platform/Docker running on 64bit Amazon Linux 2/3.4.11
  Tier: WebServer-Standard-1.0
  CNAME: UNKNOWN
  Updated: 2022-02-20 16:38:40.710000+00:00
...
2022-02-20 16:41:55    INFO    Instance deployment completed successfully.
2022-02-20 16:42:03    INFO    Application available at development.eba-quq3pqcy.us-east-1.elasticbeanstalk.com.
2022-02-20 16:42:04    INFO    Successfully launched environment: development
```

## Using the Recipes API

Create a new Recipe:

```shell
$ curl --request POST \
  --url http://localhost:8080/recipes \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "hamburguer",
	"tags": ["fastfood"],
	"ingredients": ["pickles", "meat", "bread", "cheese"]	
}'
```

List recipes:
```shell
$ curl --request GET \
  --url http://localhost:8080/recipes
```

Update recipe:

```shell
$ curl --request PUT \
  --url http://localhost:8080/recipes/c0283p3d0cvuglq85log \
  --header 'Content-Type: application/json' \
  --data '{
		"id": "c0283p3d0cvuglq85log",
	  "name": "Oregano Marinated Chicken (updated)",
		"tags": [
			"main",
			"chicken"
		],
		"ingredients": [
			"4 (6 to 7-ounce) boneless skinless chicken breasts\r",
			"10 grinds black pepper\r",
			"1/2 tsp salt\r",
			"2 tablespoon extra-virgin olive oil\r",
			"1 teaspoon dried oregano\r",
			"1 lemon, juiced"
		]
	}'
```

Search recipe by tag:

```shell
$ curl --request GET \
  --url 'http://localhost:8080/recipes/search?tag=chicken'
```

## OpenAPI specification

Go-Swagger is going to generate the API spec based on comments in the sourcecode (OpenAPI 2.0).  
To install Go-Swagger, download a binary for your platform from [github](https://github.com/go-swagger/go-swagger/releases/latest).

```shell
$ swagger version
version: v0.29.0
commit: 09ae1192ca9a941bbb534aca09e6bdc562c95ef3
```

To generate spec:
```shell
swagger generate spec -o ./api/openapi.yaml
```

To run the Swagger UI locally using the spec:
```shell
$ swagger serve .\api\openapi.yaml
2022/02/17 08:00:07 serving docs at http://localhost:53091/docs
```

To add more fields into API's metadata, go to [Swagger metadata list](https://goswagger.io/use/spec/meta.html)


References
- [GO project layout] (https://github.com/golang-standards/project-layout)
- [Go Swagger] (https://goswagger.io/generate/spec.html)
