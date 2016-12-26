package main

import (
	"log"
	"bufio"
	"os"
	"strconv"
	"fmt"
	"strings"
	"math"
)

func main() {
	log.Println("Hackerrank: Bot saves princess.")

	scanner := bufio.NewScanner(os.Stdin)
	var n int
	var size int
	var grid []string
	var err error
	for scanner.Scan() {
		if n == 0 {
			size, err = strconv.Atoi(scanner.Text())
			if err != nil {
				log.Println("First line not convertable:", err.Error())
				os.Exit(1)
			}
		} else {
			grid = append(grid, scanner.Text())
		}
		n++
	}
	if err := scanner.Err(); err != nil {
		log.Println("error reading standard input:", err.Error())
		os.Exit(1)
	}

	// Find path
	displayPathtoPrincess(size, grid)


}

func displayPathtoPrincess(size int, grid []string)  {

	// Integrity checks

	// The first line contains an odd integer N (3 <= N < 100) denoting the size of the grid.
	if size % 2 == 0 {
		log.Printf("Grid size %v is an even number, and therefore out of bounds.", size)
		os.Exit(1)
	}
	if size > 100 || size < 3 {
		log.Printf("Grid size n = %v  (3 <= n < 100) is out of bounds.", size)
		os.Exit(1)
	}

	// check that grid is the size indicated
	log.Printf("Expecting a %vx%v grid", size, size)
	log.Printf("Grid line count: %v", len(grid))
	if len(grid) != size {
		log.Println("Grid not what was expected.")
		os.Exit(1)
	}
	log.Printf("%+v", grid)

	// find bot, find princess, verify grid
	pok, bok := false, false // haven't found either princess or bot yets
	var pm, pn, bm, bn int  // m x n  grid
	for k, v := range grid {
		chars := strings.Split(v, "")
		if strings.Count(v, "m") == 1 {
			if bok == true {
				log.Printf("More than one bot found, previously at %v,%v and now at %v,%v",
					bm, bn, k, strings.Index(v, "m"))
				os.Exit(1)
			}
			bm = k
			bn = strings.Index(v, "m")
			bok = true
		}
		if strings.Count(v, "p") == 1 {
			if pok == true {
				log.Printf("More than one princess found, previously at %v,%v, and at %v,%v",
					pm, bn, k, strings.Index(v, "p"))
				os.Exit(1)
			}
			pm = k
			pn = strings.Index(v, "p")
			pok = true
		}
		if len(chars) != size {
			log.Println("This row is sized improperly", v)
			os.Exit(1)
		}

	}
	if pok {
		log.Printf("Princess found at %v,%v", pm, pn)
	} else {
		log.Println("Error finding Princess, no or multiple Princesses")
		os.Exit(1)
	}
	if bok {
		log.Printf("Bot found at %v,%v", bm, bn)
	} else {
		log.Println("Error finding Bot: No bot or multiple bots.")
		os.Exit(1)
	}

	// target
	tm := pm - bm
	tn := pn - bn

	moves := math.Abs(float64(tm)) + math.Abs(float64(tn))

	log.Printf("Bot should move %v, %v to get to Princess.", tm, tn)
	log.Printf("Bot will rescue Princess in %v moves.", moves)

	if tm > 0 {
		for i:= 0 ; i < tm; i++ {
			fmt.Println("DOWN")
		}
	} else if tm < 0 {
		for i :=0 ; i > tm; i-- {
			fmt.Println("UP")
		}
	}

	if tn > 0 {
		for i := 0; i < tn; i++ {
			fmt.Println("RIGHT")
		}
	} else if tn < 0 {
		for i := 0; i > tn; i-- {
			fmt.Println("LEFT")
		}
	}
}
