#!/bin/sh

sqlite3 db/expense-tracker-test.db < db/scripts/seed.sql
