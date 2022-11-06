#!/bin/bash
for i in {2..30}
do
   curl https://picsum.photos/800/400 -L > images/$i.png
done