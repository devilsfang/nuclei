package openapi

import (
	"github.com/devilsfang/nuclei/v3/pkg/input/formats"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/pkg/errors"
)

// OpenAPIFormat is a OpenAPI Schema File parser
type OpenAPIFormat struct {
	opts formats.InputFormatOptions
}

// New creates a new OpenAPI format parser
func New() *OpenAPIFormat {
	return &OpenAPIFormat{}
}

var _ formats.Format = &OpenAPIFormat{}

// Name returns the name of the format
func (j *OpenAPIFormat) Name() string {
	return "openapi"
}

func (j *OpenAPIFormat) SetOptions(options formats.InputFormatOptions) {
	j.opts = options
}

// Parse parses the input and calls the provided callback
// function for each RawRequest it discovers.
func (j *OpenAPIFormat) Parse(input string, resultsCb formats.ParseReqRespCallback) error {
	loader := openapi3.NewLoader()
	schema, err := loader.LoadFromFile(input)
	if err != nil {
		return errors.Wrap(err, "could not decode openapi 3.0 schema")
	}
	return GenerateRequestsFromSchema(schema, j.opts, resultsCb)
}
