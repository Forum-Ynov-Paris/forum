# Étape 1 : Préparation des dépendances et compilation de l'application
FROM golang:1.20-bullseye AS builder

WORKDIR /usr/src/app

# Copie des fichiers nécessaires pour la gestion des dépendances (anecdotique car déjà fait dans l'ensemble de la copie)
COPY go.mod go.sum ./

# Installation des dépendances si nécessaire
RUN go mod download && go mod verify

# Copie de l'ensemble du code source de l'application
COPY . .

# Compilation de l'application
RUN go build -v -o ./app-forum .
RUN ls

# Étape 2 : Construction de l'image finale
#FROM ubuntu:20.04
#FROM golang:1.20-bullseye

#WORKDIR /app

# Copie de l'exécutable de l'étape précédente
#COPY --from=builder /usr/src/app/forum-app .

# Donner les permissions d'exécution à l'exécutable
#RUN chmod +x ./forum-app

# Exposition du port
EXPOSE 8080

# Commande pour démarrer l'application
CMD ["./app-forum"]
