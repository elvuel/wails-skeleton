package main

import (
	_ "embed"
	goruntime "runtime"
	"sync"

	"github.com/getlantern/systray"
)

var (
	//go:embed build/windows/icon.ico
	windowsTrayIcon []byte

	//go:embed build/appicon.png
	defaultTrayIcon []byte
)

type trayService struct {
	app     *App
	started sync.Once
	stopped sync.Once

	mu       sync.RWMutex
	locale   string
	showItem *systray.MenuItem
	hideItem *systray.MenuItem
	quitItem *systray.MenuItem
}

func newTrayService(app *App) *trayService {
	return &trayService{app: app}
}

func (t *trayService) Start() {
	t.started.Do(func() {
		go systray.Run(t.onReady, t.onExit)
	})
}

func (t *trayService) Stop() {
	t.stopped.Do(func() {
		systray.Quit()
	})
}

func (t *trayService) onReady() {
	icon := trayIconBytes()
	if len(icon) > 0 {
		systray.SetIcon(icon)
	}

	systray.SetTitle(appName)
	systray.SetTooltip(appName)

	showItem := systray.AddMenuItem("Show Window", "Show the Wails main window")
	hideItem := systray.AddMenuItem("Hide Window", "Hide the Wails main window")
	systray.AddSeparator()
	quitItem := systray.AddMenuItem("Quit", "Quit the application")

	t.mu.Lock()
	t.showItem = showItem
	t.hideItem = hideItem
	t.quitItem = quitItem
	locale := t.locale
	t.mu.Unlock()

	t.applyLocale(locale)

	go func() {
		for {
			select {
			case <-showItem.ClickedCh:
				t.app.ShowWindow()
			case <-hideItem.ClickedCh:
				t.app.HideWindow()
			case <-quitItem.ClickedCh:
				t.app.QuitApp()
				return
			}
		}
	}()
}

func (t *trayService) onExit() {}

func (t *trayService) SetLocale(locale string) {
	t.mu.Lock()
	t.locale = locale
	t.mu.Unlock()

	t.applyLocale(locale)
}

func (t *trayService) applyLocale(locale string) {
	labels := trayLabels(locale)

	t.mu.RLock()
	showItem := t.showItem
	hideItem := t.hideItem
	quitItem := t.quitItem
	t.mu.RUnlock()

	if showItem == nil || hideItem == nil || quitItem == nil {
		return
	}

	showItem.SetTitle(labels.show)
	showItem.SetTooltip(labels.showTooltip)
	hideItem.SetTitle(labels.hide)
	hideItem.SetTooltip(labels.hideTooltip)
	quitItem.SetTitle(labels.quit)
	quitItem.SetTooltip(labels.quitTooltip)
}

func trayIconBytes() []byte {
	if goruntime.GOOS == "windows" {
		return windowsTrayIcon
	}

	return defaultTrayIcon
}

type trayLabelSet struct {
	show        string
	showTooltip string
	hide        string
	hideTooltip string
	quit        string
	quitTooltip string
}

func trayLabels(locale string) trayLabelSet {
	switch locale {
	case "zh-CN":
		return trayLabelSet{
			show:        "显示窗口",
			showTooltip: "显示 Wails 主窗口",
			hide:        "隐藏窗口",
			hideTooltip: "隐藏 Wails 主窗口",
			quit:        "退出应用",
			quitTooltip: "退出当前应用",
		}
	default:
		return trayLabelSet{
			show:        "Show Window",
			showTooltip: "Show the Wails main window",
			hide:        "Hide Window",
			hideTooltip: "Hide the Wails main window",
			quit:        "Quit",
			quitTooltip: "Quit the application",
		}
	}
}
