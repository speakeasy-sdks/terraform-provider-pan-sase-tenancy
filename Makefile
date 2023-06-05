.PHONY: *
all: speakeasy

speakeasy:
	speakeasy generate sdk --lang terraform -o . -s openapi.yaml

