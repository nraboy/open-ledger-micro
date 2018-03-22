cd ui
npm install
npm run build
cd ..
export PATH=$PATH:$GOPATH/bin
rice embed-go
env GOOS=linux GOARCH=arm GOARM=5 go build
scp open-ledger-micro coin.local:~/
rm open-ledger
