package socketio

import (
	"testing"
)

func TestDisconnect(t *testing.T) {
	m := NewDisconnect()
	if m.String() != "0::" {
		t.Errorf("Disconnect message string")
	}
}

func TestConnect(t *testing.T) {
	endpoint := NewEndpoint("/path", "Key=Value")
	m := NewConnect(endpoint)
	if m.String() != "1::/path?Key=Value" {
		t.Errorf("Connect message string")
	}
}

func TestHeartbeat(t *testing.T) {
	m := NewHeartbeat()
	if m.String() != "2::" {
		t.Errorf("Connect message string")
	}
}

func TestMessage(t *testing.T) {
	endpoint := NewEndpoint("/path", "Key=Value")
	m := NewMessage(endpoint, "This is a test message")
	if m.String() != "3::/path?Key=Value:This is a test message" {
		t.Errorf("Connect message string")
	}
}
