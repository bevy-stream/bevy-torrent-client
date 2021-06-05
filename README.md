# Bevy Torrent Client

This is a torrent client that aims to bring a modern flavor to the client scene, while also being backwards compatible with other torrent client apis.

## Development

Due to the fact that both anacrolix/torrent and gorm import sqlite3, we need to run with a flag that allows multiple definitions in the linker.

```
go run --ldflags '-extldflags "-Wl,--allow-multiple-definition"' cmd/main.g
```

## Testing
