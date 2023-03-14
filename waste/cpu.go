package waste

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/crypto/chacha20"
)

func CPU(interval time.Duration, duration time.Duration) {
	var buffer []byte
	if len(Buffers) > 0 {
		buffer = Buffers[0].B[:4*MiB]
	} else {
		buffer = make([]byte, 4*MiB)
	}
	rand.Read(buffer)

	running := true
	for {
		// startTime := time.Now()
		// endTime := startTime.Add(duration)
		running = true
		fmt.Println("[CPU] Start wasted on", time.Now())
		for i := 0; i < 16; i++ {
			go func() {

				// try to construct a new cipher
				cipher, _ := chacha20.NewUnauthenticatedCipher(buffer[:32], buffer[:24])
				for {
					for i := 0; i < 16; i++ {
						cipher.XORKeyStream(buffer, buffer)
					}
					if (!running) {
						break;
					}
				}
			}()
		}
		time.Sleep(duration)
		running = false

		fmt.Println("[CPU] Successfully wasted on", time.Now())

		time.Sleep(interval)
	}
}
