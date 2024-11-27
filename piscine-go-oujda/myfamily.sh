URL="https://learn.zone01oujda.ma/assets/superhero/all.json"
 curl -s "$URL" | jq -r --arg id "$HERO_ID" '.[] | select(.id == ($id | tonumber)) | .connections.relatives | gsub ("\\n";"\\n")'
