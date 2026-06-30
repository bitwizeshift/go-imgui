package imgui

// EndFunc ends a scope opened by a scope-returning function such as [Window],
// [Child] or [StyleColor]. It is always safe to call, including via defer, and
// for conditional scopes is a no-op when the scope did not open.
type EndFunc func()
