package api

import (
	"runtime/debug"

	"github.com/rs/zerolog/log"
)

func LogStackTrace() {
	log.Info().Msg("stack trace\n" + string(debug.Stack()))
}
