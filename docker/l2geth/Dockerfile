FROM scrolltech/l2geth:ccc-0.0.6
# FROM scrolltech/l2geth:scroll-v3.3.1

ENV LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/lib/
ENV RUST_LOG="info"

RUN mkdir -p /l2geth/keystore

COPY password /l2geth/
COPY genesis.json /l2geth/
COPY genesis-keystore /l2geth/
COPY l2geth_run.sh /bin/

ENTRYPOINT ["sh", "/bin/l2geth_run.sh"]
