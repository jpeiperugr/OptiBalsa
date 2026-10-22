package riego

const DiasEntreRecargas = 14

type PlanDistribucion struct {
	asignaciones [DiasEntreRecargas]AsignacionDiaria
}