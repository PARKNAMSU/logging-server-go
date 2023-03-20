FROM golang:1.19

WORKDIR /logging-batch-go
COPY . .

RUN mkdir ./logs
RUN mv ./.aws ~/
RUN go env -w GO111MODULE=auto
RUN go mod download
RUN go get
RUN go build logFile.com/log-file-go

# RUN wget https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm

# RUN sudo yum install google-chrome-stable_current_x86_64.rpm

EXPOSE 8000 6061

# CMD [ "go" , "run" , "main.go" ]
CMD [ "./log-file-go" ]