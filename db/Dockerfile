FROM amacneil/dbmate

WORKDIR /db

# Copy all db files
COPY . .

ENTRYPOINT [ "/usr/local/bin/dbmate", "-d", "./migrations", "--no-dump-schema", "--wait", "up"]
