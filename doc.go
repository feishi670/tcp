// Copyright 2014 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tcp implements TCP-level socket options.
//
// The package provides TCP-level socket options that allow
// manipulation of TCP connection facilities.
//
//
// Monitoring a TCP connection
//
// For now only Darwin, FreeBSD, Linux and NetBSD kernels support the
// TCP information option. A custom net.Dial function that hooks up an
// underlying transport connection must be prepared before monitoring.
//
//	import (
//		"github.com/mikioh/tcp"
//		"github.com/mikioh/tcpinfo"
//		"github.com/mikioh/tcpopt"
//	)
//
//	tr := &http.Transport{
//		Dial: func(network, host string) (net.Conn, error) {
//			d := net.Dialer{DualStack: true}
//			c, err := d.Dial(network, host)
//			if err != nil {
//				return nil, err
//			}
//			tc, err := tcp.NewConn(c)
//			if err != nil {
//				c.Close()
//				return nil, err
//			}
//			go monitor(tc)
//			return tc.Conn, nil
//		},
//		TLSClientConfig: &tls.Config{ServerName: "golang.org"},
//	}
//	client := http.Client{Transport: tr}
//	resp, err := client.Get("https://golang.org")
//	if err != nil {
//		// error handling
//	}
//
// When the underlying transport connection is established, your
// monitor goroutine can start monitoring the connection by using the
// Option method of Conn and tcpinfo package.
//
//	func monitor(tc *tcp.Conn) {
//		tc.SetOption(tcpopt.KeepAlive(true))
//		tc.SetOption(tcpopt.KeepAliveProbeCount(3))
//		var o tcpinfo.Info
//		var b [256]byte
//		for {
//			i, err := tc.Option(o.Level(), o.Name(), b[:])
//			if err != nil {
//				// error handling
//			}
//			txt, err := json.Marshal(i)
//			if err != nil {
//				// error handling
//			}
//			fmt.Println(txt)
package tcp
