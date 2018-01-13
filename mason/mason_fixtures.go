package mason

func testMetadataObj() (metadata Metadata) {
	metadata = Metadata{
		Package:     testModuleName(),
		Version:     "0.1.0",
		Description: "A tool for building and testing your project in a clean GOPATH.",
		BuildInfo:   BuildInfo{},
		SignInfo: SignInfo{
			Program: "gpg",
			Email:   "gomason-tester@foo.com",
		},
		PublishInfo: PublishInfo{
			Targets:    make([]PublishTarget, 0),
			TargetsMap: make(map[string]PublishTarget),
		},
	}

	return metadata
}

func testModuleName() string {
	return "github.com/nikogura/gomason"
}
