for f in $(gogrep -l "import \"shardmaster\" .)
do
		sed -i'' 's/import "shardmaster"/import "github.com\/chewr\/6.824-2016\/shardmaster"' $f
done
