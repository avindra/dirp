package dir

import "fmt"

// PrintHook emits shell code to wire up functionality
func PrintHook() {
	fmt.Println(`function dir
		set prog "$HOME/bin/lib/dir"
	
		if [ "$argv[1]" = "cfg" ]
			$EDITOR "$HOME/.config/dir/list"
			return $status
		else if test -d "$argv[1]"
			# driller feature
			set selection (find "$argv[1]" -maxdepth 1 -type d | $prog)
		else
			# default
			set selection ($prog)
		end
	
		if [ "x$selection" = "x" ]
			echo -n "How are we doing @ "
			uptime
			return $status
		end
	
		echo "Switching to $selection"
		pushd "$selection"
	end`)
}
