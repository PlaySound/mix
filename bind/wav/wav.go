// Package wav is direct WAV filo I/O
package wav

import (
	"io"
	"os"

	"gopkg.in/mix.v0/bind/sample"
	"gopkg.in/mix.v0/bind/spec"
)

// Load a WAV file into memory
func Load(path string) (out []sample.Sample, specs *spec.AudioSpec) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("File not found: " + path)
	}
	file, _ := os.Open(path)
	reader, err := NewReader(file)
	if err != nil {
		panic(err)
	}
	specs = &spec.AudioSpec{
		Freq:     float64(reader.Format.SampleRate),
		Format:   reader.AudioFormat,
		Channels: int(reader.Format.NumChannels),
	}
	for {
		samples, err := reader.ReadSamples()
		if err == io.EOF {
			break
		}
		out = append(out, samples...)
	}
	return
}

//
// Private
//
