// Since both the services and their consumers needs to know URL paths, single source-of-truth is needed.
// TODO: Keep this up to date as new services and endpoints defined.

package paths

var Site = Domain{"http", "deepthinker.app", "8080"}
var ApiGateway = Gateway{Site, "/api/v1.0.0"}

var (
	Customs                  = Service{ApiGateway, "/customs"}
	CustomsModuleUpload      = Endpoint{Customs, "/module", POST}
	CustomsModuleDownload    = Endpoint{Customs, "/module/{id}", GET}
	CustomsModuleList        = Endpoint{Customs, "/module/{id}/list", GET}
	CustomsModuleAstFuncDecl = Endpoint{Customs, "/module/{id}/ast/{package}/{file}/{function}", GET}
	CustomsModuleAstFile     = Endpoint{Customs, "/module/{id}/ast/{package}/{file}", GET}
	CustomsModuleAstPackage  = Endpoint{Customs, "/module/{id}/ast/{package}", GET}
	CustomsModuleContext     = Endpoint{Customs, "/module/{id}/context/package/{package}/file/{file}", GET}
)

var (
	Evolver              = Service{ApiGateway, "/evolver"}
	EvolverSessionCreate = Endpoint{Evolver, "/session", GET}
)

var (
	Runner            = Service{ApiGateway, "/runner"}
	RunnerBatchCreate = Endpoint{Runner, "/batch", GET}
)
