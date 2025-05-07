base:
	docker build -t konfigo-nginx .

build-demo:
	docker build -t konfigo-nginx-demo -f Dockerfile.demo .

rn-demo:
	docker run --env=KONFIGO_API_ENDPOINT=http://127.0.0.1:3000/api --env=KONFIGO_PATH=konfigo/gateway/dev --env=KONFIGO_API_KEY=super_secure_api_key --env=KONFIGO_INTERVAL=2 --network="host" -d konfigo-nginx-demo:latest 