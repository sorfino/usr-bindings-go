package kvsbinding_test

import (
	"bytes"
	"context"
	"github/mercadolibre/go-bindings/pkg/kvsbinding"
	"github/mercadolibre/go-bindings/pkg/kvsbinding/protocol"
	"testing"

	flatbuffers "github.com/google/flatbuffers/go"
)

func TestNewClient(t *testing.T) {
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
		t.Fatal(response.Error)
	}

	if len(response.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(response.Items))
	}

	if !bytes.Equal([]byte(`{"hello":"world"}`), response.Items[0].Value) {
		t.Fatalf("expected %s, got %s", `{"hello":"world"}`, string(response.Items[0].Value))
	}
}
