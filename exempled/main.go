//Overriding Inner Types

// Programme d'exemple montrant ce qui se passe lorsque le type outer et le type inner
// implémentent la même interface.

package main

import (
	"fmt"
)

// notifier est une interface qui définit le comportement d'un type de notification.
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

// notify implémente une méthode qui peut être appelée via une valeur de type user.
func (u *user) notify() {
	fmt.Printf("Utilisateur : Envoi d'un e-mail utilisateur à %s<%s>\n",
		u.name,
		u.email)
}

// notify implémente une méthode qui peut être appelée via une valeur de type Admin.
func (a *admin) notify() {
	fmt.Printf("Utilisateur : Envoi d'un e-mail administrateur à %s<%s>\n",
		a.name,
		a.email)
}

// main est le point d'entrée de l'application.
func main() {
	// Créer un utilisateur administrateur.
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// Envoyer une notification à l'utilisateur administrateur.
	// L'implémentation du type inner intégré à l'interface n'est PAS "promue" vers le type outer.
	sendNotification(&ad)

	// Nous pouvons accéder directement à la méthode du type inner.
	ad.user.notify()

	// La méthode du type inner n'est PAS promue.
	ad.notify()
}

// sendNotification accepte des valeurs qui implémentent l'interface notifier
// et envoie des notifications.
func sendNotification(n notifier) {
	n.notify()
}
