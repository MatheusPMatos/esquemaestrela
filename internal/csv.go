package internal

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func ReaderCsv(path string) ([]Nota, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir o arquivo: %w", err)
	}
	defer file.Close()

	// Criar um leitor CSV
	reader := csv.NewReader(file)
	reader.Comma = ';'

	// Ler todas as linhas do CSV
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o arquivo CSV: %w", err)
	}

	// Slice para armazenar os dados
	var enades = make([]Nota, 0)

	// Iterar sobre as linhas do CSV (ignorando o cabeçalho)
	for i, row := range rows {
		if i == 0 {
			continue // Ignorar a primeira linha (cabeçalho)
		}
		if len(row) < 19 {
			fmt.Printf("NUMERO DE COLUNAS: %d", len(row))
			return nil, fmt.Errorf("linha %d com número inválido de colunas", i+1)
		}

		NU_ANO, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, fmt.Errorf("erro ao converter NT_GER na linha %d: %w", i+1, err)
		}

		CO_IES, err := strconv.Atoi(row[1])
		if err != nil {
			return nil, fmt.Errorf("erro ao converter ANO_IN_GRAD na linha %d: %w", i+1, err)
		}

		CO_CATEGAD, err := strconv.Atoi(row[2])
		if err != nil {
			return nil, fmt.Errorf("erro ao converter TP_PRES na linha %d: %w", i+1, err)
		}
		CO_ORGACAD, err := strconv.Atoi(row[3])
		if err != nil {
			return nil, fmt.Errorf("erro ao converter NT_GER na linha %d: %w", i+1, err)
		}

		CO_GRUPO, err := strconv.Atoi(row[4])
		if err != nil {
			return nil, fmt.Errorf("erro ao converter ANO_IN_GRAD na linha %d: %w", i+1, err)
		}

		CO_CURSO, err := strconv.Atoi(row[5])
		if err != nil {
			return nil, fmt.Errorf("erro ao converter TP_PRES na linha %d: %w", i+1, err)
		}
		CO_MODALIDADE, err := strconv.Atoi(row[6])
		if err != nil {
			return nil, fmt.Errorf("erro ao converter NT_GER na linha %d: %w", i+1, err)
		}

		CO_MUNIC_CURSO, err := strconv.Atoi(row[7])
		if err != nil {
			return nil, fmt.Errorf("erro ao converter ANO_IN_GRAD na linha %d: %w", i+1, err)
		}

		TP_INSCRICAO, err := strconv.Atoi(row[8])
		if err != nil {
			return nil, fmt.Errorf("erro ao converter TP_PRES na linha %d: %w", i+1, err)
		}

		IN_REGULAR, err := strconv.Atoi(row[9])
		if err != nil {
			return nil, fmt.Errorf("erro ao converter ANO_IN_GRAD na linha %d: %w", i+1, err)
		}

		TP_INCRICAO_ADM, err := strconv.Atoi(row[10])
		if err != nil {
			return nil, fmt.Errorf("erro ao converter TP_PRES na linha %d: %w", i+1, err)
		}
		ANO_IN_GRAD, err := strconv.Atoi(row[11])
		if err != nil {
			return nil, fmt.Errorf("erro ao converter NT_GER na linha %d: %w", i+1, err)
		}

		TP_PRES, err := strconv.Atoi(row[12])
		if err != nil {
			return nil, fmt.Errorf("erro ao converter ANO_IN_GRAD na linha %d: %w", i+1, err)
		}

		NT_GER, err := strconv.ParseFloat(row[13], 32)
		if err != nil {
			return nil, fmt.Errorf("erro ao converter TP_PRES na linha %d: %w", i+1, err)
		}

		ANO_ENEM, err := strconv.Atoi(row[14])
		if err != nil {
			return nil, fmt.Errorf("erro ao converter TP_PRES na linha %d: %w", i+1, err)
		}
		ENEM_NT_CN, err := strconv.ParseFloat(row[15], 32)
		if err != nil {
			return nil, fmt.Errorf("erro ao converter NT_GER na linha %d: %w", i+1, err)
		}

		ENEM_NT_CH, err := strconv.ParseFloat(row[16], 32)
		if err != nil {
			return nil, fmt.Errorf("erro ao converter ANO_IN_GRAD na linha %d: %w", i+1, err)
		}

		ENEM_NT_LC, err := strconv.ParseFloat(row[17], 32)
		if err != nil {
			return nil, fmt.Errorf("erro ao converter TP_PRES na linha %d: %w", i+1, err)
		}
		ENEM_NT_MT, err := strconv.ParseFloat(row[18], 32)
		if err != nil {
			return nil, fmt.Errorf("erro ao converter TP_PRES na linha %d: %w", i+1, err)
		}

		// Criar uma nova instância de Enade
		enade := Nota{
			NT_GER: NT_GER,
			ENEM: Enem{
				ANO_ENEM:   ANO_ENEM,
				ENEM_NT_CN: ENEM_NT_CN,
				ENEM_NT_CH: ENEM_NT_CH,
				ENEM_NT_LC: ENEM_NT_LC,
				ENEM_NT_MT: ENEM_NT_MT,
			},
			CURSO: Curso{
				CO_IES:         CO_IES,
				CO_CATEGAD:     CO_CATEGAD,
				CO_ORGACAD:     CO_ORGACAD,
				CO_GRUPO:       CO_GRUPO,
				CO_CURSO:       CO_CURSO,
				CO_MODALIDADE:  CO_MODALIDADE,
				CO_MUNIC_CURSO: CO_MUNIC_CURSO,
			},
			ENADE: Enade{
				NU_ANO:          NU_ANO,
				TP_INSCRICAO:    TP_INSCRICAO,
				IN_REGULAR:      IN_REGULAR,
				TP_INCRICAO_ADM: TP_INCRICAO_ADM,
				ANO_IN_GRAD:     ANO_IN_GRAD,
				TP_PRES:         TP_PRES,
			},
		}

		// Adicionar ao slice
		enades = append(enades, enade)
	}

	return enades, nil
}
