package api

import (
	"sync"
	"time"

	"github.com/rs/zerolog/log"

	"golang.org/x/time/rate"
)

const (
	perSecondLimit = 3
	burstLimit     = 20

	evictInterval = 5 * time.Minute
	banTime       = 10 * time.Minute
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
		log.Info().Int("evicted", evictCount).Msg("evict old rate limit keys")
	}
}

func getLimiter(ip string) *limiter {
	clientLimitsMutex.Lock()
	defer clientLimitsMutex.Unlock()

	lim, ok := clientLimits[ip]
	if !ok {
		lim = &limiter{limiter: rate.NewLimiter(perSecondLimit, burstLimit)}
		clientLimits[ip] = lim
	}

	return lim
}

func isBanned(ip string) (banned bool, unbannedAt *time.Time) {
	banMutex.Lock()
	defer banMutex.Unlock()

	unbanTime, banned := bannedIps[ip]
	if !banned {
		return false, nil
	}

	// Unban if it expired
	if time.Now().After(unbanTime) {
		delete(bannedIps, ip)
		return false, nil
	}

	return true, &unbanTime
}

func ban(ip string) {
	banMutex.Lock()
	defer banMutex.Unlock()
	bannedIps[ip] = time.Now().Add(banTime)

	log.Warn().Str("ip", ip).Dur("until", banTime).Msg("client banned (too many requests)")
}

func IsRateLimited(ip string) bool {
	if banned, unbanAt := isBanned(ip); banned {
		log.Warn().Str("ip", ip).Dur("banRemaining", time.Until(*unbanAt)).Msg("blocked client request (banned)")
		return true
	}

	lim := getLimiter(ip)
	lim.touchedAt = time.Now()

	if !lim.limiter.Allow() {
		ban(ip)
		return true
	}

	return false
}
