package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// TokenBucket represents a token bucket system.
type TokenBucket struct {
	mu             sync.Mutex
	tokens         float64
	maxTokens      float64
	refillRate     float64
	lastRefillTime time.Time
}

// NewTokenBucket creates a new TokenBucket instance.
func NewTokenBucket(maxTokens float64, refillRate float64) *TokenBucket {
	return &TokenBucket{
		tokens:         maxTokens,
		maxTokens:      maxTokens,
		refillRate:     refillRate,
		lastRefillTime: time.Now(),
	}
}

// refill refills the token bucket based on the elapsed time since the last refill.
func (tb *TokenBucket) refill() {
	now := time.Now()
	duration := now.Sub(tb.lastRefillTime)
	tokensToAdd := tb.refillRate * duration.Seconds()
	tb.tokens = math.Min(tb.tokens+tokensToAdd, tb.maxTokens)
	tb.lastRefillTime = now
}

// Request checks if the token bucket has enough tokens for a request.
// If yes, it deducts the tokens and returns true, otherwise returns false.
func (tb *TokenBucket) Request(tokens float64) bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill()
	if tokens <= tb.tokens {
		tb.tokens -= tokens
		return true
	}
	return false
}

// ServiceATokenBucket represents a token bucket system for Service A.
var ServiceATokenBucket = NewTokenBucket(50, 10)

// ServiceBTokenBucket represents a token bucket system for Service B.
var ServiceBTokenBucket = NewTokenBucket(100, 5)

// RequestFromServiceA simulates a request from Service A.
func RequestFromServiceA() bool {
	return ServiceATokenBucket.Request(1)
}

// RequestFromServiceB simulates a request from Service B.
func RequestFromServiceB() bool {
	return ServiceBTokenBucket.Request(1)
}

type UserTokenBucketManager struct {
	userTokenBuckets map[string]*TokenBucket
	mutex            sync.Mutex
}

func NewUserTokenBucketManager() *UserTokenBucketManager {
	return &UserTokenBucketManager{
		userTokenBuckets: make(map[string]*TokenBucket),
	}
}
func (utbm *UserTokenBucketManager) GetUserTokenBucket(ip string) *TokenBucket {
	utbm.mutex.Lock()
	defer utbm.mutex.Unlock()

	if bucket, ok := utbm.userTokenBuckets[ip]; ok {
		return bucket
	}

	bucket := NewTokenBucket(20, 1)
	utbm.userTokenBuckets[ip] = bucket
	return bucket
}

func (utbm *UserTokenBucketManager) RequestFromUser(ip string) bool {
	userTokenBucket := utbm.GetUserTokenBucket(ip)
	utbm.mutex.Lock()
	defer utbm.mutex.Unlock()

	userTokenBucket.refill()
	if userTokenBucket.tokens >= 1 {
		userTokenBucket.tokens -= 1
		return true
	}
	return false
}

func main() {
	globalTokenBucket := NewTokenBucket(500, 1)

	for i := 0; i < 2000; i++ {
		go func() {
			if globalTokenBucket.Request(1) {
				fmt.Println("Global Request Accepted")
			} else {
				fmt.Println("Global Request Denied")
			}
		}()

		go func() {
			if RequestFromServiceA() {
				fmt.Println("Service A Request Accepted")
			} else {
				fmt.Println("Service A Request Denied")
			}
		}()

		go func() {
			if RequestFromServiceB() {
				fmt.Println("Service B Request Accepted")
			} else {
				fmt.Println("Service B Request Denied")
			}
		}()
		go func() {
			userTokenBucketManager := NewUserTokenBucketManager()
			ip := "192.168.1.1" // Replace with actual user IP
			if userTokenBucketManager.RequestFromUser(ip) {
				fmt.Println("User Request Accepted")
			} else {
				fmt.Println("User Request Denied")
			}
		}()

		time.Sleep(100 * time.Millisecond)
	}
	time.Sleep(5 * time.Second)
}
