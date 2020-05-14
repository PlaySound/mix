// Package bind is for modular binding of mix to audio interface
package bind

import (
	"io"
	"time"

	"gopkg.in/mix.v0/bind/hardware/null"
	"gopkg.in/mix.v0/bind/opt"
	"gopkg.in/mix.v0/bind/sample"
	"gopkg.in/mix.v0/bind/sox"
	"gopkg.in/mix.v0/bind/spec"
	"gopkg.in/mix.v0/bind/wav"
)

// Configure begins streaming to the bound out audio interface, via a callback function
func Configure(s spec.AudioSpec) {
	sample.ConfigureOutput(s)
	switch useOutput {
	case opt.OutputWAV:
		wav.ConfigureOutput(s)
	case opt.OutputNull:
		null.ConfigureOutput(s)
	}
}

func IsDirectOutput() bool {
	return useOutput == opt.OutputWAV
}

// SetMixNextOutFunc to stream mix out from mix
func SetOutputCallback(fn sample.OutNextCallbackFunc) {
	sample.SetOutputCallback(fn)
}

// OutputStart requires a known length
func OutputStart(length time.Duration, out io.Writer) {
	switch useOutput {
	case opt.OutputWAV:
		wav.OutputStart(length, out)
	case opt.OutputNull:
		// do nothing
	}
}

// OutputNext using the configured writer.
func OutputNext(numSamples spec.Tz) {
	switch useOutput {
	case opt.OutputWAV:
		wav.OutputNext(numSamples)
	case opt.OutputNull:
		// do nothing
	}
}

// LoadWAV into a buffer
func LoadWAV(file string) ([]sample.Sample, *spec.AudioSpec) {
	switch useLoader {
	case opt.InputWAV:
		return wav.Load(file)
	case opt.InputSOX:
		return sox.Load(file)
	default:
		return make([]sample.Sample, 0), &spec.AudioSpec{}
	}
}

// Teardown to close all hardware bindings
func Teardown() {
	switch useOutput {
	case opt.OutputWAV:
		wav.TeardownOutput()
	case opt.OutputNull:
		// do nothing
	}
}

// UseLoader to select the file loading interface
func UseLoader(opt opt.Input) {
	useLoader = opt
}

// UseLoaderString to select the file loading interface by string
func UseLoaderString(loader string) {
	switch loader {
	case string(opt.InputWAV):
		useLoader = opt.InputWAV
	case string(opt.InputSOX):
		useLoader = opt.InputSOX
	default:
		panic("No such Loader: " + loader)
	}
}

// UseOutput to select the outback interface
func UseOutput(opt opt.Output) {
	useOutput = opt
}

// UseOutputString to select the outback interface by string
func UseOutputString(output string) {
	switch output {
	case string(opt.OutputWAV):
		useOutput = opt.OutputWAV
	case string(opt.OutputNull):
		useOutput = opt.OutputNull
	default:
		panic("No such Output: " + output)
	}
}

//
// Private
//

var (
	useLoader = opt.InputWAV
	useOutput = opt.OutputNull
)
