#!/bin/sh

cd themes/default/styles
fswatch -0 --recursive --exclude=".*" --include="\\.less$" . | xargs -0 -n1 lessc index.less ../static/css/index.css
