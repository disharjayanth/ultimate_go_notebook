package main

import "fmt"

var matrix [][]int = [][]int{
	{0xFF, 0xFF, 0xFF, 0xFF},
	{0xFF, 0xFF, 0xFF, 0xFF},
	{0xFF, 0xFF, 0xFF, 0xFF},
	{0xFF, 0xFF, 0xFF, 0xFF},
}

type linkedList struct {
	v int
	p *linkedList
}

var list *linkedList = &linkedList{
	v: 0xFF,
	p: &linkedList{
		v: 0xFF,
		p: &linkedList{
			v: 0xFF,
			p: &linkedList{
				v: 0xFF,
				p: &linkedList{
					v: 0xFF,
					p: &linkedList{
						v: 0xFF,
						p: &linkedList{
							v: 0xFF,
							p: &linkedList{
								v: 0xFF,
								p: &linkedList{
									v: 0xFF,
									p: &linkedList{
										v: 0xFF,
										p: &linkedList{
											v: 0xFF,
											p: &linkedList{
												v: 0xFF,
												p: &linkedList{
													v: 0xFF,
													p: &linkedList{
														v: 0xFF,
														p: &linkedList{
															v: 0xFF,
															p: &linkedList{
																v: 0xFF,
																p: nil,
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},
}

func RowTraverse() int {
	var ctr int
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			if matrix[row][col] == 0xFF {
				ctr++
			}
		}
	}
	return ctr
}

func ColTraverse() int {
	var ctr int
	for col := 0; col < 4; col++ {
		for row := 0; row < 4; row++ {
			if matrix[row][col] == 0xFF {
				ctr++
			}
		}
	}
	return ctr
}

func LinkedListTraverse() int {
	var ctr int
	d := list
	for d != nil {
		if d.v == 0xFF {
			ctr++
		}
		d = d.p
	}
	return ctr
}

func main() {
	fmt.Println("Row Traverse:", RowTraverse())
	fmt.Println("Col Traverse:", ColTraverse())
	fmt.Println("LinkedList Traverse:", LinkedListTraverse())
}
