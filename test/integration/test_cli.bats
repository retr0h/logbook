#!/usr/bin/env bats

@test "invoke logbook without arguments prints usage" {
	run go run main.go

	[ "$status" -eq 0 ]
	[ "${lines[5]}" = "A CLI based tool to manage HAM logbook entries and/or DX contacts." ]
}

@test "invoke logbook version subcommand" {
	run go run main.go version

	[ "$status" -eq 0 ]
}

@test "invoke logbook add subcommand without arguments prints usage" {
	skip
}

@test "invoke logbook add subcommand" {
	run bash -c 'echo "testCallSign\ntestName\n" | go run main.go add'

	[ "$status" -eq 0 ]
}

@test "invoke logbook get subcommand without arguments prints usage" {
	skip
}

@test "invoke logbook get subcommand with a known call sign" {
	skip
	run bash -c 'echo "testCallSign\ntestName\n" | go run main.go add'
	run bash -c 'go run main.go get --callsign testCallSign; false'

	# [ "${lines[5]}" = "A CLI based tool to manage HAM logbook entries and/or DX contacts." ]
	echo "${lines}"
	echo "${lines[0]}"
	[ "$status" -eq 0 ]
}

@test "invoke logbook get subcommand with an unknown call sign" {
	skip
}
