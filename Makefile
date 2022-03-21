go: export GOMODULE111=off
go:
	cd 2021/$(DAY) && go run main.go

rust:
	cd 2021/rust && cargo run --bin $(DAY)