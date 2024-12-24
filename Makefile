
install:
	go install golang.org/x/mobile/cmd/gomobile@latest 
	yay -Sy android-sdk	android-studio android-ndk

run:
	go run .

build-android:
	gomobile build -target=android -androidapi 35 golang.org/x/mobile/example/basic