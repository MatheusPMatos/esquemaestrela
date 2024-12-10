package internal

type BaseModel struct {
	ID uint `gorm:"primaryKey"`
}

type Nota struct {
	BaseModel
	NT_GER float64
	ENEM   Enem  `gorm:"foreignkey:ENADE_ID"`
	CURSO  Curso `gorm:"foreignkey:ENADE_ID"`
	ENADE  Enade `gorm:"foreignkey:ENADE_ID"`
}

type Enade struct {
	BaseModel
	ENADE_ID        int
	NU_ANO          int
	TP_INSCRICAO    int
	IN_REGULAR      int
	TP_INCRICAO_ADM int
	ANO_IN_GRAD     int
	TP_PRES         int
}

type Enem struct {
	BaseModel
	ENADE_ID   int
	ANO_ENEM   int
	ENEM_NT_CN float64
	ENEM_NT_CH float64
	ENEM_NT_LC float64
	ENEM_NT_MT float64
}

type Curso struct {
	BaseModel
	ENADE_ID       int
	CO_IES         int
	CO_CATEGAD     int
	CO_ORGACAD     int
	CO_GRUPO       int
	CO_CURSO       int
	CO_MODALIDADE  int
	CO_MUNIC_CURSO int
}
