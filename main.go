package main

func main() {

}

func cachback(ammount int) int {
	const bound = 3000
	const percent = 5
	result := 0
	if ammount >= 3000{
		result = ammount * percent / 100
	}
	return result
}
