#!/bin/bash

TIME="10"
URL="https://api.telegram.org/bot$TELEGRAM_TOKEN/sendMessage"
TEXT="Deploy the status: $1 \n Project: + $CI_PROJECT_NAME  \n URL: + $ CI_PROJECT_URL / Pipelines / $CI_PIPELINE_ID / \n Branch: + $CI_COMMIT_REF_SLUG "

curl -s --max-time $TIME -d "chat_id=$TELEGRAM_CHATID&disable_web_page_preview=1&text=$TEXT"  $URL > /dev/null