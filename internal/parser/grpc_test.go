package parser_test

import (
	"context"
	"io"
	"scaper-demo/internal/parser"
	pb "scaper-demo/proto"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParserServer(t *testing.T) {
	parseFunc := func(reader io.Reader) (pb.ParsedPageResponse, error) {
		t.Helper()

		content := make([]byte, 10)
		_, err := reader.Read(content)
		if err != nil {
			return pb.ParsedPageResponse{}, err
		}

		return pb.ParsedPageResponse{
			Name: strings.Trim(string(content), "\x00"),
		}, nil
	}

	server := parser.NewParserServer(parseFunc)

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		req := &pb.RawPageData{
			HtmlContent: "test",
		}

		resp, err := server.ParsePage(context.Background(), req)
		require.NoError(t, err)
		require.Equal(t, "test", string(resp.GetName()))
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		req := &pb.RawPageData{
			HtmlContent: "",
		}

		resp, err := server.ParsePage(context.Background(), req)
		require.Error(t, err)
		require.Nil(t, resp)
	})
}
