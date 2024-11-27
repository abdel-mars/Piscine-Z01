curl -s https://learn.zone01oujda.ma/assets/superhero/all.json | \
jq -r 'map(select(.id == 70)) | .[0].name | "\""+ . + "\""' 

