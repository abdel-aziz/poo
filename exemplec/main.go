// Extending Types

// Programme d'exemple montrant comment les types intégrés fonctionnent avec les interfaces.

package main

import (
	"fmt"
)

// notifier est une interface qui définit un comportement de type de notification.
type notifier interface {
	notify()
}

// user définit un utilisateur dans le programme.
type user struct {
	name  string
	email string
}

// admin représente un utilisateur administrateur avec des privilèges.
type admin struct {
	user
	level string
}

// notify implémente une méthode qui peut être appelée via
// une valeur de type user.
func (u *user) notify() {
	fmt.Printf("utilisateur : Envoi de l'e-mail utilisateur à %s<%s>\n",
		u.name,
		u.email)
}

// main est le point d'entrée de l'application.
func main() {
	// Crée un utilisateur administrateur.
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// Envoie à l'utilisateur administrateur une notification.
	// L'implémentation du type inner intégré de l'interface est "promue" vers le type externe.
	sendNotification(&ad)
}

// sendNotification accepte des valeurs qui implémentent l'interface notifier
// et envoie des notifications.
func sendNotification(n notifier) {
	n.notify()
}
