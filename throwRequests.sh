#!/bin/bash
wget --header='Accept-Encoding:'$1 --header='X-Allow-Logs: '$2 -O/dev/null -q http://localhost:8091/
