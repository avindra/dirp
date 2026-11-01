@doskey d = @( FOR /F "tokens=*" %%i IN ('dirp') DO pushd %%i)
