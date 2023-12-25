#!/bin/bash

whoami
cd /srv || exit
touch .env
{
  echo HTTP_PORT=${{ secrets.HTTP_PORT }}
  HTTP_TOKEN_SECRET_KEY=${{ secrets.HTTP_TOKEN_SECRET_KEY }}
  PG_DSN=${{ secrets.PG_DSN }}
  S3_ACCESS_KEY_ID=${{ secrets.S3_ACCESS_KEY_ID }}
  S3_SECRET_ACCESS_KEY=${{ secrets.S3_SECRET_ACCESS_KEY }}

} >> .env