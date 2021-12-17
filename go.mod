module learn.go

go 1.16

require (
	github.com/spf13/cobra v1.3.0
	learn.go.tools v0.0.0-00010101000000-000000000000
)

replace learn.go.tools => ../learn.go.tools

replace github.com/spf13/cobra => github.com/spf13/cobra v1.2.0
