module moduletwo

go 1.12

// replace moduleone => ../module-one
replace moduleone => /home/runner/work/go-module-example/go-module-example/module-one

require moduleone v0.0.0
