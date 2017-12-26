package main

type PointQueue []Point3D

func (p *PointQueue) Push(n Point3D) {
	*p = append(*p, n)
}

func (p *PointQueue) Pop() (n Point3D) {
	n = (*p)[0]
	*p = (*p)[1:]
	return n
}
