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
		stdout=$(dirp $@)
		status=$?
		if [[ -n $stdout ]]; then
			if [[ $status -eq 2 ]]; then
				$EDITOR "$stdout"
				return $?
			fi

			echo "Switching to $stdout... "
			pushd "$stdout"
		fi
	}
	
	export -f dir`)
}

// PrintHook emits shell code for Fish
func PrintHook() {
	fmt.Println(`function dir
		set stdout (dirp $argv)
		if [ $status = 2 ]
			$EDITOR "$stdout"
			return $status
		end
	
		if [ "x$stdout" = "x" ]
			echo -n "How are we doing @ "
			uptime
			return $status
		end
	
		echo "Switching to $stdout"
		pushd "$stdout"
	end`)
}
