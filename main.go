package main

import (
    "fmt"
    "strings"
)
//Main for Main.go
var diccionario= map[string]string{"A":".-", "B":"-...", 
"C":"-.-.", "CH": "----", "D": "-..", "E":".", "F":"..-.",
"G":"--.", "H":"....", "I":"..", "J":".---", "K":"-.-", 
"L":".-..","M":"--", "N":"-.", "Ñ":"--.--", "O":"---", 
"P":".--.", "Q":"--.-", "R":".-.", "S":"...", "T":"-", "U":"..-",
"V":"...-", "W": ".--", "X":"-..-", "Y":"-.--", "Z":"--..",
"0":"-----", "1":".----", "2":"..---", "3":"...--", "4":"....-",
"5":".....", "6":"-....", "7":"--...", "8":"---..","9":"----.",
".":".-.-.-", ",":"--..--", "?":"..--..", "\"":".-..-."}


func main(){
	var saludoMorse string= from_arabic_to_morse("HOLA PAPI")
	var saludoNoMorse string= from_morse_to_arabic(saludoMorse)
	fmt.Println(saludoMorse)
	fmt.Println("===================================")
	fmt.Println(saludoNoMorse)
}

func out_text(text []string) string{
	var outed string =strings.Join(text," ")
	return outed
}
/*
*Golang has not a function built to get the key from a value as java
*To solve this uniquely we will use an obvios solution, iterate each key
*to get each valuea and to inverse it 
*
*/
func find_key_from_value_in_dic(diccionario map[string]string)map[string]string{
	Diccionario_inverted := make(map[string]string)
	for key, element := range diccionario{
		Diccionario_inverted[element]=key
	}
	return Diccionario_inverted

}

func from_morse_to_arabic(texto string)string{
	var diccionario_local=find_key_from_value_in_dic(diccionario)
	var arreglo = strings.Split(texto,"/")
	var container=""
	for y := range arreglo{
		container=container+diccionario_local[arreglo[y]]+" "
	}
	return container
}
func from_arabic_to_morse(texto string)string{
	var arreglo = strings.Split(texto,"")
	var container string=""
	for x := range(arreglo){

		value, ok := diccionario[arreglo[x]]
		if ok {
			container=container+value+"/"
		}
	}
	return container
}