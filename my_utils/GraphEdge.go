package my_utils

type GraphEdge struct {
	Start string
	End string
	Weight int
}

/*
IndexGraphEdge is the one which has Start And End Values as integer values.
 */
type IndexGraphEdge struct {
	Start int32
	End int32
	Weight int32
}
