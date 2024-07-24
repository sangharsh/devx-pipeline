#!/bin/zsh

if [[ -v INVENTORY_SERVICE_URL ]]; then
    echo "INVENTORY_SERVICE_URL is ${INVENTORY_SERVICE_URL}"
else
    echo "Please set INVENTORY_SERVICE_URL before using this script"
fi

echo "\nCalling /statusz"
curl -s ${INVENTORY_SERVICE_URL}/statusz | jq
echo "\nAdding inventory"
curl --silent -X POST -d '{"product_name": "carz", "quantity": 13}' ${INVENTORY_SERVICE_URL}/inventory | jq
echo "\nFetching inventory"
curl -s ${INVENTORY_SERVICE_URL}/inventory | jq
