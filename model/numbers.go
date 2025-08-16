package model

import (
	"fmt"
	"sort"
	"strconv"
)

type NumbersSolution struct {
	Expression string `json:"expression"`
	Result     int    `json:"result"`
	Distance   int    `json:"distance"`
}

type NumbersResponse struct {
	Target    int               `json:"target"`
	Numbers   []int             `json:"numbers"`
	Solutions []NumbersSolution `json:"solutions"`
	Exact     bool              `json:"exact"`
}

func SolveNumbers(numbers []int, target int) NumbersResponse {
	bestSolutions := findSolutions(numbers, target)
	
	// Sort solutions by distance from target, then by result value
	sort.Slice(bestSolutions, func(i, j int) bool {
		if bestSolutions[i].Distance != bestSolutions[j].Distance {
			return bestSolutions[i].Distance < bestSolutions[j].Distance
		}
		return abs(bestSolutions[i].Result-target) < abs(bestSolutions[j].Result-target)
	})
	
	// Keep only the best solutions (same distance as the best)
	if len(bestSolutions) > 0 {
		bestDistance := bestSolutions[0].Distance
		var filtered []NumbersSolution
		for _, sol := range bestSolutions {
			if sol.Distance == bestDistance && len(filtered) < 10 {
				filtered = append(filtered, sol)
			}
		}
		bestSolutions = filtered
	}
	
	exact := len(bestSolutions) > 0 && bestSolutions[0].Distance == 0
	
	return NumbersResponse{
		Target:    target,
		Numbers:   numbers,
		Solutions: bestSolutions,
		Exact:     exact,
	}
}

func findSolutions(numbers []int, target int) []NumbersSolution {
	var solutions []NumbersSolution
	solutionSet := make(map[string]bool) // To avoid duplicates
	
	// Try all combinations of numbers and operations
	findSolutionsRecursive(numbers, target, "", 0, &solutions, &solutionSet)
	
	return solutions
}

func findSolutionsRecursive(numbers []int, target int, expression string, currentResult int, solutions *[]NumbersSolution, solutionSet *map[string]bool) {
	if len(numbers) == 0 {
		if expression != "" {
			distance := abs(currentResult - target)
			if distance <= 50 { // Only keep solutions within 50 of target
				sol := NumbersSolution{
					Expression: expression,
					Result:     currentResult,
					Distance:   distance,
				}
				
				// Check for duplicate expressions
				if !(*solutionSet)[expression] {
					(*solutionSet)[expression] = true
					*solutions = append(*solutions, sol)
				}
			}
		}
		return
	}
	
	// Try each remaining number
	for i, num := range numbers {
		remaining := make([]int, len(numbers)-1)
		copy(remaining, numbers[:i])
		copy(remaining[i:], numbers[i+1:])
		
		if expression == "" {
			// First number
			findSolutionsRecursive(remaining, target, strconv.Itoa(num), num, solutions, solutionSet)
		} else {
			// Try each operation
			operations := []struct {
				symbol string
				calc   func(int, int) (int, bool)
			}{
				{"+", func(a, b int) (int, bool) { return a + b, true }},
				{"-", func(a, b int) (int, bool) { return a - b, true }},
				{"*", func(a, b int) (int, bool) { return a * b, true }},
				{"/", func(a, b int) (int, bool) {
					if b != 0 && a%b == 0 {
						return a / b, true
					}
					return 0, false
				}},
			}
			
			for _, op := range operations {
				if result, valid := op.calc(currentResult, num); valid {
					newExpression := fmt.Sprintf("(%s %s %d)", expression, op.symbol, num)
					findSolutionsRecursive(remaining, target, newExpression, result, solutions, solutionSet)
				}
				
				// Also try the operation in reverse (for subtraction and division)
				if op.symbol == "-" || op.symbol == "/" {
					if result, valid := op.calc(num, currentResult); valid {
						newExpression := fmt.Sprintf("(%d %s %s)", num, op.symbol, expression)
						findSolutionsRecursive(remaining, target, newExpression, result, solutions, solutionSet)
					}
				}
			}
		}
	}
}

// Enhanced algorithm with better expression building
func findSolutionsEnhanced(numbers []int, target int) []NumbersSolution {
	type state struct {
		values      []int
		expressions []string
	}
	
	var solutions []NumbersSolution
	solutionSet := make(map[string]bool)
	
	// Initialize with the given numbers
	initial := state{
		values:      make([]int, len(numbers)),
		expressions: make([]string, len(numbers)),
	}
	copy(initial.values, numbers)
	for i, num := range numbers {
		initial.expressions[i] = strconv.Itoa(num)
	}
	
	queue := []state{initial}
	
	// BFS to try all combinations
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		
		// Check if any current value matches target or is close
		for i, val := range current.values {
			distance := abs(val - target)
			if distance <= 50 {
				sol := NumbersSolution{
					Expression: current.expressions[i],
					Result:     val,
					Distance:   distance,
				}
				
				if !solutionSet[current.expressions[i]] {
					solutionSet[current.expressions[i]] = true
					solutions = append(solutions, sol)
				}
			}
		}
		
		// Generate new states by combining pairs
		if len(current.values) > 1 {
			for i := 0; i < len(current.values); i++ {
				for j := i + 1; j < len(current.values); j++ {
					val1, val2 := current.values[i], current.values[j]
					exp1, exp2 := current.expressions[i], current.expressions[j]
					
					// Create new combinations
					newCombos := []struct {
						val int
						exp string
						valid bool
					}{
						{val1 + val2, fmt.Sprintf("(%s + %s)", exp1, exp2), true},
						{val1 - val2, fmt.Sprintf("(%s - %s)", exp1, exp2), true},
						{val2 - val1, fmt.Sprintf("(%s - %s)", exp2, exp1), true},
						{val1 * val2, fmt.Sprintf("(%s * %s)", exp1, exp2), true},
						{val1 / val2, fmt.Sprintf("(%s / %s)", exp1, exp2), val2 != 0 && val1%val2 == 0},
						{val2 / val1, fmt.Sprintf("(%s / %s)", exp2, exp1), val1 != 0 && val2%val1 == 0},
					}
					
					for _, combo := range newCombos {
						if combo.valid && combo.val > 0 { // Only positive results
							// Create new state with this combination
							newState := state{
								values:      make([]int, 0, len(current.values)-1),
								expressions: make([]string, 0, len(current.expressions)-1),
							}
							
							// Add the new combined value
							newState.values = append(newState.values, combo.val)
							newState.expressions = append(newState.expressions, combo.exp)
							
							// Add remaining values
							for k := 0; k < len(current.values); k++ {
								if k != i && k != j {
									newState.values = append(newState.values, current.values[k])
									newState.expressions = append(newState.expressions, current.expressions[k])
								}
							}
							
							if len(newState.values) > 0 {
								queue = append(queue, newState)
							}
						}
					}
				}
			}
		}
	}
	
	return solutions
}

func SolveNumbersEnhanced(numbers []int, target int) NumbersResponse {
	solutions := findSolutionsEnhanced(numbers, target)
	
	// Sort solutions by distance from target
	sort.Slice(solutions, func(i, j int) bool {
		if solutions[i].Distance != solutions[j].Distance {
			return solutions[i].Distance < solutions[j].Distance
		}
		return len(solutions[i].Expression) < len(solutions[j].Expression)
	})
	
	// Keep only the best solutions
	if len(solutions) > 0 {
		bestDistance := solutions[0].Distance
		var filtered []NumbersSolution
		for _, sol := range solutions {
			if sol.Distance == bestDistance && len(filtered) < 5 {
				filtered = append(filtered, sol)
			}
		}
		solutions = filtered
	}
	
	exact := len(solutions) > 0 && solutions[0].Distance == 0
	
	return NumbersResponse{
		Target:    target,
		Numbers:   numbers,
		Solutions: solutions,
		Exact:     exact,
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}