package strucs

type RTS900 struct {
	RouteID      int
	LoadIDS      []int
	CountOfLoads int
}

type RTS900SS struct {
	TotalLoadIDs int
	Routes       []RTS900
}
