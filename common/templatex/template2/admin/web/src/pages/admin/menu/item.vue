<template>
  <q-tr v-show="menu.parent.expand">
    <q-td key="index" :props="menuProps" :auto-width="true">
      <q-btn
        dense
        round
        flat
        v-if="menu.children"
        :icon="menu.expand ? 'remove' : 'add'"
        @click="toggleExpand(menu)"
      />
    </q-td>
    <q-td :auto-width="true" :style="paddingLeft">{{ menu.title }}</q-td>
    <q-td :auto-width="true" text-color="red">
      <q-badge color="blue" :label="menu.name" @click="copyIcon(menu)" />
    </q-td>
    <q-td :auto-width="true" key="icon" :props="menuProps">
      <q-icon :name="menu.icon" size="sm" color="t-grey" />
    </q-td>

    <q-td key="type" :auto-width="true">{{ showDictType(menu.type) }}</q-td>
    <q-td :auto-width="true" class="ellipsis url-class" key="url">
      <div class="text-left" style="white-space: normal">{{ menu.path }}</div>
    </q-td>
    <q-td :auto-width="true" class="ellipsis url-class" key="component">{{
      menu.component
    }}</q-td>
    <q-td key="hidden" :props="menuProps" :auto-width="true">
      <q-chip
        dense
        text-color="white"
        :color="menu.hidden ? 'grey' : 'positive'"
      >
        {{ showDictHidden(menu.hidden) }}
      </q-chip>
    </q-td>
    <q-td key="keepAlive">
      <q-chip
        dense
        text-color="white"
        :color="menu.keepAlive ? 'positive' : 'grey'"
      >
        {{ showDictBool(menu.keepAlive) }}
      </q-chip>
    </q-td>
    <q-td class="ellipsis url-class" key="button">
      <div
        v-for="(item, index) in menu.button"
        :key="index"
        class="text-left q-ml-sm"
        style="white-space: normal"
      >
        <q-chip
          color="positive"
          :icon="item.icon"
          clickable
          @click="copyIcon(item)"
        >
          {{ item.title }}-{{ item.name }}</q-chip
        >
      </div>
    </q-td>
    <q-td key="opt" :props="menuProps" :auto-width="true">
      <q-btn
        flat
        round
        dense
        color="primary"
        icon="edit"
        @click.stop="edit(menu)"
      >
        <q-tooltip>编辑</q-tooltip>
      </q-btn>
      <q-btn
        flat
        round
        dense
        color="primary"
        :disable="menu.type === 'M'"
        icon="add"
        @click.stop="addChild(menu)"
      >
        <q-tooltip>添加</q-tooltip>
      </q-btn>
      <com-del label="菜单" @confirm="del(menu)" />
    </q-td>
  </q-tr>
</template>

<script>
export default {
  name: 'menu-item',
}
</script>
<script setup>
import { reactive, ref, computed, toRefs, onMounted } from 'vue'
import { DictOptions, getDictLabel } from 'src/utils/dict'
import { copyToClipboard, useQuasar } from 'quasar'
import { useI18n } from 'vue-i18n'
const $q = useQuasar()
const { t } = useI18n()
const props = defineProps({
  menu: {
    type: Object,
    required: true,
  },
  menuProps: {
    type: Object,
    required: true,
  },
})
const dictOptions = ref({})
onMounted(async () => {
  dictOptions.value = await DictOptions()
})
const showDictType = (val) => {
  if (!val) {
    val = false
  }
  return getDictLabel(dictOptions.value.menu_type, val)
}
const showDictHidden = (val) => {
  if (!val) {
    val = false
  }
  return getDictLabel(dictOptions.value.hidden, val)
}
const showDictBool = (val) => {
  if (!val) {
    val = false
  }
  return getDictLabel(dictOptions.value.sys_enabled, val)
}
const paddingLeft = computed(() => {
  if (props.menu.level > 1) {
    return {
      paddingLeft: `${(props.menu.level - 1) * 40}px`,
    }
  }
  return {}
})

const emit = defineEmits(['addChild', 'edit', 'del', 'toggleExpand'])
const addChild = (menu) => {
  emit('addChild', menu)
}
const edit = (menu) => {
  emit('edit', menu)
}
const del = (menu) => {
  emit('del', menu)
}

const toggleExpand = (menu) => {
  emit('toggleExpand', menu)
}
const copyIcon = (item) => {
  if (item) {
    copyToClipboard(item.name)
      .then(() => {
        $q.notify({
          type: 'positive',
          message: t('CopyToClipboard') + ': ' + item.name,
        })
      })
      .catch(() => {
        $q.notify({
          type: 'negative',
          message: t('CopyToClipboard') + ' ' + t('Failed'),
        })
      })
  }
}
</script>
