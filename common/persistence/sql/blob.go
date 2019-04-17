package sql

import (
	"bytes"

	"github.com/uber/cadence/.gen/go/sqlblobs"
	"go.uber.org/thriftrw/protocol"
	"go.uber.org/thriftrw/wire"
)

func thriftEncode(wireBytes wire.Value) ([]byte, error) {
	var b bytes.Buffer
	err := protocol.Binary.Encode(wireBytes, &b)
	return b.Bytes(), err
}

func thriftDecode(b []byte) (wire.Value, error) {
	wireBytes, err := protocol.Binary.Decode(bytes.NewReader(b), wire.TStruct)
	if err != nil {
		return wire.Value{}, err
	}
	return wireBytes, err
}

func serializeWorkflowExecutionInfo(info *sqlblobs.WorkflowExecutionInfo) ([]byte, error) {
	value, err := info.ToWire()
	if err != nil {
		return nil, err
	}
	return thriftEncode(value)
}

func deserializeWorkflowExecutionInfo(b []byte, result *sqlblobs.WorkflowExecutionInfo) error {
	value, err := thriftDecode(b)
	if err != nil {
		return err
	}
	return result.FromWire(value)
}
