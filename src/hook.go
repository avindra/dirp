package dirp

import (
	"fmt"
	"strings"
)

// PrintHook emits shell code for Bash, ZSH, sh, BusyBox, etc
func PrintHook() {
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

// PrintFishHook emits shell code for Fish
func PrintFishHook() {
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

// PrintRcHook emits code for rc, the plan 9 shell
func PrintRcHook() {
	src := `fn dir {
		stdout=` + "`" + `{dirp $*};
		if (~ $bqstatus 2 ) {
			$EDITOR $stdout;
			return $status;
		};
	
		if (~ "x$stdout" "x" ) {
			echo -n How are we doing @;
			uptime;
			return $status;
		};
	
		echo Switching to $stdout;
		cd $stdout
	}`

	fmt.Println(strings.ReplaceAll(src, "\t", "  "))
}

// PrintEsHook emits code for es, a shell based on rc
func PrintEsHook() {
	fmt.Println(`fn dir {
	stdout=` + "`" + `{dirp $*};
	if {~ $bqstatus 2 } {
		$EDITOR $stdout;
		return $status;
	};
	
	if {~ "x$stdout" "x" } {
		echo -n "How are we doing @";
		uptime;
		return $status;
	};
	
	echo Switching to $stdout; 
	cd $stdout
}`)
}
