package modelos

type Mujer struct {
	Hombre
}

func (m *Mujer) Respirar() {
	m.respirando = true
}
func (m *Mujer) Pensar() {
	m.respirando = true
}
func (m *Mujer) Comer() {
	m.respirando = true
}
func (m *Mujer) Sexo() string {
	return "Mujer"
}
