go build -o $1
dlv --listen=:2345 --headless=true --api-version=2 exec ./$1