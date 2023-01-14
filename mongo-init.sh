#!/bin/bash
set -e

mongosh <<EOF
use admin
db.createUser(
  {
    user: "${DB_USER}",
    pwd: "${DB_PASSWORD}",
    roles: [
      {
        role: "readWrite",
        db: "${DB_NAME}"
      }
    ]
  }
);
EOF