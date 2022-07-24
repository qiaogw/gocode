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
         {{- if .IsPk -}}
          {{.DataTypeProto}} {{.FieldName}}={{.Indexs}};
         {{- end }}
       {{- end }}
    }
    //提取 {{.Table}} ({{.TableComment}}) Response
    message Get{{.Table}}Response {
      {{- range  .Columns }}
        {{.DataTypeProto}} {{.FieldName}}={{.Indexs}};
      {{- end }}
    }

    //列表 {{.Table}} ({{.TableComment}}) Request
    message List{{.Table}}Request {
    {{- range  .Columns }}
        {{- if .IsPk }}
        {{- else}}
            {{.DataTypeProto}} {{.FieldName}}={{.Indexs}};
        {{- end}}
    {{- end }}
    }
    //列表 {{.Table}} ({{.TableComment}}) Response
    message List{{.Table}}Response {
      repeated Get{{.Table}}Response data = 1;
    }
    //创建 {{.Table}} ({{.TableComment}}) Request
    message Create{{.Table}}Request {
       {{- range  .Columns }}
           {{- if .IsPk }}
           {{- else}}
               {{.DataTypeProto}} {{.FieldName}}={{.Indexs}};
           {{- end}}
       {{- end }}
    }

    //创建 {{.Table}} ({{.TableComment}}) Response
    message Create{{.Table}}Response {
      {{- range  .Columns }}
       {{- if .IsPk -}}
        {{.DataTypeProto}} {{.FieldName}}={{.Indexs}};
       {{- end }}
     {{- end }}
    }

    //修改 {{.Table}} ({{.TableComment}}) Request
    message Update{{.Table}}Request {
       {{- range  .Columns }}
         {{.DataTypeProto}} {{.FieldName}}={{.Indexs}};
       {{- end }}
    }
    //修改 {{.Table}} ({{.TableComment}}) Response
    message Update{{.Table}}Response {
       {{- range  .Columns }}
         {{- if .IsPk -}}
          {{.DataTypeProto}} {{.FieldName}}={{.Indexs}};
         {{- end }}
       {{- end }}
    }

    //删除 {{.Table}} ({{.TableComment}}) Request
    message Delete{{.Table}}Request {
       {{- range  .Columns }}
         {{- if .IsPk -}}
          {{.DataTypeProto}} {{.FieldName}}={{.Indexs}};
         {{- end }}
       {{- end }}
    }

    //删除 {{.Table}} ({{.TableComment}}) Response
    message Delete{{.Table}}Response {
    }
{{ end}}


service {{.Package}} {
{{range .Tables }}
  rpc Get{{.Table}}(Get{{.Table}}Request) returns(Get{{.Table}}Response);
  rpc List{{.Table}}(List{{.Table}}Request) returns(List{{.Table}}Response);
  rpc Create{{.Table}}(Create{{.Table}}Request) returns(Create{{.Table}}Response);
  rpc Update{{.Table}}(Update{{.Table}}Request) returns(Update{{.Table}}Response);
  rpc Delete{{.Table}}(Delete{{.Table}}Request) returns(Delete{{.Table}}Response);
{{ end}}
}
