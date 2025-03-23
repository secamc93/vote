package domainerrors

import "errors"

var (
	ErrVoteGroupNotFound            = errors.New("grupo de votación no encontrado")
	ErrHouseNotFound                = errors.New("casa no encontrada")
	ErrHouseAndVoteGroupInUserExist = errors.New("usuario ya registrado en este grupo de votación")
	ErrHouseFound                   = errors.New("casa ya existe")
	ErrHouseNameIsRequired          = errors.New("el nombre de la casa es requerido")
	ErrVoteGroupIDIsRequired        = errors.New("el id del grupo de votación es requerido")
)
