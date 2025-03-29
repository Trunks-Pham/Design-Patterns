package main

import (
    "fmt"
    "math"
)

type Shape struct {
    Type   string
    Radius float64
    Side   float64
    Base   float64 
    Height float64 
    Area   float64 
}

func (s *Shape) CalculateArea() {
    switch s.Type {
    case "circle":
        s.Area = math.Pi * s.Radius * s.Radius
    case "square":
        s.Area = s.Side * s.Side
    case "triangle":
        s.Area = 0.5 * s.Base * s.Height
    }
}

func printAreas(shapes []Shape) {
    for _, shape := range shapes {
        if shape.Type == "circle" {
            fmt.Printf("Diện tích: %.2f\n", math.Pi*shape.Radius*shape.Radius)
        } else if shape.Type == "square" {
            fmt.Printf("Diện tích: %.2f\n", shape.Side*shape.Side)
        }
    }
}

func main() {
    shapes := []Shape{
        {Type: "circle", Radius: 2.0},
        {Type: "square", Side: 3.0},
        {Type: "triangle", Base: 4.0, Height: 5.0}, 
    }

    for i := range shapes {
        shapes[i].CalculateArea()
    }

    for _, shape := range shapes {
        if shape.Type == "triangle" {
            fmt.Printf("Diện tích: %.2f\n", shape.Area)
        }
    }

    printAreas(shapes)
}