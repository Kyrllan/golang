package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type timeStruct struct {
	id         uint
	nome       string
	pontos     uint
	jogos      uint
	vitorias   uint
	empates    uint
	derrotas   uint
	saldo      int
	golsPro    int
	golsContra int
}

type jogo struct {
	mandante      *timeStruct
	visitante     *timeStruct
	golsMandante  int
	golsVisitante int
}

func main() {
	times := []*timeStruct{
		{1, "Atlético-GO", 0, 0, 0, 0, 0, 0, 0, 0},
		{2, "Atlético-MG", 0, 0, 0, 0, 0, 0, 0, 0},
		{3, "Athletico-PR", 0, 0, 0, 0, 0, 0, 0, 0},
		{4, "Bahia", 0, 0, 0, 0, 0, 0, 0, 0},
		{5, "Botafogo", 0, 0, 0, 0, 0, 0, 0, 0},
		{6, "Bragantino", 0, 0, 0, 0, 0, 0, 0, 0},
		{7, "Corinthians", 0, 0, 0, 0, 0, 0, 0, 0},
		{8, "Cruzeiro", 0, 0, 0, 0, 0, 0, 0, 0},
		{9, "Cuiabá", 0, 0, 0, 0, 0, 0, 0, 0},
		{10, "Flamengo", 0, 0, 0, 0, 0, 0, 0, 0},
		{11, "Fluminense", 0, 0, 0, 0, 0, 0, 0, 0},
		{12, "Fortaleza", 0, 0, 0, 0, 0, 0, 0, 0},
		{13, "Grêmio", 0, 0, 0, 0, 0, 0, 0, 0},
		{14, "Internacional", 0, 0, 0, 0, 0, 0, 0, 0},
		{15, "Juventude", 0, 0, 0, 0, 0, 0, 0, 0},
		{16, "Palmeiras", 0, 0, 0, 0, 0, 0, 0, 0},
		{17, "São Paulo", 0, 0, 0, 0, 0, 0, 0, 0},
		{18, "Vasco", 0, 0, 0, 0, 0, 0, 0, 0},
		{19, "Vitória", 0, 0, 0, 0, 0, 0, 0, 0},
		{20, "Criciúma", 0, 0, 0, 0, 0, 0, 0, 0},
	}

	// Simular campeonato completo
	simularCampeonato(times)

	// Mostrar classificação final
	fmt.Println("\n================ CLASSIFICAÇÃO FINAL ================")
	ordenarPorPontuacao(times)
	for i, t := range times {
		fmt.Printf("%2d. %-12s %3d pts | %2dV %2dE %2dD | GP: %2d GC: %2d SG: %+d\n",
			i+1, t.nome, t.pontos, t.vitorias, t.empates, t.derrotas, t.golsPro, t.golsContra, t.saldo)
	}
}

func simularCampeonato(times []*timeStruct) {
	rand.Seed(time.Now().UnixNano())
	totalTimes := len(times)

	for i := 0; i < totalTimes; i++ {
		for j := 0; j < totalTimes; j++ {
			if i == j {
				continue // Um time não pode jogar contra si mesmo
			}

			// Jogo de ida: times[i] (mandante) vs times[j] (visitante)
			simularJogo(times[i], times[j])
		}
	}
}

func simularJogo(mandante, visitante *timeStruct) {
	golsMandante := rand.Intn(6)
	golsVisitante := rand.Intn(6)

	// Atualiza stats
	mandante.jogos++
	visitante.jogos++
	mandante.golsPro += golsMandante
	visitante.golsPro += golsVisitante
	mandante.golsContra += golsVisitante
	visitante.golsContra += golsMandante
	mandante.saldo += golsMandante - golsVisitante
	visitante.saldo += golsVisitante - golsMandante

	if golsMandante > golsVisitante {
		mandante.pontos += 3
		mandante.vitorias++
		visitante.derrotas++
	} else if golsMandante < golsVisitante {
		visitante.pontos += 3
		visitante.vitorias++
		mandante.derrotas++
	} else {
		mandante.pontos++
		visitante.pontos++
		mandante.empates++
		visitante.empates++
	}
}

func ordenarPorPontuacao(times []*timeStruct) {
	sort.Slice(times, func(i, j int) bool {
		if times[i].pontos != times[j].pontos {
			return times[i].pontos > times[j].pontos
		}
		if times[i].saldo != times[j].saldo {
			return times[i].saldo > times[j].saldo
		}
		if times[i].golsPro != times[j].golsPro {
			return times[i].golsPro > times[j].golsPro
		}
		return times[i].nome < times[j].nome
	})
}
