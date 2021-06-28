package main

import "fmt"
import "sync"

type tree struct {
	value int
	left *tree
	right *tree
	branchId string
}

func findValueInTree(tree *tree, searchValue int) chan string {
	out := make(chan string)

	if tree == nil {
		return out
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		wg.Add(1)
		go searchWorker(tree, &wg, out, searchValue)

		wg.Done()
	}()

	go func() {
		// defer close(out)
		wg.Wait()
	}()

	return out
}

func searchWorker(treeBranch *tree, wg *sync.WaitGroup, out chan <- string, searchValue int)  {
	if treeBranch == nil {
		wg.Done()
		return
	}

	if treeBranch.value == searchValue {
		fmt.Println("found in ", treeBranch.branchId)
		out <- treeBranch.branchId
	}

	wg.Add(1)

	go func() {
		wg.Add(2)
		go searchWorker(treeBranch.left, wg, out, searchValue)
		go searchWorker(treeBranch.right, wg, out, searchValue)

		wg.Done()
	}()

	wg.Done()
}

func readChannel(in chan string, closeOnFirstFinding bool) <-chan struct{} {
	done := make(chan struct{})

	go func() {
		defer close(done)
		for data := range in {
			fmt.Println(data)
			if closeOnFirstFinding {
				close(in)
			}
		}
	}()

	return done
}

func main() {
	tree := tree {
		8,
		&tree {
			8,
			&tree {
				8,
				nil,
				nil,
				"BA",
			},
			&tree {
				8,
				&tree {
					8,
					&tree {
						101,
						nil,
						nil,
						"BBAA",
					},
					&tree {
						8,
						nil,
						nil,
						"BBAB",
					},
					"BBA",
				},
				nil,
				"BB",
			},
			"B",
		},
		&tree {
			8,
			&tree {
				101,
				nil,
				nil,
				"CA",
			},
			&tree {
				8,
				&tree {
					101,
					nil,
					nil,
					"in88",
				},
				&tree {
					81,
					&tree {
						101,
						nil,
						nil,
						"CB 101",
					},
					nil,
					"in77",
				},
				"CB",
			},
			"C",
		},
		"Top",
	}

	closeOnFirstFinding :=  true

	in := findValueInTree(&tree, 101)

	<-readChannel(in, closeOnFirstFinding)
}

