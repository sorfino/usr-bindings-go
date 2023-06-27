package kvsbinding_test

import (
	"bytes"
	"context"
	"github/mercadolibre/go-bindings/pkg/kvsbinding"
	"github/mercadolibre/go-bindings/pkg/kvsbinding/internal/kvsprotocol"
	"testing"

	flatbuffers "github.com/google/flatbuffers/go"
)

//go:generate flatc --go --gen-onefile --go-namespace kvsprotocol -o internal/kvsprotocol --gen-object-api ../../ubnified-sdk-runtime/flatbuffers/kvs.fbs
func TestNewClient(t *testing.T) {
	client, err := kvsbinding.NewClient("test")
	if err != nil {
		t.Fatal(err)
	}

	defer client.Close()
	fb := flatbuffers.NewBuilder(0)
	r := kvsprotocol.RequestT{
		Keys: []*kvsprotocol.ItemT{{Key: "a_key"}},
	}
	fb.Finish(r.Pack(fb))

	root, err := client.Call(context.Background(), uint32(kvsprotocol.MethodGET), fb.FinishedBytes())
	if err != nil {
		t.Fatal(err)
	}

	res := kvsprotocol.GetRootAsResponse(root, 0)
	response := res.UnPack()

	if response.Error != nil {
		t.Fatal(response.Error)
	}

	if len(response.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(response.Items))
	}

	if !bytes.Equal([]byte(`{"hello":"world"}`), response.Items[0].Value) {
		t.Fatalf("expected %s, got %s", `{"hello":"world"}`, string(response.Items[0].Value))
	}
}
