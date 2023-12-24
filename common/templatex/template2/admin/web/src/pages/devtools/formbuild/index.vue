<template>
  <q-layout style="z-index: 1" class="bg-page" view="hhh lpr lfr">
    <q-header class="bg-white text-black" height-hint="98">
      <q-toolbar>
        <!--  <q-btn dense flat round icon="menu" @click="left = !left" /> -->

        <!--  <q-btn dense flat round icon="menu" @click="right = !right" /> -->
      </q-toolbar>
      <q-separator />
    </q-header>
    <q-drawer show-if-above v-model="left" side="left" bordered>
      <draggable
        class="dragArea list-group"
        :list="unselectedList"
        group="people"
        @change="log"
        itemKey="column.fieldName"
      >
        <template #item="{ element }">
          <div class="list-group-item">
            <q-list bordered separator>
              <q-item clickable v-ripple>
                <q-item-section>
                  <q-item-label lines="1">{{
                    element.column.fieldName
                  }}</q-item-label>
                  <q-item-label caption>{{
                    element.column.columnComment
                  }}</q-item-label>
                </q-item-section>
              </q-item>
            </q-list>
          </div>
        </template>
      </draggable>
    </q-drawer>
    <q-drawer show-if-above v-model="right" side="right" bordered>
      <div v-if="!!currentElement.column" class="q-pa-md">
        <div class="q-pa-md">
          字段:
          {{
            currentElement.column.caption + ', ' + currentElement.column.name
          }}
        </div>

        <div v-show="type != 'list'">
          <div class="q-pa-md">栅格宽度:</div>
          <div class="q-pa-md">
            <q-slider
              v-model="currentElement.width"
              color="primary"
              :min="1"
              :step="1"
              :max="12"
              label
              label-always
            />
          </div>
        </div>
        <!-- <div class="col-3">
          <pre>{{ currentElement | jsonFormat }}</pre>
        </div> -->
      </div>
    </q-drawer>

    <q-page-container>
      <div class="q-pa-md form-build-body" :class="device">
        <div class="q-pb-md row justify-right">
          <div class="q-px-md">
            <q-radio
              @input="deviceChange"
              v-model="device"
              val="pc"
              label="电脑"
            />
            <q-radio
              @input="deviceChange"
              v-model="device"
              val="pad"
              label="平板"
            />
            <q-radio
              @input="deviceChange"
              v-model="device"
              val="phone"
              label="手机"
            />
          </div>

          <div class="q-px-md">
            <q-radio
              @input="typeChange"
              v-model="type"
              val="create"
              label="创建"
            />
            <q-radio
              @input="typeChange"
              v-model="type"
              val="update"
              label="编辑"
            />
            <q-radio
              @input="typeChange"
              v-model="type"
              val="list"
              label="列表"
            />
          </div>

          <div class="q-px-md">
            <q-btn
              class="q-mx-md"
              unelevated
              @click="onSubmitClick"
              color="primary"
              label="保存"
            />
            <q-btn
              unelevated
              @click="onDeleteClick"
              color="negative"
              label="删除"
            />
          </div>
        </div>

        <q-separator />

        <draggable
          class="dragArea list-group row"
          :list="selectedList"
          group="people"
          @change="log"
          itemKey="columnId"
        >
          <template #item="{ element }">
            <div
              class="list-group-item col-6 q-pa-md"
              :class="element"
              @click="selectForEdit(element)"
            >
              <div>
                <div v-bind:class="{ required: !element.column.nullable }">
                  {{ element.column.columnComment }}:
                </div>
                <q-input
                  v-if="isStringType(element)"
                  readonly
                  outlined
                  dense
                  :placeholder="element.column.description"
                  :type="element.isPwd ? 'password' : 'text'"
                  v-model="element.column.value"
                >
                  <template v-slot:append v-if="!element.isText">
                    <q-icon
                      :name="element.isPwd ? 'visibility_off' : 'visibility'"
                      class="cursor-pointer"
                      @click="element.isPwd = !element.isPwd"
                    />
                  </template>
                </q-input>

                <q-editor
                  readonly
                  v-else-if="isTextType(element)"
                  v-model="textValue"
                  :placeholder="element.column.description"
                >
                </q-editor>

                <q-input
                  v-else-if="isDateTimeType(element)"
                  outlined
                  dense
                  readonly
                >
                  <template v-slot:prepend>
                    <q-icon name="event" class="cursor-pointer">
                      <q-popup-proxy
                        ref="qDateProxy"
                        transition-show="scale"
                        transition-hide="scale"
                      >
                        <q-date
                          mask="YYYY-MM-DD HH:mm:ss"
                          @input="hideRefPopProxyAction('qDateProxy')"
                        />
                      </q-popup-proxy>
                    </q-icon>
                  </template>

                  <template v-slot:append>
                    <q-icon name="access_time" class="cursor-pointer">
                      <q-popup-proxy
                        ref="qTimeProxy"
                        transition-show="scale"
                        transition-hide="scale"
                      >
                        <q-time
                          mask="YYYY-MM-DD HH:mm:ss"
                          format24h
                          with-seconds
                          @input="hideRefPopProxyAction('qTimeProxy')"
                        />
                      </q-popup-proxy>
                    </q-icon>
                  </template>
                </q-input>

                <q-input
                  v-else-if="isDateType(element)"
                  outlined
                  dense
                  readonly
                >
                  <template v-slot:append>
                    <q-icon name="event" class="cursor-pointer">
                      <q-popup-proxy
                        ref="qDateProxy"
                        transition-show="scale"
                        transition-hide="scale"
                      >
                        <q-date
                          mask="YYYY-MM-DD"
                          @input="hideRefPopProxyAction('qDateProxy')"
                        />
                      </q-popup-proxy>
                    </q-icon>
                  </template>
                </q-input>

                <q-input
                  v-else-if="isTimeType(element)"
                  outlined
                  dense
                  readonly
                >
                  <template v-slot:append>
                    <q-icon name="access_time" class="cursor-pointer">
                      <q-popup-proxy
                        ref="qTimeProxy"
                        transition-show="scale"
                        transition-hide="scale"
                      >
                        <q-time
                          mask="HH:mm:ss"
                          format24h
                          with-seconds
                          @input="hideRefPopProxyAction('qTimeProxy')"
                        />
                      </q-popup-proxy>
                    </q-icon>
                  </template>
                </q-input>

                <q-toggle
                  v-else-if="isBoolType(element)"
                  readonly
                  v-model="element.column.value"
                >
                </q-toggle>

                <q-input
                  readonly
                  outlined
                  dense
                  v-else-if="isNumberType(element)"
                  :placeholder="element.column.description"
                  type="number"
                  v-model="element.column.value"
                >
                </q-input>

                <CFile
                  v-else-if="isAttachmentType(element)"
                  v-model="element.column.value"
                >
                </CFile>

                <q-input
                  v-else
                  readonly
                  outlined
                  dense
                  :placeholder="element.column.description"
                  :type="element.isPwd ? 'password' : 'text'"
                  v-model="element.column.value"
                >
                  <template v-slot:append v-if="!element.isText">
                    <q-icon
                      :name="element.isPwd ? 'visibility_off' : 'visibility'"
                      class="cursor-pointer"
                      @click="element.isPwd = !element.isPwd"
                    />
                  </template>
                </q-input>
              </div>
              <div class="row reverse editable-element-action-buttons">
                <div class="justify-end q-pt-xs">
                  <q-btn
                    @click="deleteElement(element)"
                    v-if="isSelectedForEdit(element)"
                    class="editable-element-button"
                    color="red"
                    icon="delete"
                    round
                    unelevated
                    size="xs"
                  >
                    <q-tooltip>移除</q-tooltip>
                  </q-btn>
                </div>
              </div>
            </div>
          </template>
        </draggable>
      </div>
    </q-page-container>
  </q-layout>
</template>

<script setup>
import {
  computed,
  onMounted,
  onActivated,
  onDeactivated,
  onUpdated,
  onUnmounted,
  ref,
  reactive,
  watch,
  getCurrentInstance,
} from 'vue'
import draggable from 'vuedraggable'

import { requiredRule } from 'src/utils/inputRule'
import { useQuasar, format } from 'quasar'
import { useRoute, useRouter } from 'vue-router'

import {
  DictOptions,
  getOptionsByList,
  getDictLabel,
  getDict,
} from 'src/utils/dict'
import { listDict } from 'src/api/admin/dictType'
import { downloadAction } from 'src/api/manage'

import {
  listDatasource,
  getTables,
  getTable,
  getColumns,
  previewCode,
  genCode,
  gendb,
  genCoverCode,
} from 'src/api/devtools/datasource'
import { getGenTable } from 'src/api/devtools/table'
import { extend } from 'quasar'

const list1 = ref([
  { name: 'John', id: 1 },
  { name: 'Joao', id: 2 },
  { name: 'Jean', id: 3 },
  { name: 'Gerard', id: 4 },
])

const route = useRoute()
const $q = useQuasar()
let { proxy } = getCurrentInstance()
const router = useRouter()
const tag = ref('')

const db = ref(null)
const sourceList = ref([])
const tableList = ref([])
const fieldList = ref([])
const fieldOptions = ref([])
const htmlOptions = ref([])

const dictList = ref([])
const columnIndex = ref(0)
const dbVisable = ref(true)

const tableDetail = ref({})

const dataSource = ref('')
const dataSourceUrl = ref('')
const left = ref(true)
const right = ref(true)
const unselectedList = ref([])
const selectedList = ref([])
const loading = ref(true)
const type = ref('create')
const device = ref('pc')
const table = ref({})
const formBuilders = ref([])
const currentElement = ref({})
const textValue = ref('')
const fileValue = ref('')

const init = async () => {
  htmlOptions.value = await getDict('htmlType')
  fieldOptions.value = await getDict('filed_type')

  const queryReq = {
    pageIndex: 0,
    pageSize: 9999,
  }
  let res = await listDict(queryReq)
  dictList.value = getOptionsByList(res.list, 'name', 'type')

  // onRequest()
  listsource()
  let id = route.query.id
  let tablePrefix = route.query.tablePrefix
  if (id) {
    dbVisable.value = false
    getTableById(id, tablePrefix)
  }
}
const listsource = async () => {
  let res = await listDatasource()
  sourceList.value = res.list
}
const listTables = async (val) => {
  let data = {
    id: val.id,
  }
  let res = await getTables(data)
  tableList.value = res.list
}
onMounted(() => {
  if (route.name === 'formbuild') {
    if (!route.query.id) {
      router.push({ name: 'gencode_history' })
    }
  }
  init()
})

onActivated(() => {
  console.info('activated')
})

onDeactivated(() => {
  console.info('deactivated')
})

onUpdated(() => {
  console.info('updated')
})

onUnmounted(() => {
  console.info('destroyed')
})

const iconFormat = (formElement) => {
  const value = formElement.column
  let icon = 'text_rotate_vertical'
  if (value.dataType === 'PASSWORD') {
    icon = 'lock'
  } else if (value.dataType === 'DATETIME') {
    icon = 'calendar_today'
  } else if (value.dataType === 'DATE') {
    icon = 'event'
  } else if (value.dataType === 'TIME') {
    icon = 'schedule'
  } else if (value.dataType === 'TIME') {
    icon = 'schedule'
  } else if (value.dataType === 'ATTACHMENT') {
    icon = 'upload_file'
  } else if (value.dataType === 'LONGBLOB') {
    icon = 'attachment'
  } else if (value.dataType === 'DECIMAL') {
    icon = 'format_list_numbered'
  }

  return icon
}

const classFormat = (formElement, currentElement) => {
  let value = ''
  if (formElement.width) {
    value = 'col-' + formElement.width
  }

  if (
    currentElement &&
    currentElement.column &&
    currentElement.column.id === formElement.column.id
  ) {
    value += ' selected'
  }

  //console.log(formElement.column.name + ": " + value);
  return value
}

const jsonFormat = (value) => {
  return JSON.stringify(value, null, 2)
}

const log = (evt) => {
  window.console.log(evt)
}

const isTextType = (formElement) => {
  if (
    formElement.column.dataType === 'TEXT' ||
    formElement.column.dataType === 'LONGTEXT'
  ) {
    return true
  } else {
    return false
  }
}

const isStringType = (formElement) => {
  if (
    formElement.column.dataType === 'CHAR' ||
    formElement.column.dataType === 'VARCHAR'
  ) {
    return true
  } else {
    return false
  }
}

const isNumberType = (formElement) => {
  if (
    formElement.column.dataType === 'TINYINT' ||
    formElement.column.dataType === 'SMALLINT' ||
    formElement.column.dataType === 'MEDIUMINT' ||
    formElement.column.dataType === 'INT' ||
    formElement.column.dataType === 'BIGINT'
  ) {
    return true
  } else {
    return false
  }
}

const isDateType = (formElement) => {
  if (formElement.column.dataType === 'DATE') {
    return true
  } else {
    return false
  }
}

const isTimeType = (formElement) => {
  if (formElement.column.dataType === 'TIME') {
    return true
  } else {
    return false
  }
}

const isDateTimeType = (formElement) => {
  if (formElement.column.dataType === 'DATETIME') {
    return true
  } else {
    return false
  }
}

const isBoolType = (formElement) => {
  if (formElement.column.dataType === 'bool') {
    return true
  } else {
    return false
  }
}

const isAttachmentType = (formElement) => {
  if (formElement.column.dataType === 'ATTACHMENT') {
    return true
  } else {
    return false
  }
}

const hideRefPopProxyAction = (ref) => {}

const selectForEdit = (formElement) => {
  currentElement.value = formElement
}
const isSelectedForEdit = (formElement) => {
  return (
    currentElement.value &&
    currentElement.value.column &&
    currentElement.value.column.id === formElement.column.id
  )
}
const currentElementWidthInput = (value) => {
  // console.log('currentElementWidthInput')
  proxy.$forceUpdate()
}
const deleteElement = (formElement) => {
  const index = selectedList.value.findIndex(
    (t) => t.column.id === formElement.column.id
  )
  if (index >= 0) {
    selectedList.value = [
      ...selectedList.value.slice(0, index),
      ...selectedList.value.slice(index + 1),
    ]
    unselectedList.value.push(formElement)
  }
  currentElement.value = {}
  proxy.$forceUpdate()
}

const getTableById = async (id, tablePrefix) => {
  tableDetail.value = {}
  let req = {
    id: id,
  }
  let res = await getGenTable(req)
  tableDetail.value = res
  table.value = res
  tableDetail.value.tablePrefix = tablePrefix
  tableDetail.value.columns.forEach((row, index) => {
    row.sort = index
  })
  db.value = {}
  db.value.id = res.sourceId
  let reqTables = {
    id: res.sourceId,
  }
  await listTables(reqTables)
  setFormBuilder()
}
const loadData = async (id) => {
  $q.loading.show({
    message: '加载中',
  })
  try {
    loading.value = true
    const tableId = id || route.query.id
    let req = {
      id: id,
    }
    const tabled = await getGenTable(req)
    table.value = tabled

    let query = {
      tableId: tableId,
    }
    // formBuilders.value = await tableService.list(
    //   dataSource.value,
    //   'tableFormBuilder',
    //   0,
    //   999,
    //   null,
    //   query,
    //   null
    // )

    setFormBuilder()

    loading.value = false
    $q.loading.hide()
  } catch (error) {
    console.error(error)
    loading.value = false
    $q.loading.hide()
    $q.notify(error)
  }
}
const deviceChange = (value, evt) => {
  // console.log(value)
  currentElement.value = {}
  setFormBuilder()
}
const typeChange = (value, evt) => {
  // console.log(value)
  currentElement.value = {}
  setFormBuilder()
}
const setFormBuilder = () => {
  const columns = table.value.columns
  let formBuilder = formBuilders.value.find(
    (t) => t.device === device.value && t.type === type.value
  )

  let unselectedLs = []
  let selectedLs = []
  // console.log(columns, formBuilder)
  if (!formBuilder) {
    columns.forEach((column) => {
      let formElement = {
        columnId: column.id,
        column: column,
        width: 12,
      }

      // if (column.dataType === 'PASSWORD') {
      //   formElement.isText = false
      //   formElement.isPwd = true
      // } else {
      //   formElement.isText = true
      //   formElement.isPwd = false
      // }

      if (column.isEdit) {
        selectedLs.push(formElement)
      } else {
        unselectedLs.push(formElement)
      }
    })
  } else {
    const rowSelectedList = JSON.parse(formBuilder.body)
    console.dir(rowSelectedList)

    rowSelectedList.forEach((formElement) => {
      if (columns.find((t) => t.id === formElement.columnId)) {
        selectedLs.push(formElement)
      }
    })
    console.dir(selectedLs)

    selectedLs.forEach((formElement) => {
      formElement.column = columns.find((t) => t.id === formElement.columnId)
    })

    columns.forEach((column) => {
      let formElement = {
        columnId: column.id,
        column: column,
        width: 12,
      }

      // if (column.dataType === 'PASSWORD') {
      //   formElement.isText = false
      //   formElement.isPwd = true
      // } else {
      //   formElement.isText = true
      //   formElement.isPwd = false
      // }

      if (
        selectedLs.findIndex((t) => t.columnId === formElement.columnId) < 0
      ) {
        unselectedLs.push(formElement)
      }
    })
  }

  unselectedList.value = unselectedLs
  selectedList.value = selectedLs
  // console.log(unselectedList.value)
  // console.log(selectedList.value)
}

const onSubmitClick = async () => {
  $q.loading.show({
    message: '提交中',
  })
  try {
    let selectedLs = extend(true, [], selectedList.value)
    selectedLs.forEach((t) => {
      delete t.column
    })

    const data = {
      name: device.value + ' ' + type.value,
      device: device.value,
      type: type.value,
      body: JSON.stringify(selectedLs),
      tableId: table.value.id,
    }

    let formBuilder = formBuilders.value.find(
      (t) => t.device === device.value && t.type === type.value
    )
    // if (!formBuilder) {
    //   await tableService.create(dataSource.value, 'tableFormBuilder', data)
    // } else {
    //   await tableService.update(
    //     dataSource.value,
    //     'tableFormBuilder',
    //     formBuilder.id,
    //     data
    //   )
    // }

    $q.loading.hide()
    $q.notify('保存成功')
    await loadData(table.value.id)
  } catch (error) {
    $q.loading.hide()
    console.info(error)
  }
}
const onDeleteClick = async () => {
  // $q.loading.show({
  //   message: '提交中',
  // })
  // try {
  //   let formBuilder = this.formBuilders.find((t) => t.type === this.type)
  //   if (formBuilder) {
  //     await tableService.delete(
  //       dataSource,
  //       'tableFormBuilder',
  //       formBuilder.id
  //     )
  //     $q.loading.hide()
  //     $q.notify('删除成功')
  //     await loadData(this.table.id)
  //   } else {
  //     $q.loading.hide()
  //     $q.notify('已经为空')
  //   }
  // } catch (error) {
  //   $q.loading.hide()
  //   console.info(error)
  // }
}
</script>

<style lang="scss" scoped>
// SCSS equivalent of the given Stylus code

.dragArea {
  min-height: 50px;
}

.q-item {
  cursor: move !important;
}

.list-group {
  min-height: 20px;

  &-item {
    cursor: move;
  }
}

.sortable-chosen {
  opacity: 0.3;
  background: $primary;
}

.sortable-ghost {
  opacity: 0.5;
  background: #c8ebfb;
}

.selected {
  background: #c8ebfb;
}

.form-build-body.phone {
  max-width: 767px;
}

.form-build-body.pad {
  max-width: 979px;
}
</style>
