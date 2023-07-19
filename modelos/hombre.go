package modelos

type Hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
}

func (h *Hombre) Respirar() {
	h.respirando = true
}
func (h *Hombre) Pensar() {
	h.respirando = true
}
func (h *Hombre) Comer() {
	h.respirando = true
}
func (h *Hombre) Sexo() string {
	return "Hombre"
}
