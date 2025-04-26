function dir
    set stdout (dirp $argv)
    if [ $status = 2 ]
        $EDITOR "$stdout"
        return $status
    end

    if [ "x$stdout" = x ]
        echo -n "How are we doing @ "
        uptime
        return $status
    end

    echo "Switching to $stdout"
    pushd "$stdout"
end

abbr -a -g d dir
abbr -a -g d. dir .
abbr -a -g d.. dir ..
