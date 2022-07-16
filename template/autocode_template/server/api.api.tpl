type (
	Create{{.StructName}}Req {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
	}
	Create{{.StructName}}Reply {
		Id           int64  `json:"id"`
		Name         string `json:"name"`
		Gender       string `json:"gender"`
		AccessToken  string `json:"accessToken"`
		AccessExpire int64  `json:"accessExpire"`
		RefreshAfter int64  `json:"refreshAfter"`
	}
	// / 用户注册
	RegisterRequest {
		Name     string `json:"name"`
		Gender   int64  `json:"gender"`
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
	}
	RegisterResponse {
		Id     int64  `json:"id"`
		Name   string `json:"name"`
		Gender int64  `json:"gender"`
		Mobile string `json:"mobile"`
	}
	// 用户注册

	// 用户信息
	UserInfoResponse {
		Id     int64  `json:"id"`
		Name   string `json:"name"`
		Gender int64  `json:"gender"`
		Mobile string `json:"mobile"`
	}
	// 用户信息
	SearchReq {
		// 图书名称
		Name string `form:"name"`
	}

	SearchReply {
		Name  string `json:"name"`
		Count int    `json:"count"`
	}
)

service admin-api {
	@handler login
	post /admin/login (LoginReq) returns (LoginReply)

	@handler Register
	post /admin/register (RegisterRequest) returns (RegisterResponse)

}

@server(
	jwt: Auth
)

service admin-api {
	@handler search
	get /admin/search (SearchReq) returns (SearchReply)

	@handler userinfo
	get /admin/userinfo () returns (UserInfoResponse)

}