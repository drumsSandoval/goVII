package main

import (
	"fmt"
	"os"
	"sort"
)


func  main()  {
	opc := ""
	for opc != "4" {
		fmt.Printf("\t\tMenu\n\t1.- Ingresar strings\n" +
			"\t2.- Mostrar archivo Ascendente\n" +
			"\t3.- Mostrar archivo Descendente\n" +
			"\t4.- Salir\n" +
			"\topc: ")
		fmt.Scanln(&opc)
		switch opc {
			case "1":
				var n int
				fmt.Println(" Cantidad de  strings : ")
				fmt.Scanln(&n)
				data := make([] string, n)
				for i := 0; i<n; i++ {
					fmt.Printf("String %d: ",i+1)
					fmt.Scanln(&data[i])
				}
				sort.Strings(data)
				writeFile("ascendente.txt",data)
				sort.Sort(sort.Reverse(sort.StringSlice(data)))
				writeFile("descendente.txt", data)
			case "2":
				readFile("ascendente.txt")
			case "3":
				readFile("descendente.txt")
			case "4":
				fmt.Println("Adios uwu")
			default:
				fmt.Println("Opcion no valida")
		}
	}


}

func writeFile(name string, data [] string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Err")
		return
	}
	defer file.Close()
	for _, str := range data  {
		file.WriteString(str+"\n")
	}
}

func readFile(name string) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	total := stat.Size()
	bs := make([]byte, total)
	count, err := file.Read(bs)
	if err != nil {
		fmt.Println(err)
		return
	}
	str := string(bs)
	fmt.Println(str, "\nBytes: ", count)
}
