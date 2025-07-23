# NeuroCoach Monorepo

## Structure

- `apps/frontend`: React frontend (Vite)
- `apps/rest-api`: Go REST API
- `apps/ws-api`: Go WebSocket API
- `apps/android-app`: Android app (Flutter)
- `libs/`: Shared libraries and utilities

## Development
- To access frontend, just go to [blazz1t.online](https://blazz1t.online:8443/) in your browser
- To access rest api, go to [api.blazz1t.online/health](https://api.blazz1t.online:8443/health)
- To access ws-api, use `wscat -c ws://ws.blazz1t.online/ws/`
- To access the android app, navigate to `apps/android-app/releases` and download the `.apk` file on your Android device.
- To build your own local version of the project, do the following steps:
  - Edit `nginx/nginx.conf` so all instances of `blazz1t.online` become `localhost`
  - In the same file, replace all `listen 443 ssl` instances with `listen 80` and remove all ssl related lines
  - Also remove the server {} block that returns 301 with redirect.
  - Add this line to your `hosts` file:  
    `127.0.0.1 localhost api.localhost ws.localhost`  
    - On **Windows**, it's located in `C:\Windows\System32\drivers\etc\hosts`  
    - On **Linux/macOS**, it's located in `/etc/hosts`
  - Your machine may reserve port 80. In that case, edit `docker-compose.yml` so the `nginx` service uses the following ports:  
    `"8080:80"`  
    Then access endpoints via [http://localhost:8080](http://localhost:8080) instead of [http://localhost](http://localhost)
  - Make sure you have all Docker prerequisites installed and run:  
    `docker compose up --build -d`
  - Access parts of the project as described above, just replacing the domain with `localhost`
  - If you get 404s on subdomains, confirm your `hosts` file is edited correctly and the `nginx` container is listening on the right domain names.
