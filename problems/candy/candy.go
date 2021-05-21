package main

import "fmt"

func main()  {
	candy([]int{89, 54, 36, 54, 54, 99, 64})
}

type Children struct {
	Rating int `json:"rating"`
	Candy int `json:"candy"`
}


func candy(ratings []int) int {
	n := len(ratings)
	var childes []Children

	// Give all children 1 candy
	for _, rating := range ratings  {
		childes = append(childes, Children{
			Rating: rating,
			Candy:  1,
		})
	}

	// check if rating of current children is greater than next children, that current children will be given 1 candy
	// there are 2 directions, left -> right and left <- right
	for i := 0 ; i < n ; i ++ {
		if i + 1 <= n - 1 {
			if i == 0 {
				if childes[i].Rating > childes[i + 1].Rating {
					childes[i].Candy += 1
				}
			} else {
				if childes[i].Rating >= childes[i + 1].Rating  {
					if childes[i].Rating == childes[i + 1].Rating {
						if childes[i].Rating > childes[i - 1].Rating {
							childes[i].Candy = childes[i - 1].Candy
						}
					} else {
						if childes[i].Candy > childes[i - 1].Candy {
							childes[i].Candy = childes[i - 1].Candy + 1
						} else {
							childes[i].Candy = childes[i + 1].Candy + 1
						}
					}
				} else {
					if childes[i].Rating < childes[i + 1].Rating {
						childes[i + 1].Candy = childes[i].Candy + 1
					} else {
						childes[i + 1].Candy += 1
					}
				}
			}
		}
	}


	fmt.Println(childes)

	for j := n - 1 ; j > 0 ; j -- {
		if j - 1 > 0 {
			if j == n - 1 {
				if childes[j].Rating > childes[j - 1].Rating {
					childes[j].Candy = childes[j - 1].Candy + 1
				}
			} else {
				if childes[j].Rating >= childes[j - 1].Rating && childes[j].Rating >= childes[j + 1].Rating {
					if childes[j].Rating == childes[j - 1].Rating {
						if childes[j].Rating > childes[j + 1].Rating {
							childes[j].Candy = childes[j + 1].Candy
						}
					} else {
						if childes[j].Rating > childes[j - 1].Rating {
							if childes[j + 1].Rating > childes[j - 1].Rating {
								childes[j].Candy = childes[j + 1].Candy + 1
							} else {
								if childes[j - 1].Candy > childes[j + 1].Candy{
									childes[j].Candy = childes[j - 1].Candy + 1
								} else {
									childes[j].Candy = childes[j + 1].Candy + 1
								}
							}
						}
					}
				} else {
					if childes[j].Rating < childes[j -1].Rating {
						childes[j - 1].Candy = childes[j].Candy + 1
					}
				}
			}
		}
	}

	fmt.Println(childes)

	for j := n - 1; j >= 0; j -- {
		if j == 0 {
			if n > 2 {
				if childes[j + 1].Rating > childes[j].Rating && childes[j + 1].Rating > childes[j + 2].Rating {
					childes[j + 1].Candy = childes[j + 2].Candy + 1
				}
			}
		} else {
			if childes[j - 1].Rating > childes[j].Rating && childes[j - 1].Candy == childes[j].Candy {
				childes[j - 1].Candy = childes[j].Candy + 1
			}
		}
	}

	fmt.Println(childes)

	for i := 0; i < n; i ++ {
		if i == n - 1 {
			if n > 2 {
				if childes[i - 1].Rating > childes[i].Rating && childes[i - 1].Rating > childes[i - 2].Rating {
					childes[i - 1].Candy = childes[i - 2].Candy + 1
				}
			}
		} else {
			if childes[i + 1].Rating > childes[i].Rating && childes[i + 1].Candy == childes[i].Candy {
				childes[i + 1].Candy = childes[i].Candy + 1
			}
		}
	}

	fmt.Println(childes)

	totalCandy := 0

	for  _, children := range childes {
		totalCandy += children.Candy
	}

	return totalCandy
}