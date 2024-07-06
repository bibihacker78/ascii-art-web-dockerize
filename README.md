*******   DESCRIPTION  ***********
Ascii-art-web-dockerize consiste à utiliser et paramétrer Docker :
Services et dépendances.
Conteneuriser une application.
Compatibilité/Dépendance.
Création d'images.


*************  AUTEURS  *************
  Obintou
  Bballaki
  Skaridja

************* COMMENT EXECUTER *******
docker image build -f Dockerfile -t "ascii-art-web-dockerize" .

echo "----------------------------------------------------------------"

docker images

echo "----------------------------------------------------------------"

docker container run -p 9090:8080 --detach --name ascii-art-web ascii-art-web-dockerize

echo "----------------------------------------------------------------"

docker ps -a

echo "----------------------------------------------------------------"

docker exec -it ascii-art-web /bin/bash
 