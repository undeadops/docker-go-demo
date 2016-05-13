# docker-go-demo
Testing Go - quick demo

Simple showing of Docker Containers, with a minimal Go webserver.

Dockerfile is built from Scratch, Only includes compiled go binary and runs it.

To build, run:

    make docker

To run built container:

    make run
    
Also have a quick way of starting MongoDB with:

    make mongo
