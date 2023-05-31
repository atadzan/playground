FROM minio/minio

ENV MINIO_ACCESS_KEY=hello
ENV MINIO_SECRET_KEY=world

COPY configure_minio.sh /usr/local/bin/configure_minio.sh
COPY custom-policy.json /usr/local/bin/custom-policy.json
COPY custom-policy.json /usr/local/bin/read-only-policy.json

RUN chmod +x /usr/local/bin/configure_minio.sh

RUN /usr/local/bin/configure_minio.sh

CMD ["minio", "server", "/data", "--console-address", ":9090"]