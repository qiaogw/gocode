<!-- eslint-disable vue/no-side-effects-in-computed-properties -->
<template>
  <q-card-section>
    <q-scroll-area style="width: 100%; height: 70vh" :thumb-style="thumbStyle">
      <q-splitter v-model="splitterModel">
        <template v-slot:before>
          <q-tabs
            v-model="apiTab"
            dense
            vertical
            class="text-grey"
            active-color="primary"
            indicator-color="primary"
          >
            <q-tab
              v-for="(item, index) in apiData"
              :name="item.module"
              :label="item.module + getThisTickedNumber(item)"
              :key="index"
            />
          </q-tabs>
        </template>
        <template v-slot:after>
          <q-tab-panels
            v-model="apiTab"
            animated
            swipeable
            vertical
            transition-prev="jump-up"
            transition-next="jump-up"
          >
            <q-tab-panel
              v-for="(item, index) in apiData"
              :name="item.module"
              :key="index"
            >
              <q-tree
                dense
                :nodes="item.children"
                default-expand-all
                node-key="id"
                selected-color="primary"
                v-if="item.children.length !== 0"
                tick-strategy="strict"
                v-model:ticked="ticked"
              >
                <template v-slot:default-header="prop">
                  <div class="row items-center">
                    <q-chip text-color="white" dense color="primary">{{
                      prop.node.title
                    }}</q-chip>
                    <q-chip color="positive" dense>
                      {{ prop.node.method }}
                    </q-chip>
                    <q-chip dense text-color="white" color="accent">
                      {{ prop.node.path }}
                    </q-chip>
                  </div>
                </template>
              </q-tree>
            </q-tab-panel>
          </q-tab-panels>
        </template>
      </q-splitter>
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
import { listApi, getApiByRole } from 'src/api/admin/api'
import { setRoleApi } from 'src/api/admin/role'
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

const ticked = ref([])
const expanded = ref([])
const dataTree = ref([])
const dataList = ref([])
const splitterModel = ref(20)
const apiTab = ref('')

onMounted(async () => {
  let queryReq = {
    sortBy: 'id',
    descending: false,
    pageIndex: 1,
    pageSize: 99999,
  }

  let table = await listApi(queryReq)
  expanded.value = []
  dataList.value = table.list
  if (table.list) {
    table.list.forEach((item) => {
      expanded.value.push(item.id)
    })
  }
  dataTree.value = ArrayToTree(table.list, 'id', 'parentId')
  // onRequest();
  if (row.value.apis) {
    row.value.apis.forEach((item) => {
      ticked.value.push(item.id)
      if (item.button) {
        item.button.forEach((o) => {
          ticked.value.push(o.id)
        })
      }
    })
  }
})

const apiData = computed(() => {
  if (dataList.value.length) {
    const data = dataList.value
    for (let item of data) {
      item.trueId = 'g:' + item.module + 'p:' + item.path + 'm:' + item.method
    }
    const apiTree = []
    for (let d of data) {
      if (apiTree.find((item) => item.module === d.module) === undefined) {
        apiTree.push({
          module: d.module,
          children: [],
        })
      }
    }
    for (let d of data) {
      for (let a of apiTree) {
        if (a.module === d.module) {
          a.children.push(d)
        }
      }
    }
    // eslint-disable-next-line vue/no-side-effects-in-computed-properties
    apiTab.value = apiTree[0].module

    return apiTree
  }
  return []
})

const getThisTickedNumber = computed(() => {
  return (api) => {
    const allNumber = api.children.length
    var tickedNumber = 0
    for (let t of ticked.value) {
      if (api.children.find((item) => item.id === t) !== undefined) {
        tickedNumber++
      }
    }
    return '(' + tickedNumber + '/' + allNumber + ')'
  }
})
const handleAll = () => {
  ticked.value = []
  dataList.value.forEach((item) => {
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
  row.value.permission = ticked.value
  await setRoleApi(row.value)
}
</script>
