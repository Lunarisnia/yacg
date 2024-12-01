debug:
	go run ./cmd debug | less -S
render:
	go run ./cmd > test.ppm
