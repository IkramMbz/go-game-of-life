# Game of Life

Ce projet implémente le célèbre jeu de la vie de Conway en utilisant la bibliothèque Ebiten pour Go. Le jeu de la vie est un automate cellulaire, un modèle mathématique utilisé pour simuler l'évolution d'une population de cellules.

## Fonctionnalités

- Grille personnalisable avec des dimensions carrées définies par l'utilisateur.
- Simulation du jeu de la vie avec des règles classiques.
- Contrôle de la vitesse de la simulation avec un slider.
- Zoom et déplacement dans la grille avec la souris et les touches fléchées.
- Sauvegarde automatique (Touche S) et chargement de la dernière partie (Touche L).

## Prérequis

- [Go](https://golang.org/doc/install) (version 1.16 ou ultérieure)
- [Ebiten](https://ebiten.org/documents/install.html) (version 2.1.0 ou ultérieure)

## Installation


1. Installez les dépendances :
    ```sh
    go get -u github.com/hajimehoshi/ebiten/v2
    ```

## Utilisation

1. Lancez le jeu :
    ```sh
    go run main.go utils.go game.go
    ```

    or

    ```sh
    go run .
    ```

2. Suivez les instructions à l'écran pour entrer les dimensions de la grille entre 15 et 60.

## Commandes

- **Molette de la souris** : Zoom avant/arrière.
- **S** : Sauvegarder l'état actuel du jeu.
- **L** : Charger la dernière partie.
- **Clique gauche** : Activer une cellule et déplace une celulle.

## Structure du projet

- `main.go` : Point d'entrée du programme. Configure la fenêtre de jeu et initialise la simulation.
- `utils.go` : Contient la définition du slider utilisé pour ajuster la vitesse de la simulation.
- `game.go` : Contient la logique principale du jeu de la vie, y compris le rendu, la mise à jour de l'état du jeu, et la gestion des sauvegardes.


## Contributeurs
Lionel NGOLO / Ben MEITE / Ikram MBECHEZI
