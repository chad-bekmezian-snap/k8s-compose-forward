package-windows-amd64:
	mkdir composeForward_windows_x86_64
	cp -r ./bin/* composeForward_windows_x86_64
	env GOARCH=amd64 GOOS=windows go build -o ./composeForward_windows_x86_64/composeForward ./cmd/compose-forward/
	tar --exclude="*.DS_Store" -czf composeForward_windows_x86_64.tgz composeForward_windows_x86_64/
	rm -r composeForward_windows_x86_64
.PHONY: package-windows-amd64

package-windows-arm64:
	mkdir composeForward_windows_arm64
	cp -r ./bin/* composeForward_windows_arm64
	env GOARCH=arm64 GOOS=windows go build -o ./composeForward_windows_arm64/composeForward ./cmd/compose-forward/
	tar --exclude="*.DS_Store" -czf composeForward_windows_arm64.tgz composeForward_windows_arm64/
	rm -r composeForward_windows_arm64
.PHONY: package-windows-amd64

package-darwin-amd64:
	mkdir composeForward_darwin_amd64
	cp -r ./bin/* composeForward_darwin_amd64
	env GOARCH=amd64 GOOS=darwin go build -o ./composeForward_darwin_amd64/composeForward ./cmd/compose-forward/
	tar --exclude="*.DS_Store" -czf composeForward_darwin_amd64.tgz composeForward_darwin_amd64/
	rm -r composeForward_darwin_amd64
.PHONY: package-windows-amd64

package-darwin-arm64:
	mkdir composeForward_darwin_arm64
	cp -r ./bin/* composeForward_darwin_arm64
	env GOARCH=arm64 GOOS=darwin go build -o ./composeForward_darwin_arm64/composeForward ./cmd/compose-forward/
	tar --exclude="*.DS_Store" -czf composeForward_darwin_arm64.tgz composeForward_darwin_arm64/
	rm -r composeForward_darwin_arm64
.PHONY: package-windows-arm64

package: package-windows-amd64 package-windows-arm64 package-darwin-amd64 package-darwin-arm64