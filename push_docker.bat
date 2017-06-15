SET VERSION=v0.0.1
SET REPO=godevops/gen-linq:%VERSION%

docker build -t gen:%VERSION% . ^
 && docker tag gen:%VERSION% %REPO% ^
 && docker push %REPO%