module tc-static-routes-file-generator

go 1.19

replace file-persistence-mod => ../file-persistence-mod

replace std-io-mod => ../std-io-mod

replace tc-util-mod => ../tc-util-mod

require (
	file-persistence-mod v0.0.0-00010101000000-000000000000
	std-io-mod v0.0.0-00010101000000-000000000000
	tc-util-mod v0.0.0-00010101000000-000000000000
)
