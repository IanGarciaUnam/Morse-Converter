package main

import (
    "fmt"
    "strings"
    "os"
    "io/ioutil"
    "log"
)
//Main for Main.go
/*Map designated to save the equivalences in both alphabets*/
var diccionario= map[string]string{"A":".-", "B":"-...", 
"C":"-.-.", "CH": "----", "D": "-..", "E":".", "F":"..-.",
"G":"--.", "H":"....", "I":"..", "J":".---", "K":"-.-", 
"L":".-..","M":"--", "N":"-.", "Ñ":"--.--", "O":"---", 
"P":".--.", "Q":"--.-", "R":".-.", "S":"...", "T":"-", "U":"..-",
"V":"...-", "W": ".--", "X":"-..-", "Y":"-.--", "Z":"--..",
"0":"-----", "1":".----", "2":"..---", "3":"...--", "4":"....-",
"5":".....", "6":"-....", "7":"--...", "8":"---..","9":"----.",
".":".-.-.-", ",":"--..--", "?":"..--..", "\"":".-..-."}
/*
*
*/



func main(){
	var files []string=reader_num_file()
	fmt.Print(files)
	checker(files)
	fmt.Println("Finish checker")
	apply_for_each_file(files)
}


func apply_for_each_file(files []string){
	for x := range files{

		if files[x]!= "main.go"{
		var toWrite= encrypt(files[x])
		write_File(files[x], toWrite)
	
		}
	}
}

func reader_num_file()[]string{
	var file []string =os.Args
	return file
}

/**
*Check if each file is correct or exists
*/
func checker(files []string){
	for x := range files{
		if _, err := os.Stat(files[x]); err==nil{
			print("---")
		}else if os.IsNotExist(err){
			print(files[x], "No Such file or directory")
		
		}
	}
}

/**
*Encrypt the file and return the message converted to Morse
*
*/
func encrypt(files string)string{
	fmt.Println("Encriptando")
		content, err := ioutil.ReadFile(files)
		if err != nil {
			log.Fatal(err)
		}
		/*Read file*/

	var knowledge string=from_arabic_to_morse(strings.ToUpper(string(content)))
	return knowledge

}
func check(e error){
	if e != nil{
		panic(e)
	}


}

func write_File(file string,content string){

f, err :=os.Create("m-"+file)
b:= []byte(content)
n2, err:= f.Write(b)
check(err)
fmt.Println("Wrote in " , n2)
defer f.Close()	


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