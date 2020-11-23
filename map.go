type RouterGroup struct{
	prefix        string
	middlewares   []HandlerFunc
	parent        *RouterGroup
	engine        *Engine

	Group(prefix string) *RouterGroup
	Use(middlewares ...HandlerFunc)
	addRoute(method string, comp string, handler HandlerFunc)
	GET(pattern string, handler HandlerFunc)
	POST(pattern string, handler HandlerFunc)
	createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc
	Static(relativePath string, root string)
}

type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	Params     map[string]string
	StatusCode int
	handlers   []HandlerFunc
	index      int
	engine     *Engine

	Next()
	Fail(code int, err string)
	Param(key string) string
	PostForm(key string) string
	Query(key string) string
	Status(code int)
	SetHeader(key string, value string)
	String(code int, format string, values ...interface{})
	JSON(code int, obj interface{})
	Data(code int, data []byte)
	HTML(code int, name string, data interface{})
	Write(b []byte)
}

type Engine struct {
	*RouterGroup
	router        *router
	groups        []*RouterGroup     
	htmlTemplates *template.Template 
	funcMap       template.FuncMap   

	SetFuncMap(funcMap template.FuncMap)
	LoadHTMLGlob(pattern string)
	Run(addr string) (err error)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

type router struct {
	roots    map[string]*TrieNode
	handlers map[string]HandlerFunc

	addRoute(method string, pattern string, handler HandlerFunc)
	getRoute(method string, path string) (*TrieNode, map[string]string)
	getRoutes(method string) []*TrieNode
	handle(c *Context)
}