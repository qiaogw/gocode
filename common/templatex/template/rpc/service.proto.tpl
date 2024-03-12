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
{{- $j = add $j }}
GetFlowInstanceResponse flowInstance={{- $j -}};

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


//提取 FlowInstance (工作流实例) Response
message GetFlowInstanceResponse {
    string id=1;
    string flowId=2;
    string name=3;
    string busyName=4;
    string busyId=5;
    string type=6;
    string start=7;
    string end=8;
    string status=9;
    int64 sort=10;
    string startTime=11;
    string endTime=12;
    bool enabled=13;
    string remark=14;
    string createBy=15;
    string updateBy=16;
    string createdAt=17;
    string updatedAt=18;
    string deletedAt=19;
    repeated GetNodeInstanceResponse nodeList = 20;
    repeated GetLinkInstanceResponse linkList = 21;
    repeated GetMessageResponse messageList = 22;
}

//提取 NodeInstance (节点实例) Response
message GetNodeInstanceResponse {
    string id=1;
    string nodeId=2;
    string flowInstanceId=3;
    string userId=4;
    string name=5;
    string content=6;
    int64 height=7;
    int64 width=8;
    int64 x=9;
    int64 y=10;
    string icon=11;
    string type=12;
    string startTime=13;
    string endTime=14;
    string status=15;
    int64 sort=16;
    bool enabled=17;
    string remark=18;
    string createBy=19;
    string updateBy=20;
    string createdAt=21;
    string updatedAt=22;
    string deletedAt=23;
}

//提取 LinkInstance (状态转换实例) Response
message GetLinkInstanceResponse {
    string id=1;
    string finkId=2;
    string flowInstanceId=3;
    string name=4;
    string sourceId=5;
    string targetId=6;
    string status=7;
    int64 sort=8;
    bool enabled=9;
    string startTime=10;
    string endTime=11;
    string remark=12;
    string createBy=13;
    string updateBy=14;
    string createdAt=15;
    string updatedAt=16;
    string deletedAt=17;
}

//提取 Message (消息) Response
message GetMessageResponse {
    string id=1;
    string flowInstanceId=2;
    string userId=3;
    string busyName=4;
    string busyId=5;
    string title=6;
    string content=7;
    string type=8;
    bool readed=9;
    string status=10;
    int64 sort=11;
    bool enabled=12;
    string remark=13;
    string createBy=14;
    string updateBy=15;
    string createdAt=16;
    string updatedAt=17;
    string deletedAt=18;
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
}
{{ end}}