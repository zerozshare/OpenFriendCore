> ## ⚠️ Unofficial — not affiliated with Microsoft, Mojang, or the Xbox brand
>
> OpenFriend is an **independent, community-built** project. It is **not** developed, endorsed, supported, sponsored, certified, or otherwise officially connected to Microsoft Corporation, Mojang AB, Mojang Studios, or the Xbox brand. "Minecraft", "Xbox", "Xbox Live", "Microsoft", and "Mojang" are trademarks of their respective owners. Use OpenFriend on accounts you control, on servers you operate or have permission to operate on. You assume all risk associated with running this software.

> ## 🚧 Current scope: offline-mode servers only
>
> OpenFriend bridges Friends-List joins **only to offline-mode Minecraft servers** at this time. The online-mode bypass (Floodgate-style auth skip) is **implemented but not yet verified end-to-end** because Paper / Spigot have not released a build matching snapshot 26.2. Set `online-mode=false` on the backend server you bridge to until the bypass is certified.

---

# OpenFriend Core

Standalone CLI. Authenticates with your Microsoft account, broadcasts presence on the Minecraft Java Edition Friends List (snapshot 26.2+), and bridges incoming WebRTC joins to a TCP Minecraft server.

Use this when you want OpenFriend without a Bukkit/Spigot/Velocity plugin.

## Binaries

| File | Platform |
|---|---|
| `openfriend-darwin-arm64` | macOS Apple Silicon |
| `openfriend-darwin-amd64` | macOS Intel |
| `openfriend-linux-arm64` | Linux ARM64 |
| `openfriend-linux-amd64` | Linux x86_64 |
| `openfriend-windows-amd64.exe` | Windows x86_64 |

Mark executable on Unix: `chmod +x openfriend-*`.

## Quick start

Host mode (accept incoming joins, bridge to local server on 25565):

```
./openfriend --target 127.0.0.1:25565
```

Join mode (let your local MC client join a friend's world via OpenFriend):

```
./openfriend --join FRIENDNAME --listen 127.0.0.1:25577
```

Then open Minecraft and add `localhost:25577` as a server.

First run prints a Microsoft device code at `https://www.microsoft.com/link`. Authenticate once; the token is encrypted to your machine and reused. Run `--reset` to wipe the saved credential.

## Common flags

| Flag | Default | Purpose |
|---|---|---|
| `--target host:port` | `127.0.0.1:25565` | Host mode: backend Minecraft server |
| `--join NAME` | — | Join mode: friend name or PMID to join |
| `--listen host:port` | `127.0.0.1:25565` | Join mode: local TCP for your MC client |
| `--skin file.png` | — | Upload PNG as Minecraft skin (sets the friend-list icon) |
| `--bypass-key path` | `<data-dir>/bypass.pem` | Enable online-mode bypass via shared key |
| `--data-dir dir` | next to binary | Where `auth.pem` and `status.json` live |
| `--reset` | — | Send OFFLINE, delete saved credentials, exit |
| `--no-update` | — | Skip self-update check |
| `--verbose` | — | Debug logging |
| `--version` | — | Print version and exit |

Full reference: `./openfriend --help`.

## License

MIT. See `LICENSE` in this directory.
