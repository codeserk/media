package auth

import "time"

const sessionCacheKey = "auth.user.session"
const sessionDuration time.Duration = 24 * 10 * time.Hour
