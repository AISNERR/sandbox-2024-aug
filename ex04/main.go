package main

import (
	"bufio"
	"fmt"
	"os"
)

// prettify принимает интерфейс и удаляет пустые списки и словари рекурсивно.
func prettify(input interface{}) interface{} {
	switch v := input.(type) {
	case []interface{}:
		// Обрабатываем список
		newList := []interface{}{}
		for _, item := range v {
			if item != nil {
				// Рекурсивно очищаем элемент
				cleanedItem := prettify(item)
				if cleanedItem != nil {
					newList = append(newList, cleanedItem)
				}
			}
		}
		if len(newList) > 0 {
			return newList
		}
		return nil // Возвращаем nil для пустого списка

	case map[string]interface{}:
		// Обрабатываем словарь
		newMap := make(map[string]interface{})
		for key, value := range v {
			if value != nil {
				// Рекурсивно очищаем значение
				cleanedValue := prettify(value)
				if cleanedValue != nil {
					newMap[key] = cleanedValue
				}
			}
		}
		if len(newMap) > 0 {
			return newMap
		}
		return nil // Возвращаем nil для пустого словаря

	default:
		// Возвращаем все остальные типы данных как есть
		return input
	}
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var size int
		fmt.Fscan(in, &size)
		var node int
		nodes := make([]int, 0, size)
		for j := 0; j < size; j++ {
			fmt.Fscan(in, &node)
			nodes = append(nodes, node)
		}
		fmt.Println(prettify(nodes))
	}
}
