package main

import (
    "fmt"
    "math"
)

type Shape interface {
    Area() float64
    GetType() string
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) GetType() string {
    return "circle"
}

type Square struct {
    Side float64
}

func (s Square) Area() float64 {
    return s.Side * s.Side
}

func (s Square) GetType() string {
    return "square"
}

type Triangle struct {
    Base   float64
    Height float64
}

func (t Triangle) Area() float64 {
    return 0.5 * t.Base * t.Height
}

func (t Triangle) GetType() string {
    return "triangle"
}

func printAreas(shapes []Shape) {
    for _, shape := range shapes {
        fmt.Printf("Diện tích của %s: %.2f\n", shape.GetType(), shape.Area())
    }
}

func main() {
    shapes := []Shape{
        Circle{Radius: 2.0},
        Square{Side: 3.0},
        Triangle{Base: 4.0, Height: 5.0},
    }
    printAreas(shapes)
}