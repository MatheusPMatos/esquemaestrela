# DICIONÁRIO DE VARIÁVEIS - IDD 2022

## PARTE 1 - INFORMAÇÕES DA INSTITUIÇÃO DE ENSINO SUPERIOR E DO CURSO

| Nº | Nome              | Tipo     | Tamanho | Descrição                                          | Categorias |
|-----|-------------------|----------|---------|--------------------------------------------------|------------|
| 1   | NU_ANO           | Numérica | 4       | Ano de realização do Enade                      | 2022       |
| 2   | CO_IES           | Numérica | 5       | Código da IES (e-Mec)                           |            |
| 3   | CO_CATEGAD       | Numérica | 1       | Código da categoria administrativa da IES        | "1. Pública Federal<br>2. Pública Estadual<br>3. Pública Municipal<br>4. Privada com fins lucrativos<br>5. Privada sem fins lucrativos<br>7. Especial" |
| 4   | CO_ORGACAD       | Numérica | 5       | Código da organização acadêmica da IES          | "10019 = Centro Federal de Educação Tecnológica<br>10020 = Centro Universitário<br>10022 = Faculdade<br>10026 = Instituto Federal de Educação, Ciência e Tecnologia<br>10028 = Universidade" |
| 5   | CO_GRUPO         | Numérica | 4       | Código da Área de enquadramento do curso no Enade | "1 = Administração<br>2 = Direito<br>13 = Ciências Econômicas<br>18 = Psicologia<br>22 = Ciências Contábeis<br>29 = Turismo<br>38 = Serviço Social<br>67 = Secretariado Executivo<br>81 = Relações Internacionais<br>83 = Tecnologia em Design de Moda<br>84 = Tecnologia em Marketing<br>85 = Tecnologia em Processos Gerenciais<br>86 = Tecnologia em Gestão de Recursos Humanos<br>87 = Tecnologia em Gestão Financeira<br>88 = Tecnologia em Gastronomia<br>93 = Tecnologia em Gestão Comercial<br>94 = Tecnologia em Logística<br>100 = Administração Pública<br>101 = Teologia<br>102 = Tecnologia em Comércio Exterior<br>103 = Tecnologia em Design de Interiores<br>104 = Tecnologia em Design Gráfico<br>105 = Tecnologia em Gestão da Qualidade<br>106 = Tecnologia em Gestão Pública<br>803 = Jornalismo<br>804 = Publicidade e Propaganda" |
| 6   | CO_CURSO         | Numérica | 7       | Código do curso (e-Mec)                          |            |
| 7   | CO_MODALIDADE    | Numérica | 1       | Modalidade de Ensino                             | "0 = EaD<br>1 = Presencial" |
| 8   | CO_MUNIC_CURSO   | Numérica | 7       | Código do município de funcionamento do curso     | Entre 1100023 e 5300108 (Identificação do Município conforme IBGE - [Ir para a Planilha de Municípios](https://www.ibge.gov.br)) |

## PARTE 2 - INFORMAÇÕES DO ESTUDANTE

| Nº  | Nome              | Tipo     | Tamanho | Descrição                                          | Categorias |
|------|-------------------|----------|---------|--------------------------------------------------|------------|
| 9    | TP_INSCRICAO     | Numérica | 1       | Tipo de inscrição                                | 1 = Concluinte |
| 10   | IN_REGULAR       | Numérica | 1       | Indicativo da condição da inscrição               | 1 = Regular |
| 11   | TP_INSCRICAO_ADM | Numérica | 1       | Tipo da origem da inscrição                      | "0 = Tradicional<br>2 = Administrativo" |
| 12   | ANO_IN_GRAD      | Numérica | 4       | Ano de início da graduação                     | Entre 2009 e 2022 |
| 13   | TP_PRES          | Numérica | 3       | Tipo de presença                                  | 555 = Presente com resultado válido |
| 14   | NT_GER           | Numérica | 4       | Nota bruta da prova - Média ponderada da formação geral (25%) e componente específico (75%) (0 a 100) |            |
| 15   | ANO_ENEM         | Numérica | 4       | Identificação do ano de Enem selecionado         | Entre 2009 e 2022 |
| 16   | ENEM_NT_CN       | Numérica | 5       | Nota da prova de Ciências da Natureza            |            |
| 17   | ENEM_NT_CH       | Numérica | 5       | Nota da prova de Ciências Humanas                |            |
| 18   | ENEM_NT_LC       | Numérica | 5       | Nota da prova de Linguagens e Códigos            |            |
| 19   | ENEM_NT_MT       | Numérica | 5       | Nota da prova de Matemática                      |            |
