
type CommonResponse {
    Code uint32 `json:"code"`
    Data interface{} `json:"data,omitempty"`
    Msg string `json:"msg"`
}

//导入请求
type ImportRequest {
    UpFile interface{} `json:"upFile"`
}


type ExportResponse {
    Byte []byte `json:"byte"`
}
// 空请求
type NullRequest {}

// 空回复
type NullResponse {}

// 流程
type (
    //UpdateNodeInstanceRequest 修改-节点实例请求
    UpdateNodeInstanceRequest {
        Id string `json:"id"`
        NodeId string `json:"nodeId,optional"`
        FlowInstanceId string `json:"flowInstanceId,optional"`
        UserId string `json:"userId,optional"`
        Name string `json:"name,optional"`
        Content string `json:"content,optional"`
        Height int64 `json:"height,optional"`
        Width int64 `json:"width,optional"`
        X int64 `json:"x,optional"`
        Y int64 `json:"y,optional"`
        Icon string `json:"icon,optional"`
        Type string `json:"type,optional"`
        StartTime string `json:"startTime,optional"`
        EndTime string `json:"endTime,optional"`
        Status string `json:"status,optional"`
        Sort int64 `json:"sort,optional"`
        Enabled bool `json:"enabled,optional"`
        Remark string `json:"remark,optional"`
        UserName string `json:"userName,optional"`
    }

    //GetNodeInstanceResponse 提取-节点实例结果
    GetNodeInstanceResponse {
        Id string `json:"id,omitempty"`
        NodeId string `json:"nodeId,omitempty"`
        FlowInstanceId string `json:"flowInstanceId,omitempty"`
        UserId string `json:"userId,omitempty"`
        Name string `json:"name,omitempty"`
        Content string `json:"content"`
        Height int64 `json:"height,omitempty"`
        Width int64 `json:"width,omitempty"`
        X int64 `json:"x,omitempty"`
        Y int64 `json:"y,omitempty"`
        Icon string `json:"icon,omitempty"`
        Type string `json:"type,omitempty"`
        StartTime string `json:"startTime,omitempty"`
        EndTime string `json:"endTime,omitempty"`
        Status string `json:"status,omitempty"`
        Sort int64 `json:"sort,omitempty"`
        Enabled bool `json:"enabled,omitempty"`
        Remark string `json:"remark,omitempty"`
        CreateBy string `json:"createBy,omitempty"`
        UpdateBy string `json:"updateBy,omitempty"`
        CreatedAt string `json:"createdAt,omitempty"`
        UpdatedAt string `json:"updatedAt,omitempty"`
        DeletedAt string `json:"deletedAt,omitempty"`
        UserName string `json:"userName,omitempty"`
    }

    //GetMessageResponse 提取-消息结果
    GetMessageResponse {
        Id string `json:"id,omitempty"`
        FlowInstanceId string `json:"flowInstanceId,omitempty"`
        UserId string `json:"userId,omitempty"`
        BusyName string `json:"busyName,omitempty"`
        BusyId string `json:"busyId,omitempty"`
        Title string `json:"title,omitempty"`
        Content string `json:"content,omitempty"`
        Type string `json:"type,omitempty"`
        Readed bool `json:"readed,omitempty"`
        Status string `json:"status,omitempty"`
        Sort int64 `json:"sort,omitempty"`
        Enabled bool `json:"enabled,omitempty"`
        Remark string `json:"remark,omitempty"`
        CreateBy string `json:"createBy,omitempty"`
        UpdateBy string `json:"updateBy,omitempty"`
        CreatedAt string `json:"createdAt,omitempty"`
        UpdatedAt string `json:"updatedAt,omitempty"`
        DeletedAt string `json:"deletedAt,omitempty"`
    }
    //GetFlowInstanceResponse 提取-工作流实例结果
    GetFlowInstanceResponse {
        Id string `json:"id,omitempty"`
        FlowId string `json:"flowId,omitempty"`
        Name string `json:"name,omitempty"`
        BusyName string `json:"busyName,omitempty"`
        BusyId string `json:"busyId,omitempty"`
        Type string `json:"type,omitempty"`
        FlowStatus string `json:"flowStatus,omitempty"`
        Status string `json:"status,omitempty"`
        Sort int64 `json:"sort,omitempty"`
        StartTime string `json:"startTime,omitempty"`
        EndTime string `json:"endTime,omitempty"`
        Enabled bool `json:"enabled,omitempty"`
        Remark string `json:"remark,omitempty"`
        CreateBy string `json:"createBy,omitempty"`
        UpdateBy string `json:"updateBy,omitempty"`
        CreatedAt string `json:"createdAt,omitempty"`
        UpdatedAt string `json:"updatedAt,omitempty"`
        DeletedAt string `json:"deletedAt,omitempty"`
        NodeList []*GetNodeInstanceResponse `json:"nodeList,omitempty"`
        LinkList []*GetLinkInstanceResponse `json:"linkList,omitempty"`
        messageList []*GetMessageResponse `json:"messageList,omitempty"`
    }

    //GetLinkInstanceResponse 提取-状态转换实例结果
    GetLinkInstanceResponse {
        Id string `json:"id,omitempty"`
        LinkId string `json:"linkId,omitempty"`
        FlowInstanceId string `json:"flowInstanceId,omitempty"`
        Name string `json:"name,omitempty"`
        SourceId string `json:"sourceId,omitempty"`
        TargetId string `json:"targetId,omitempty"`
        Status string `json:"status,omitempty"`
        Sort int64 `json:"sort,omitempty"`
        Enabled bool `json:"enabled,omitempty"`
        StartTime string `json:"startTime,omitempty"`
        EndTime string `json:"endTime,omitempty"`
        Remark string `json:"remark,omitempty"`
        CreateBy string `json:"createBy,omitempty"`
        UpdateBy string `json:"updateBy,omitempty"`
        CreatedAt string `json:"createdAt,omitempty"`
        UpdatedAt string `json:"updatedAt,omitempty"`
        DeletedAt string `json:"deletedAt,omitempty"`
        LinkType string `json:"linkType,omitempty"`
    }
)