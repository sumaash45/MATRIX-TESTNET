// Copyright 2018 The MATRIX Authors as well as Copyright 2014-2017 The go-ethereum Authors
// This file is consisted of the MATRIX library and part of the go-ethereum library.
//
// The MATRIX-ethereum library is free software: you can redistribute it and/or modify it under the terms of the MIT License.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, 
//and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject tothe following conditions:
//
//The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, 
//WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISINGFROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
//OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
// Copyright 2011 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import (
	"crypto/md5"
	"crypto/sha1"
	"hash"
)

// Well known Name Space IDs and UUIDs
var (
	NameSpace_DNS  = Parse("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	NameSpace_URL  = Parse("6ba7b811-9dad-11d1-80b4-00c04fd430c8")
	NameSpace_OID  = Parse("6ba7b812-9dad-11d1-80b4-00c04fd430c8")
	NameSpace_X500 = Parse("6ba7b814-9dad-11d1-80b4-00c04fd430c8")
	NIL            = Parse("00000000-0000-0000-0000-000000000000")
)

// NewHash returns a new UUID derived from the hash of space concatenated with
// data generated by h.  The hash should be at least 16 byte in length.  The
// first 16 bytes of the hash are used to form the UUID.  The version of the
// UUID will be the lower 4 bits of version.  NewHash is used to implement
// NewMD5 and NewSHA1.
func NewHash(h hash.Hash, space UUID, data []byte, version int) UUID {
	h.Reset()
	h.Write(space)
	h.Write([]byte(data))
	s := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, s)
	uuid[6] = (uuid[6] & 0x0f) | uint8((version&0xf)<<4)
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // RFC 4122 variant
	return uuid
}

// NewMD5 returns a new MD5 (Version 3) UUID based on the
// supplied name space and data.
//
//  NewHash(md5.New(), space, data, 3)
func NewMD5(space UUID, data []byte) UUID {
	return NewHash(md5.New(), space, data, 3)
}

// NewSHA1 returns a new SHA1 (Version 5) UUID based on the
// supplied name space and data.
//
//  NewHash(sha1.New(), space, data, 5)
func NewSHA1(space UUID, data []byte) UUID {
	return NewHash(sha1.New(), space, data, 5)
}
