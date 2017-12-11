package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	gente := [3]string{"Santiago Doc", "Cleri Gorun", "Romi Goget"}
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	perm := random.Perm(len(gente))
	for i, v := range perm {
		fmt.Printf( "%s → %s\n", gente[v], gente[i] );
	}
}
