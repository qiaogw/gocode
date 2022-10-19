syntax = "proto3";

{{ if .HasTimer }}
    import "google/protobuf/timestamp.proto";
{{- end }}


package {{.Package}};
option go_package="./{{.Package}}";

{{- range .Tables }}
    //提取 {{.Table}} ({{.TableComment}}) Request
    message Get{{.Table}}Request {
    {{- range  .Columns }}
        {{- if .IsPk }}
            {{.DataTypeProto}} {{.FieldJson}}=1;
        {{- end }}
    {{- end }}
    }
    //提取 {{.Table}} ({{.TableComment}}) Response
    message Get{{.Table}}Response {
    {{- $j :=0 -}}
    {{- range  .Columns }}
        {{- if .IsPage }}
        {{- else}}
            {{- $j = add $j }}
            {{.DataTypeProto}} {{.FieldJson}}={{- $j -}};
        {{- end}}
    {{- end }}
    }

    //列表 {{.Table}} ({{.TableComment}}) Request
    message List{{.Table}}Request {
    {{- $i :=0 -}}
    {{- range  .Columns }}
        {{- if .IsPk }}
        {{- else}}
            {{- $i = add $i }}
            {{.DataTypeProto}} {{.FieldJson}}={{- $i -}};
        {{- end}}
    {{- end }}
    {{- $i = add $i }}
    string searchKey ={{- $i -}};
    {{- $i = add $i }}
    string sortBy ={{- $i -}};
    {{- $i = add $i }}
    bool descending = {{- $i -}};
    }
    //列表 {{.Table}} ({{.TableComment}}) Response
    message List{{.Table}}Response {
    repeated Get{{.Table}}Response list = 1;
    int64 count=2;
    }
    //创建 {{.Table}} ({{.TableComment}}) Request
    message Create{{.Table}}Request {
    {{- $x :=0 -}}
    {{- range  .Columns }}
        {{- if .IsPk }}
        {{- else}}
            {{- if .IsPage }}
            {{- else}}
                {{- $x = add $x }}
                {{.DataTypeProto}} {{.FieldJson}}={{- $x -}};
            {{- end}}
        {{- end}}
    {{- end }}
    }

    //创建 {{.Table}} ({{.TableComment}}) Response
    message Create{{.Table}}Response {
    {{- range  .Columns }}
        {{- if .IsPk }}
            {{.DataTypeProto}} {{.FieldJson}}=1;
        {{- end }}
    {{- end }}
    }

    //修改 {{.Table}} ({{.TableComment}}) Request
    message Update{{.Table}}Request {
    {{- $e :=0 -}}
    {{- range  .Columns }}
        {{- if .IsPage }}
        {{- else}}
            {{- $e = add $e }}
            {{.DataTypeProto}} {{.FieldJson}}={{- $e -}};
        {{- end}}
    {{- end }}
    }
    //修改 {{.Table}} ({{.TableComment}}) Response
    message Update{{.Table}}Response {
    {{- range  .Columns }}
        {{- if .IsPk }}
            {{.DataTypeProto}} {{.FieldJson}}=1;
        {{- end }}
    {{- end }}
    }

    //删除 {{.Table}} ({{.TableComment}}) Request
    message Delete{{.Table}}Request {
    {{- range  .Columns }}
        {{- if .IsPk }}
            {{.DataTypeProto}} {{.FieldJson}}=1;
        {{- end }}
    {{- end }}
    }

    //删除 {{.Table}} ({{.TableComment}}) Response
    message Delete{{.Table}}Response {
    }
{{ end}}
//NullRequest 空 Request
message NullRequest {
}
//NullResponse 空 Response
message  NullResponse {
}
//登录 Login (用户) Request
message LoginRequest {
string mobile = 1;
string password = 2;
}

//登录 Login (用户) Response
message LoginResponse {
int64 id = 1;
}

service {{.Package}} {
rpc Login(LoginRequest) returns(LoginResponse);
{{range .Tables }}

    rpc Get{{.Table}}(Get{{.Table}}Request) returns(Get{{.Table}}Response);
    rpc List{{.Table}}(List{{.Table}}Request) returns(List{{.Table}}Response);
    rpc Create{{.Table}}(Create{{.Table}}Request) returns(Create{{.Table}}Response);
    rpc Update{{.Table}}(Update{{.Table}}Request) returns(Update{{.Table}}Response);
    rpc Delete{{.Table}}(Delete{{.Table}}Request) returns(Delete{{.Table}}Response);
{{ end}}
}
