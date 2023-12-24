<template>
  <q-card-section>
    <q-scroll-area :thumb-style="thumbStyle" style="width: 100%; height: 70vh">
      <q-tree
        :nodes="menuTree"
        node-key="id"
        label-key="title"
        tick-strategy="leaf"
        v-model:selected="selected"
        v-model:ticked="ticked"
        v-model:expanded="expanded"
      >
        <template v-slot:default-header="prop">
          <div class="row q-gutter-xs no-wrap full-width q-pr-lg">
            <q-icon :name="prop.node.icon" />
            <q-chip>{{ prop.node.title }}</q-chip>
            <q-space />
            <q-badge
              v-if="prop.node.button"
              transparent
              color="warning"
              class="q-mr-md"
              ><q-icon name="radio_button_checked" />按钮:</q-badge
            >
            <q-checkbox
              v-for="item in prop.node.button"
              :key="item.id"
              dense
              v-model="ticked"
              :label="item.title"
              :value="item.id"
              :val="item.id"
              color="cyan"
            />
          </div>
        </template>
      </q-tree>
    </q-scroll-area>
  </q-card-section>
  <div class="justify-center row q-gutter-x-xs" style="width: 100%">
    <q-btn color="primary" @click="handleClear">
      {{ $t('Clear') + $t('All') }}
    </q-btn>
    <q-btn color="primary" @click="handleAll">
      {{ $t('Select') + $t('All') }}
    </q-btn>
    <q-btn
      outline
      color="primary"
      icon="mdi-close-thick"
      label="关闭"
      v-close-popup
    />
    <q-btn color="primary" @click="handleSave">
      {{ $t('Save') }}
    </q-btn>
  </div>
</template>

<script setup>
import {
  computed,
  onMounted,
  ref,
  toRefs,
  watch,
  getCurrentInstance,
} from 'vue'
import { listMenu } from 'src/api/admin/menu'
import { setRoleMenu } from 'src/api/admin/role'
import { ArrayToTree } from 'src/utils/arrayAndTree'
const thumbStyle = ref({
  right: '2px',
  borderRadius: '5px',
  backgroundColor: '#027be3',
  width: '5px',
  opacity: 0.75,
})

const props = defineProps({
  row: {
    type: Object,
    required: true,
  },
})
const { row } = toRefs(props)

const selected = ref('')
const ticked = ref([])
const expanded = ref([])
const menuTree = ref([])
const menuList = ref([])
onMounted(async () => {
  let queryReq = {
    sortBy: 'sort',
    descending: false,
    pageIndex: 1,
    pageSize: 99999,
  }

  let table = await listMenu(queryReq)
  expanded.value = []
  menuList.value = table.list
  if (table.list) {
    table.list.forEach((item) => {
      expanded.value.push(item.id)
    })
  }
  menuTree.value = ArrayToTree(table.list, 'id', 'parentId')
  onRequest()
})
const onRequest = async () => {
  ticked.value = []
  if (row.value.menus) {
    row.value.menus.forEach((item) => {
      if (!item.id) {
        return
      }
      ticked.value.push(item.id)
      if (item.button) {
        item.button.forEach((o) => {
          ticked.value.push(o.id)
        })
      }
    })
  }
}
const handleAll = () => {
  ticked.value = []
  menuList.value.forEach((item) => {
    ticked.value.push(item.id)
    if (item.button) {
      item.button.forEach((o) => {
        ticked.value.push(o.id)
      })
    }
  })
}
const handleClear = () => {
  ticked.value = []
}
const emit = defineEmits(['submit', 'update:row'])

const handleSave = async () => {
  // let list = [];
  // ticked.value.forEach((o) => {
  //   list.push({ id: o });
  // });
  row.value.permission = ticked.value
  await setRoleMenu(row.value)
}
</script>
