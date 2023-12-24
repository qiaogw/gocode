<template>
  <q-page>
    <div class="col shadow-2 q-pa-md">
      <q-table
        dense
        color="primary"
        class="cross_table"
        separator="cell"
        virtual-scroll
        :columns="columns"
        :rows="menuTree"
        row-key="id"
        @request="onRequest"
        hide-bottom
        v-model:filter="searchKey"
        :rows-per-page-options="[0]"
        :grid="$q.screen.xs"
      >
        <template v-slot:top="table">
          <div class="row no-wrap full-width">
            <com-search @query="onRequest" v-model:filter="searchKey" />
            <q-space />
            <div class="q-gutter-xs">
              <q-btn
                color="primary"
                label="切换全屏"
                no-wrap
                v-show="$q.screen.gt.md"
                @click="table.toggleFullscreen"
                :icon="table.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
              />
              <q-btn
                icon="add"
                no-wrap
                color="primary"
                label="新建"
                @click="add"
              />
            </div>
          </div>
        </template>

        <template v-slot:body="props">
          <mentItem
            :menu="props.row"
            :menuProps="props"
            @edit="edit"
            @addChild="addChild"
            @del="del"
            @toggleExpand="toggleExpand"
          />
        </template>
      </q-table>
    </div>

    <q-dialog v-model="dialogVisible" persistent>
      <q-card style="width: 800px; min-width: 60vw">
        <q-bar class="bg-primary text-white">
          <span>{{ formType }}菜单</span>
          <q-space />
          <q-btn dense flat icon="close" v-close-popup>
            <q-tooltip class="bg-white text-primary">Close</q-tooltip>
          </q-btn>
        </q-bar>
        <q-card-section>
          <q-form @submit="submit">
            <div>
              <div class="row q-col-gutter-x-md dialog_form q-pt-md">
                <q-select
                  class="col-6 q-pb-md"
                  outlined
                  dense
                  filled
                  use-chips
                  emit-value
                  map-options
                  option-value="id"
                  option-label="title"
                  v-model="form.parentId"
                  :options="parentList"
                  label="上级菜单"
                />
                <q-select
                  class="col-6"
                  outlined
                  dense
                  filled
                  use-chips
                  emit-value
                  map-options
                  v-model="form.type"
                  :options="menu_typeOptions"
                  label="菜单类型"
                />
                <q-input
                  dense
                  outlined
                  class="col-6"
                  v-model="form.title"
                  label="显示"
                  :rules="[requiredRule]"
                />
                <q-input
                  dense
                  outlined
                  class="col-6"
                  v-model="form.name"
                  label="名称"
                  :rules="[requiredRule]"
                />
                <q-input
                  dense
                  outlined
                  class="col-6"
                  v-model="form.icon"
                  label="图标"
                  ><template v-slot:prepend>
                    <q-icon :name="form.icon" />
                  </template>
                  <template v-slot:append>
                    <q-icon name="insert_emoticon" class="cursor-pointer">
                      <q-popup-proxy v-model="iconVisible">
                        <Iconpicker v-model:getIcon="form.icon" inSet />
                      </q-popup-proxy>
                    </q-icon>
                  </template>
                </q-input>

                <q-input
                  dense
                  outlined
                  class="col-6"
                  v-model="form.path"
                  :rules="[requiredRule]"
                  label="路由路径"
                />
                <q-input
                  dense
                  outlined
                  class="col-6"
                  v-model="form.component"
                  :rules="[requiredRule]"
                  label="组件路径"
                />
                <q-input
                  dense
                  outlined
                  class="col-6"
                  v-model.number="form.sort"
                  type="number"
                  :rules="[requiredRule]"
                  label="菜单排序"
                />
              </div>

              <div class="row q-gutter-md justify-center">
                <q-toggle label="外链" color="green" v-model="form.isFrame" />
                <q-toggle label="隐藏" color="green" v-model="form.hidden" />
                <q-toggle
                  label="缓存"
                  v-model="form.keepAlive"
                  :options="[
                    { label: '缓存', value: false },
                    { label: '显示', value: true },
                  ]"
                />
              </div>
              <div class="row q-gutter-md justify-center">
                <q-input
                  dense
                  outlined
                  class="col-6"
                  v-model="form.remark"
                  label="目录说明"
                  type="textarea"
                />
              </div>

              <q-separator class="q-ma-sm" />
              <q-btn
                label="添加crud按钮"
                icon="add"
                dense
                class="q-mb-xs"
                size="xs"
                color="secondary"
                @click="addCrud"
              />
              <q-btn
                label="添加按钮"
                icon="edit_road"
                dense
                class="q-mb-xs"
                size="xs"
                color="primary"
                @click="addButton"
              />
              <div
                v-for="(item, index) in form.button"
                :key="index"
                class="row q-col-gutter-x-xs"
              >
                <q-input
                  dense
                  outlined
                  class="col-3"
                  v-model="item.title"
                  label="显示"
                  :rules="[requiredRule]"
                />
                <q-input
                  dense
                  outlined
                  class="col-3"
                  v-model="item.name"
                  label="名称"
                  :rules="[requiredRule]"
                />
                <q-input
                  dense
                  outlined
                  class="col-3"
                  v-model="item.icon"
                  label="图标"
                  ><template v-slot:prepend>
                    <q-icon :name="item.icon" />
                  </template>
                  <template v-slot:append>
                    <q-icon name="insert_emoticon" class="cursor-pointer">
                      <q-popup-proxy v-model="buttonIconVisible[index]">
                        <Iconpicker v-model:getIcon="item.icon" inSet />
                      </q-popup-proxy>
                    </q-icon>
                  </template>
                </q-input>
                <com-del
                  class="q-mb-md"
                  label="按钮"
                  @confirm="delButton(item)"
                />
              </div>
            </div>
            <q-separator class="q-ma-sm" />
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

<script setup>
import { computed, onMounted, ref, reactive, watch } from 'vue'
import {
  listMenu,
  createMenu,
  deleteMenu,
  updateMenu,
} from 'src/api/admin/menu'
import { ArrayToTree } from 'src/utils/arrayAndTree'
import mentItem from './item.vue'
import { requiredRule } from 'src/utils/inputRule'
import { getDict, getDictLabel } from 'src/utils/dict'

import Iconpicker from 'src/components/comIcon/index.vue'

const dialogVisible = ref(false)
const iconVisible = ref(false)
const buttonIconVisible = ref([])
const menuList = ref([])
const parentList = ref([])

const orginMenuList = ref([])
const menuTree = ref([])
const parent = ref({})
const form = ref({})
const formType = ref('')
const searchKey = ref('')
const pagination = ref({
  sortBy: 'sort',
  descending: false,
  page: 1,
  rowsPerPage: 99999,
  rowsNumber: 0,
})
const columns = computed(() => {
  return [
    {
      name: 'index',
      align: 'center',
      label: '展开/收缩',
      field: 'index',
    },
    {
      name: 'title',
      align: 'left',
      label: '菜单标题',
      field: 'title',
    },
    {
      name: 'name',
      align: 'left',
      label: '菜单名称',
      field: 'name',
    },
    {
      name: 'icon',
      align: 'center',
      label: '图标',
      field: 'icon',
    },
    {
      name: 'type',
      align: 'left',
      label: '菜单类型',
      field: 'type',
    },
    {
      name: 'path',
      align: 'left',
      label: '访问路径',
      field: 'path',
    },
    {
      name: 'component',
      align: 'left',
      label: '组件路径',
      field: 'component',
    },
    {
      name: 'hidden',
      align: 'center',
      label: '是否隐藏',
      field: 'hidden',
    },
    {
      name: 'keepAlive',
      align: 'center',
      label: '缓存',
      field: 'keepAlive',
    },
    {
      name: 'button',
      align: 'center',
      label: '按钮',
      field: 'button',
    },
    {
      name: 'opt',
      align: 'center',
      label: '操作',
      field: 'opt',
    },
  ]
})
const menu_typeOptions = ref({})

onMounted(async () => {
  menu_typeOptions.value = await getDict('menu_type')
  menuList.value = []
  // pagination.value.rowsPerPage = 99999;
  onRequest()
})

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

  let table = await listMenu(queryReq)
  orginMenuList.value = menuTree.value
  menuTree.value = []
  parentList.value = table.list.filter((o) => o.type === 'D')
  menuList.value = ArrayToTree(table.list, 'id', 'parentId')
  initMenuList(menuList.value, 1, { expand: true })
}

const reset = () => {
  pagination.value = {
    sortBy: 'sort',
    descending: false,
    page: 1,
    rowsPerPage: 99999,
    rowsNumber: 0,
  }
  form.value = {
    enabled: true,
    sort: 10,
  }
}

const add = () => {
  reset()
  formType.value = '新建'
  dialogVisible.value = true
}
const edit = (p) => {
  form.value = {
    ...p,
  }
  formType.value = '编辑'
  dialogVisible.value = true
}
const editBefore = (row) => {
  getMenuById(menuList.value, row.parentId)
  return true
}
const getMenuById = (menuList, id) => {
  // 根据Id把parent找出来
  menuList.forEach((menu) => {
    if (menu.id === id) {
      parent.value = menu
    }
    if (menu.children && menu.children.length > 0) {
      getMenuById(menu.children, id)
    }
  })
}
const del = async (p) => {
  let res = await deleteMenu(p)
  onRequest()
}
const addChild = (p) => {
  reset()
  formType.value = '新建'
  form.value.parentId = p.id
  form.value.type = 'M'
  dialogVisible.value = true
}
const initMenuList = (list, level, parent) => {
  list.forEach((m) => {
    // 尝试还原下原来的状态
    const orgin = orginMenuList.value.find((o) => m.id === o.id)
    const expand = orgin && orgin.expand
    menuTree.value.push({
      ...m,
      level,
      expand,
      parent,
    })
    if (m.children && m.children.length > 0) {
      initMenuList(m.children, level + 1, { expand })
    }
  })
}
const toggleExpand = (menu) => {
  const expand = !menu.expand
  menuTree.value.forEach((m) => {
    if (m.id === menu.id) {
      m.expand = expand
    }
    // 找到子节点
    if (m.parentId === menu.id) {
      m.parent.expand = expand
      if (expand) {
        expandAllSon(menu.id)
      }
      // 折叠的话需要把所有后代后收缩起来
      if (!expand) {
        folderAllSon(menu.id)
      }
    }
  })
}
const expandAllSon = (id) => {
  // 展开所有子节点
  menuTree.value.forEach((m) => {
    if (m.parentId === id) {
      m.parent.expand = true
      m.expand = true
      expandAllSon(m.id)
    }
  })
}
const folderAllSon = (id) => {
  menuTree.value.forEach((m) => {
    if (m.parentId === id) {
      m.parent.expand = false
      folderAllSon(m.id)
    }
  })
}
const submit = async () => {
  // const res = undefined;
  if (formType.value === '编辑') {
    let res = await updateMenu(form.value)
  } else if (formType.value === '新建') {
    let res = await createMenu(form.value)
  } else {
    proxy.$error('路径错误')
  }
  dialogVisible.value = false
  reset()
  onRequest({
    pagination: pagination.value,
  })
}
const delButton = (b) => {
  form.value.button.pop(b)
}
const addButton = () => {
  let button = { type: 'B', path: '/', name: form.value.name + ':' }
  if (!form.value.button) {
    form.value.button = []
  }
  form.value.button.push(button)
}
const addCrud = () => {
  let button1 = {
    type: 'B',
    path: '/',
    title: '添加',
    name: form.value.name + ':add',
    icon: 'add',
  }
  let button2 = {
    type: 'B',
    path: '/',
    title: '编辑',
    name: form.value.name + ':edit',
    icon: 'edit',
  }
  let button3 = {
    type: 'B',
    path: '/',
    title: '删除',
    name: form.value.name + ':del',
    icon: 'delete',
  }
  if (!form.value.button) {
    form.value.button = []
  }
  form.value.button.push(button1, button2, button3)
}

watch(
  form,
  (oldItem, newItem) => {
    iconVisible.value = false
    buttonIconVisible.value = []
  },
  { deep: true } //开启deep:true模式
)
</script>
