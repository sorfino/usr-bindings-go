package kvsbinding_test

import (
	"bytes"
	"context"
	"fmt"
	"github/mercadolibre/go-bindings/pkg/kvsbinding"
	"github/mercadolibre/go-bindings/pkg/kvsbinding/protocol"
	"net/http"
	"net/http/httptest"
	"testing"

	flatbuffers "github.com/google/flatbuffers/go"
)

const (
	container = "test"
	key       = "a_key"
	value     = `{"greetings":"ciao ragazzi!!"}`
)

func TestNewClient(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle(fmt.Sprintf("/%s/%s", container, key), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(value))
	}))

	server := httptest.NewServer(mux)
	defer server.Close()

	t.Setenv("KEY_VALUE_STORE_TEST_END_POINT_READ", server.URL)
	client, err := kvsbinding.NewClient("test")
	if err != nil {
		t.Fatal(err)
	}

	defer client.Close()
	fb := flatbuffers.NewBuilder(0)
	r := protocol.RequestT{
		Keys: []*protocol.ItemT{{Key: "a_key"}},
	}
	fb.Finish(r.Pack(fb))

	root, err := client.Call(context.Background(), uint32(protocol.MethodGET), fb.FinishedBytes())
	if err != nil {
		t.Fatal(err)
	}

	res := protocol.GetRootAsResponse(root, 0)
	response := res.UnPack()

	if response.Error != nil {
		t.Fatal(response.Error.Message)
	}

	if len(response.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(response.Items))
	}

	if !bytes.Equal([]byte(value), response.Items[0].Value) {
		t.Fatalf("expected %s, got %s", value, string(response.Items[0].Value))
	}
}
