#!/usr/bin/env bash
cd /app/cmd/go-ambassador/
dlv debug --headless --log -l 0.0.0.0:2345 --api-version=2