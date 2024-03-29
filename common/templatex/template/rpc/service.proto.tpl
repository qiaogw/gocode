syntax = "proto3";

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
{{- $j = add $j }}
string busyName={{- $j -}};
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
    {{- $i = add $i }}
    string token={{- $i -}};
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
    {{- else if .IsModelTime}}
    {{- else if .IsControl}}
    {{- else}}
    {{- $x = add $x }}
    {{.DataTypeProto}} {{.FieldJson}}={{- $x -}};
    {{- end}}
{{- end}}
{{- end }}
    {{- $x = add $x }}
    string createBy={{- $x -}};
}

//创建 {{.Table}} ({{.TableComment}}) Response
message Create{{.Table}}Response {
{{- range  .Columns }}
{{- if .IsPk }}
    {{.DataTypeProto}} {{.FieldJson}}=1;
{{- end }}
{{- end }}
    string busyName=2;
}

//修改 {{.Table}} ({{.TableComment}}) Request
message Update{{.Table}}Request {
{{- $e :=0 -}}
{{- range  .Columns }}
    {{- if .IsPage }}
    {{- else if .IsModelTime}}
    {{- else if .IsControl}}
    {{- else}}
    {{- $e = add $e }}
    {{.DataTypeProto}} {{.FieldJson}}={{- $e -}};
    {{- end}}
{{- end }}
{{- $e = add $e }}
    string updateBy={{- $e -}};
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
    string updateBy=2;
}
//批量删除 {{.Table}} ({{.TableComment}}) Request
    message DeleteList{{.Table}}Request {
    repeated string list = 1;
}
{{ end}}

//NullRequest 空 Request
message NullRequest {
}
//NullResponse 空 Response
message  NullResponse {
}
{{- if .IsFlow }}
//BusyNameResponse 业务名称 Response
message  BusyNameResponse {
    string name=1;
}
{{- end}}

//导出 Export  Request
message ExportRequest {
    int64 pageIndex = 1;
    int64 pageSize = 2;
    string searchKey = 3;
    string sortBy = 4;
    bool descending = 5;
}

//导出 Export  Response
message ExportResponse {
    bytes data=1;
}


{{range .Tables }}

service {{.Table}} {
    rpc Get{{.Table}}(Get{{.Table}}Request) returns(Get{{.Table}}Response);
    rpc List{{.Table}}(List{{.Table}}Request) returns(List{{.Table}}Response);
    rpc Create{{.Table}}(Create{{.Table}}Request) returns(Create{{.Table}}Response);
    rpc Update{{.Table}}(Update{{.Table}}Request) returns(Update{{.Table}}Response);
    rpc Delete{{.Table}}(Delete{{.Table}}Request) returns(NullResponse);
    rpc DeleteList{{.Table}}(DeleteList{{.Table}}Request) returns(NullResponse);
{{- if .IsImport}}
    rpc Export{{.Table}}(ExportRequest) returns(ExportResponse);
    rpc ExportTemplate{{.Table}}(NullRequest) returns(ExportResponse);
    rpc Import{{.Table}}(ExportResponse) returns(NullResponse);
{{ end}}
{{- if .IsFlow }}
    rpc GetBusyName{{.Table}}(NullRequest) returns(BusyNameResponse);
{{- end}}
}
{{ end}}