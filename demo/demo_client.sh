#!/usr/bin/env bash

# Load the common demo lib
demo_dir="$(dirname "$0")"
. "$demo_dir/common.sh"

#####################################################################
# Start the good stuff

p  "# Show CLI usage"
p  ""
pe "goapi --help"
sleep $SLEEP_LONG
clear

p  "# Start the goapi server"
sleep $SLEEP_LONG
clear

p  "# Obtain a client auth token and save it to the configuration"
p  ""
pe "goapi auth login -u admin -p password"
sleep $SLEEP_SHORT
clear

p  "# View the current configuration"
p  ""
pe "goapi config list"
sleep $SLEEP_SHORT
clear

p  "# Create two new key/value pairs in the db"
p  ""
pe "goapi keyval create demokey1 demovalue1"
p  ""
pe "goapi keyval create demokey2 demovalue2"
sleep $SLEEP_SHORT
clear


p  "# Note: The CLI client uses the GRPC protocol"
p  "#   Demonstrate that the JSON REST interface works as well"
p  ""
pe 'TOKEN=$(goapi config get token)'
pe 'echo $TOKEN'
pe 'curl -X GET -k https://localhost:10080/v1/keyval/demokey1 -H "Content-Type: text/plain" -H "Authorization: Bearer $TOKEN"'
p  "\n"
p  "# Here's the equivalent client using the GRPC protocol"
pe "goapi keyval read demokey1"
sleep $SLEEP_LONG
clear

p  "# Update a key from the db"
p  ""
pe "goapi keyval update demokey1 demovalue1_modified"
p  ""
p  "# Read back the updated key from the db (should succeed)"
p  ""
pe "goapi keyval read demokey1"
sleep $SLEEP_SHORT
clear

p  "# Delete a key from the db"
p  ""
pe "goapi keyval delete demokey2"
p  ""
p  "# Read back the deleted key from the db, (should fail)"
p  ""
pe "goapi keyval read demokey2"
sleep $SLEEP_SHORT
clear
