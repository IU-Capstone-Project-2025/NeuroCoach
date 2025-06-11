# NeuroCoach Monorepo

## Structure

- `apps/frontend`: React frontend (Vite)
- `apps/rest-api`: Go REST API
- `apps/ws-api`: Go WebSocket API
- `libs/`: Shared libraries and utilities

## Development

- Use `docker-compose up --build` to start all services.
- To access frontend, just go to 127.0.0.1/ in your browser
- To access rest api, go to 127.0.0.1/api/hello
- To access ws-api, use `wscat -c ws://127.0.0.1/ws/`
- To access the android app, navigate to `apps/android-app/releases` and download .apk file on your android device.