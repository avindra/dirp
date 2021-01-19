package dirp

import "fmt"

// PrintBashHook emits shell code for Bash, ZSH, sh, etc
func PrintBashHook() {
	// the first part removes existing "dir" aliases, if any
	fmt.Println(`
alias dir &> /dev/null
if [[ $? -eq 0 ]]; then
	unalias dir
fi

function dir() {
		dir=$(dirp $@)
		if [[ -n $dir ]]; then
			echo "Switching to $dir... "
			pushd "$dir"
		fi
	}
	
	export -f dir`)
}

// PrintHook emits shell code for Fish
func PrintHook() {
	fmt.Println(`function dir
		if [ "$argv[1]" = "cfg" ]
			$EDITOR "$HOME/.config/dir/list"
			return $status
		else
			# default
			set selection (dirp $argv)
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
