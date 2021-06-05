# Bevy Torrent Client

This is a torrent client that aims to bring a modern flavor to the client scene, while also being backwards compatible with other torrent client apis.

Some goals for this project:

- REST api 
  - CRUD
  - Serve torrent files over HTTP
  - Stream media files as they are downloading
- HTTP Dashboard
- Powerfule auto-removal rules and scripting
- Run scripts on events such as torrents finishing

## Development

Due to the fact that both anacrolix/torrent and gorm import sqlite3, we need to run with a flag that allows multiple definitions in the linker.

```
go run --ldflags '-extldflags "-Wl,--allow-multiple-definition"' cmd/main.g
```

## Testing
