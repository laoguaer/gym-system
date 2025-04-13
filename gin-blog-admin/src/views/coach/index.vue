<script setup>
import { h, onMounted, ref } from 'vue'
import { NButton, NCheckbox, NCheckboxGroup, NForm, NFormItem, NImage, NInput, NSelect, NSpace, NSwitch, NTag } from 'naive-ui'

import CommonPage from '@/components/common/CommonPage.vue'
import QueryItem from '@/components/crud/QueryItem.vue'
import CrudModal from '@/components/crud/CrudModal.vue'
import CrudTable from '@/components/crud/CrudTable.vue'

import { convertImgUrl, formatDate } from '@/utils'
import { useCRUD } from '@/composables'
import api from '@/api'

defineOptions({ name: '教练列表' })

const $table = ref(null)
const queryItems = ref({
  name: '',
  // 根据CoachQuery结构更新查询参数
})

const {
  modalVisible,
  modalLoading,
  handleSave,
  handleEdit,
  modalForm,
  modalFormRef,
} = useCRUD({
  name: '教练',
  doUpdate: api.updateCoach,
  refresh: () => $table.value?.handleSearch(),
})

// 教练相关配置已根据AdminCoachVO结构更新

onMounted(() => {
  $table.value?.handleSearch()
})

const columns = [
  {
    title: '头像',
    key: 'avatar',
    width: 30,
    align: 'center',
    render(row) {
      return h(NImage, {
        'height': 30,
        'imgProps': { style: { 'border-radius': '3px' } },
        'src': convertImgUrl(row.avatar),
        'fallback-src': 'http://dummyimage.com/400x400', // 加载失败
        'show-toolbar-tooltip': true,
      })
    },
  },
  {
    title: '姓名',
    key: 'name',
    width: 60,
    align: 'center',
    ellipsis: { tooltip: true },
  },
  {
    title: '性别',
    key: 'gender',
    width: 40,
    align: 'center',
    render(row) {
      return h('span', row.gender === 1 ? '男' : '女')
    },
  },
  {
    title: '电话',
    key: 'phone',
    width: 70,
    align: 'center',
    ellipsis: { tooltip: true },
  },
  {
    title: '描述',
    key: 'desc',
    width: 100,
    align: 'center',
    ellipsis: { tooltip: true },
  },
  {
    title: '加入时间',
    key: 'join_time',
    width: 120,
    align: 'center',
    render(row) {
      return h('span', formatDate(row.join_time))
    },
  },
  {
    title: '创建时间',
    key: 'created_at',
    align: 'center',
    width: 70,
    render(row) {
      return h(
        NButton,
        { size: 'small', type: 'text', ghost: true },
        {
          default: () => formatDate(row.created_at),
          icon: () => h('i', { class: 'i-mdi:update' }),
        },
      )
    },
  },
  {
    title: '禁用',
    key: 'is_disable',
    width: 30,
    align: 'center',
    fixed: 'left',
    render(row) {
      return h(NSwitch, {
        size: 'small',
        rubberBand: false,
        value: row.is_disable,
        loading: !!row.publishing,
        onUpdateValue: () => handleUpdateDisable(row),
      })
    },
  },
  {
    title: '操作',
    key: 'actions',
    width: 60,
    align: 'center',
    fixed: 'right',
    render(row) {
      return [
        h(
          NButton,
          {
            size: 'small',
            type: 'primary',
            onClick: () => handleEdit(row),
          },
          {
            default: () => '编辑',
            icon: () => h('i', { class: 'i-material-symbols:edit-outline' }),
          },
        ),
      ]
    },
  },
]

// 修改教练禁用状态
async function handleUpdateDisable(row) {
  if (!row.id) {
    return
  }
  row.publishing = true
  row.is_disable = !row.is_disable
  try {
    await api.updateCoachDisable(row.id, row.is_disable)
    $message?.success(row.is_disable ? '已禁用该教练' : '已取消禁用该教练')
    $table.value?.handleSearch()
  }
  catch (err) {
    row.is_disable = !row.is_disable
    console.error(err)
  }
  finally {
    row.publishing = false
  }
}
</script>

<template>
  <CommonPage title="教练列表">
    <CrudTable
      ref="$table"
      v-model:query-items="queryItems"
      :columns="columns"
      :get-data="api.getCoaches"
    >
      <template #queryBar>
        <QueryItem label="教练姓名" :label-width="80">
          <NInput v-model:value="queryItems.name" clearable placeholder="请输入教练姓名" />
        </QueryItem>
      </template>
    </CrudTable>

    <CrudModal
      v-model:visible="modalVisible"
      :loading="modalLoading"
      title="教练详情"
      @save="handleSave"
    >
      <NForm
        ref="modalFormRef"
        :model="modalForm"
        label-placement="left"
        label-width="auto"
        require-mark-placement="right-hanging"
        :style="{
          maxWidth: '640px',
        }"
      >
        <NFormItem label="姓名" path="name">
          <NInput v-model:value="modalForm.name" />
        </NFormItem>
        <NFormItem label="电话" path="phone">
          <NInput v-model:value="modalForm.phone" />
        </NFormItem>
        <NFormItem label="头像" path="avatar">
          <NInput v-model:value="modalForm.avatar" />
        </NFormItem>
        <NFormItem label="描述" path="desc">
          <NInput
            v-model:value="modalForm.desc"
            type="textarea"
            placeholder="请输入描述信息"
          />
        </NFormItem>
        <NFormItem label="加入时间" path="join_time">
          <NInput v-model:value="modalForm.join_time" disabled />
        </NFormItem>
      </NForm>
    </CrudModal>
  </CommonPage>
</template>
