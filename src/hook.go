package dirp

import "fmt"

// PrintBashHook emits shell code for Bash, ZSH, sh, etc
func PrintBashHook() {
	fmt.Println(`function dir() {
		dir=$(dirp)
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
			set selection (dirp)
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
