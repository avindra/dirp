package dirp

import "fmt"

// PrintBashHook emits shell code for Bash, ZSH, sh, BusyBox, etc
func PrintBashHook() {
	// 1) Remove existing "dir" aliases, if any exit
	// 2) Detect and prefer pushd over cd
	// 3) Provide "dir" function
	// 4) export dir function for Bash users
	fmt.Println(`
alias dir &> /dev/null
if [[ $? -eq 0 ]]; then
	unalias dir
fi

_DIRP_CD=cd
type pushd &> /dev/null && _DIRP_CD=pushd

function dir() {
		stdout=$(dirp $@)
		status=$?
		if [[ -n $stdout ]]; then
			if [[ $status -eq 2 ]]; then
				$EDITOR "$stdout"
				return $?
			fi

			echo "Switching to $stdout... "
			$_DIRP_CD "$stdout"
		fi
	}
	
	export -f dir &> /dev/null`)
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
