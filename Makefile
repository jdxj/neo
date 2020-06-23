file=neo.out

send: build
	scp ./$(file) root@hd1h.ssh.aaronkir.xyz:neo
	scp ./config.json root@hd1h.ssh.aaronkir.xyz:neo
build:
	go build -o $(file) *.go
clean:
	rm -vf ./$(file)
