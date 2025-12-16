package middleware

var rolePermissions = map[string][]string{
	"mahasiswa": {
		"achievement:create",
		"achievement:submit",
		"achievement:view",
	},
	"dosen": {
		"achievement:view",
		"achievement:verify",
		"achievement:reject",
	},
	"admin": {
		"achievement:view_all",
	},
}
