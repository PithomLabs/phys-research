package paper

import (
	"crypto/tls"
	"testing"
)

// NOTE: http.ProxyFromEnvironment caches its configuration once per process,
// so tests must not mutate proxy variables to exercise both branches. Each
// constructor is tested directly; the dispatcher is tested for consistency
// with the live environment probe. No test here performs network I/O.

func TestProxyTransport(t *testing.T) {
	tr := proxyTransport(&tls.Config{InsecureSkipVerify: false})
	if tr.Proxy == nil {
		t.Fatal("proxy transport must honor ProxyFromEnvironment")
	}
	if tr.DialContext != nil {
		t.Fatal("proxy transport must not use the validating direct dialer")
	}
	if tr.TLSClientConfig == nil || tr.TLSClientConfig.InsecureSkipVerify {
		t.Fatal("proxy transport must enforce TLS verification")
	}
}

func TestDirectTransport(t *testing.T) {
	tr := directTransport(&tls.Config{InsecureSkipVerify: false})
	if tr.Proxy != nil {
		t.Fatal("direct transport must not set a proxy func")
	}
	if tr.DialContext == nil {
		t.Fatal("direct transport must use the validating direct dialer")
	}
	if tr.TLSClientConfig == nil || tr.TLSClientConfig.InsecureSkipVerify {
		t.Fatal("direct transport must enforce TLS verification")
	}
}

func TestSafeTransportMatchesEnv(t *testing.T) {
	tr := safeTransport()
	if tr.TLSClientConfig == nil || tr.TLSClientConfig.InsecureSkipVerify {
		t.Fatal("safeTransport must enforce TLS verification in every branch")
	}
	if envProxyForHTTPS() && tr.Proxy == nil {
		t.Fatal("proxy configured: safeTransport must honor it")
	}
	if !envProxyForHTTPS() && tr.DialContext == nil {
		t.Fatal("no proxy: safeTransport must use the validating direct dialer")
	}
}
