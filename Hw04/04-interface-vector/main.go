package main

import (
	"fmt"
)

type Vector3 interface {
	Add(v Vector3) Vector3
	Subtract(u Vector3) Vector3
	Multiply(scalar float64) Vector3
	Dot(u Vector3) Vector3
	Length() float64
	X() float64
	Y() float64
	Z() float64
}

var _ Vector3 = arrayVector3{}

type arrayVector3 [3]float64

func (a arrayVector3) X() float64 {
	return a[0]
}

func (a arrayVector3) Y() float64 {
	return a[1]
}

func (a arrayVector3) Z() float64 {
	return a[2]
}

// Сложение векторов и возврат нового вектора {a.X + u.X, a.Y + u.Y, a.Z + u.Z}
func (a arrayVector3) Add(v Vector3) Vector3 {
	// TODO implement me
	panic("implement me")
}

// Вычитаение векторов и возврат нового вектора {a.X - u.X, a.Y - u.Y, a.Z - u.Z}
func (a arrayVector3) Subtract(u Vector3) Vector3 {
	// TODO implement me
	panic("implement me")
}

// Умножение вектора на число. Умножьте кажду каоордианту на число и верните веткор
func (a arrayVector3) Multiply(scalar float64) Vector3 {
	// TODO implement me
	panic("implement me")
}

// Скалярное произведение векторов. Перемножьте координаты вектора v на координаты вектора u
// Пример a.x * u.x
func (a arrayVector3) Dot(u Vector3) Vector3 {
	// TODO implement me
	panic("implement me")
}

// Вычисление длины вектора math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
func (a arrayVector3) Length() float64 {
	// TODO implement me
	panic("implement me")
}

// Аналогично сделайте для вектора на основе структуры. Формулы аналогичные
var _ Vector3 = structVector3{}

type structVector3 struct {
	x float64
	y float64
	z float64
}

func (s structVector3) X() float64 {
	return s.x
}

func (s structVector3) Y() float64 {
	return s.y
}

func (s structVector3) Z() float64 {
	return s.z
}

func (s structVector3) Add(v Vector3) Vector3 {
	// TODO implement me
	panic("implement me")
}

func (s structVector3) Subtract(u Vector3) Vector3 {
	// TODO implement me
	panic("implement me")
}

func (s structVector3) Multiply(scalar float64) Vector3 {
	// TODO implement me
	panic("implement me")
}

func (s structVector3) Dot(u Vector3) Vector3 {
	// TODO implement me
	panic("implement me")
}

func (s structVector3) Length() float64 {
	// TODO implement me
	panic("implement me")
}

func printVec(v, v1 Vector3) {
	fmt.Println(v.Add(v1))
	fmt.Println(v.Subtract(v1))
	fmt.Println(v.Multiply(5))
	fmt.Println(v.Dot(v1))
	fmt.Println(v.Length())
	fmt.Println(v1.Length())
}

func main() {
	vec := arrayVector3{0: 1, 1: 1, 2: 1}
	vec1 := structVector3{
		x: 0,
		y: 0,
		z: 0,
	}
	printVec(vec, vec1)
}
