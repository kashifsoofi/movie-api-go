FROM amacneil/dbmate

WORKDIR /db

# Copy all db files
COPY . .

ENTRYPOINT [ "/usr/local/bin/dbmate", "--no-dump-schema", "--wait", "up"]
