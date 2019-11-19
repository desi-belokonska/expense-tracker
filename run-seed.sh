#!/bin/sh

sqlite3 db/expense-tracker.db < db/seed.sql
