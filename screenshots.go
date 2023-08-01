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
		Name:         "iphone-11",
		Width:        1242,
		Height:       2688,
		MockupOffset: 316,
		PhotoWidth:   283,
		PhotoOffset:  433,
		Filename:     "Apple iPhone 11 Black.png",
	},
	1: {
		Name:         "ipad-pro",
		Width:        2048,
		Height:       2732,
		MockupOffset: 200,
		PhotoWidth:   378,
		PhotoOffset:  364,
		Filename:     "Apple iPad Pro 13 Space Gray.png",
	},
	2: {
		Name:         "iphone-11-alt",
		Width:        1242,
		Height:       2208,
		MockupOffset: 180,
		PhotoWidth:   280,
		PhotoOffset:  290,
		Filename:     "Apple iPhone 11 Black.png",
	},
}
