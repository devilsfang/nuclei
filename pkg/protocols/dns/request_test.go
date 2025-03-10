package dns

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/devilsfang/nuclei/v3/pkg/model"
	"github.com/devilsfang/nuclei/v3/pkg/model/types/severity"
	"github.com/devilsfang/nuclei/v3/pkg/operators"
	"github.com/devilsfang/nuclei/v3/pkg/operators/extractors"
	"github.com/devilsfang/nuclei/v3/pkg/operators/matchers"
	"github.com/devilsfang/nuclei/v3/pkg/output"
	"github.com/devilsfang/nuclei/v3/pkg/protocols/common/contextargs"
	"github.com/devilsfang/nuclei/v3/pkg/testutils"
)

func TestDNSExecuteWithResults(t *testing.T) {
	options := testutils.DefaultOptions

	recursion := false
	testutils.Init(options)
	templateID := "testing-dns"
	executerOpts := testutils.NewMockExecuterOptions(options, &testutils.TemplateInfo{
		ID:   templateID,
		Info: model.Info{SeverityHolder: severity.Holder{Severity: severity.Low}, Name: "test"},
	})
	request := &Request{
		RequestType: DNSRequestTypeHolder{DNSRequestType: A},
		Class:       "INET",
		Retries:     5,
		ID:          templateID,
		Recursion:   &recursion,
		Name:        "{{FQDN}}",
		Operators: operators.Operators{
			Matchers: []*matchers.Matcher{{
				Name:  "test",
				Part:  "raw",
				Type:  matchers.MatcherTypeHolder{MatcherType: matchers.WordsMatcher},
				Words: []string{"93.184.215.14"},
			}},
			Extractors: []*extractors.Extractor{{
				Part:  "raw",
				Type:  extractors.ExtractorTypeHolder{ExtractorType: extractors.RegexExtractor},
				Regex: []string{"[0-9]+\\.[0-9]+\\.[0-9]+\\.[0-9]+"},
			}},
		},
		options: executerOpts,
	}
	err := request.Compile(executerOpts)
	require.Nil(t, err, "could not compile dns request")

	var finalEvent *output.InternalWrappedEvent
	t.Run("domain-valid", func(t *testing.T) {
		metadata := make(output.InternalEvent)
		previous := make(output.InternalEvent)
		ctxArgs := contextargs.NewWithInput(context.Background(), "example.com")
		err := request.ExecuteWithResults(ctxArgs, metadata, previous, func(event *output.InternalWrappedEvent) {
			finalEvent = event
		})
		require.Nil(t, err, "could not execute dns request")
	})
	require.NotNil(t, finalEvent, "could not get event output from request")
	require.Equal(t, 1, len(finalEvent.Results), "could not get correct number of results")
	require.Equal(t, "test", finalEvent.Results[0].MatcherName, "could not get correct matcher name of results")
	require.Equal(t, 1, len(finalEvent.Results[0].ExtractedResults), "could not get correct number of extracted results")
	require.Equal(t, "93.184.215.14", finalEvent.Results[0].ExtractedResults[0], "could not get correct extracted results")
	finalEvent = nil
	// Note: changing url to domain is responsible at tmplexec package and is implemented there
}
