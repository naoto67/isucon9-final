.PHONY: frontend webapp payment bench

all: frontend webapp payment bench

frontend:
	cd webapp/frontend && make
	cd webapp/frontend/dist && tar zcvf ../../../ansible/files/frontend.tar.gz .

webapp:
	tar zcvf ansible/files/webapp.tar.gz \
	--exclude webapp/frontend \
	webapp

payment:
	cd blackbox/payment && make && cp bin/payment_linux ../../ansible/roles/benchmark/files/payment

bench:
	cd bench && make && cp -av bin/bench_linux ../ansible/roles/benchmark/files/bench && cp -av bin/benchworker_linux ../ansible/roles/benchmark/files/benchworker

up:
	docker-compose -f webapp/docker-compose.yml -f webapp/docker-compose.go.yml up -d

down:
	docker-compose -f webapp/docker-compose.yml -f webapp/docker-compose.go.yml down

run_bench:
	bench/bin/bench_linux run --assetdir=webapp/frontend/dist --target=http://0.0.0.0:8080 --payment=http://0.0.0.0:5000

restart:
	docker-compose -f webapp/docker-compose.yml -f webapp/docker-compose.go.yml restart
