package zapfancyencoder

import (
	"fmt"
	"time"
	"unicode"

	"go.uber.org/zap/zapcore"
)

type fieldList struct {
	idx         int
	fieldValues []fieldValue
}

func (f *fieldList) AddArray(key string, marshaler zapcore.ArrayMarshaler) error {
	fl := fieldList{
		fieldValues: make([]fieldValue, 0),
	}
	err := marshaler.MarshalLogArray(&fl)
	f.fieldValues = append(f.fieldValues, fieldValue{
		key:    key,
		fields: fl.fieldValues,
	})
	return err
}

func (f *fieldList) AddObject(key string, marshaler zapcore.ObjectMarshaler) error {
	fl := fieldList{
		fieldValues: make([]fieldValue, 0),
	}
	err := marshaler.MarshalLogObject(&fl)
	f.fieldValues = append(f.fieldValues, fieldValue{
		key:    key,
		fields: fl.fieldValues,
	})
	return err
}

func isPrintable(s []rune) bool {
	for i := 0; i < len(s); i++ {
		if !unicode.IsPrint(s[i]) {
			return false
		}
	}
	return true
}

func (f *fieldList) AddBinary(key string, value []byte) {
	if isPrintable([]rune(string(value))) {
		f.AddString(key, string(value))
		return
	}
	f.AddString(key, fmt.Sprintf("%x", value))
}

func (f *fieldList) AddByteString(key string, value []byte) {
	f.AddString(key, string(value))
}

func (f *fieldList) AddBool(key string, value bool) {
	f.AddString(key, fmt.Sprintf("%t", value))
}

func (f *fieldList) AddComplex128(key string, value complex128) {
	f.AddString(key, fmt.Sprintf("%f", value))
}

func (f *fieldList) AddComplex64(key string, value complex64) {
	f.AddString(key, fmt.Sprintf("%f", value))
}

func (f *fieldList) AddDuration(key string, value time.Duration) {
	f.AddString(key, value.String())
}

func (f *fieldList) AddFloat64(key string, value float64) {
	f.AddString(key, fmt.Sprintf("%f", value))
}

func (f *fieldList) AddFloat32(key string, value float32) {
	f.AddString(key, fmt.Sprintf("%f", value))
}

func (f *fieldList) AddInt(key string, value int) {
	f.AddString(key, fmt.Sprintf("%d", value))
}

func (f *fieldList) AddInt64(key string, value int64) {
	f.AddString(key, fmt.Sprintf("%d", value))
}

func (f *fieldList) AddInt32(key string, value int32) {
	f.AddString(key, fmt.Sprintf("%d", value))
}

func (f *fieldList) AddInt16(key string, value int16) {
	f.AddString(key, fmt.Sprintf("%d", value))
}

func (f *fieldList) AddInt8(key string, value int8) {
	f.AddString(key, fmt.Sprintf("%d", value))
}

func (f *fieldList) AddString(key, value string) {
	f.fieldValues = append(f.fieldValues, fieldValue{key: key, value: value})
}

func (f *fieldList) AddTime(key string, value time.Time) {
	f.AddString(key, value.Format(time.RFC3339))
}

func (f *fieldList) AddUint(key string, value uint) {
	f.AddString(key, fmt.Sprintf("%d", value))
}

func (f *fieldList) AddUint64(key string, value uint64) {
	f.AddString(key, fmt.Sprintf("%d", value))
}

func (f *fieldList) AddUint32(key string, value uint32) {
	f.AddString(key, fmt.Sprintf("%d", value))
}

func (f *fieldList) AddUint16(key string, value uint16) {
	f.AddString(key, fmt.Sprintf("%d", value))
}

func (f *fieldList) AddUint8(key string, value uint8) {
	f.AddString(key, fmt.Sprintf("%d", value))
}

func (f *fieldList) AddUintptr(key string, value uintptr) {
	f.AddString(key, fmt.Sprintf("%x", value))
}

func (f *fieldList) AddReflected(key string, value interface{}) error {
	f.AddString(key, fmt.Sprintf("%+v", value))
	return nil
}

func (f *fieldList) AppendBool(b bool) {
	f.AppendString(fmt.Sprintf("%t", b))
}

func (f *fieldList) AppendByteString(bytes []byte) {
	f.AppendString(string(bytes))
}

func (f *fieldList) AppendComplex128(c complex128) {
	f.AppendString(fmt.Sprintf("%f", c))
}

func (f *fieldList) AppendComplex64(c complex64) {
	f.AppendString(fmt.Sprintf("%f", c))
}

func (f *fieldList) AppendFloat64(f2 float64) {
	f.AppendString(fmt.Sprintf("%f", f2))
}

func (f *fieldList) AppendFloat32(f2 float32) {
	f.AppendString(fmt.Sprintf("%f", f2))
}

func (f *fieldList) AppendInt(i int) {
	f.AppendString(fmt.Sprintf("%d", i))
}

func (f *fieldList) AppendInt64(i int64) {
	f.AppendString(fmt.Sprintf("%d", i))
}

func (f *fieldList) AppendInt32(i int32) {
	f.AppendString(fmt.Sprintf("%d", i))
}

func (f *fieldList) AppendInt16(i int16) {
	f.AppendString(fmt.Sprintf("%d", i))
}

func (f *fieldList) AppendInt8(i int8) {
	f.AppendString(fmt.Sprintf("%d", i))
}

func (f *fieldList) AppendString(s string) {
	f.fieldValues = append(f.fieldValues, fieldValue{
		key:   fmt.Sprintf("[%d]", f.idx),
		value: s,
	})
	f.idx++
}

func (f *fieldList) AppendUint(u uint) {
	f.AppendString(fmt.Sprintf("%d", u))
}

func (f *fieldList) AppendUint64(u uint64) {
	f.AppendString(fmt.Sprintf("%d", u))
}

func (f *fieldList) AppendUint32(u uint32) {
	f.AppendString(fmt.Sprintf("%d", u))
}

func (f *fieldList) AppendUint16(u uint16) {
	f.AppendString(fmt.Sprintf("%d", u))
}

func (f *fieldList) AppendUint8(u uint8) {
	f.AppendString(fmt.Sprintf("%d", u))
}

func (f *fieldList) AppendUintptr(u uintptr) {
	f.AppendString(fmt.Sprintf("%x", u))
}

func (f *fieldList) AppendDuration(duration time.Duration) {
	f.AppendString(duration.String())
}

func (f *fieldList) AppendTime(t time.Time) {
	f.AppendString(t.Format(time.RFC3339))
}

func (f *fieldList) AppendArray(marshaler zapcore.ArrayMarshaler) error {
	fv := fieldValue{
		key: fmt.Sprintf("[%d]", f.idx),
	}
	f.idx++
	fl := fieldList{
		fieldValues: make([]fieldValue, 0),
	}
	err := marshaler.MarshalLogArray(&fl)
	fv.fields = fl.fieldValues
	f.fieldValues = append(f.fieldValues, fv.fields...)
	return err
}

func (f *fieldList) AppendObject(marshaler zapcore.ObjectMarshaler) error {
	fv := fieldValue{
		key: fmt.Sprintf("[%d]", f.idx),
	}
	f.idx++
	fl := fieldList{
		fieldValues: make([]fieldValue, 0),
	}
	err := marshaler.MarshalLogObject(f)
	fv.fields = fl.fieldValues
	f.fieldValues = append(f.fieldValues, fv.fields...)
	return err
}

func (f *fieldList) AppendReflected(value interface{}) error {
	f.AppendString(fmt.Sprintf("%+v", value))
	return nil
}

func (f *fieldList) OpenNamespace(_ string) {
	// ???
}
