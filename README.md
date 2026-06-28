# Dear Imgui for Go

[![Continuous Integration][ci-badge]][ci-link]

[ci-badge]: https://github.com/bitwizeshift/go-imgui/actions/workflows/continuous-integration.yaml/badge.svg
[ci-link]: https://github.com/bitwizeshift/go-imgui/actions/workflows/continuous-integration.yaml

> [!WARNING]
>
> Until this project reaches a stable 1.0 release, the API is not guaranteed to
> be stable and may change at any time.

This project provides Go-bindings for the [Dear Imgui][imgui] library. In
addition to providing raw access to the underlying C++ API, it also provides a
more Go-idiomatic abstraction leveraging a `Widget` hierarchy to allow easy
construction of GUI layouts.

This project leverages the [cimgui] project as FFI bindings.

[cimgui]: https://github.com/cimgui/cimgui
[imgui]: https://github.com/ocornut/imgui

## Why yet another `go-imgui` wrapper?

Every existing `imgui` library either:

* Is outdated/not maintained/abandoned
* Forces a transitive dependency on unnecessary packages (e.g. `testify` or
  other CGO wrapper tools)
* Doesn't feel Go-idiomatic (mostly 1:1 exposing of the C++ API)

Ultimately, rolling my own version is not that hard, and prevents pulling in
dependencies I can't trust.
