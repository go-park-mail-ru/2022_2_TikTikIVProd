#!/bin/bash
for i in {1..30}
do
   curl https://picsum.photos/800/400 -L > attachments/$i.png
done