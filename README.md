# Wails Svelte Skeleton

This project is a Wails v2 skeleton based on the official Svelte template and extended with:

- bilingual UI scaffolding (`zh-CN`, `en-US`)
- GORM + SQLite3 persistence for app settings
- system tray support with show, hide, and quit actions

## Structure

- `app.go`: Wails bindings, lifecycle hooks, database bootstrap, window controls, and tray wiring
- `db.go`: GORM models and SQLite3 persistence helpers
- `tray.go`: systray service and localized menu wiring
- `frontend/src/App.svelte`: starter dashboard UI
- `frontend/src/lib/backend.js`: lightweight backend bridge for Svelte

## Development

Install dependencies first:

```powershell
go mod tidy
cd frontend
npm install
cd ..
```

Start live development:

```powershell
wails dev
```

Build the desktop app:

```powershell
wails build
```

## Notes

- Locale defaults are detected from environment variables (`LC_ALL`, `LC_MESSAGES`, `LANG`) with fallback to `zh-CN`.
- Locale updates are persisted in SQLite via GORM.
- Closing the main window hides it to the system tray (`HideWindowOnClose: true`).
- `github.com/getlantern/systray` requires `CGO_ENABLED=1`.
