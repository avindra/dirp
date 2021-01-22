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
		response=$(dirp $@)
		status=$?
		if [[ -n $response ]]; then
			if [[ $status -eq 2 ]]; then
				$EDITOR "$response"
				return $?
			fi

			echo "Switching to $response... "
			pushd "$response"
		fi
	}
	
	export -f dir`)
}

// PrintHook emits shell code for Fish
func PrintHook() {
	fmt.Println(`function dir
		set response (dirp $argv)
		if [ $status = 2 ]
			$EDITOR "$response"
			return $status
		end
	
		if [ "x$response" = "x" ]
			echo -n "How are we doing @ "
			uptime
			return $status
		end
	
		echo "Switching to $response"
		pushd "$selection"
	end`)
}
