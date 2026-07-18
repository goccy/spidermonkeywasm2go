# spidermonkeywasm2go

wasm2go sources for `spidermonkey.wasm` — a build of
[SpiderMonkey](https://spidermonkey.dev/), the JavaScript engine from Mozilla
Firefox 147.0.4, transpiled to Go by
[wasm2go](https://github.com/goccy/wasm2go) and consumed by
[go-spidermonkey](https://github.com/goccy/go-spidermonkey).

The bundle is generated and released by
[spidermonkey-wasm](https://github.com/goccy/spidermonkey-wasm) and vendored
here with `make bundle`, which downloads the release tarball and verifies it
against its sha256 manifest and the upstream GitHub SLSA attestation before
extracting it in place. Nothing here is hand-written or meant to be called
directly — [go-spidermonkey](https://github.com/goccy/go-spidermonkey) is the
API you want.

## License

This bundle (`base/`, `p0/`–`p7/`, `data.bin`, and the generated Go glue) is a
derivative work of SpiderMonkey, so it is distributed under SpiderMonkey's
license — the Mozilla Public License, Version 2.0 — reproduced in full in
[LICENSE](./LICENSE). SpiderMonkey is part of mozilla-central and is
Copyright (c) the Mozilla project and its contributors. See
<https://mozilla.org/MPL/2.0/>.
