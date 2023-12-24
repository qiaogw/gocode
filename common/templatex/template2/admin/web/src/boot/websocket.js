import { boot } from 'quasar/wrappers'

import { io } from 'socket.io-client'
// 链接 服务端
const socket = io(process.env.BASE_WS_URL, {
  query: { fd: '就看见了就立刻' },
  transports: ['websocket', 'polling'],
})

// "async" is optional;
// more info on params: https://v2.quasar.dev/quasar-cli/boot-files
export default boot(({ app }) => {
  app.config.globalProperties.$socket = socket
})
export { socket }
