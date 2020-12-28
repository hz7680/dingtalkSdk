package model

import (
	"sync"
	"time"
)

type AccessTokenResp struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type AccessToken struct {
	accessToken string
	expiredAt   time.Time
	locker      *sync.RWMutex
}

func (a *AccessToken) IsExpired() bool {
	return a.accessToken == "" || a.expiredAt.Unix()-100 > time.Now().Unix()
}

func (a *AccessToken) RLock() {
	a.locker.RLock()
}

func (a *AccessToken) RUnlock() {
	a.locker.RUnlock()
}

func (a *AccessToken) Lock() {
	a.locker.Lock()
}

func (a *AccessToken) UnLock() {
	a.locker.Unlock()
}

func (a *AccessToken) SetAccessToken(token string, expiresIn int) {
	a.accessToken = token
	a.expiredAt = time.Now().Add(time.Duration(expiresIn) * time.Second)
}

func (a *AccessToken) GetAccessToken() string {
	return a.accessToken
}

func NewAccessToken() *AccessToken {
	return &AccessToken{
		accessToken: "",
		expiredAt:   time.Now().Add(-1*time.Second),
		locker:      new(sync.RWMutex),
	}
}
