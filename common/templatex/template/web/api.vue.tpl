{{- $db :=.Db }}
{{- $isTableFlow :=.IsFlow }}
<template>
  <q-page class="q-pa-xs">
    <div class="shadow-2 q-pa-xs">
      <q-table
              dense
              flat
              bordered
              separator="cell"
              :columns="columns"
              :rows="dataList"
              row-key="id"
              v-model:filter="pagination.searchKey"
              @request="onRequest"
              :grid="$q.screen.xs"
              v-model:pagination="pagination"
              binary-state-sort
              :selected-rows-label="getSelectedString"
              selection="multiple"
              v-model:selected="selected"
      >
        <template v-slot:top="table">
          <div class="row no-wrap full-width">
            <com-filter
                    @query="onRequest"
                    @reset="resetQuery"
                    v-model:filter="pagination.searchKey"
            />
            <q-btn
                    flat
                    dense
                    no-wrap
                    class="q-ml-md"
                    color="primary"
                    @click="filterVisible = !filterVisible"
                    :icon="filterVisible ? 'arrow_drop_up' : 'arrow_drop_down'"
            ><q-tooltip>更多查询条件</q-tooltip></q-btn
            >
            <q-space />
            <q-btn-group class="q-gutter-xs">
              <q-btn
                      v-permission="'{{.TableUrl}}:add'"
                      flat
                      dense
                      glossy
                      icon="add"
                      no-wrap
                      color="primary"
                      @click="create"
              ><q-tooltip>新建</q-tooltip></q-btn
              >

              <com-del
                      :disable="selected.length < 1"
                      v-permission="'{{.TableUrl}}:del'"
                      flat
                      glossy
                      label="{{.TableComment}}"
                      @confirm="delList"
              />

              {{- if .IsImport }}
                <q-btn
                        v-permission="'{{.TableUrl}}:export'"
                        flat
                        dense
                        glossy
                        icon="download"
                        no-wrap
                        color="primary"
                        @click="handleExport"
                ><q-tooltip>导出</q-tooltip></q-btn
                >
                <q-btn
                        v-permission="'{{.TableUrl}}:export'"
                        flat
                        dense
                        glossy
                        icon="downloading"
                        no-wrap
                        color="primary"
                        @click="handleExportTemplate"
                ><q-tooltip>导出模板</q-tooltip></q-btn
                >
                <com-upload
                        v-permission="'{{.TableUrl}}:import'"
                        flat
                        dense
                        glossy
                        @upload="resetQuery"
                        title="导入"
                        :uploadUrl="uploadUrl"
                        fileType=".xlsx,.xls"
                />
              {{- end}}
              <q-btn
                      color="primary"
                      flat
                      dense
                      glossy
                      no-wrap
                      v-show="$q.screen.gt.md"
                      @click="table.toggleFullscreen"
                      :icon="table.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
              >
                <q-tooltip>切换全屏</q-tooltip>
              </q-btn>
              <q-btn
                      color="deep-orange"
                      flat
                      dense
                      glossy
                      no-wrap
                      v-show="$q.screen.gt.md"
                      icon="help_outline"
              >
                <q-tooltip>使用说明</q-tooltip>
                <q-popup-proxy :offset="[10, 10]">
                  <q-banner class="bg-purple text-white">
                    <template v-slot:avatar>
                      <q-icon name="help" />
                    </template>
                    <div v-html="route.meta.remark"></div>
                  </q-banner>
                </q-popup-proxy>
              </q-btn>
            </q-btn-group>
          </div>
          <div v-show="filterVisible" class="row no-wrap full-width">
            {{- range  .Columns }}
                {{- if .DictType }}
                  <q-select
                          class="col-2 q-pr-sm q-py-sm"
                          outlined
                          dense
                          clearable
                          map-options
                          emit-value
                          option-value="value"
                          option-label="label"
                          v-model="pagination.{{.FieldJson}}"
                          :options="{{.DictType}}Options"
                          label="{{.ColumnComment}}"
                          @update:model-value="filterSearch"
                  />
                {{- else if  .FkTable}}
                  <q-select
                          class="col-2 q-pr-sm q-py-sm"
                          outlined
                          dense
                          map-options
                          emit-value
                          clearable
                          v-model="pagination.{{.FieldJson}}"
                          :options="{{.FkTablePackage}}Options"
                          label="{{.ColumnComment}}"
                          @update:model-value="filterSearch"
                  />
                {{- else if  eq  .DataType "bool" }}
                  <q-toggle
                          label="{{.ColumnComment}}"
                          color="green"
                          class="col-2 q-pr-sm q-py-sm"
                          v-model="pagination.{{.FieldJson}}"
                          @update:model-value="filterSearch"
                  />
                {{- end }}
            {{- end }}
          </div>
        </template>
        <template v-slot:body-cell-remark="props">
          <q-td key="remark" :props="props">
            <div>
              <q-tooltip anchor="bottom middle" self="top middle">
                {{"{{"}}props.value{{"}}"}}
              </q-tooltip>
              <div class="ellipsis">
                {{"{{"}} props.value {{"}}"}}
              </div>
            </div>
          </q-td>
        </template>
        {{- range  .Columns }}
          {{- if .DictType }}
            <template v-slot:body-cell-{{.FieldJson}}="props">
              <q-td :props="props">
                <q-chip
                        dense
                        text-color="white"
                        :color="props.value ? 'positive' : 'grey'"
                >{{"{{"}}format{{.DictType}}(props.value){{"}}"}}
                </q-chip
                >
              </q-td>
            </template>
          {{- else if  .FkTable}}
            <template v-slot:body-cell-{{.FieldJson}}="props">
              <q-td :props="props">
                <q-chip
                        dense
                        text-color="white"
                        :color="props.value ? 'positive' : 'grey'"
                >{{"{{"}}format{{.FkTableClass}}(props.value){{"}}"}}
                </q-chip
                >
              </q-td>
            </template>
          {{- else if  eq  .DataType "bool" }}
            <template v-slot:body-cell-{{.FieldJson}}="props">
              <q-td :props="props">
                <q-chip
                        dense
                        text-color="white"
                        :color="props.value ? 'positive' : 'grey'"
                >{{"{{"}} format_sys_enabled(props.value) {{"}}"}}
                </q-chip>
              </q-td>
            </template>
          {{- end }}
        {{- end }}
          {{- if .IsFlow}}
            <template v-slot:body-cell-status="props">
              <q-td :props="props">
                <q-chip
                        dense
                        text-color="white"
                        :color="getFlowStatusColor(props.value)"
                >{{"{{"}} formatFlowStatus(props.value) {{"}}"}}
                </q-chip>
              </q-td>
            </template>
          {{- end }}
        <template v-slot:body-cell-actions="props">
          <q-td :props="props">
            <div class="q-gutter-xs">
                {{- if .IsFlow }}
                  <q-btn
                      flat
                      round
                      dense
                      color="primary"
                      icon="approval"
                      @click="triggerHandle(props.row)"
              >
                <q-tooltip>工作流审批</q-tooltip></q-btn
              >
                {{- end}}

              <com-edit
                      {{- if .IsFlow }}
                      v-if="checkEdit(props.row)"
                      {{- end}}
                      v-permission="'{{.TableUrl}}:edit'"
                      label="{{.TableComment}}"
                      @confirm="edit(props.row)"
              />
              <com-del
                      {{- if .IsFlow }}
                        v-if="checkEdit(props.row)"
                      {{- end}}
                      v-permission="'{{.TableUrl}}:del'"
                      label="{{.TableComment}}"
                      @confirm="del(props.row)"
              />
            </div>
          </q-td>
        </template>
      </q-table>
    </div>
    <q-dialog  v-model="dialogVisible" persistent :full-width="isFlow">
      <q-card :class="isFlow ? 'flow-card' : 'edit-card'">
        <q-bar class="bg-primary text-white">
          <span> {{"{{"}}formType{{"}}"}}{{.TableComment}} </span>
          <q-space />
          <q-btn dense flat icon="close" v-close-popup>
            <q-tooltip class="bg-white text-primary">关闭</q-tooltip>
          </q-btn>
        </q-bar>
        <q-card-section>
            {{- if .IsFlow }}
          <q-chip v-if="isFlow" size="18px" icon="bookmark">业务数据</q-chip>
            {{- end}}
          <q-form @submit="submit">
            <div class="row q-col-gutter-x-md dialog_form q-pt-md">
              {{- range  .Columns }}
                {{- if .IsPk }}
                {{else if .IsPage}}
                {{- else if .IsEdit}}
                  {{- if .DictType }}
                    {{-  if eq .FormType "Toggle"}}
                      <q-toggle
                              label="{{.ColumnComment}}"
                              color="green"
                              class="{{.FormClass}}"
                              v-model="form.{{.FieldJson}}"
                      {{- if $isTableFlow }}
                        :disable="isFlow"
                      {{- end}}
                      />
                    {{- else  }}
                      <q-select
                              class="{{.FormClass}}"
                              outlined
                              dense
                              map-options
                              emit-value
                              option-value="value"
                              option-label="label"
                              v-model="form.{{.FieldJson}}"
                              :options="{{.DictType}}Options"
                              label="{{.ColumnComment}}"
                              {{- if .Require }}
                                :rules="[requiredRule]"
                              {{- end }}
                      {{- if $isTableFlow }}
                        :disable="isFlow"
                      {{- end}}
                      />
                    {{- end }}
                  {{- else if .FkTable}}
                      <q-select
                              class="{{.FormClass}}"
                              outlined
                              dense
                              map-options
                              emit-value
                              option-value="value"
                              option-label="label"
                              v-model="form.{{.FieldJson}}"
                              :options="{{.FkTablePackage}}Options"
                              label="{{.ColumnComment}}"
                              {{- if .Require }}
                                :rules="[requiredRule]"
                              {{- end }}
                      {{- if $isTableFlow }}
                        :disable="isFlow"
                      {{- end}}
                      />
                  {{- else if eq .FormType "Toggle"}}
                    <q-toggle
                            label="{{.ColumnComment}}"
                            color="green"
                            class="{{.FormClass}}"
                            v-model="form.{{.FieldJson}}"
                    {{- if $isTableFlow }}
                      :disable="isFlow"
                    {{- end}}/>
                  {{- else if eq .FormType "FilePick"}}
                     <q-file
                            v-model="form.{{.FieldJson}}"
                            dense
                            class="{{.FormClass}}"
                            label="上传文件-{{.ColumnComment}}"
                            @update:model-value="updateFile"
                     {{- if $isTableFlow }}
                       :disable="isFlow"
                     {{- end}}
                    />
                  {{- else if eq .FormType "Input"}}
                    <q-input
                            outlined
                            dense
                            {{- if eq .HtmlType "number" }} type="number"{{- end }}
                            class="{{.FormClass}}"
                            v-model{{- if eq .HtmlType "number" -}}.number{{- end -}}="form.{{.FieldJson}}"
                            label="{{.ColumnComment}}"
                            {{- if .Require }}
                              :rules="[requiredRule]"
                            {{- end }}
                    {{- if $isTableFlow }}
                      :disable="isFlow"
                    {{- end}}
                    />
                  {{- else if eq .FormType "Editor"}}
                    <q-input
                            outlined
                            dense
                            type="textarea"
                            class="{{.FormClass}}"
                            label="{{.ColumnComment}}"
                            v-model="form.{{.FieldJson}}"
                            {{- if .Require }}
                              :rules="[requiredRule]"
                            {{- end }}
                            {{- if $isTableFlow }}
                              :disable="isFlow"
                            {{- end}}
                    />
                  {{- else if eq .FormType "DatePick"}}
                    <q-input
                            filled
                             mask="date"
                            v-model="form.{{.FieldJson}}"
                            label="{{.ColumnComment}}"
                            {{- if .Require }}
                            :rules="[requiredRule]"
                            {{- end }}
                    {{- if $isTableFlow }}
                      :disable="isFlow"
                    {{- end}}
                    >
                      <template v-slot:append>
                        <q-icon name="event" class="cursor-pointer">
                          <q-popup-proxy cover transition-show="scale" transition-hide="scale">
                            <q-date v-model="form.{{.FieldJson}}">
                              <div class="row items-center justify-end">
                                <q-btn v-close-popup label="Close" color="primary" flat />
                              </div>
                            </q-date>
                          </q-popup-proxy>
                        </q-icon>
                      </template>
                    </q-input>
                  {{- else if eq .FormType "TimePick"}}
                    <q-input
                            filled
                            mask="fulltime"
                            v-model="form.{{.FieldJson}}"
                            label="{{.ColumnComment}}"
                            {{- if .Require }}
                      :rules="[requiredRule]"
                            {{- end }}
                    {{- if $isTableFlow }}
                      :disable="isFlow"
                    {{- end}}
                    >
                      <template v-slot:append>
                        <q-icon name="access_time" class="cursor-pointer">
                          <q-popup-proxy cover transition-show="scale" transition-hide="scale">
                            <q-time
                                    v-model="form.{{.FieldJson}}"
                                    with-seconds
                                    format24h
                            >
                              <div class="row items-center justify-end">
                                <q-btn v-close-popup label="Close" color="primary" flat />
                              </div>
                            </q-time>
                          </q-popup-proxy>
                        </q-icon>
                      </template>
                    </q-input>
                  {{- else if eq .FormType "DateTimePick"}}
                    <q-input
                            filled
                            dense
                            outlined
                            v-model="form.{{.FieldJson}}"
                            label="{{.ColumnComment}}"
                    {{- if $isTableFlow }}
                      :disable="isFlow"
                    {{- end}}
                    >
                      <template v-slot:prepend>
                        <q-icon name="event" class="cursor-pointer">
                          <q-popup-proxy
                                  cover
                                  transition-show="scale"
                                  transition-hide="scale"
                          >
                            <q-date v-model="form.{{.FieldJson}}" mask="YYYY-MM-DD HH:mm">
                              <div class="row items-center justify-end">
                                <q-btn
                                        v-close-popup
                                        label="Close"
                                        color="primary"
                                        flat
                                />
                              </div>
                            </q-date>
                          </q-popup-proxy>
                        </q-icon>
                      </template>

                      <template v-slot:append>
                        <q-icon name="access_time" class="cursor-pointer">
                          <q-popup-proxy
                                  cover
                                  transition-show="scale"
                                  transition-hide="scale"
                          >
                            <q-time
                                    v-model="form.{{.FieldJson}}"
                                    with-seconds
                                    mask="YYYY-MM-DD HH:mm:ss"
                                    format24h
                            >
                              <div class="row items-center justify-end">
                                <q-btn
                                        v-close-popup
                                        label="Close"
                                        color="primary"
                                        flat
                                />
                              </div>
                            </q-time>
                          </q-popup-proxy>
                        </q-icon>
                      </template>
                    </q-input>
                  {{- end }}
                {{- end }}
              {{- end }}
            </div>
            <div
                    {{- if .IsFlow }}
                      v-if="!isFlow"
              {{- end}}
              class="row justify-center q-pa-md">
              <q-btn
                      outline
                      color="primary"
                      icon="mdi-close-thick"
                      label="关闭"
                      v-close-popup
              />
              <q-btn
                      class="q-mx-sm"
                      color="primary"
                      icon="mdi-check-bold"
                      label="提交"
                      type="submit"
              />
            </div>
          </q-form>
        </q-card-section>
          {{- if .IsFlow }}
            <q-card-section v-if="isFlow">
              <flowTriger :flowData="form" @submitTrigger="submitTrigger" />
            </q-card-section>
          {{- end}}

      </q-card>
    </q-dialog>
  </q-page>
</template>
<script>
  export default {
    name: '{{- $db -}}{{.Table}}',
  }
</script>
<script setup>
  import {
    computed,
    onMounted,
    ref,
    reactive,
    watch,
    getCurrentInstance,
  } from "vue"
  import {
    list{{.Table}},
    create{{.Table}},
    update{{.Table}},
    delete{{.Table}},
    get{{.Table}},
    deleteList{{.Table}},
    {{- if .IsFlow }}
      trigger,
  {{- end}}

  } from "src/api/{{- $db -}}/{{.TableUrl}}"

  {{- range  .Columns }}
  {{- if  .FkTable}}
  import {
    list{{.FkTableClass}},
  } from "src/api/{{- $db -}}/{{.FkTable}}"
  {{- end }}
  {{- end }}



  import { getIds } from 'src/utils/arrayOrObject'
  import { useQuasar } from "quasar"
  import { requiredRule } from "src/utils/inputRule"
  import { DictOptions,getOptionsByList, getDictLabel,getDict } from "src/utils/dict"
  import { downloadAction } from 'src/api/manage'
  import { useRoute } from 'vue-router'

  {{- if .IsFlow }}
  import { commonEvents } from 'src/components/comfsm/script/link'
  import flowTriger from 'src/components/comfsm/components/flowTriger.vue'
  const flowStatus = ref('')
  const formatFlowStatus = (prop) => {
      return getDictLabel(dictOptions.value.flowStatus, prop)
  }
  const flowStatusOptions = ref([])
  {{- end}}

  const route = useRoute()

  const $q = useQuasar()
  let { proxy } = getCurrentInstance()
  const filterVisible = ref(false)
  const dialogVisible = ref(false)
  const dataList = ref([])
  const formType = ref("")
  const dictId = ref(0)
  const form = ref({})
  const dictOptions = ref({})
  {{- range  .Columns }}
  {{- if .DictType }}
  const {{.DictType}}Options = ref([])
  const format{{.DictType}} = (prop) => {
    if (!prop) {
      prop = false
    }
    return getDictLabel(dictOptions.value.{{.DictType}}, prop)
  }
  {{- else if .FkTable}}
  const {{.FkTablePackage}}Options = ref([])
  const format{{.FkTableClass}}= (prop) => {
    return getDictLabel({{.FkTablePackage}}Options.value, prop)
  }
  {{- end }}
  {{- end }}

  const format_sys_enabled = (prop) => {
      if (!prop) {
          prop = false
      }
      return getDictLabel(dictOptions.value.sys_enabled, prop)
  }
  const searchKey = ref('')

  const selected = ref([])
  const pagination = ref({
    sortBy: 'createdAt',
    descending: true,
    page: 1,
    rowsPerPage: 10,
    rowsNumber: 0,
    searchKey: '',
  {{- if .IsFlow }}
    flowStatus: '',
  {{- end }}
    {{- range  .Columns }}
    {{- if .DictType }}
    {{.FieldJson}}: '',
    {{- else if  .FkTable}}
    {{.FieldJson}}: '',
    {{- else if  eq  .DataType "bool" }}
    {{.FieldJson}}: '',
    {{- end }}
    {{- end }}
  })
  const columns = computed(() => {
    return [
      {{- range  .Columns -}}
      {{- if .IsPk -}}
      {{- else if .IsPage }}
      {{else if .IsList}}
      {
        name: "{{.FieldJson}}",
        align: "left",
        label: "{{.ColumnComment}}",
        field: "{{.FieldJson}}",
        sortable: true,
        classes: "ellipsis",
      },
      {{- end -}}
      {{- end -}}
      { name: 'actions', align: 'center', label: '操作', field: 'actions' },
    ]
  })


  onMounted(async () => {
    dictOptions.value = await DictOptions()
    {{- range  .Columns }}
    {{- if .DictType }}
    {{.DictType}}Options.value = await getDict('{{.DictType}}')
    {{- else if .FkTable}}
    const queryReq{{.FkTableClass}} = {
      pageIndex: 0,
      pageSize: 9999,
    }
    let res{{.FkTableClass}} = await list{{.FkTableClass}}(queryReq{{.FkTableClass}})
    {{.FkTablePackage}}Options.value = getOptionsByList(res{{.FkTableClass}}.list, "{{.FkLabelId}}", "{{.FkLabelName}}")
    {{- end }}
    {{- end }}
      {{- if .IsFlow }}
      flowStatusOptions.value = await getDict('flowStatus')
      {{- end}}
      await onRequest()
  })

  const reset = () => {
    pagination.value = {
      sortBy: 'createdAt',
      descending: true,
      page: 1,
      rowsPerPage: 10,
      rowsNumber: 0,
      searchKey: '',
      {{- if .IsFlow }}
        flowStatus: '',
      {{- end }}
      {{- range  .Columns }}
      {{- if .DictType }}
      {{.FieldJson}}: '',
      {{- else if  .FkTable}}
      {{.FieldJson}}: '',
      {{- else if  eq  .DataType "bool" }}
      {{.FieldJson}}: '',
      {{- end }}
      {{- end }}
    }
    form.value = {
      enabled: true,
    }
    dictId.value = 0
      searchKey.value = undefined
      formType.value = ''
      {{- if .IsFlow }}
      flowStatus.value = undefined
      {{- end}}
  }
  const clear = () => {
      reset()
      onRequest()
  }
  const onRequest = async () => {
    let val
    if (!val) {
      val = { pagination: pagination.value }
    }
    if (!val.pagination) {
      val.pagination = pagination.value
    }
    if (!val.filter) {
        val.filter = pagination.value.searchKey
    } else {
        pagination.value.searchKey = val.filter
    }
    const { page, rowsPerPage } = val.pagination
    let queryReq = pagination.value
    queryReq.pageSize = rowsPerPage
    queryReq.pageIndex = page

    queryReq.searchKey = val.filter


    let table = await list{{.Table}}(queryReq)

    pagination.value = val.pagination

    pagination.value.rowsNumber = table.count
    if (table.list) {
      dataList.value = table.list
    } else {
        dataList.value = []
    }
  }
  const getSelectedString = () => {
    return selected.value.length === 0
            ? ''
            : `${selected.value.length} record${
                    selected.value.length > 1 ? 's' : ''
            } selected of ${dataList.value.length}`
  }

  const filterSearch = () => {
      let val = {}
      val.pagination = pagination.value
      onRequest(val)
  }

  const create = () => {
    reset()
    formType.value = "新建"
    dialogVisible.value = true
  }
  const edit = async (p) => {
    reset()
    let req = {
        id: p.id,
    }
    form.value = await get{{.Table}}(req)

    formType.value = "编辑"
    dialogVisible.value = true
  }

  const del = async (p) => {
    await delete{{.Table}}(p)
    await onRequest()
  }

  const delList = async () => {
    let req = {
        idList: getIds(selected.value),
    }
    await deleteList{{.Table}}(req)
    await onRequest()
  }
  const submit = async () => {
    // const res = undefined
    if (formType.value === "编辑") {
      await update{{.Table}}(form.value)
    } else if (formType.value === "新建") {
      await create{{.Table}}(form.value)
    } else {
      proxy.$error("请求错误")
    }
    dialogVisible.value = false
    reset()
    await onRequest()
  }
  const updateFile = (val) => {
    form.value.upFileName = val.name
  }
  {{- if .IsImport }}
  const uploadUrl = process.env.BASE_URL + '/{{.Db}}/{{.TableUrl}}/import'
  const exportUrl = '/{{.Db}}/{{.TableUrl}}/export'
  const exportTemplateUrl = '/{{.Db}}/{{.TableUrl}}/exportTemplate'
  const handleExport = () => {
    let queryReq = {}
    let val = pagination.value

    queryReq.pageSize = val.pagination.rowsPerPage
    queryReq.pageIndex = val.pagination.page

    downloadAction(exportUrl, '{{.TableComment}}-导出.xlsx', queryReq)
  }

  const handleExportTemplate = () => {
    downloadAction(exportTemplateUrl, '{{.TableComment}}模板.xlsx')
  }

  const resetQuery = async () => {
    reset()
    await onRequest()
  }
  {{- end }}

  {{- if .IsFlow }}
  const checkEdit = (val) => {
      let ct = false
      if (val.status === 'start' || val.status === 'goback') {
          ct = true
      }
      return ct
  }
  const getFlowStatusColor = (prop) => {
      let fc = ''
      if (prop === 'end') {
          fc = 'teal'
      }
      if (prop === 'working') {
          fc = 'primary'
      }
      if (prop === 'goback') {
          fc = 'red'
      }
      if (prop === 'start') {
          fc = 'positive'
      }
      return fc
  }
const triggerHandle = async (p) => {
    let req = {
        id: p.id,
    }
    form.value = = await get{{.Table}}(req)
    formType.value = '流程审批'
    dialogVisible.value = true
}
const submitTrigger = async () => {
  await trigger(form.value.currentNode)
  dialogVisible.value = false
  reset()
  await onRequest()
}
const isFlow = computed(() => {
  let val = false
  if (formType.value === '流程审批') {
    val = true
  }
  return val
})
  {{- end}}
</script>