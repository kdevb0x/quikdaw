package main

import (
	"io"
	"io/ioutil"
	"time"

	"github.com/therecipe/qt/core"
)

func init() {

}

// Track represents the core track object;
// It handles the signal flow and low-level interactions.
type track struct {
	input   Source
	plugins []Plugin
	output  Sync
}

// Source represents an abstract signal flow start point.
type Source interface {
	io.ReaderClose

	Device() string
	Format() string
	SampleRate() int
	Latency() time.Duration
}

// NullSource is a source object whose Read method streams silence.
type NullSource struct {
	s ioutil.NopCloser
}

func NewNullSource() Source {
	return nullSource
}

type Sync interface {
	io.WriterCloser
}

// Plugin is an abstract node in the signal chain.
type Plugin struct {
	In     Source
	Plugin Processor
	Out    Sync
}

// Processor represents a dsp process.
type Processor struct {
	Status ProcessorStatus
}

// ProcessorStatus describes the current state of the Processor
type ProcessorStatus int

const (
	_ ProcessorStatus = iota
	Active
	Ready
	Offline
	Error
)

type ChannelStrip struct {
	core.QAbstractItemModel
	track    *Track
	playlist Playlist
}

type Playlist struct {
	core.QAbstractListModel

	modelData []PlaylistItem
}
type PlaylistItem struct {
	Name   string
	Length float32 // in seconds
	Index  int     // this items index in the playlist

}
