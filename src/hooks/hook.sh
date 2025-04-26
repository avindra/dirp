# 1) Remove existing "dir" aliases, if any exist
# 2) Detect and prefer pushd over cd
# 3) Provide "dir" function
# 4) export dir function for Bash users

# important: do not use $status as a variable
# because it's a readonly reserved special var in ZSH

unalias dir >/dev/null 2>&1

_DIRP_CD=cd
type pushd >/dev/null 2>&1 && _DIRP_CD=pushd

dir () {
	stdout=$(dirp $@)
	stat=$?
	if [ -n $stdout ]; then
		if [ $stat -eq 2 ]; then
			$EDITOR "$stdout"
			return $?
		fi

		echo "Switching to $stdout... "
		$_DIRP_CD "$stdout"
	fi
}

alias d=dir
alias d.="dir ."
alias d..="dir .."

export -f dir >/dev/null 2>&1