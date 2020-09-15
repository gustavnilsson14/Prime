Dockerized golang webserver providing a servicen to calculate the closest previous prime number to a user number.

Requires docker to run.
Please set your firewall to allow HTTP traffic.

To build, and run:
docker build -t primes .
docker run --name primes-container -d -p 80:80 primes

To shutdown
docker stop primes-container
docker rm primes-container

A demo of the service is available at http://35.187.224.88/ (an ephemeral, non-static ip which is subject to change unfortunately)