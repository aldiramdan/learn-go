package src

import "fmt"

func FlightLoop(flight_len int, movie_len []int) (result string) {
	if flight_len <= 0 {
		return ("flight duration is required!")
	}

	var n_movie int = len(movie_len)
	for i := 0; i < n_movie; i++ {
		for j := 0; j < n_movie; j++ {
			if movie_len[i]+movie_len[j] == flight_len {
				fmt.Printf("Movie : %d and Movie : %d", movie_len[i], movie_len[j])
				return
			}
		}
	}
	return ("not found")
}

func FlightRange(flight_len int, movie_len []int) (result string) {
	if flight_len <= 0 {
		return ("flight duration is required!")
	}

	for _, v := range movie_len {
		for _, j := range movie_len {
			if v+j == flight_len {
				fmt.Printf("Movie : %d and Movie : %d", v, j)
				return
			}
		}
	}
	return ("not found")
}
