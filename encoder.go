package zapfancyencoder

import (
	"fmt"

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
	fieldList

	buf             *buffer.Buffer
	maxFieldNameLen int
	openNamespaces  int
	idx             int
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
	fl := fieldList{
		fieldValues: make([]fieldValue, 0, len(fields)+len(f.fieldValues)),
	}
	fl.fieldValues = append(fl.fieldValues, f.fieldValues...)
	for _, field := range fields {
		field.AddTo(&fl)
	}

	colorMessageValue := colorMessage
	if entry.Level == zapcore.WarnLevel {
		colorMessageValue = colorLabelWarn
	} else if entry.Level >= zapcore.ErrorLevel {
		colorMessageValue = colorLabelError
	}
	f.printLabel(colorLabel(labelLevel), entry.Level.CapitalString())
	f.printLabel(colorLabel(labelMessage), colorMessageValue(entry.Message))
	f.printLabel(colorLabel(labelTimestamp), entry.Time.Format("2006-01-02 15:04:05 Z07:00"))

	if len(fl.fieldValues) == 0 {
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
	f.printFields("", fl.fieldValues)
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
	_, _ = fmt.Fprintf(f.buf, fmt.Sprintf("%%-%ds: %%s\n", len(labelLonger)+2), label, value)
}

func (f *FancyEncoder) printFields(prefix string, fields []fieldValue) {
	maxLen := 0
	for _, field := range fields {
		if len(field.key) > maxLen {
			maxLen = len(field.key)
		}
	}
	format := fmt.Sprintf("      %%s %%-%ds: %%s\n", maxLen)
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
