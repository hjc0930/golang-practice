package main

import (
	"fmt"
	"sync"
	"time"
)

type UserSession struct {
	UserID       string
	LoginTime    time.Time
	LastActivity time.Time
	IPAddress    string
}

type SessionManager struct {
	sessions map[string]*UserSession
	mutex    sync.RWMutex
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: make(map[string]*UserSession),
	}
}

func (sm *SessionManager) AddSession(token string, userID, ipAddress string) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	userSession := &UserSession{
		UserID:       userID,
		LoginTime:    time.Now(),
		LastActivity: time.Now(),
		IPAddress:    ipAddress,
	}

	sm.sessions[token] = userSession
}

func (sm *SessionManager) GetSession(token string) (*UserSession, bool) {
	// 读锁
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()

	userSession, exists := sm.sessions[token]

	if exists {
		userSession.LastActivity = time.Now()
	}

	return userSession, exists
}

func (sm *SessionManager) RemoveSession(token string) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	delete(sm.sessions, token)
}

func (sm *SessionManager) CleanupExpiredSessions(timeout time.Duration) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	now := time.Now()
	for token, session := range sm.sessions {
		if now.Sub(session.LastActivity) > timeout {
			delete(sm.sessions, token)
		}
	}
}

func main() {
	token := "token1"
	sessionManager := NewSessionManager()
	sessionManager.AddSession(token, "user1", "192.168.1.1")

	if session, ok := sessionManager.GetSession(token); ok {
		fmt.Printf("Find user session: %ss, Login time is: %v\n", session.UserID, session.LoginTime.Format(time.DateTime))
	}

	// 定时清理
	sessionManager.CleanupExpiredSessions(time.Hour)
}
