package main

import "fmt"

type Cachorro struct {
	Raca  string
	Cor   string
	Patas int
}

type Gato struct {
	Cor   string
	Patas int
}

func (c Cachorro) Barulho() string {
	return "au au"
}

func (c Cachorro) NumeroDePatas() int {
	return c.Patas
}

func (g Gato) Barulho() string {
	return "miau"
}

func (g Gato) NumeroDePatas() int {
	return g.Patas
}

type Animal interface {
	Barulho() string
	NumeroDePatas() int
}

func FazBarulho(animal Animal) {
	fmt.Println(animal.Barulho())

}

type Pessoa struct {
	Nome              string
	Idade             int
	Prof              string
	TempoDeProffissao int
}

type Crianca struct {
	Nome  string
	Idade int
}

func (c Crianca) FalaBomdia() string {
    return c.Nome + "Deseja bom dia pra você"
}

func (p Pessoa) FalaBomdia() string {
	return fmt.Sprintf("%s deseja bom dia pra você", p.Nome)

}

func (p Pessoa) Profissao() string {
	return fmt.Sprintf("%s te, % d anos como %s", p.Nome, p.TempoDeProffissao, p.Prof)
}

type Adulto interface {
	FalaBomdia() string
	Profissao() string
}

func Bomdia(adulto Adulto) string {
    return adulto.FalaBomdia()
}


func main() {

	cachorro := Cachorro{
		Raca:  "spitz alemão",
		Cor:   "preto",
		Patas: 4,
	}

	gato := Gato{
		Cor:   "branco",
		Patas: 4,
	}

	FazBarulho(gato)
	FazBarulho(cachorro)


    adulto := Pessoa{
        Nome: "Luis",
        Idade: 21,
        Prof: "dev",
        TempoDeProffissao: 1,
    }

    crianca := Crianca{
        Nome: "bento",
        Idade: 4,
    }

    crianca.FalaBomdia()
    adulto.FalaBomdia()
    Bomdia(adulto)

}
