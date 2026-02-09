package graphql

import "fmt"

// Semaphore eşzamanlılık kontrolü için semaphore pattern implementasyonu
type Semaphore struct {
	ch chan struct{}
}

// NewSemaphore belirtilen kapasitede yeni bir semaphore oluşturur
func NewSemaphore(max int) *Semaphore {
	return &Semaphore{
		ch: make(chan struct{}, max),
	}
}

// Acquire semaphore'dan bir slot alır. Eğer tüm slotlar doluysa bekler.
func (s *Semaphore) Acquire() {
	fmt.Println("🔒 Semaphore acquiring...")
	s.ch <- struct{}{}
	fmt.Println("✅ Semaphore acquired")
}

// Release semaphore'a bir slot geri verir
func (s *Semaphore) Release() {
	<-s.ch
	fmt.Println("🔓 Semaphore released")
}
