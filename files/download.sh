#!/bin/bash
for i in {1..10}
do
   curl https://www.lipsum.com/feed/html -L > files/$i.html
done