package opac

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"sync"
)

type Result map[string]any

type DataResult struct {
	Result     Result `json:"result"`
	DecisionId string `json:"decision_id"`
}

type DataInput struct {
	Input any `json:"input"`
}

type DataApi struct {
	baseURL string
	client  *http.Client
}

func NewOPADataClient(baseURL string) *DataApi {
	return &DataApi{
		baseURL: baseURL,
		client:  http.DefaultClient,
	}
}

func (c *DataApi) WithInput(ctx context.Context, path string, input DataInput) (result DataResult, err error) {
	b := getBuf()
	defer putBuf(b)

	var (
		req  *http.Request
		resp *http.Response
	)

	enc := json.NewEncoder(b)
	if err = enc.Encode(input); err != nil {
		return
	}

	if req, err = http.NewRequestWithContext(ctx, http.MethodPost, getUrl(c.baseURL, path), b); err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json")

	if resp, err = c.client.Do(req); err != nil {
		return
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	if err = dec.Decode(&result); err != nil {
		return
	}

	return
}

func getUrl(base, path string) string {
	return base + "/v1/data/" + path
}

var xpool = sync.Pool{
	New: func() any {
		b := &bytes.Buffer{}
		b.Grow(256)
		return b
	},
}

func getBuf() *bytes.Buffer {
	b := xpool.Get().(*bytes.Buffer)
	b.Reset()
	return b
}

func putBuf(b *bytes.Buffer) {
	xpool.Put(b)
}
