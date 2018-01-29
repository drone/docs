#!/bin/bash
set -euo pipefail

REGEX='s/   (--(\S* \S+)|(\S*))\s+(\b(.*?))? \B(\(default\: (.*?)\))? ?\B\[(\$.*?)\]/`$1$3` <br><br> `$8` | $5 <br><br> `$7 `/g'

docker run -it drone/drone:latest -h | perl -i -pe "$REGEX" | grep '<br>' > server.txt

docker run -it drone/agent:latest -h | perl -i -pe "$REGEX" | grep '<br>' > agent.txt
