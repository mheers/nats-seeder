package models

import "time"

type RuntimeInfo struct {
	Version    string
	GoVersion  string
	DirtyBuild bool
	BuildTime  string
	GitTag     string
	GitBranch  string
	CommitHash string
	CommitTime time.Time
}
