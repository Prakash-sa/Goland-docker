# Goland-Docker

CMDs:

- go mod init github.com/Prakash-sa
- go mod tidy
- docker build -t docker-test .
- docker image ls
- docker run -p 8080:8081 -it docker-test


## AWS cmds

- ssh-keygen -t rsa -b 4096
- docker-machine create -d amazonec2 \
> --amazonec2-region us-east-2 \
> --amazonec2-instance-type "t2.micro" \
> --amazonec2-ssh-keypath ~/.ssh/id_rsa \
> aws-new-cli-test 

- eval $(docker-machine env aws-new-cli-test)
- docker-test % docker build -t docker-test .
- docker run -p 8080:8081 -it docker-test