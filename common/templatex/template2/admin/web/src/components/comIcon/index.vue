<template>
  <div :class="inSet ? 'q-gutter-xs  myIcon' : 'q-gutter-xs'">
    <div class="row">
      <div class="col q-ma-sm">
        <q-input
          v-model="filter"
          label="搜索"
          dense
          outlined
          clearable
          style="width: 100%"
        />
      </div>
      <div v-if="!inSet" class="col">
        <q-chip dense color="teal" size="lg" :icon="icon"
          >已选图标：{{ icon }}</q-chip
        >
      </div>
    </div>
    <q-btn
      v-for="item in iconData"
      flat
      dense
      color="primary"
      :key="item.name"
      :icon="item.name"
      size="md"
      @click="copyIcon(item)"
    />
    <div class="q-pa-xs flex flex-center">
      <q-pagination
        v-model="current"
        :max="max"
        input
        input-class="text-orange-10"
      />
    </div>
  </div>
</template>

<script setup>
import materialIcons from '@quasar/quasar-ui-qiconpicker/src/components/icon-set/material-icons'
import { ref, computed, watch, toRefs } from 'vue'
import { copyToClipboard, useQuasar } from 'quasar'
import { useI18n } from 'vue-i18n'
import { searchReg } from 'src/utils/str'

const props = defineProps({
  getIcon: String,
  inSet: Boolean,
})
const { getIcon, inSet } = toRefs(props)
const emit = defineEmits(['update:getIcon'])

const $q = useQuasar()

const icon = ref('')
const { t } = useI18n()
const filter = ref('')

const pageSize = computed(() => {
  if (inSet.value) {
    return 56
  }
  return 300
})
const current = ref(1)
const max = computed(() => {
  return Math.ceil(materialIcons.icons.length / pageSize.value)
})
const iconData = computed(() => {
  let list = materialIcons.icons
  if (filter.value && filter.value.length > 1) {
    list = materialIcons.icons.filter((o) => searchReg(o.name, filter.value))
  }
  return list.slice(
    (current.value - 1) * pageSize.value,
    current.value * pageSize.value
  )
})
const copyIcon = (item) => {
  if (item) {
    icon.value = item.name
    emit('update:getIcon', item.name)
    copyToClipboard(item.name)
      .then(() => {
        $q.notify({
          type: 'positive',
          message: t('CopyToClipboard') + ' ' + t('Success') + ': ' + item.name,
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
<style lang="scss" scoped>
.myIcon {
  width: 300px;
  max-width: 300px;
}
</style>
