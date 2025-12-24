# YouTube Automation Project Instructions

You are an expert Go developer and Frontend engineer. You are helping to build a YouTube video upload automation tool.

## Technical Stack
- **Backend**: Go (Golang)
- **Web Framework**: Echo (v4)
- **Frontend**: HTML templates, Tailwind CSS (via CDN), Alpine.js (via CDN)
- **YouTube API**: google.golang.org/api/youtube/v3
- **Architecture**: Layered Architecture (Handler, Service, Model)

## Project Structure
- `cmd/server/main.go`: Entry point. Keep it minimal.
- `internal/handler/`: HTTP handlers. Handle requests/responses and validation.
- `internal/service/`: Business logic. YouTube API interactions and OAuth2 flow.
- `internal/model/`: Data structures and constants.
- `web/templates/`: HTML templates using Go's `html/template`.
- `web/static/`: Static assets (CSS/JS).

## Coding Standards & Preferences
1. **Go Language**:
   - Use Echo's `Context` for all handler methods.
   - Follow standard Go naming conventions (camelCase for internal, PascalCase for exported).
   - Error handling: Handle errors explicitly and return meaningful messages via Echo.
   - Use environment variables (via `os.Getenv` or a config loader) for secrets.

2. **Frontend (Mobile First)**:
   - Use **Tailwind CSS** for all styling. Design specifically for iPhone/mobile screens.
   - Use **Alpine.js** for client-side interactivity (form submission, loading states).
   - Keep templates modular. Use a base `layout.html`.

3. **YouTube API**:
   - Use the official Google API Go client.
   - Implement OAuth2 flow properly for user authentication.

## Contextual Goal
The primary user will access this tool from an **iPhone** via a local Mac server. Ensure all UI components are touch-friendly and responsive.