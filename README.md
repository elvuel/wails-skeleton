# Wails Svelte Skeleton

This project is a Wails v2 skeleton based on the official Svelte template and extended with:

- bilingual UI scaffolding (`zh-CN`, `en-US`)
- GORM + SQLite3 persistence for app settings

## Structure

- `app.go`: Wails bindings, lifecycle hooks, database bootstrap, and window controls
- `db.go`: GORM models and SQLite3 persistence helpers
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
