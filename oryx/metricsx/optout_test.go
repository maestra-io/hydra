package metricsx

import (
	"sync"
	"testing"

	"github.com/ory/analytics-go/v5"
)

type recordingClient struct {
	mu     sync.Mutex
	queued []analytics.Message
}

func (c *recordingClient) Enqueue(m analytics.Message) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.queued = append(c.queued, m)
	return nil
}

func (c *recordingClient) Close() error { return nil }

func (c *recordingClient) count() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.queued)
}

func TestOptOutSendsNothing(t *testing.T) {
	c := &recordingClient{}
	sw := &Service{optOut: true, c: c, o: &Options{}, mem: new(MemoryStatistics)}

	sw.Identify()
	sw.Track()

	if n := c.count(); n != 0 {
		t.Fatalf("opted out: %d telemetry message(s) enqueued, want 0", n)
	}
}

func TestOptInSendsStartupIdentify(t *testing.T) {
	c := &recordingClient{}
	sw := &Service{optOut: false, c: c, o: &Options{}, mem: new(MemoryStatistics)}

	IdentifySend(sw, true)

	if n := c.count(); n != 1 {
		t.Fatalf("opted in: %d message(s) enqueued, want 1 startup Identify", n)
	}
}
