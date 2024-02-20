# Stress Test

## Building Docker container

``docker build -t andre2ar/stress-test .``

## Running the stress test

``docker run andre2ar/stress-test --url=http://google.com --requests=50 --concurrency=10``