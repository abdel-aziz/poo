//Interfaces

// Programme d'exemple montrant comment utiliser une interface en Go. Dans ce cas,
// un pointeur est utilisé pour prendre en charge l'appel de l'interface.
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

// notify implémente une méthode qui peut être appelée via
// une valeur de type user.
func (u *user) notify() {
	fmt.Printf("utilisateur : Envoi de l'e-mail utilisateur à %s<%s>\n",
		u.name,
		u.email)
}

// main est le point d'entrée de l'application.
func main() {
	// Crée une valeur de type user.
	u := &user{"Jill", "jill@email.com"}

	// Passe un pointeur de type user à la fonction.
	sendNotification(u)
}

// sendNotification accepte des valeurs qui implémentent l'interface notifier
// et envoie des notifications.
func sendNotification(notify notifier) {
	notify.notify()
}
