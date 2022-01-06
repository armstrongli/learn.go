package main

import (
	"fmt"
	"math/rand"
)

type Math = int
type English = int
type Chinese = int

func main() {
	var mathScore int = 100
	var mathScore2 Math = 100

	mathScore2 = mathScore
	fmt.Println(mathScore2)
	getScoresOfStudent("")

	vg := &voteGame{students: []*student{
		&student{name: fmt.Sprintf("%d", 1)},
		&student{name: fmt.Sprintf("%d", 2)},
		&student{name: fmt.Sprintf("%d", 3)},
		&student{name: fmt.Sprintf("%d", 4)},
		&student{name: fmt.Sprintf("%d", 5)},
	}}
	leader := vg.goRun()
	fmt.Println(leader)
}

func getScoresOfStudent(name string) (Math, Chinese, English) {
	// todo
	return 0, 0, 0
}

type voteGame struct {
	students []*student
}

func (g *voteGame) goRun() *Leader {
	// todo ....
	for _, item := range g.students {
		randInt := rand.Int()
		if randInt%2 == 0 {
			item.voteA(g.students[randInt%len(g.students)])
		} else {
			item.voteD(g.students[randInt%len(g.students)])
		}
	}

	maxScore := -1
	maxScoreIndex := -1
	for i, item := range g.students {
		if maxScore < item.agree {
			maxScore = item.agree
			maxScoreIndex = i
		}
	}
	if maxScoreIndex >= 0 { // 判断是否大于0，因为如果没有学生，那么index就是默认值-1.
		return g.students[maxScoreIndex]
	}
	return nil
}

type Leader = student

type student struct {
	name     string
	agree    int
	disagree int
}

func (std *student) voteA(target *student) {
	target.agree++
}

func (std *student) voteD(target *student) {
	target.disagree++
}
