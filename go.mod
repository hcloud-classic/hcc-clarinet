module hcc/clarinet

go 1.13

require (
	github.com/Terry-Mao/goconf v0.0.0-20161115082538-13cb73d70c44
	github.com/go-openapi/strfmt v0.19.11 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gojp/goreportcard v0.0.0-20201106142952-232d912e513e // indirect
	github.com/jedib0t/go-pretty/v6 v6.2.1
	github.com/spf13/cobra v1.1.1
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
	golang.org/x/lint v0.0.0-20201208152925-83fdc39ff7b5 // indirect
	golang.org/x/tools v0.1.0 // indirect
	innogrid.com/hcloud-classic/hcc_errors v0.0.0
)

replace innogrid.com/hcloud-classic/hcc_errors => ../hcc_errors
