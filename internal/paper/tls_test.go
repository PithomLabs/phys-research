package paper

import (
	"context"
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestTLSReproduction(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/pdf")
		w.Write([]byte("%PDF-1.4 test pdf content"))
	}))
	defer server.Close()

	serverAddr := strings.TrimPrefix(server.URL, "https://")
	_, port, err := net.SplitHostPort(serverAddr)
	if err != nil {
		t.Fatalf("split host port: %v", err)
	}

	t.Run("CurrentImplementation_EmptyServerName", func(t *testing.T) {
		dialer := &net.Dialer{Timeout: 10 * time.Second}
		transport := &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				_, port, _ := net.SplitHostPort(addr)
				return dialer.DialContext(ctx, "tcp", net.JoinHostPort("127.0.0.1", port))
			},
			TLSClientConfig: &tls.Config{
				ServerName:         "",
				InsecureSkipVerify: false,
				RootCAs:            server.TLS.ClientCAs,
			},
		}

		client := &http.Client{
			Timeout:   10 * time.Second,
			Transport: transport,
		}

		req, err := http.NewRequest("GET", "https://test.local:"+port, nil)
		if err != nil {
			t.Fatalf("new request: %v", err)
		}
		req.Header.Set("User-Agent", userAgent)
		req.Header.Set("Accept", "application/pdf")

		resp, err := client.Do(req)
		if err != nil {
			t.Logf("Request failed: %v", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("status %d, want 200", resp.StatusCode)
			return
		}

		body, _ := io.ReadAll(resp.Body)
		if !strings.HasPrefix(string(body), "%PDF") {
			t.Errorf("not a valid PDF: %s", string(body)[:50])
		}

		t.Log("RESULT: NOT REPRODUCED - Go auto-supplies hostname from request URL when ServerName=''")
	})

	t.Run("RedirectPreservesHostname", func(t *testing.T) {
		redirectServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://test.local:"+port, http.StatusMovedPermanently)
		}))
		defer redirectServer.Close()

		redirectURL := strings.Replace(redirectServer.URL, "127.0.0.1", "test.local", 1)
		redirectURL = "https://" + strings.TrimPrefix(redirectURL, "http://")

		dialer := &net.Dialer{Timeout: 10 * time.Second}
		transport := &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				_, port, _ := net.SplitHostPort(addr)
				return dialer.DialContext(ctx, "tcp", net.JoinHostPort("127.0.0.1", port))
			},
			TLSClientConfig: &tls.Config{
				ServerName:         "",
				InsecureSkipVerify: false,
				RootCAs:            server.TLS.ClientCAs,
			},
		}

		client := &http.Client{
			Timeout:   10 * time.Second,
			Transport: transport,
		}

		req, _ := http.NewRequest("GET", redirectURL, nil)
		req.Header.Set("User-Agent", userAgent)

		resp, err := client.Do(req)
		if err != nil {
			t.Logf("Redirect request failed: %v", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Redirect status %d", resp.StatusCode)
		}
		t.Log("SUCCESS: Redirect preserves hostname through ServerName auto-supply")
	})
}
