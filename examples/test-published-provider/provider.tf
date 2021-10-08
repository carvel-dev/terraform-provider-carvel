provider "carvel" {
	kapp {
	    diff_output_file = "kapp.diff.log"
		kubeconfig {
			from_env = true
		}
	}
}
