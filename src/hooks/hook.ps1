Remove-Alias dir

function dir {
	$dir = dirp @args

	if ($LASTEXITCODE -eq 2) {
		notepad $dir
	}

	pushd $dir
}


Set-Alias -Name d -Value dir
#Set-Alias -Name d. -Value dir .
#Set-Alias -Name d.. -Value dir ..