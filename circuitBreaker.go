package main

import (
	"errors"
	"time"
)

type CircuitBreaker struct {
	state           int
	failureThreshold int
	retryTime       time.Duration
	lastFailureTime time.Time
}

const (
	StateClosed = 0
	StateOpen = 1
	StateHalfOpen = 2
)
func (cb *CircuitBreaker) CallService() error {
	if cb.state == StateOpen {
		return errors.New("Circuit breaker is open")
	} else if cb.state == StateHalfOpen {
		// Make a single call to the service
		// If the call succeeds, close the circuit breaker
		// If the call fails, trip the circuit breaker
		return nil
	} else {
		// Make the call to the service
		// If the call fails, increment the failure count
		// If the failure count exceeds the threshold, trip the circuit breaker
		return nil
	}
}

func (cb *CircuitBreaker) Trip() {
	cb.state = StateOpen
	cb.lastFailureTime = time.Now()
}

func (cb *CircuitBreaker) Reset() {
	cb.state = StateClosed
	cb.lastFailureTime = time.Time{}
}

func (cb *CircuitBreaker) MonitorService(serviceURL string) {
	// Make a call to the service
	// If the call fails, increment the failure count
	// If the failure count exceeds the threshold, trip the circuit breaker
	// If the call succeeds, reset the failure count
	// Sleep for a period of time before checking again
	time.Sleep(cb.retryTime)
}

