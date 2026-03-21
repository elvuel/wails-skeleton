package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"sync/atomic"

	wailsruntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

const (
	appName          = "wails-skeleton"
	defaultLocale    = "zh-CN"
	readyMessage     = "ready"
	localeSettingKey = "locale"
)

var supportedLocales = []string{"zh-CN", "en-US"}

type BootstrapData struct {
	AppName           string   `json:"appName"`
	Locale            string   `json:"locale"`
	SupportedLocales  []string `json:"supportedLocales"`
	DatabaseDriver    string   `json:"databaseDriver"`
	DatabasePath      string   `json:"databasePath"`
	DatabaseReady     bool     `json:"databaseReady"`
	HideWindowOnClose bool     `json:"hideWindowOnClose"`
	Message           string   `json:"message"`
	StartupError      string   `json:"startupError"`
}

type App struct {
	ctx          context.Context
	locale       string
	db           *databaseState
	startupError string
	quitting     atomic.Bool
	mu           sync.RWMutex
}

func NewApp() *App {
	return &App{
		locale: detectPreferredLocale(),
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	db, err := openDatabase(appName)
	if err != nil {
		a.setStartupError(err)
		return
	}

	storedLocale, err := db.GetSetting(localeSettingKey)
	if err != nil {
		a.setStartupError(err)
	}

	a.mu.Lock()
	a.db = db
	if normalized, ok := normalizeLocale(storedLocale); ok {
		a.locale = normalized
	}
	a.mu.Unlock()
}

func (a *App) shutdown(context.Context) {
	a.quitting.Store(true)

	a.mu.Lock()
	db := a.db
	a.db = nil
	a.mu.Unlock()

	if err := db.Close(); err != nil {
		a.setStartupError(err)
	}
}

func (a *App) Bootstrap() BootstrapData {
	return a.snapshot()
}

func (a *App) SetLocale(locale string) (BootstrapData, error) {
	normalized, ok := normalizeLocale(locale)
	if !ok {
		return a.snapshot(), fmt.Errorf("unsupported locale: %s", locale)
	}

	a.mu.Lock()
	a.locale = normalized
	db := a.db
	a.mu.Unlock()

	if err := db.SetSetting(localeSettingKey, normalized); err != nil {
		return a.snapshot(), fmt.Errorf("persist locale: %w", err)
	}

	return a.snapshot(), nil
}

func (a *App) ShowWindow() {
	if a.ctx == nil {
		return
	}

	wailsruntime.Show(a.ctx)
	wailsruntime.WindowUnminimise(a.ctx)
	wailsruntime.WindowShow(a.ctx)
}

func (a *App) HideWindow() {
	if a.ctx == nil {
		return
	}

	wailsruntime.WindowHide(a.ctx)
}

func (a *App) QuitApp() {
	a.quitting.Store(true)

	if a.ctx != nil {
		wailsruntime.Quit(a.ctx)
	}
}

func (a *App) snapshot() BootstrapData {
	a.mu.RLock()
	defer a.mu.RUnlock()

	message := readyMessage
	dbReady := false
	dbPath := ""
	if a.db != nil {
		dbReady = true
		dbPath = a.db.Path()
	}
	if a.startupError != "" {
		message = a.startupError
	}

	return BootstrapData{
		AppName:           appName,
		Locale:            a.locale,
		SupportedLocales:  append([]string(nil), supportedLocales...),
		DatabaseDriver:    "sqlite",
		DatabasePath:      dbPath,
		DatabaseReady:     dbReady,
		HideWindowOnClose: false,
		Message:           message,
		StartupError:      a.startupError,
	}
}

func (a *App) setStartupError(err error) {
	if err == nil {
		return
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	if a.startupError == "" {
		a.startupError = err.Error()
		return
	}

	a.startupError = a.startupError + "; " + err.Error()
}

func normalizeLocale(locale string) (string, bool) {
	switch strings.ToLower(strings.TrimSpace(locale)) {
	case "zh", "zh-cn", "zh_hans", "zh-hans":
		return "zh-CN", true
	case "en", "en-us", "en-gb":
		return "en-US", true
	default:
		return "", false
	}
}

func detectPreferredLocale() string {
	for _, key := range []string{"LC_ALL", "LC_MESSAGES", "LANG"} {
		value := strings.TrimSpace(os.Getenv(key))
		if value == "" {
			continue
		}

		if normalized, ok := normalizeLocale(value); ok {
			return normalized
		}
	}

	return defaultLocale
}
