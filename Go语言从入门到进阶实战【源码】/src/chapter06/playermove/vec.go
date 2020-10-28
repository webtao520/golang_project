package main

import "math"

type Vec2 struct {
	X, Y float32
}

// 加
func (v Vec2) Add(other Vec2) Vec2 {

	return Vec2{
		v.X + other.X,
		v.Y + other.Y,
	}

}

// 减
func (v Vec2) Sub(other Vec2) Vec2 {

	return Vec2{
		v.X - other.X,
		v.Y - other.Y,
	}
}

// 乘
func (v Vec2) Scale(s float32) Vec2 {

	return Vec2{v.X * s, v.Y * s}
}

// 距离
func (v Vec2) DistanceTo(other Vec2) float32 {
	dx := v.X - other.X
	dy := v.Y - other.Y

	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

// 插值
func (v Vec2) Normalize() Vec2 {
	mag := v.X*v.X + v.Y*v.Y
	if mag > 0 {
		oneOverMag := 1 / float32(math.Sqrt(float64(mag)))
		return Vec2{v.X * oneOverMag, v.Y * oneOverMag}
	}

	return Vec2{0, 0}
}
