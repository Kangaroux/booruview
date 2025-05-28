package api

import (
	"log"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

const (
	perSecondLimit = 3
	perSecondBurst = 5

	evictInterval = 5 * time.Minute
	banTime       = 15 * time.Minute
)

var (
	bannedIps         map[string]time.Time = make(map[string]time.Time)
	banMutex          sync.Mutex
	clientLimits      map[string]*limiter = make(map[string]*limiter)
	clientLimitsMutex sync.Mutex
)

type limiter struct {
	limiter   *rate.Limiter
	touchedAt time.Time
}

func init() {
	// Cleanup rate limits to avoid leaking
	go func() {
		ticker := time.NewTicker(evictInterval)
		defer ticker.Stop()

		for range ticker.C {
			evictKeys()
		}
	}()
}

func evictKeys() {
	clientLimitsMutex.Lock()
	defer clientLimitsMutex.Unlock()

	evictCount := 0
	evictIfOlder := time.Now().Add(-evictInterval)
	for k, v := range clientLimits {
		if v.touchedAt.Before(evictIfOlder) {
			delete(clientLimits, k)
			evictCount++
		}
	}

	if evictCount > 0 {
		log.Println("[rate limit] evicted", evictCount, "old rate limit keys")
	}
}

func getLimiter(ip string) *limiter {
	clientLimitsMutex.Lock()
	defer clientLimitsMutex.Unlock()

	lim, ok := clientLimits[ip]
	if !ok {
		lim = &limiter{limiter: rate.NewLimiter(perSecondLimit, perSecondBurst)}
		clientLimits[ip] = lim
	}

	return lim
}

func isBanned(ip string) bool {
	banMutex.Lock()
	defer banMutex.Unlock()

	unbanTime, banned := bannedIps[ip]
	if !banned {
		return false
	}

	// Unban if it expired
	if time.Now().After(unbanTime) {
		delete(bannedIps, ip)
		return false
	}

	return true
}

func ban(ip string) {
	banMutex.Lock()
	defer banMutex.Unlock()
	bannedIps[ip] = time.Now().Add(banTime)
}

func IsRateLimited(ip string) bool {
	if isBanned(ip) {
		log.Println("[rate limit] blocked:", ip)
		return true
	}

	lim := getLimiter(ip)
	lim.touchedAt = time.Now()

	if !lim.limiter.Allow() {
		ban(ip)
		log.Println("[rate limit] too many requests, banning:", ip)
		return true
	}

	return false
}
