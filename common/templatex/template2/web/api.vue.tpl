{{- $db :=.Db }}
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
              v-model:filter="searchKey"
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
            <com-search @query="onRequest" v-model:filter="searchKey"/>
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
                        @upload="uploadFn"
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
              <q-icon name="help_outline" class="text-purple cursor-pointer">
                <q-popup-proxy :offset="[10, 10]">
                  <q-banner class="bg-purple text-white">
                    <template v-slot:avatar>
                      <q-icon name="help" />
                    </template>
                    {{`{{ route.meta.remark }}`}}
                  </q-banner>
                </q-popup-proxy>
              </q-icon>
            </q-btn-group>
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
          {{- end }}
        {{- end }}

        <template v-slot:body-cell-actions="props">
          <q-td :props="props">
            <div class="q-gutter-xs">
              <com-edit
                      v-permission="'{{.TableUrl}}:edit'"
                      label="{{.TableComment}}"
                      @confirm="edit(props.row)"
              />
              <com-del
                      v-permission="'{{.TableUrl}}:del'"
                      label="{{.TableComment}}"
                      @confirm="del(props.row)"
              />
            </div>
          </q-td>
        </template>
      </q-table>
    </div>
    <q-dialog v-model="dialogVisible" persistent>
      <q-card style="width: 800px; min-width: 40vw">
        <q-bar class="bg-primary text-white">
          <span> {{"{{"}}formType{{"}}"}}{{.TableComment}} </span>
          <q-space />
          <q-btn dense flat icon="close" v-close-popup>
            <q-tooltip class="bg-white text-primary">关闭</q-tooltip>
          </q-btn>
        </q-bar>
        <q-card-section>
          <q-form @submit="submit">
            <div class="row q-col-gutter-x-md dialog_form q-pt-md">
              {{- range  .Columns }}
                {{- if .IsPk }}
                {{else if .IsPage}}
                {{- else if .IsEdit}}
                  {{- if .DictType }}
                    {{-  if eq .DataType "bool"}}
                      <q-toggle label="{{.ColumnComment}}" color="green" v-model="form.{{.FieldJson}}" />
                    {{- else  }}
                      <q-select
                              class="col-6 q-pb-md"
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
                      />
                    {{- end }}
                  {{- else if .FkTable}}
                    {{-  if eq .DataType "bool"}}
                      <q-toggle label="{{.ColumnComment}}" color="green" v-model="form.{{.FieldJson}}" />
                    {{- else  }}
                      <q-select
                              class="col-6 q-pb-md"
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
                      />
                    {{- end }}
                  {{- else if eq .DataType "bool"}}
                    <q-toggle label="{{.ColumnComment}}" color="green" v-model="form.{{.FieldJson}}" />
                  {{- else }}
                    <q-input
                            outlined
                            dense
                            {{- if eq .HtmlType "number" }} type="number"{{- end }}
                            class="col-6 q-pb-md"
                            v-model{{- if eq .HtmlType "number" -}}.number{{- end -}}="form.{{.FieldJson}}"
                            label="{{.ColumnComment}}"
                            {{- if .Require }}
                              :rules="[requiredRule]"
                            {{- end }}
                    />
                  {{- end }}
                {{- end }}
              {{- end }}
            </div>
            <div class="row justify-center q-pa-md">
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
  } from "src/api/{{- $db -}}/{{.TableUrl}}"

  {{- range  .Columns }}
  {{- if  .FkTable}}
  import {
    list{{.FkTableClass}},
  } from "src/api/{{- $db -}}/{{.FkTable}}"
  {{- end }}
  {{- end }}

  import { useQuasar } from "quasar"
  import { requiredRule } from "src/utils/inputRule"
  import { DictOptions,getOptionsByList, getDictLabel,getDict } from "src/utils/dict"
  import { downloadAction } from 'src/api/manage'
  import { useRoute } from 'vue-router'
  const route = useRoute()

  const $q = useQuasar()
  let { proxy } = getCurrentInstance()
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

  const searchKey = ref('')

  const selected = ref([])
  const pagination = ref({
    sortBy: "id",
    descending: false,
    page: 1,
    rowsPerPage: 10,
    rowsNumber: 0,
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
    onRequest()
  })

  const reset = () => {
    pagination.value = {
      sortBy: "id",
      descending: false,
      page: 1,
      rowsPerPage: 10,
      rowsNumber: 0,
    }
    form.value = {
      enabled: true,
    }
    dictId.value = 0
  }

  const onRequest = async (val) => {
    if (!val) {
      val = { pagination: pagination.value }
    }
    if (!val.pagination) {
      val.pagination = pagination.value
    }
    if (!val.filter) {
      val.filter = searchKey.value
    }
    const { page, rowsPerPage, sortBy, descending } = val.pagination
    let queryReq = {}
    queryReq.pageSize = rowsPerPage
    queryReq.pageIndex = page
    queryReq.sortBy = sortBy
    queryReq.descending = descending
    queryReq.searchKey = val.filter

    let table = await list{{.Table}}(queryReq)

    pagination.value = val.pagination

    pagination.value.rowsNumber = table.count
    if (table.list) {
      dataList.value = table.list
    }
  }
  const getSelectedString = () => {
    return selected.value.length === 0
            ? ''
            : `${selected.value.length} record${
                    selected.value.length > 1 ? 's' : ''
            } selected of ${dataList.value.length}`
  }

  const create = () => {
    reset()
    formType.value = "新建"
    dialogVisible.value = true
  }
  const edit = (p) => {
    reset()
    form.value = {
      ...p,
    }
    formType.value = "编辑"
    dialogVisible.value = true
  }

  const del = async (p) => {
    let res = await delete{{.Table}}(p)
    onRequest()
  }

  const delList = async () => {
    // let res = await deleteDatasource()
    // onRequest()
    let req = {
      ids: getIds(selected.value),
    }
    let res = await delete{{.Table}}(req)
    onRequest()
  }
  const submit = async () => {
    // const res = undefined
    if (formType.value === "编辑") {
      let res = await update{{.Table}}(form.value)
    } else if (formType.value === "新建") {
      let res = await create{{.Table}}(form.value)
    } else {
      proxy.$error("请求错误")
    }
    dialogVisible.value = false
    reset()
    onRequest()
  }
  {{- if .IsImport }}
  const uploadUrl = process.env.BASE_URL + '/{{.Db}}/{{.TableUrl}}/import'
  const exportUrl = '/{{.Db}}/{{.TableUrl}}/export'
  const exportTemplateUrl = '/{{.Db}}/{{.TableUrl}}/exportTemplate'
  const handleExport = () => {
    let queryReq = {}
    let val = {}
    val.pagination = pagination.value
    queryReq.pageSize = val.pagination.rowsPerPage
    queryReq.pageIndex = val.pagination.page
    queryReq.sortBy = val.pagination.sortBy
    queryReq.descending = val.pagination.descending
    queryReq.searchKey = searchKey.value
    downloadAction(exportUrl, '{{.TableComment}}-导出.xlsx', queryReq)
  }

  const handleExportTemplate = () => {
    downloadAction(exportTemplateUrl, '{{.TableComment}}模板.xlsx')
  }

  const uploadFn = async (val) => {
    reset()
    onRequest()
  }
  {{- end }}
</script>