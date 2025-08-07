package example

import (
	"sync/atomic"
)

// GlobalCounter - это глобальная переменная счетчика.
// Используем int64 для совместимости с atomic операциями.
var GlobalCounter int64

// Increment увеличивает счетчик на 1 и возвращает новое значение.
func (m *module) Increment() int64 {
	return atomic.AddInt64(&GlobalCounter, 1)
}

// Decrement уменьшает счетчик на 1 и возвращает новое значение.
func (m *module) Decrement() int64 {
	return atomic.AddInt64(&GlobalCounter, -1)
}

// Get возвращает текущее значение счетчика.
func (m *module) Get() int64 {
	return atomic.LoadInt64(&GlobalCounter)
}

// Set устанавливает значение счетчика.
func (m *module) Set(val int64) {
	atomic.StoreInt64(&GlobalCounter, val)
}

// Add увеличивает счетчик на заданное значение и возвращает новое значение.
func (m *module) Add(delta int64) int64 {
	return atomic.AddInt64(&GlobalCounter, delta)
}

// Reset сбрасывает счетчик в 0 и возвращает предыдущее значение.
func (m *module) Reset() int64 {
	return atomic.SwapInt64(&GlobalCounter, 0)
}
