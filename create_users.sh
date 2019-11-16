#!/bin/bash

containerId=$(docker ps | grep mongo |awk  '{print $1}')
VAR=""

start=1
size_lote=2000
size_acc=$size_lote
total=0
lots=50

for z in $(seq 1 $lots);do
	echo $start"===="$size_acc
		VAR=""
		for x in $(seq $start $size_acc);do
			total=$(($total+1))
			VAR+="db.users.insert({'name':'Teste$x', 'age':$x, 'id':$x});"
		done
	start=$(($total+1))
	size_acc=$(($total+$size_lote))
	docker exec -it $containerId mongo poc --eval "$VAR"
done
