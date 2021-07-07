`docker build -t version_controll:0.1 -f Dockerfile .`

`docker run --publish 8080:8080 -e WS_ORIGIN=localhost:8080 -e WS_SECRET_KEY=1234 version_controll:0.1`