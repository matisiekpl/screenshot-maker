package main

type Screenshot struct {
	Name         string
	Width        uint
	Height       uint
	MockupOffset uint
	PhotoWidth   uint
	PhotoOffset  uint
	Filename     string
}

var screenshots = map[int]Screenshot{
	0: {
		Name:         "iphone-12",
		Width:        1242,
		Height:       2688,
		MockupOffset: 316,
		PhotoWidth:   283,
		PhotoOffset:  433,
		Filename:     "Apple iPhone 11 Black.png",
	},
}
