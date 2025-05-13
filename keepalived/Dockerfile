FROM debian:bullseye-slim

RUN apt update && \
    apt install -y keepalived iproute2 iputils-ping neovim && \
    apt clean

COPY keepalived.conf /etc/keepalived/keepalived.conf

CMD ["keepalived", "--dont-fork", "--log-console"]
