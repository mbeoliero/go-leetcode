# [] => {}
cat temp.txt | sed 's/\[/\{/g;s/\]/\}/g'

# ls -l | grep ".go" | wc -l