<template>
  <q-layout view="lHh lpR lFr" class="shadow-2 rounded-borders">
    <q-header elevated :class="darkTheme">
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          @click="toggleLeftDrawer = !toggleLeftDrawer"
        />
        <q-toolbar-title
          shrink
          class="text-bold text-italic cursor-pointer text-deep-purple-1"
          style="padding: 0 5px"
        >
          细胞分子结构管理
        </q-toolbar-title>
        <q-separator vertical inset spaced />

        <q-breadcrumbs
          v-if="findCurrentTopMenu"
          active-color="white"
          class="text-lime justify-start items-start"
        >
          <q-breadcrumbs-el icon="home" to="/" />
          <q-breadcrumbs-el
            :label="selectOptionLabel(findCurrentTopMenu)"
            :class="darkTheme"
          />
          <q-breadcrumbs-el :label="selectRouteLabel(route)" />
        </q-breadcrumbs>
        <q-space />
        <div class="q-gutter-sm row items-center no-wrap">
          <q-btn
            dense
            round
            flat
            :icon="$q.fullscreen.isActive ? 'fullscreen_exit' : 'fullscreen'"
            @click="$q.fullscreen.toggle()"
          />
          <settings />
          <UserMenu />
        </div>
      </q-toolbar>
      <!-- header下面的标签页 -->
      <div class="row bg-white">
        <TabMenu />
      </div>
      <q-img src="images/dna6.jpg" class="absolute-top header-image" />
    </q-header>
    <q-drawer
      v-model="toggleLeftDrawer"
      show-if-above
      :breakpoint="500"
      :width="250"
      content-class="bg-grey-3"
    >
      <SideBarLeft :topMenuChildren="topMenu" />
      <q-img class="absolute-top" src="images/dna1.jpg" style="height: 150px">
        <div class="absolute-bottom bg-transparent">
          <q-avatar size="56px" class="q-mb-sm">
            <q-img src="images/avatar.png" class="" />
          </q-avatar>
          <div class="text-weight-bold">{{ userInfo.nickName }}</div>
        </div>
      </q-img>
    </q-drawer>

    <q-page-container>
      <router-view v-slot="{ Component, route }">
        <keep-alive>
          <component
            :is="Component"
            :key="route.name"
            v-if="route.meta.keepAlive"
          />
        </keep-alive>
        <component
          :is="Component"
          :key="route.name"
          v-if="!route.meta.keepAlive"
        />
      </router-view>
      <!-- <router-view /> -->
      <q-page-sticky position="bottom-right" :offset="fabPos" class="column">
        <q-page-scroller
          position="bottom-right"
          :scroll-offset="150"
          :offset="[0, -80]"
        >
          <q-btn
            push
            fab
            glossy
            padding="xs"
            rounded
            icon="keyboard_arrow_up"
            color="primary"
            v-touch-pan.prevent.mouse="moveFab"
          >
            <q-tooltip content-class="bg-accent">返回顶部</q-tooltip>
          </q-btn>
        </q-page-scroller>
      </q-page-sticky>
    </q-page-container>
  </q-layout>
</template>

<script>
export default {
  name: 'MainLayout',
}
</script>
<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { usePermissionStore } from 'src/stores/permission'
import XEUtils from 'xe-utils'
import { useQuasar, LocalStorage } from 'quasar'
import TabMenu from './tabMenu.vue'
import SideBarLeft from './SideBarLeft/index.vue'
import settings from './settings.vue'
import useDarkTheme from 'src/composables/useDarkTheme'
// import userProfile from "./UserProfile/index.vue";
import UserMenu from './UserMenu.vue'
import useCommon from 'src/composables/useCommon'
import { useUserStore } from 'src/stores/user'
const userStore = useUserStore()
const userInfo = ref({})

const { selectOptionLabel, selectRouteLabel } = useCommon()

const toggleLeftDrawer = ref(false)
const route = useRoute()
const router = useRouter()
const topMenuChildren = ref({})
const currentTopMenu = ref('')
const permissionStore = usePermissionStore()
const { darkTheme } = useDarkTheme()
const fabPos = ref([3, 80])
const $q = useQuasar()

onMounted(async () => {
  const localDark = LocalStorage.getItem('sub-dark-theme') || false
  $q.dark.set(localDark)
  currentTopMenu.value = findCurrentTopMenu.value?.name
  topMenuChildren.value = topMenu.value.filter((item) => {
    return item.name === currentTopMenu.value
  })[0]?.children
  userInfo.value = await userStore.GetInfo()
})
watch(route, () => {
  currentTopMenu.value = findCurrentTopMenu.value?.name
  topMenuChildren.value = topMenu.value.filter((item) => {
    return item.name === currentTopMenu.value
  })[0]?.children
})

const topMenu = computed(() => {
  return permissionStore.topMenu
})
const findCurrentTopMenu = computed(() => {
  const name = route.name
  return XEUtils.searchTree(topMenu.value, (item) => item.name === name)[0]
})

const moveFab = (ev) => {
  fabPos.value = [fabPos.value[0] - ev.delta.x, fabPos.value[1] - ev.delta.y]
}
</script>
<style lang="scss">
.header-image {
  height: 100%;
  z-index: -1;
  opacity: 0.2;
  filter: grayscale(100%);
}
</style>
