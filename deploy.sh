#!/usr/bin/bash

if [ -f /usr/local/go/bin/go  ]
then

  /usr/local/go/bin/go install && sudo systemctl stop blog && sudo systemctl start blog
  echo "go install done"
  cd app && /usr/bin/npm run build && cd -
  echo "npm run build done"
fi
