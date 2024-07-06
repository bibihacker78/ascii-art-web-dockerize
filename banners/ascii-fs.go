package asciifs

import (
	"bufio"
	"os"
	"strings"
)

func Asciifs(mot, banners string) string {

	file, err := os.Open("./banners/" + banners + ".txt") // Ouvre le fichier "standard.txt"
	if err != nil {
		// 	// Vérifie si une erreur est survenue pendant l'ouverture du fichier
		return "Error"
		// Retourne la fonction main
	}
	defer file.Close()           // fermeture automatique du fichier a
	var table []string           //déclare une variable de type slice de chaînes de caractères appelée "table"
	nb := bufio.NewScanner(file) //crée un nouveau scanner bufio qui sera utilisé pour lire le fichier file
	for nb.Scan() {              //Cette boucle for lit chaque ligne du fichier une par une à l'aide de la méthode Scan() du scanner bufio
		table = append(table, nb.Text())
	}
	input := mot                         //  récupère le premier argument passé à la ligne de commande et l'assigne à la variable "kadi".
	ouput := strings.Split(input, "\\n") //divise la chaîne de caractères "kadi" en plusieurs lignes
	chaine := AsciiArt(ouput, table)     //appelle la fonction "chiffrermot" avec sanogo et table
	return chaine
}
func AsciiArt(ouput []string, table []string) string { // crée une nouvelle instance de la structure de données "Builder"  "strings"
	var dok strings.Builder // commence une boucle qui itère à travers chaque élément du tableau "sanogo"
	for _, mots := range ouput {
		for j := 0; j < 8; j++ { // commence une boucle qui itère de 0 à 7.
			for _, caracters := range mots { //commence une boucle qui itère à travers chaque caractère de la variable "mots"
				bebe := 1 + (int(caracters)-32)*9 + j // calcule un indice dans le tableau "table" en utilisant la valeur ASCII du caractère, et l'indice  "j" de la boucle précédente.
				dok.WriteString(table[bebe])          // ajoute la chaîne de caractères située à l'indice "bebe" du tableau "table"
			}
			if mots != "" {
				dok.WriteString("\n")
			}
		}
		if mots == "" {
			dok.WriteString("\n")
		}
	}
	return dok.String()
}
