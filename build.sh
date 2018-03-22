cd ui
npm install
npm run build
cd ..
export PATH=$PATH:$GOPATH/bin
rice embed-go
go build