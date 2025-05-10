package main

import (
	"fmt"
	"mymodule/experiments/survey_practice/Assessment"
)

func main() {
	{
		// Assessment should be open after creation
		var assessment, _ = Assessment.New("term 1", 0.3)
		if assessment.IsOpen() {
			fmt.Println("La Asignatura no tiene una calificación asignada")
		}
	}

	{
		// Assessment should be closed after assigning an score
		var assessment, _ = Assessment.New("term 1", 0.3)
		assessment.AssignScore(8.0)
		if assessment.IsClosed() {
			fmt.Printf("La Asignatura tiene una calificación asignada: %f \n", assessment.UnwrapScore())
		}
	}
}
