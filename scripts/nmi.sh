#! /bin/bash

. ./prelude.sh

query_url='https://secure.networkmerchants.com/api/query.php'
transact_url='https://secure.networkmerchants.com/api/transact.php'

get() {
	curl --data "security_key=$NMI_TOKEN" $query_url
}

new_customer() {
	file=$dir/data/signUpQueue.txt
	while read -r line 
	do
		params=`echo $line | awk '{print $6}'`
		curl -d "security_key=$NMI_TOKEN&customer_vault=add_customer&recurring=add_subscription&$params" $transact_url
	done < $file
}

$*
