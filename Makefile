ci_run_test:
	export __AMINO_SID__=$(./scripts/GetSID)
	go test github.com/AminoJS/AminoGo/test