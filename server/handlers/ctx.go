package handlers

import (
	"context"
	"net/http"
	"net/url"

	"github.com/go-chi/jwtauth"
	"github.com/sirupsen/logrus"
)

// CtxKey - context key
type CtxKey int

const (
	logCtxKey CtxKey = iota
	httpCtxKey
	dbCtxKey
	jwtCtxKey
)

// CtxLog handler
func CtxLog(entry *logrus.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

// Log handler
func Log(r *http.Request) *logrus.Entry {
	return r.Context().Value(logCtxKey).(*logrus.Entry)
}

// CtxHTTP handler
func CtxHTTP(http *url.URL) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, httpCtxKey, http)
	}
}

// HTTP Handler
func HTTP(r *http.Request) *url.URL {
	return r.Context().Value(httpCtxKey).(*url.URL)
}

// CtxJWT Handler
func CtxJWT(entry *jwtauth.JWTAuth) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, jwtCtxKey, entry)
	}
}

// JWT handler
func JWT(r *http.Request) *jwtauth.JWTAuth {
	return r.Context().Value(jwtCtxKey).(*jwtauth.JWTAuth)
}
