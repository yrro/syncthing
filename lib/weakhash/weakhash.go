// Copyright (C) 2016 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// The existence of this file means we get 0% test coverage rather than no
// test coverage at all. Remove when implementing an actual test.

package weakhash

import (
	"hash"
	"io"
)

const (
	Size = 4
)

func New(size int) hash.Hash32 {
	return &digest{
		buf:  make([]byte, size),
		size: size,
	}
}

func RollMap(r io.Reader, size int) (map[uint32][]int64, error) {
	if r == nil {
		return nil, nil
	}
	hf := New(size)

	n, err := io.CopyN(hf, r, int64(size))
	if err == io.EOF {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if n != int64(size) {
		return nil, io.ErrShortBuffer
	}

	hashes := make(map[uint32][]int64)
	hashes[hf.Sum32()] = []int64{0}

	buf := make([]byte, 1)

	var i int64 = 1
	var hash uint32
	for {
		_, err := r.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			return hashes, err
		}
		hf.Write(buf)
		hash = hf.Sum32()
		hashes[hash] = append(hashes[hash], i)
		i++
	}
	return hashes, nil
}

type digest struct {
	buf  []byte
	size int
	a    uint16
	b    uint16
	j    int
}

func (d *digest) Write(data []byte) (int, error) {
	for _, c := range data {
		d.a = d.a - uint16(d.buf[d.j]) + uint16(c)
		d.b = d.b - uint16(d.size)*uint16(d.buf[d.j]) + d.a
		d.buf[d.j] = c
		d.j = (d.j + 1) % d.size
	}
	return len(data), nil
}

func (d *digest) Reset() {
	for i := range d.buf {
		d.buf[i] = 0x0
	}
	d.a = 0
	d.b = 0
	d.j = 0
}

func (d *digest) Sum(b []byte) []byte {
	r := d.Sum32()
	return append(b, byte(r>>24), byte(r>>16), byte(r>>8), byte(r))
}

func (d *digest) Sum32() uint32 { return uint32(d.a) | (uint32(d.b) << 16) }
func (digest) Size() int        { return Size }
func (digest) BlockSize() int   { return 1 }
