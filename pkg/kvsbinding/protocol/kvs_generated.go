// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package protocol

import (
	flatbuffers "github.com/google/flatbuffers/go"
	"strconv"
)

type Method int8

const (
	MethodGET    Method = 0
	MethodSET    Method = 1
	MethodDELETE Method = 2
)

var EnumNamesMethod = map[Method]string{
	MethodGET:    "GET",
	MethodSET:    "SET",
	MethodDELETE: "DELETE",
}

var EnumValuesMethod = map[string]Method{
	"GET":    MethodGET,
	"SET":    MethodSET,
	"DELETE": MethodDELETE,
}

func (v Method) String() string {
	if s, ok := EnumNamesMethod[v]; ok {
		return s
	}
	return "Method(" + strconv.FormatInt(int64(v), 10) + ")"
}

type ErrorCode int8

const (
	ErrorCodeOK              ErrorCode = 0
	ErrorCodeNOT_FOUND       ErrorCode = 1
	ErrorCodeSTALE_VALUE     ErrorCode = 2
	ErrorCodeOVER_QUOTA      ErrorCode = 3
	ErrorCodeTRANSPORT_ERROR ErrorCode = 10
	ErrorCodeTIMEOUT         ErrorCode = 11
	ErrorCodeSERVICE_ERROR   ErrorCode = 100
	ErrorCodeOTHER           ErrorCode = 101
)

var EnumNamesErrorCode = map[ErrorCode]string{
	ErrorCodeOK:              "OK",
	ErrorCodeNOT_FOUND:       "NOT_FOUND",
	ErrorCodeSTALE_VALUE:     "STALE_VALUE",
	ErrorCodeOVER_QUOTA:      "OVER_QUOTA",
	ErrorCodeTRANSPORT_ERROR: "TRANSPORT_ERROR",
	ErrorCodeTIMEOUT:         "TIMEOUT",
	ErrorCodeSERVICE_ERROR:   "SERVICE_ERROR",
	ErrorCodeOTHER:           "OTHER",
}

var EnumValuesErrorCode = map[string]ErrorCode{
	"OK":              ErrorCodeOK,
	"NOT_FOUND":       ErrorCodeNOT_FOUND,
	"STALE_VALUE":     ErrorCodeSTALE_VALUE,
	"OVER_QUOTA":      ErrorCodeOVER_QUOTA,
	"TRANSPORT_ERROR": ErrorCodeTRANSPORT_ERROR,
	"TIMEOUT":         ErrorCodeTIMEOUT,
	"SERVICE_ERROR":   ErrorCodeSERVICE_ERROR,
	"OTHER":           ErrorCodeOTHER,
}

func (v ErrorCode) String() string {
	if s, ok := EnumNamesErrorCode[v]; ok {
		return s
	}
	return "ErrorCode(" + strconv.FormatInt(int64(v), 10) + ")"
}

type ItemT struct {
	Key string `json:"key"`
	Value []byte `json:"value"`
	Version int32 `json:"version"`
	Timestamp uint64 `json:"timestamp"`
}

func (t *ItemT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	keyOffset := flatbuffers.UOffsetT(0)
	if t.Key != "" {
		keyOffset = builder.CreateString(t.Key)
	}
	valueOffset := flatbuffers.UOffsetT(0)
	if t.Value != nil {
		valueOffset = builder.CreateByteString(t.Value)
	}
	ItemStart(builder)
	ItemAddKey(builder, keyOffset)
	ItemAddValue(builder, valueOffset)
	ItemAddVersion(builder, t.Version)
	ItemAddTimestamp(builder, t.Timestamp)
	return ItemEnd(builder)
}

func (rcv *Item) UnPackTo(t *ItemT) {
	t.Key = string(rcv.Key())
	t.Value = rcv.ValueBytes()
	t.Version = rcv.Version()
	t.Timestamp = rcv.Timestamp()
}

func (rcv *Item) UnPack() *ItemT {
	if rcv == nil {
		return nil
	}
	t := &ItemT{}
	rcv.UnPackTo(t)
	return t
}

type Item struct {
	_tab flatbuffers.Table
}

func GetRootAsItem(buf []byte, offset flatbuffers.UOffsetT) *Item {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Item{}
	x.Init(buf, n+offset)
	return x
}

func FinishItemBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsItem(buf []byte, offset flatbuffers.UOffsetT) *Item {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Item{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedItemBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *Item) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Item) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Item) Key() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Item) Value(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *Item) ValueLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Item) ValueBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Item) MutateValue(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *Item) Version() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Item) MutateVersion(n int32) bool {
	return rcv._tab.MutateInt32Slot(8, n)
}

func (rcv *Item) Timestamp() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Item) MutateTimestamp(n uint64) bool {
	return rcv._tab.MutateUint64Slot(10, n)
}

func ItemStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func ItemAddKey(builder *flatbuffers.Builder, key flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(key), 0)
}
func ItemAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(value), 0)
}
func ItemStartValueVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func ItemAddVersion(builder *flatbuffers.Builder, version int32) {
	builder.PrependInt32Slot(2, version, 0)
}
func ItemAddTimestamp(builder *flatbuffers.Builder, timestamp uint64) {
	builder.PrependUint64Slot(3, timestamp, 0)
}
func ItemEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type ErrorT struct {
	Code ErrorCode `json:"code"`
	Message string `json:"message"`
}

func (t *ErrorT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	messageOffset := flatbuffers.UOffsetT(0)
	if t.Message != "" {
		messageOffset = builder.CreateString(t.Message)
	}
	ErrorStart(builder)
	ErrorAddCode(builder, t.Code)
	ErrorAddMessage(builder, messageOffset)
	return ErrorEnd(builder)
}

func (rcv *Error) UnPackTo(t *ErrorT) {
	t.Code = rcv.Code()
	t.Message = string(rcv.Message())
}

func (rcv *Error) UnPack() *ErrorT {
	if rcv == nil {
		return nil
	}
	t := &ErrorT{}
	rcv.UnPackTo(t)
	return t
}

type Error struct {
	_tab flatbuffers.Table
}

func GetRootAsError(buf []byte, offset flatbuffers.UOffsetT) *Error {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Error{}
	x.Init(buf, n+offset)
	return x
}

func FinishErrorBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsError(buf []byte, offset flatbuffers.UOffsetT) *Error {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Error{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedErrorBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *Error) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Error) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Error) Code() ErrorCode {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return ErrorCode(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Error) MutateCode(n ErrorCode) bool {
	return rcv._tab.MutateInt8Slot(4, int8(n))
}

func (rcv *Error) Message() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func ErrorStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func ErrorAddCode(builder *flatbuffers.Builder, code ErrorCode) {
	builder.PrependInt8Slot(0, int8(code), 0)
}
func ErrorAddMessage(builder *flatbuffers.Builder, message flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(message), 0)
}
func ErrorEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type RequestT struct {
	Keys []*ItemT `json:"keys"`
}

func (t *RequestT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	keysOffset := flatbuffers.UOffsetT(0)
	if t.Keys != nil {
		keysLength := len(t.Keys)
		keysOffsets := make([]flatbuffers.UOffsetT, keysLength)
		for j := 0; j < keysLength; j++ {
			keysOffsets[j] = t.Keys[j].Pack(builder)
		}
		RequestStartKeysVector(builder, keysLength)
		for j := keysLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(keysOffsets[j])
		}
		keysOffset = builder.EndVector(keysLength)
	}
	RequestStart(builder)
	RequestAddKeys(builder, keysOffset)
	return RequestEnd(builder)
}

func (rcv *Request) UnPackTo(t *RequestT) {
	keysLength := rcv.KeysLength()
	t.Keys = make([]*ItemT, keysLength)
	for j := 0; j < keysLength; j++ {
		x := Item{}
		rcv.Keys(&x, j)
		t.Keys[j] = x.UnPack()
	}
}

func (rcv *Request) UnPack() *RequestT {
	if rcv == nil {
		return nil
	}
	t := &RequestT{}
	rcv.UnPackTo(t)
	return t
}

type Request struct {
	_tab flatbuffers.Table
}

func GetRootAsRequest(buf []byte, offset flatbuffers.UOffsetT) *Request {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Request{}
	x.Init(buf, n+offset)
	return x
}

func FinishRequestBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsRequest(buf []byte, offset flatbuffers.UOffsetT) *Request {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Request{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedRequestBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *Request) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Request) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Request) Keys(obj *Item, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Request) KeysLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func RequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func RequestAddKeys(builder *flatbuffers.Builder, keys flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(keys), 0)
}
func RequestStartKeysVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func RequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type ResponseT struct {
	Items []*ItemT `json:"items"`
	Error *ErrorT `json:"error"`
}

func (t *ResponseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	itemsOffset := flatbuffers.UOffsetT(0)
	if t.Items != nil {
		itemsLength := len(t.Items)
		itemsOffsets := make([]flatbuffers.UOffsetT, itemsLength)
		for j := 0; j < itemsLength; j++ {
			itemsOffsets[j] = t.Items[j].Pack(builder)
		}
		ResponseStartItemsVector(builder, itemsLength)
		for j := itemsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(itemsOffsets[j])
		}
		itemsOffset = builder.EndVector(itemsLength)
	}
	errorOffset := t.Error.Pack(builder)
	ResponseStart(builder)
	ResponseAddItems(builder, itemsOffset)
	ResponseAddError(builder, errorOffset)
	return ResponseEnd(builder)
}

func (rcv *Response) UnPackTo(t *ResponseT) {
	itemsLength := rcv.ItemsLength()
	t.Items = make([]*ItemT, itemsLength)
	for j := 0; j < itemsLength; j++ {
		x := Item{}
		rcv.Items(&x, j)
		t.Items[j] = x.UnPack()
	}
	t.Error = rcv.Error(nil).UnPack()
}

func (rcv *Response) UnPack() *ResponseT {
	if rcv == nil {
		return nil
	}
	t := &ResponseT{}
	rcv.UnPackTo(t)
	return t
}

type Response struct {
	_tab flatbuffers.Table
}

func GetRootAsResponse(buf []byte, offset flatbuffers.UOffsetT) *Response {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Response{}
	x.Init(buf, n+offset)
	return x
}

func FinishResponseBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsResponse(buf []byte, offset flatbuffers.UOffsetT) *Response {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Response{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedResponseBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *Response) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Response) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Response) Items(obj *Item, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Response) ItemsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Response) Error(obj *Error) *Error {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Error)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func ResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func ResponseAddItems(builder *flatbuffers.Builder, items flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(items), 0)
}
func ResponseStartItemsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ResponseAddError(builder *flatbuffers.Builder, error flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(error), 0)
}
func ResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}