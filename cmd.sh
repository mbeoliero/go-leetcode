# [] => {}
cat temp.txt | sed 's/\[/\{/g;s/\]/\}/g'