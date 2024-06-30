#! /bin/sh
cd ./sql/schema
goose postgres postgres://maxkopitz:@localhost:5432/blogator up

