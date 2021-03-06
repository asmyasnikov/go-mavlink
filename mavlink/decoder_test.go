package mavlink

import (
	"bytes"
	"fmt"
	"github.com/asmyasnikov/go-mavlink/mavlink/mock"
	"github.com/asmyasnikov/go-mavlink/mavlink/version"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func BenchmarkDecoder(b *testing.B) {
	var buffer bytes.Buffer
	dec := NewDecoder(&buffer)
	enc := NewEncoder(&buffer)

	rand.Seed(123)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	rand.Seed(time.Now().UnixNano())

	sended := uint32(0)
	received := uint32(0)
	go func() {
		defer wg.Done()
		fmt.Println("to send", b.N)
		for i := 0; i < b.N; i++ {
			ping := mock.Ping{
				Seq: rand.Uint32(),
			}
			p, err := NewPacket(version.MAVLINK_VERSION(i%2), uint8(rand.Uint32()%uint32(^uint8(0))), uint8(rand.Uint32()%uint32(^uint8(0))), &ping)
			if err != nil {
				b.Fatal(err)
			}
			if err := enc.Encode(p); err != nil {
				b.Fatal(err)
			}
			sended++
		}
	}()

	go func() {
		defer wg.Done()
		for {
			packet, err := dec.Decode()
			if err != nil {
				b.Fatal(err)
			}
			if packet.IsNil() {
				b.Fatal("nil packet")
			}
			received++
		}
	}()
	wg.Wait()
	if received != sended {
		b.Fatalf("Sended (%d) and received (%d) packets not equal", sended, received)
	}
}
