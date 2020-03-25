ci_run_test:
	export __AMINO_SID__=$(CI=true go run scripts/GetSID.go)
	go test github.com/AminoJS/AminoGo/test