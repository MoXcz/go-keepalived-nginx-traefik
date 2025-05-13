FROM golang:1.24-bookworm

RUN apt update && \
    apt install -y neovim tmux nginx && \
    apt clean

WORKDIR /root

COPY nginx.conf /etc/nginx/nginx.conf

COPY main.go .
