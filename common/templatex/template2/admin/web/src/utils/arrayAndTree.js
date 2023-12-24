import XEUtils from 'xe-utils'
// import { uniqueId } from 'lodash'
import _ from 'lodash'

// export const HandleAsideMenu = function (menuData, key, parentKey) {
//   // 将列表数据转换为树形数据
//   // 处理菜单成树，key值为name，parentKey为parentCode
//   const menu = ArrayToTree(menuData, key, parentKey)
//   return checkPathAndChildren(menu)
// }

// function checkPathAndChildren(menu) {
//   return menu.map((m) => ({
//     ...m,
//     path: m.path || uniqueId('sub-null-path-'),
//     ...(m.children ? { children: checkPathAndChildren(m.children) } : {}),
//   }))
// }

export const ArrayToTree = (arrayData, key, parentKey) => {
  let arr = XEUtils.clone(arrayData, true)
  const data = XEUtils.toArrayTree(arr, {
    key: key,
    parentKey: parentKey,
    strict: true,
    sortKey: 'sort',
  })

  return data
}

export const Array2Tree = (
  arrayData,
  key = 'id',
  parentKey = 'parentId',
  children = 'children',
  sortKey
) => {
  const treeData = _.cloneDeep(arrayData)
  const tree = []
  const childrenOf = {}

  treeData.forEach((item) => {
    const newItem = { ...item }
    const { [key]: id, [parentKey]: parentId } = newItem
    childrenOf[id] = childrenOf[id] || []
    newItem[children] = childrenOf[id]
    if (parentId) {
      childrenOf[parentId] = childrenOf[parentId] || []
      childrenOf[parentId].push(newItem)
    } else {
      tree.push(newItem)
    }
    // console.log('newItem:', newItem)
  })

  // console.log('tree:', tree)
  if (sortKey) {
    tree.forEach((node) => {
      node[children] = _.sortBy(node[children], sortKey)
    })
  }

  return tree
}

export const TreeToArray = (treeData) => {
  const data = XEUtils.toTreeArray(treeData)
  return data
}

export const ChangeNullChildren2Array = (data) => {
  if (!data) {
    return data
  }
  const index = data.length
  for (let i = index - 1; i >= 0; i--) {
    if (data[i].children === null) {
      data[i].children = []
    } else {
      ChangeNullChildren2Array(data[i].children)
    }
  }
  return data
}
