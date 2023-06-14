// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64

package tracer

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type tcpRTTHist struct {
	Latency uint64
	Cnt     uint64
	Slots   [27]uint32
	_       [4]byte
}

type tcpRTTHistKey struct {
	Family uint16
	Addr   [16]uint8
}

// loadTcpRTT returns the embedded CollectionSpec for tcpRTT.
func loadTcpRTT() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_TcpRTTBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load tcpRTT: %w", err)
	}

	return spec, err
}

// loadTcpRTTObjects loads tcpRTT and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*tcpRTTObjects
//	*tcpRTTPrograms
//	*tcpRTTMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadTcpRTTObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadTcpRTT()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// tcpRTTSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type tcpRTTSpecs struct {
	tcpRTTProgramSpecs
	tcpRTTMapSpecs
}

// tcpRTTSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type tcpRTTProgramSpecs struct {
	IgTcprcvestKp *ebpf.ProgramSpec `ebpf:"ig_tcprcvest_kp"`
}

// tcpRTTMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type tcpRTTMapSpecs struct {
	Hists *ebpf.MapSpec `ebpf:"hists"`
}

// tcpRTTObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadTcpRTTObjects or ebpf.CollectionSpec.LoadAndAssign.
type tcpRTTObjects struct {
	tcpRTTPrograms
	tcpRTTMaps
}

func (o *tcpRTTObjects) Close() error {
	return _TcpRTTClose(
		&o.tcpRTTPrograms,
		&o.tcpRTTMaps,
	)
}

// tcpRTTMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadTcpRTTObjects or ebpf.CollectionSpec.LoadAndAssign.
type tcpRTTMaps struct {
	Hists *ebpf.Map `ebpf:"hists"`
}

func (m *tcpRTTMaps) Close() error {
	return _TcpRTTClose(
		m.Hists,
	)
}

// tcpRTTPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadTcpRTTObjects or ebpf.CollectionSpec.LoadAndAssign.
type tcpRTTPrograms struct {
	IgTcprcvestKp *ebpf.Program `ebpf:"ig_tcprcvest_kp"`
}

func (p *tcpRTTPrograms) Close() error {
	return _TcpRTTClose(
		p.IgTcprcvestKp,
	)
}

func _TcpRTTClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed tcprtt_bpfel_arm64.o
var _TcpRTTBytes []byte
