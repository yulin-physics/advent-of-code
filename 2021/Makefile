rustfile=$(shell cd rust/src/bin && find . -maxdepth 1 -type f -name "day*" | sed 's|.rs||')
gofile=$(shell find . -maxdepth 1 -type d -name "day*" | sed s'|./day-||')

$(gofile):
	@echo "Running Go solution"
	@cd day-$@ && GOMODULE=off go run main.go

$(rustfile):
	cargo build --manifest-path rust/Cargo.toml
	@if [ -f rust/src/bin/inputs/$@.txt ]; then cat rust/src/bin/inputs/$@.txt | RUST_BACKTRACE=1 rust/target/debug/$@; fi