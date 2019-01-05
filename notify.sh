#!/bin/bash

TIME="10"
URL="https://api.telegram.org/bot$TELEGRAM_TOKEN/sendMessage"
TEXT="Deploy the status: $1 %0A%0A Project: + $CI_PROJECT_NAME  %0A%0A URL: + $CI_PROJECT_URL / Pipelines / $CI_PIPELINE_ID / %0A%0A Branch: + $CI_COMMIT_REF_SLUG "

curl -s --max-time $TIME -d "chat_id=$TELEGRAM_CHATID&disable_web_page_preview=1&text=$TEXT"  $URL > /dev/null