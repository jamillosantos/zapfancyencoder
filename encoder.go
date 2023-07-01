package zapfancyencoder

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

func newColor(p ...color.Attribute) *color.Color {
	c := color.Set(p...)
	c.EnableColor()
	return c
}

var (
	colorLabel      = fmt.Sprint
	colorLabelError = newColor(color.FgRed, color.Bold).Sprint
	colorMessage    = newColor(color.FgHiWhite, color.Bold).Sprint
	colorError      = newColor(color.FgRed).Sprint
	colorLabelWarn  = newColor(color.FgYellow, color.Bold).Sprint
	colorWarn       = newColor(color.FgYellow, color.Bold).Sprint
	colorLabelTree  = newColor(color.FgHiBlack).Sprint
)

const (
	labelLevel     = "Level"
	labelTimestamp = "Timestamp"
	labelMessage   = "Message"
	labelFields    = "Fields"

	labelLonger = labelTimestamp
)

type fieldValue struct {
	key    string
	value  string
	fields []fieldValue
}

type FancyEncoder struct {
	buf             *buffer.Buffer
	maxFieldNameLen int
	openNamespaces  int
	fieldValues     []fieldValue
}

func (f *FancyEncoder) AddArray(key string, marshaler zapcore.ArrayMarshaler) error {
	//TODO implement me
	panic("implement me")
}

func (f *FancyEncoder) AddObject(key string, marshaler zapcore.ObjectMarshaler) error {
	fv := fieldValue{
		key: key,
	}
	fvs := f.fieldValues
	f.fieldValues = make([]fieldValue, 0)
	f.openNamespaces++
	err := marshaler.MarshalLogObject(f)
	fv.fields = f.fieldValues
	f.fieldValues = append(fvs, fv)
	f.openNamespaces--
	return err
}

func (f *FancyEncoder) AddBinary(key string, value []byte) {
	f.AddString(key, fmt.Sprintf("%x", value))
}

func (f *FancyEncoder) AddByteString(key string, value []byte) {
	f.AddString(key, string(value))
}

func (f *FancyEncoder) AddBool(key string, value bool) {
	f.AddString(key, fmt.Sprintf("%t", value))
}

func (f *FancyEncoder) AddComplex128(key string, value complex128) {
	f.AddString(key, fmt.Sprintf("%f", value))
}

func (f *FancyEncoder) AddComplex64(key string, value complex64) {
	f.AddString(key, fmt.Sprintf("%f", value))
}

func (f *FancyEncoder) AddDuration(key string, value time.Duration) {
	f.AddString(key, value.String())
}

func (f *FancyEncoder) AddFloat64(key string, value float64) {
	f.AddString(key, fmt.Sprintf("%f", value))
}

func (f *FancyEncoder) AddFloat32(key string, value float32) {
	f.AddString(key, fmt.Sprintf("%f", value))
}

func (f *FancyEncoder) AddInt(key string, value int) {
	f.AddString(key, fmt.Sprintf("%d", value))
}

func (f *FancyEncoder) AddInt64(key string, value int64) {
	f.AddString(key, fmt.Sprintf("%d", value))
}

func (f *FancyEncoder) AddInt32(key string, value int32) {
	f.AddString(key, fmt.Sprintf("%d", value))
}

func (f *FancyEncoder) AddInt16(key string, value int16) {
	f.AddString(key, fmt.Sprintf("%d", value))
}

func (f *FancyEncoder) AddInt8(key string, value int8) {
	f.AddString(key, fmt.Sprintf("%d", value))
}

func (f *FancyEncoder) AddString(key, value string) {
	f.fieldValues = append(f.fieldValues, fieldValue{key: key, value: value})
}

func (f *FancyEncoder) AddTime(key string, value time.Time) {
	f.AddString(key, value.Format(time.RFC3339))
}

func (f *FancyEncoder) AddUint(key string, value uint) {
	f.AddString(key, fmt.Sprintf("%d", value))
}

func (f *FancyEncoder) AddUint64(key string, value uint64) {
	f.AddString(key, fmt.Sprintf("%d", value))
}

func (f *FancyEncoder) AddUint32(key string, value uint32) {
	f.AddString(key, fmt.Sprintf("%d", value))
}

func (f *FancyEncoder) AddUint16(key string, value uint16) {
	f.AddString(key, fmt.Sprintf("%d", value))
}

func (f *FancyEncoder) AddUint8(key string, value uint8) {
	f.AddString(key, fmt.Sprintf("%d", value))
}

func (f *FancyEncoder) AddUintptr(key string, value uintptr) {
	f.AddString(key, fmt.Sprintf("%x", value))
}

func (f *FancyEncoder) AddReflected(key string, value interface{}) error {
	f.AddString(key, fmt.Sprintf("%+v", value))
	return nil
}

func (f *FancyEncoder) OpenNamespace(key string) {
	f.openNamespaces++
}

func (f *FancyEncoder) Clone() zapcore.Encoder {
	buf := _pool.Get()
	_, _ = buf.Write(f.buf.Bytes())
	return &FancyEncoder{
		buf: buf,
	}
}

var (
	_pool = buffer.NewPool()
)

func (f *FancyEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	f.fieldValues = make([]fieldValue, 0)
	colorMessageValue := colorMessage
	if entry.Level == zapcore.WarnLevel {
		colorMessageValue = colorLabelWarn
	} else if entry.Level >= zapcore.ErrorLevel {
		colorMessageValue = colorLabelError
	}
	f.printLabel(colorLabel(labelLevel), entry.Level.CapitalString())
	f.printLabel(colorLabel(labelMessage), colorMessageValue(entry.Message))
	f.printLabel(colorLabel(labelTimestamp), entry.Time.Format("2006-01-02 15:04:05 Z07:00"))

	if len(fields) == 0 {
		f.printHL()
		buf := f.buf
		f.Free()
		return buf, nil
	}

	f.printLabel(colorLabel(labelFields), "")
	f.maxFieldNameLen = 0
	for _, field := range fields {
		if len(field.Key) > f.maxFieldNameLen {
			f.maxFieldNameLen = len(field.Key)
		}
	}
	for _, field := range fields {
		field.AddTo(f)
	}
	f.printFields("", f.fieldValues)
	f.printHL()
	buf := f.buf
	f.Free()
	return buf, nil
}

// Free ...
// TODO Not being cleaned up. Ignoring this for now as this is a PoC.
func (f *FancyEncoder) Free() {
	f.buf = _pool.Get()
}

func (f *FancyEncoder) printLabel(label string, value string) {
	_, _ = fmt.Fprintf(f.buf, fmt.Sprintf("%%%ds: %%s\n", len(labelLonger)+2), label, value)
}

func (f *FancyEncoder) printFields(prefix string, fields []fieldValue) {
	maxLen := 0
	for _, field := range fields {
		if len(field.key) > maxLen {
			maxLen = len(field.key)
		}
	}
	format := fmt.Sprintf("      %%s %%%ds: %%s\n", maxLen)
	for i, field := range fields {
		p := "├─"
		if i == len(fields)-1 {
			p = "└─"
		}
		_, _ = fmt.Fprintf(f.buf, format, colorLabelTree(prefix+p), field.key, field.value)
		if len(field.fields) > 0 {
			p := "|   "
			if i == len(fields)-1 {
				p = "    "
			}
			f.printFields(prefix+p, field.fields)
		}
	}
}

func (f *FancyEncoder) printHL() {
	_, _ = fmt.Fprintf(f.buf, colorLabelTree("------------------------------------\n"))
}

func init() {
	err := zap.RegisterEncoder("fancy", func(cfg zapcore.EncoderConfig) (zapcore.Encoder, error) {
		return &FancyEncoder{
			buf: _pool.Get(),
		}, nil
	})
	if err != nil {
		panic(err)
	}
}
