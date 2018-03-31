package core

import (
    "sync"
)

// Rules
var (
    // File mutex
    mutex              sync.Mutex

    // Allow rules
    AllowFtpProtocol   bool = true
    AllowCreateUser    bool = true
    AllowCreateUserDb  bool = true
    AllowCreateUserSys bool = true
    AllowShellAccess   bool = true
    AllowSudo          bool = true

    // Limit rules
    MaxTotalStorage     int = 1000
    MaxStoragePerUser   int = 50
    MaxEmailsPerUser    int = 50

    // Default definitions
    Ulimit              int = 8192
    FTPGroup         string = "sftp"
)

// Lock file
func LockRules() {mutex.Lock()}

// Unlock file
func UnlockRules() {mutex.Unlock()}

// Live stats is a mutex protected struct
// Only one goroutine can update the stats at onces
type LiveStats struct {
    mutex          sync.Mutex
    Storage        int        // Mbs
    FTPConnections int
}

// Update livestats
func (rt *LiveStats) Update(storage, ftpc int) {
    rt.Lock()
    defer rt.Unlock()
    rt.Storage = storage
    rt.FTPConnections = ftpc
}

// Lock mutex
func (rt *LiveStats) Lock() {rt.mutex.Lock()}

// Unlockmutex
func (rt *LiveStats) Unlock() {rt.mutex.Unlock()}

// RealTimeStats
var RealTimeStats = &LiveStats{}
