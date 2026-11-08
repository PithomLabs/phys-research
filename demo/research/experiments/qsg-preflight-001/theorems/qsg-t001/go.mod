module github.com/PithomLabs/qsg-t001

go 1.25.7

// Resolved by the colocated go.work (same pattern as phys-lib's
// require + workspace comment); no replace directives.
require (
	github.com/PithomLabs/phys-artifact v0.0.0
	github.com/PithomLabs/phys-math v0.0.0
)
